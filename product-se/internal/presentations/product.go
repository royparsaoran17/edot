package presentations

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ProductUpdate struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Unit        string  `json:"unit"`
	Sku         string  `json:"sku"`
	Category    string  `json:"category"`
	IsActive    bool    `json:"is_active"`
}

func (r *ProductUpdate) Validate() error {
	return validation.Errors{
		"name":        validation.Validate(&r.Name, validation.Required),
		"description": validation.Validate(&r.Description, validation.Required),
		"price":       validation.Validate(&r.Price, validation.Required),
		"unit":        validation.Validate(&r.Unit, validation.Required),
		"sku":         validation.Validate(&r.Sku, validation.Required),
		"category":    validation.Validate(&r.Category, validation.Required),
	}.Filter()
}

type ProductCreate struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Unit        string  `json:"unit"`
	Sku         string  `json:"sku"`
	Category    string  `json:"category"`
	IsActive    bool    `json:"is_active"`
}

func (r *ProductCreate) Validate() error {
	return validation.Errors{
		"name":        validation.Validate(&r.Name, validation.Required),
		"description": validation.Validate(&r.Description, validation.Required),
		"price":       validation.Validate(&r.Price, validation.Required),
		"unit":        validation.Validate(&r.Unit, validation.Required),
		"sku":         validation.Validate(&r.Sku, validation.Required),
		"category":    validation.Validate(&r.Category, validation.Required),
	}.Filter()
}
