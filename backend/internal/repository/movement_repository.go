package repository

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/user/controle-estoque/backend/internal/model"
)

type MovementRepository interface {
	Register(movement *model.Movement) error
	ListByProduct(productID int64) ([]model.Movement, error)
}

type sqliteMovementRepository struct {
	db *sqlx.DB
}

func NewMovementRepository(db *sqlx.DB) MovementRepository {
	return &sqliteMovementRepository{db: db}
}

func (r *sqliteMovementRepository) Register(m *model.Movement) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// 1. Verificar estoque atual
	var currentQty int
	err = tx.Get(&currentQty, "SELECT quantity FROM products WHERE id = ?", m.ProductID)
	if err != nil {
		return fmt.Errorf("produto não encontrado: %w", err)
	}

	// 2. Calcular novo saldo
	newQty := currentQty
	if m.Type == "IN" {
		newQty += m.Quantity
	} else {
		if currentQty < m.Quantity {
			return errors.New("estoque insuficiente para esta saída")
		}
		newQty -= m.Quantity
	}

	// 3. Atualizar saldo do produto
	_, err = tx.Exec("UPDATE products SET quantity = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?", newQty, m.ProductID)
	if err != nil {
		return err
	}

	// 4. Inserir movimento
	query := `INSERT INTO movements (product_id, type, quantity, note)
	          VALUES (:product_id, :type, :quantity, :note)`
	result, err := tx.NamedExec(query, m)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	m.ID = id

	return tx.Commit()
}

func (r *sqliteMovementRepository) ListByProduct(productID int64) ([]model.Movement, error) {
	movements := []model.Movement{}
	err := r.db.Select(&movements, "SELECT * FROM movements WHERE product_id = ? ORDER BY created_at DESC", productID)
	return movements, err
}
