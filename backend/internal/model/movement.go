package model

import "time"

type Movement struct {
	ID        int64     `db:"id" json:"id"`
	ProductID int64     `db:"product_id" json:"product_id" binding:"required"`
	Type      string    `db:"type" json:"type" binding:"required,oneof=IN OUT"`
	Quantity  int       `db:"quantity" json:"quantity" binding:"required,gt=0"`
	Note      string    `db:"note" json:"note"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
