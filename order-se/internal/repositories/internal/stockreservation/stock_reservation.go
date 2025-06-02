package stockreservation

import (
	"order-se/pkg/databasex"
)

type stockReservation struct {
	db databasex.Adapter
}

func NewStockReservation(db databasex.Adapter) *stockReservation {
	return &stockReservation{
		db: db,
	}
}

func (c *stockReservation) Sortable(field string) bool {
	switch field {
	case "created_at", "updated_at", "name":
		return true
	default:
		return false
	}

}

func (c *stockReservation) Searchable(field string) bool {
	switch field {
	case "name":
		return true
	default:
		return false
	}

}
