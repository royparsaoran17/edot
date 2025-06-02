package presentations

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type WarehouseUpdate struct {
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

func (r *WarehouseUpdate) Validate() error {
	return validation.Errors{
		"name": validation.Validate(&r.Name, validation.Required),
	}.Filter()
}

type WarehouseCreate struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

func (r *WarehouseCreate) Validate() error {
	return validation.Errors{
		"name": validation.Validate(&r.Name, validation.Required),
	}.Filter()
}

type WarehouseCreateMoveStock struct {
	ProductID       string `db:"product_id,omitempty" json:"product_id"`
	FromWarehouseID string `db:"from_warehouse_id,omitempty" json:"from_warehouse_id"`
	ToWarehouseID   string `db:"to_warehouse_id,omitempty" json:"to_warehouse_id"`
	Quantity        int    `db:"quantity,omitempty" json:"quantity"`
	Notes           string `db:"notes,omitempty" json:"notes"`
}

func (r *WarehouseCreateMoveStock) Validate() error {
	return validation.Errors{
		"product_id":        validation.Validate(&r.ProductID, validation.Required),
		"from_warehouse_id": validation.Validate(&r.FromWarehouseID, validation.Required),
		"to_warehouse_id":   validation.Validate(&r.ToWarehouseID, validation.Required),
		"quantity":          validation.Validate(&r.Quantity, validation.Required),
		"notes":             validation.Validate(&r.Notes, validation.Required),
	}.Filter()
}

type WarehouseStock struct {
	ID        string `json:"id"`
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

func (r *WarehouseStock) Validate() error {
	return validation.Errors{
		"id":         validation.Validate(&r.ID, validation.Required),
		"product_id": validation.Validate(&r.ProductID, validation.Required),
		"quantity":   validation.Validate(&r.Quantity, validation.Required),
	}.Filter()
}
