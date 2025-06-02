// Package entity
// Automatic generated
package entity

import (
	"time"
)

// Order entity
type Order struct {
	ID         string     `db:"id,omitempty" json:"id"`
	UserID     string     `db:"user_id,omitempty" json:"user_id"`
	Status     string     `db:"status,omitempty" json:"status"`
	TotalPrice float64    `db:"total_price,omitempty" json:"total_price"`
	CreatedAt  time.Time  `db:"created_at,omitempty" json:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at,omitempty" json:"updated_at"`
	DeletedAt  *time.Time `db:"deleted_at,omitempty" json:"deleted_at"`
}

type OrderDetail struct {
	Order
	Items []OrderItem `db:"items,omitempty" json:"items"`
}
