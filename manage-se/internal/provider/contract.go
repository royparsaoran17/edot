package provider

import (
	"context"
	"manage-se/internal/common"
	"manage-se/internal/presentations"
	"manage-se/internal/provider/auth"
)

type Auth interface {
	Login(ctx context.Context, input presentations.Login) (*auth.UserDetailToken, error)
	Verify(ctx context.Context, input presentations.Verify) (*auth.UserDetail, error)

	CreateUser(ctx context.Context, input presentations.UserCreate) (*auth.UserDetail, error)
	GetListUsers(ctx context.Context, meta *common.Metadata) ([]auth.User, error)

	GetListRoles(ctx context.Context) ([]auth.Role, error)
	CreateRole(ctx context.Context, input presentations.RoleCreate) (*auth.Role, error)
}
