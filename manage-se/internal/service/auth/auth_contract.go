package auth

import (
	"context"
	"manage-se/internal/presentations"
	"manage-se/internal/provider/user"
)

type Auth interface {
	VerifyToken(ctx context.Context, input presentations.Verify) (*user.UserDetail, error)
	Login(ctx context.Context, input presentations.Login) (*user.UserDetailToken, error)
	Register(ctx context.Context, input presentations.Register) (*user.UserDetail, error)
}
