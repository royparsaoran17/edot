package shop

import (
	"shop-se/pkg/databasex"
)

type shop struct {
	db databasex.Adapter
}

func NewShop(db databasex.Adapter) *shop {
	return &shop{
		db: db,
	}
}

func (c *shop) Sortable(field string) bool {
	switch field {
	case "created_at", "updated_at", "name":
		return true
	default:
		return false
	}

}

func (c *shop) Searchable(field string) bool {
	switch field {
	case "name":
		return true
	default:
		return false
	}

}
