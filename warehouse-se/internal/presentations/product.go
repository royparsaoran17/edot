package presentations

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ProductUpdateStock struct {
	Quantity int `json:"price"`
}

func (r *ProductUpdateStock) Validate() error {
	return validation.Errors{
		"Quantity": validation.Validate(&r.Quantity, validation.Required),
	}.Filter()
}
