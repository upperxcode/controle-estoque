package model

import "time"

type Product struct {
	ID          int64     `db:"id" json:"id"`
	Name        string    `db:"name" json:"name" binding:"required"`
	Description string    `db:"description" json:"description"`
	SKU         string    `db:"sku" json:"sku" binding:"required"`
	Quantity    int       `db:"quantity" json:"quantity" binding:"min=0"`
	Price       float64   `db:"price" json:"price" binding:"required,gt=0"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
