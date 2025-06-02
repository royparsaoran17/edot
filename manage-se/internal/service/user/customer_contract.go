package user

import (
	"context"
	"manage-se/internal/common"
	"manage-se/internal/provider/auth"
)

type User interface {
	GetAllUser(ctx context.Context, meta *common.Metadata) ([]auth.User, error)
}
