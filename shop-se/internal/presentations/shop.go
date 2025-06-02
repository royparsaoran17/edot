package presentations

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ShopUpdate struct {
	Name    string `json:"name"`
	OwnerID string `json:"owner_id"`
}

func (r *ShopUpdate) Validate() error {
	return validation.Errors{
		"name": validation.Validate(&r.Name, validation.Required),
	}.Filter()
}

type ShopCreate struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	OwnerID bool   `json:"owner_id"`
}

func (r *ShopCreate) Validate() error {
	return validation.Errors{
		"name":     validation.Validate(&r.Name, validation.Required),
		"owner_id": validation.Validate(&r.OwnerID, validation.Required),
	}.Filter()
}
