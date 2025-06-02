package order

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
