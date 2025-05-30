package auth

import (
	config "auth-se/internal/appctx"
	"auth-se/internal/common"
	"auth-se/internal/consts"
	"auth-se/internal/entity"
	"auth-se/internal/presentations"
	"auth-se/internal/repositories"
	"auth-se/pkg/jwt"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
)

type service struct {
	repo *repositories.Repository
	cfg  config.Common
}

func NewService(repo *repositories.Repository, cfg config.Common) Auth {
	return &service{repo: repo, cfg: cfg}
}

func (s *service) Verify(input presentations.Verify) (*entity.UserDetail, error) {
	if err := input.Validate(); err != nil {
		return nil, errors.Wrap(err, "validation(s) error")
	}

	claim, err := jwt.VerifyToken(input.Token, s.cfg.JWTKey)
	if err != nil {
		return nil, errors.Wrap(err, "verify token")
	}

	claims := *claim

	var user entity.UserDetail
	jsonb, err := json.Marshal(claims["data"].(map[string]interface{}))
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	if err := json.Unmarshal(jsonb, &user); err != nil {
		return nil, errors.Wrap(err, "unmarshal")

	}

	return &user, nil
}

func (s *service) Login(ctx context.Context, input presentations.Login) (*entity.UserDetailToken, error) {
	if err := input.Validate(); err != nil {
		return nil, errors.Wrap(err, "validation(s) error")
	}

	_, err := s.repo.User.FindUserByPhoneEmail(ctx, input.Username)
	if err != nil {
		return nil, consts.ErrUserNotFound
	}

	hashString := common.Hasher(input.Password)
	users, err := s.repo.User.FindUserByPhoneEmailPassword(ctx, input.Username, hashString)
	if err != nil {
		return nil, consts.ErrWrongPassword
	}

	token, err := jwt.CreateToken(users, jwt.RequestJwt{
		ID:        users.ID.String(),
		JWTKey:    s.cfg.JWTKey,
		CreatedAt: users.CreatedAt,
	})
	if err != nil {
		return nil, errors.Wrap(err, "generate token")
	}

	return &entity.UserDetailToken{
		UserDetail: *users,
		Token:      *token,
	}, nil
}
