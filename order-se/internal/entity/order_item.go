package entity

import "time"

type OrderItem struct {
	ID          string     `db:"id,omitempty" json:"id"`
	OrderID     string     `db:"order_id,omitempty" json:"order_id"`
	ProductID   string     `db:"product_id,omitempty" json:"product_id"`
	Product     Product    `db:"product,omitempty" json:"product"`
	WarehouseID string     `db:"warehouse_id,omitempty" json:"warehouse_id"`
	Warehouse   Warehouse  `db:"warehouse,omitempty" json:"warehouse"`
	Quantity    int        `db:"quantity,omitempty" json:"quantity"`
	Price       float64    `db:"price,omitempty" json:"price"`
	CreatedAt   time.Time  `db:"created_at,omitempty" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at,omitempty" json:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at,omitempty" json:"deleted_at"`
}
