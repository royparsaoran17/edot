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
	ShopID    string     `db:"shop_id,omitempty" json:"shop_id"`
	IsActive  bool       `db:"is_active,omitempty" json:"is_active"`
	CreatedAt time.Time  `db:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at,omitempty" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at,omitempty" json:"deleted_at"`
}
