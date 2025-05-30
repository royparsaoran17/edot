package user

import (
	"auth-se/internal/common"
	"context"
	"github.com/google/uuid"

	"auth-se/internal/entity"
	"auth-se/internal/presentations"
	"auth-se/internal/repositories"
	"github.com/pkg/errors"
)

type service struct {
	repo *repositories.Repository
}

func NewService(repo *repositories.Repository) User {
	return &service{repo: repo}
}

func (s *service) GetAllUser(ctx context.Context, meta *common.Metadata) ([]entity.User, error) {
	users, err := s.repo.User.GetAllUser(ctx, meta)
	if err != nil {
		return nil, errors.Wrap(err, "getting all users on ")
	}

	return users, nil
}

func (s *service) GetUserByID(ctx context.Context, userID string) (*entity.UserDetail, error) {
	users, err := s.repo.User.FindUserByID(ctx, userID)
	if err != nil {
		return nil, errors.Wrapf(err, "getting user id %s", userID)
	}

	return users, nil
}

func (s *service) UpdateUserByID(ctx context.Context, userID string, input presentations.UserUpdate) error {
	if err := input.Validate(); err != nil {
		return errors.Wrap(err, "validation(s) error")
	}

	_, err := s.repo.User.FindUserByID(ctx, userID)
	if err != nil {
		return errors.Wrapf(err, "getting user id %s", userID)
	}

	_, err = s.repo.Role.FindRoleByID(ctx, input.RoleID)
	if err != nil {
		return errors.Wrap(err, "creating user")

	}

	if err := s.repo.User.UpdateUser(ctx, userID, input); err != nil {
		return errors.Wrap(err, "updating user")

	}

	return nil
}

func (s *service) CreateUser(ctx context.Context, input presentations.UserCreate) (*entity.UserDetail, error) {
	if err := input.Validate(); err != nil {
		return nil, errors.Wrap(err, "validation(s) error")
	}

	_, err := s.repo.Role.FindRoleByID(ctx, input.RoleID)
	if err != nil {
		return nil, errors.Wrap(err, "creating user")

	}

	userID := uuid.NewString()
	input.ID = userID
	err = s.repo.User.CreateUser(ctx, input)
	if err != nil {
		return nil, errors.Wrap(err, "creating user")

	}

	user, err := s.repo.User.FindUserByID(ctx, userID)
	if err != nil {
		return nil, errors.Wrapf(err, "getting user id %s", userID)
	}

	return user, nil
}
