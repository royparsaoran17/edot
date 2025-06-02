package user

import (
	"context"
	"manage-se/internal/common"
	"manage-se/internal/provider/user"
)

type User interface {
	GetAllUser(ctx context.Context, meta *common.Metadata) ([]user.User, error)
}
