package user

import (
	"order-se/pkg/databasex"
)

type user struct {
	db databasex.Adapter
}

func NewUser(db databasex.Adapter) *user {
	return &user{
		db: db,
	}
}

func (c *user) Sortable(field string) bool {
	switch field {
	case "created_at", "updated_at", "name", "email", "phone":
		return true
	default:
		return false
	}

}

func (c *user) Searchable(field string) bool {
	switch field {
	case "name", "role_id", "phone", "email":
		return true
	default:
		return false
	}

}
