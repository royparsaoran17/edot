package product

import "product-se/pkg/postgres"

type product struct {
	db postgres.Adapter
}

func NewProduct(db postgres.Adapter) *product {
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
