// Package entity
// Automatic generated
package entity

import (
	"time"
)

// Warehouse entity
type Warehouse struct {
	ID        string     `db:"id,omitempty" json:"id"`
	Name      string     `db:"name,omitempty" json:"name"`
	IsActive  bool       `db:"is_active,omitempty" json:"is_active"`
	CreatedAt time.Time  `db:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at,omitempty" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at,omitempty" json:"deleted_at"`
}
type WarehouseDetail struct {
	ID        string         `db:"id,omitempty" json:"id"`
	Name      string         `db:"name,omitempty" json:"name"`
	IsActive  bool           `db:"is_active,omitempty" json:"is_active"`
	CreatedAt time.Time      `db:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time      `db:"updated_at,omitempty" json:"updated_at"`
	DeletedAt *time.Time     `db:"deleted_at,omitempty" json:"deleted_at"`
	Products  []ProductStock `json:"products"`
}

type StockMovement struct {
	ID              string     `db:"id,omitempty" json:"id"`
	ProductID       string     `db:"product_id,omitempty" json:"product_id"`
	Product         Product    `db:"product,omitempty" json:"product"`
	FromWarehouseID string     `db:"from_warehouse_id,omitempty" json:"from_warehouse_id"`
	FromWarehouse   Warehouse  `db:"from_warehouse,omitempty" json:"from_warehouse"`
	ToWarehouseID   string     `db:"to_warehouse_id,omitempty" json:"to_warehouse_id"`
	ToWarehouse     Warehouse  `db:"to_warehouse,omitempty" json:"to_warehouse"`
	Quantity        string     `db:"quantity,omitempty" json:"quantity"`
	Notes           string     `db:"notes,omitempty" json:"notes"`
	CreatedAt       time.Time  `db:"created_at,omitempty" json:"created_at"`
	UpdatedAt       time.Time  `db:"updated_at,omitempty" json:"updated_at"`
	DeletedAt       *time.Time `db:"deleted_at,omitempty" json:"deleted_at"`
}
