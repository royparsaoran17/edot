package presentations

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"
)

type StockReservationCreate struct {
	ID          string    `json:"id"`
	OrderID     string    `json:"order_id"`
	ProductID   string    `json:"product_id"`
	WarehouseID string    `json:"warehouse_id"`
	Quantity    int       `json:"quantity"`
	Price       float64   `json:"amount"`
	ReservedAt  time.Time `json:"reserved_at"`
	ExpiresAt   time.Time `json:"expires_at"`
}

func (r *StockReservationCreate) Validate() error {
	return validation.Errors{
		"order_id":     validation.Validate(&r.OrderID, validation.Required),
		"product_id":   validation.Validate(&r.ProductID, validation.Required),
		"warehouse_id": validation.Validate(&r.WarehouseID, validation.Required),
		"quantity":     validation.Validate(&r.Quantity, validation.Required),
		"amount":       validation.Validate(&r.Price, validation.Required),
		"reserved_at":  validation.Validate(&r.ReservedAt, validation.Required),
		"expires_at":   validation.Validate(&r.ExpiresAt, validation.Required),
	}.Filter()
}
