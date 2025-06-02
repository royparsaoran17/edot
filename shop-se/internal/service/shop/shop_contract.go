package shop

import (
	"context"
	"shop-se/internal/common"
	"shop-se/internal/presentations"

	"shop-se/internal/entity"
)

type Shop interface {
	GetAllShop(ctx context.Context, meta *common.Metadata) ([]entity.Shop, error)
	GetShopByID(ctx context.Context, shopID string) (*entity.Shop, error)
	UpdateShopByID(ctx context.Context, shopID string, input presentations.ShopUpdate) error
	CreateShop(ctx context.Context, input presentations.ShopCreate) (*entity.Shop, error)
}
