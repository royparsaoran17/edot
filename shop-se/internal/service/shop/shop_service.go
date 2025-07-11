package shop

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"shop-se/internal/common"
	"shop-se/internal/entity"
	"shop-se/internal/presentations"
	"shop-se/internal/repositories"
)

type service struct {
	repo *repositories.Repository
}

func NewService(repo *repositories.Repository) Shop {
	return &service{repo: repo}
}

func (s *service) GetAllShop(ctx context.Context, meta *common.Metadata) ([]entity.Shop, error) {
	shops, err := s.repo.Shop.GetAllShop(ctx, meta)
	if err != nil {
		return nil, errors.Wrap(err, "getting all shops on ")
	}

	return shops, nil
}

func (s *service) GetShopByID(ctx context.Context, shopID string) (*entity.ShopDetail, error) {
	shops, err := s.repo.Shop.FindShopByID(ctx, shopID)
	if err != nil {
		return nil, errors.Wrapf(err, "getting shop id %s", shopID)
	}

	owner, err := s.repo.User.FindUserByID(ctx, shops.OwnerID)
	if err != nil {
		return nil, errors.Wrapf(err, "getting shop id %s", shopID)
	}

	warehouses, err := s.repo.Warehouse.GetAllWarehouse(ctx, shops.ID)
	if err != nil {
		return nil, errors.Wrapf(err, "getting shop id %s", shopID)
	}

	return &entity.ShopDetail{
		ID:         shops.ID,
		Name:       shops.Name,
		OwnerID:    shops.OwnerID,
		Owner:      *owner,
		CreatedAt:  shops.CreatedAt,
		UpdatedAt:  shops.UpdatedAt,
		Warehouses: warehouses,
	}, nil
}

func (s *service) UpdateShopByID(ctx context.Context, shopID string, input presentations.ShopUpdate) error {
	if err := input.Validate(); err != nil {
		return errors.Wrap(err, "validation(s) error")
	}

	_, err := s.repo.Shop.FindShopByID(ctx, shopID)
	if err != nil {
		return errors.Wrapf(err, "getting shop id %s", shopID)
	}

	if err := s.repo.Shop.UpdateShop(ctx, shopID, input); err != nil {
		return errors.Wrap(err, "updating shop")
	}

	return nil
}

func (s *service) CreateShop(ctx context.Context, input presentations.ShopCreate) (*entity.Shop, error) {
	if err := input.Validate(); err != nil {
		return nil, errors.Wrap(err, "validation(s) error")
	}

	shopID := uuid.NewString()
	input.ID = shopID
	err := s.repo.Shop.CreateShop(ctx, input)
	if err != nil {
		return nil, errors.Wrap(err, "creating shop")

	}

	shop, err := s.repo.Shop.FindShopByID(ctx, shopID)
	if err != nil {
		return nil, errors.Wrapf(err, "getting shop id %s", shopID)
	}

	return shop, nil
}
