package auth

import (
	"auth-se/internal/entity"
	"auth-se/internal/presentations"
	"context"
)

type Auth interface {
	Verify(input presentations.Verify) (*entity.UserDetail, error)
	Login(ctx context.Context, input presentations.Login) (*entity.UserDetailToken, error)
}
