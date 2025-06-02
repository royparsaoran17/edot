package role

import (
	"context"
	"manage-se/internal/provider/user"
)

type Role interface {
	GetAllRole(ctx context.Context) ([]user.Role, error)
}
