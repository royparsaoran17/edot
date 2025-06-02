// Package entity
// Automatic generated
package entity

import (
	"time"
)

// Payment entity
type Payment struct {
	ID        string     `db:"id,omitempty" json:"id"`
	OrderID   string     `db:"order_id,omitempty" json:"order_id"`
	Method    string     `db:"method,omitempty" json:"method"`
	Status    string     `db:"status,omitempty" json:"status"`
	Amount    float64    `db:"amount,omitempty" json:"amount"`
	PaidAt    time.Time  `db:"paid_at,omitempty" json:"paid_at"`
	CreatedAt time.Time  `db:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at,omitempty" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at,omitempty" json:"deleted_at"`
}
