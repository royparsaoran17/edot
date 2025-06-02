package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"manage-se/internal/common"
	"manage-se/internal/consts"
	"manage-se/internal/presentations"
	"manage-se/internal/provider"
	"manage-se/internal/provider/user"
	"time"
)

type service struct {
	provider *provider.Provider
	rdb      redis.Cmdable
}

func NewService(provider *provider.Provider, rdb redis.Cmdable) Auth {
	return &service{provider: provider, rdb: rdb}
}

func (s *service) Login(ctx context.Context, input presentations.Login) (*user.UserDetailToken, error) {
	if err := input.Validate(); err != nil {
		return nil, errors.Wrap(err, "validation(s) error")
	}

	auth, err := s.provider.User.Login(ctx, input)
	if err != nil {
		return nil, errors.Wrap(err, "provider error")
	}

	return auth, nil
}

func (s *service) Register(ctx context.Context, input presentations.Register) (*user.UserDetail, error) {
	if err := input.Validate(); err != nil {
		return nil, errors.Wrap(err, "validation(s) error")
	}

	password := common.RandomString(10)
	user, err := s.provider.User.CreateUser(ctx, presentations.UserCreate{
		ID:       uuid.NewString(),
		Name:     input.Name,
		Phone:    input.Phone,
		Password: password,
		RoleID:   input.RoleID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "provider error")
	}

	user.Password = password

	return user, nil
}

func (s *service) VerifyToken(ctx context.Context, input presentations.Verify) (*user.UserDetail, error) {

	var (
		keyRedis = fmt.Sprintf(consts.FormatStringAuthCache, input.Token)
	)

	userCacheBytes, err := s.rdb.Get(ctx, keyRedis).Bytes()
	if err != nil {
		switch err {
		case redis.Nil:

			verify, err := s.provider.User.Verify(ctx, input)
			if err != nil {
				return nil, errors.Wrap(err, "provider error")
			}

			userCacheBytes, err = json.Marshal(verify)
			if err != nil {
				return nil, errors.Wrap(err, "marshal user to bytes")
			}

			err = s.rdb.Set(ctx, keyRedis, userCacheBytes, time.Hour*1).Err()
			if err != nil {
				return nil, errors.Wrap(err, "set user user cache on redis")
			}

			return verify, nil

		default:
			return nil, errors.Wrap(err, "redis get user user cache")
		}
	}

	var user user.UserDetail
	err = json.Unmarshal(userCacheBytes, &user)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal user cache bytes to struct")
	}

	return &user, nil

}
