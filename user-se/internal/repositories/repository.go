package repositories

import (
	"auth-se/internal/repositories/internal/role"
	"auth-se/internal/repositories/internal/user"
	"auth-se/pkg/postgres"
)

type Repository struct {
	Role Role
	User User
}

func NewRepository(db postgres.Adapter) *Repository {
	return &Repository{
		Role: role.NewRole(db),
		User: user.NewUser(db),
	}
}
