package presentations

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type OrderItemCreate struct {
	ID          string  `json:"id"`
	OrderID     string  `json:"order_id"`
	ProductID   string  `json:"product_id"`
	WarehouseID string  `json:"warehouse_id"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}

func (r *OrderItemCreate) Validate() error {
	return validation.Errors{
		"order_id":     validation.Validate(&r.OrderID, validation.Required),
		"product_id":   validation.Validate(&r.ProductID, validation.Required),
		"warehouse_id": validation.Validate(&r.WarehouseID, validation.Required),
		"quantity":     validation.Validate(&r.Quantity, validation.Required),
		"price":        validation.Validate(&r.Price, validation.Required),
	}.Filter()
}
