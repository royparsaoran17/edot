package user

import (
	"auth-se/internal/common"
	"auth-se/internal/presentations"
	"context"

	"auth-se/internal/entity"
)

type User interface {
	GetAllUser(ctx context.Context, meta *common.Metadata) ([]entity.User, error)
	GetUserByID(ctx context.Context, userID string) (*entity.UserDetail, error)
	UpdateUserByID(ctx context.Context, userID string, input presentations.UserUpdate) error
	CreateUser(ctx context.Context, input presentations.UserCreate) (*entity.UserDetail, error)
}
