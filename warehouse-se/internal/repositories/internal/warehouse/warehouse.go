package warehouse

import (
	"warehouse-se/pkg/databasex"
)

type warehouse struct {
	db databasex.Adapter
}

func NewWarehouse(db databasex.Adapter) *warehouse {
	return &warehouse{
		db: db,
	}
}

func (c *warehouse) Sortable(field string) bool {
	switch field {
	case "created_at", "updated_at", "name":
		return true
	default:
		return false
	}

}

func (c *warehouse) Searchable(field string) bool {
	switch field {
	case "name":
		return true
	default:
		return false
	}

}
