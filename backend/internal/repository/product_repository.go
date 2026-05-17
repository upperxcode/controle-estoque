package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/user/controle-estoque/backend/internal/model"
)

type ProductRepository interface {
	Create(product *model.Product) error
	GetByID(id int64) (*model.Product, error)
	List() ([]model.Product, error)
	Update(product *model.Product) error
	Delete(id int64) error
}

type sqliteProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) ProductRepository {
	return &sqliteProductRepository{db: db}
}

func (r *sqliteProductRepository) Create(p *model.Product) error {
	query := `INSERT INTO products (name, description, sku, quantity, price)
	          VALUES (:name, :description, :sku, :quantity, :price)`
	result, err := r.db.NamedExec(query, p)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	p.ID = id
	return nil
}

func (r *sqliteProductRepository) GetByID(id int64) (*model.Product, error) {
	var p model.Product
	err := r.db.Get(&p, "SELECT * FROM products WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *sqliteProductRepository) List() ([]model.Product, error) {
	products := []model.Product{}
	err := r.db.Select(&products, "SELECT * FROM products ORDER BY created_at DESC")
	return products, err
}

func (r *sqliteProductRepository) Update(p *model.Product) error {
	query := `UPDATE products SET name=:name, description=:description, sku=:sku,
	          quantity=:quantity, price=:price, updated_at=CURRENT_TIMESTAMP
	          WHERE id=:id`
	_, err := r.db.NamedExec(query, p)
	return err
}

func (r *sqliteProductRepository) Delete(id int64) error {
	_, err := r.db.Exec("DELETE FROM products WHERE id = ?", id)
	return err
}
