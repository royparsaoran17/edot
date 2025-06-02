package product

import (
	"order-se/pkg/databasex"
)

type product struct {
	db databasex.Adapter
}

func NewProduct(db databasex.Adapter) *product {
	return &product{
		db: db,
	}
}

func (c *product) Sortable(field string) bool {
	switch field {
	case "created_at", "updated_at", "name":
		return true
	default:
		return false
	}

}

func (c *product) Searchable(field string) bool {
	switch field {
	case "name":
		return true
	default:
		return false
	}

}
