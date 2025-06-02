package order

import (
	"order-se/pkg/databasex"
)

type order struct {
	db databasex.Adapter
}

func NewOrder(db databasex.Adapter) *order {
	return &order{
		db: db,
	}
}

func (c *order) Sortable(field string) bool {
	switch field {
	case "created_at", "updated_at", "name":
		return true
	default:
		return false
	}

}

func (c *order) Searchable(field string) bool {
	switch field {
	case "name":
		return true
	default:
		return false
	}

}
