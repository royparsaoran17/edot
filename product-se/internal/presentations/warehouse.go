package presentations

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type WarehouseUpdate struct {
	Name string `json:"name"`
}

func (r *WarehouseUpdate) Validate() error {
	return validation.Errors{
		"name": validation.Validate(&r.Name, validation.Required),
	}.Filter()
}

type WarehouseCreate struct {
	Name string `json:"name"`
}

func (r *WarehouseCreate) Validate() error {
	return validation.Errors{
		"name": validation.Validate(&r.Name, validation.Required),
	}.Filter()
}
