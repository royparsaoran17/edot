package repositories

import (
	"auth-se/internal/common"
	"auth-se/internal/entity"
	"context"

	"auth-se/internal/presentations"
)

type Role interface {
	CreateRole(ctx context.Context, input presentations.RoleCreate) error
	UpdateRole(ctx context.Context, roleID string, input presentations.RoleUpdate) error
	FindRoleByID(ctx context.Context, roleID string) (*entity.Role, error)
	GetAllRole(ctx context.Context) ([]entity.Role, error)
}

type User interface {
	CreateUser(ctx context.Context, input presentations.UserCreate) error
	FindUserByPhoneEmailPassword(ctx context.Context, value, password string) (*entity.UserDetail, error)
	FindUserByPhoneEmail(ctx context.Context, value string) (*entity.UserDetail, error)
	UpdateUser(ctx context.Context, roleID string, input presentations.UserUpdate) error
	GetAllUser(ctx context.Context, meta *common.Metadata) ([]entity.User, error)
	FindUserByID(ctx context.Context, userID string) (*entity.UserDetail, error)
}
