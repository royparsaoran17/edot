package user

import "auth-se/pkg/postgres"

type user struct {
	db postgres.Adapter
}

func NewUser(db postgres.Adapter) *user {
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
