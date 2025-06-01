// Package entity
// Automatic generated
package entity

import (
	"time"
)

// Product entity
type Product struct {
	ID          string     `db:"id,omitempty" json:"id"`
	Name        string     `db:"name,omitempty" json:"name"`
	Description string     `db:"description,omitempty" json:"description"`
	Price       float64    `db:"price,omitempty" json:"price"`
	Unit        string     `db:"unit,omitempty" json:"unit"`
	Sku         string     `db:"sku,omitempty" json:"sku"`
	Category    string     `db:"category,omitempty" json:"category"`
	IsActive    bool       `db:"is_active,omitempty" json:"is_active"`
	CreatedAt   time.Time  `db:"created_at,omitempty" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at,omitempty" json:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at,omitempty" json:"deleted_at"`
}

type ProductStock struct {
	ID          string     `db:"id,omitempty" json:"id"`
	ProductID   string     `db:"product_id,omitempty" json:"product_id"`
	WarehouseID string     `db:"warehouse_id,omitempty" json:"warehouse_id"`
	Quantity    int        `db:"quantity,omitempty" json:"quantity"`
	CreatedAt   time.Time  `db:"created_at,omitempty" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at,omitempty" json:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at,omitempty" json:"deleted_at"`
}
