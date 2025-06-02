package orderitem

import (
	"order-se/pkg/databasex"
)

type orderItem struct {
	db databasex.Adapter
}

func NewOrderItem(db databasex.Adapter) *orderItem {
	return &orderItem{
		db: db,
	}
}

func (c *orderItem) Sortable(field string) bool {
	switch field {
	case "created_at", "updated_at", "name":
		return true
	default:
		return false
	}

}

func (c *orderItem) Searchable(field string) bool {
	switch field {
	case "name":
		return true
	default:
		return false
	}

}
