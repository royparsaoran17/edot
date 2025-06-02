package presentations

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
)

type OrderCreate struct {
	ID         string  `json:"id"`
	UserID     string  `json:"user_id"`
	Status     string  `json:"status"`
	TotalPrice float64 `json:"total_price"`
}

func (r *OrderCreate) Validate() error {
	return validation.Errors{
		"user_id":     validation.Validate(&r.UserID, validation.Required),
		"status":      validation.Validate(&r.Status, validation.Required),
		"total_price": validation.Validate(&r.TotalPrice, validation.Required),
	}.Filter()
}

type Order struct {
	UserID string          `json:"user_id"`
	Items  []OrderItemData `json:"items"`
}

type OrderItemData struct {
	ProductID   string `json:"product_id"`
	WarehouseID string `json:"warehouse_id"`
	Quantity    int    `json:"quantity"`
}

func (o *Order) Validate() error {
	return validation.Errors{
		"user_id": validation.Validate(o.UserID, validation.Required),
		"items": validation.Validate(o.Items,
			validation.Required,
			validation.Each(validation.By(validateItem))),
	}.Filter()
}

func validateItem(value interface{}) error {
	item, ok := value.(OrderItemData)
	if !ok {
		return errors.New("invalid item data")
	}
	return validation.Errors{
		"product_id":   validation.Validate(item.ProductID, validation.Required),
		"warehouse_id": validation.Validate(item.WarehouseID, validation.Required),
		"quantity":     validation.Validate(item.Quantity, validation.Required, validation.Min(1)),
	}.Filter()
}

type OrderUpdate struct {
	Status string `json:"status"`
}
