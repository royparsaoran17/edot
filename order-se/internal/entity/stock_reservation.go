package entity

import "time"

type StockReservation struct {
	ID          string     `db:"id,omitempty" json:"id"`
	OrderID     string     `db:"order_id,omitempty" json:"order_id"`
	ProductID   string     `db:"product_id,omitempty" json:"product_id"`
	WarehouseID string     `db:"warehouse_id,omitempty" json:"warehouse_id"`
	Quantity    int        `db:"quantity,omitempty" json:"quantity"`
	Price       float64    `db:"price,omitempty" json:"price"`
	ReservedAt  time.Time  `db:"reserved_at,omitempty" json:"reserved_at"`
	ExpiresAt   time.Time  `db:"expires_at,omitempty" json:"expires_at"`
	CreatedAt   time.Time  `db:"created_at,omitempty" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at,omitempty" json:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at,omitempty" json:"deleted_at"`
}
