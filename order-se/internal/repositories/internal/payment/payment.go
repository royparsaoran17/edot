package payment

import (
	"order-se/pkg/databasex"
)

type payment struct {
	db databasex.Adapter
}

func NewPayment(db databasex.Adapter) *payment {
	return &payment{
		db: db,
	}
}

func (c *payment) Sortable(field string) bool {
	switch field {
	case "created_at", "updated_at", "name":
		return true
	default:
		return false
	}

}

func (c *payment) Searchable(field string) bool {
	switch field {
	case "name":
		return true
	default:
		return false
	}

}
