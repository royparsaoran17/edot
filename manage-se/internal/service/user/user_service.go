package user

import (
	"context"
	"manage-se/internal/common"
	"manage-se/internal/provider"
	"manage-se/internal/provider/user"

	"github.com/pkg/errors"
)

type service struct {
	provider *provider.Provider
}

func NewService(provider *provider.Provider) User {
	return &service{provider: provider}
}

func (s *service) GetAllUser(ctx context.Context, meta *common.Metadata) ([]user.User, error) {
	users, err := s.provider.User.GetListUsers(ctx, meta)
	if err != nil {
		return nil, errors.Wrap(err, "getting all users ")
	}

	return users, nil
}
