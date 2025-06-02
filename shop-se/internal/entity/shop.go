// Package entity
// Automatic generated
package entity

import (
	"time"
)

// Shop entity
type Shop struct {
	ID        string     `db:"id,omitempty" json:"id"`
	Name      string     `db:"name,omitempty" json:"name"`
	OwnerID   string     `db:"owner_id,omitempty" json:"owner_id"`
	CreatedAt time.Time  `db:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at,omitempty" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at,omitempty" json:"deleted_at"`
}

type ShopDetail struct {
	ID        string     `db:"id,omitempty" json:"id"`
	Name      string     `db:"name,omitempty" json:"name"`
	OwnerID   string     `db:"owner_id,omitempty" json:"owner_id"`
	Owner     User       `db:"owner,omitempty" json:"owner"`
	CreatedAt time.Time  `db:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at,omitempty" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at,omitempty" json:"deleted_at"`

	Warehouses []Warehouse `db:"warehouses" json:"warehouses"`
}
