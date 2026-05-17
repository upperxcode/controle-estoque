package repository

import (
	"os"
	"testing"

	"github.com/user/controle-estoque/backend/internal/database"
	"github.com/user/controle-estoque/backend/internal/model"
)

func TestMovementRepository(t *testing.T) {
	os.Remove("./estoque.db")
	db, _ := database.Connect()
	defer db.Close()

	productRepo := NewProductRepository(db)
	movementRepo := NewMovementRepository(db)

	p := &model.Product{
		Name:  "Cadeira Gamer",
		SKU:   "CAD-001",
		Price: 1200.0,
		Quantity: 5,
	}
	productRepo.Create(p)

	t.Run("Register IN", func(t *testing.T) {
		m := &model.Movement{
			ProductID: p.ID,
			Type:      "IN",
			Quantity:  10,
			Note:      "Chegada de carga",
		}
		err := movementRepo.Register(m)
		if err != nil {
			t.Fatalf("Erro ao registrar entrada: %v", err)
		}

		updated, _ := productRepo.GetByID(p.ID)
		if updated.Quantity != 15 {
			t.Errorf("Esperava estoque 15, recebeu %d", updated.Quantity)
		}
	})

	t.Run("Register OUT Success", func(t *testing.T) {
		m := &model.Movement{
			ProductID: p.ID,
			Type:      "OUT",
			Quantity:  3,
		}
		err := movementRepo.Register(m)
		if err != nil {
			t.Fatalf("Erro ao registrar saída: %v", err)
		}

		updated, _ := productRepo.GetByID(p.ID)
		if updated.Quantity != 12 {
			t.Errorf("Esperava estoque 12, recebeu %d", updated.Quantity)
		}
	})

	t.Run("Register OUT Insufficient", func(t *testing.T) {
		m := &model.Movement{
			ProductID: p.ID,
			Type:      "OUT",
			Quantity:  20,
		}
		err := movementRepo.Register(m)
		if err == nil {
			t.Fatal("Esperava erro de estoque insuficiente, recebeu nil")
		}

		// Validar que o saldo não mudou (rollback implícito)
		updated, _ := productRepo.GetByID(p.ID)
		if updated.Quantity != 12 {
			t.Errorf("Esperava estoque mantido em 12, recebeu %d", updated.Quantity)
		}
	})
}
