package warehouse

import "product-se/pkg/postgres"

type warehouse struct {
	db postgres.Adapter
}

func NewWarehouse(db postgres.Adapter) *warehouse {
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
