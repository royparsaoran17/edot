package repositories

import (
	"context"
	"shop-se/internal/common"
	"shop-se/internal/entity"
	"shop-se/internal/presentations"
	"shop-se/internal/repositories/repooption"
)

type Shop interface {
	CreateShop(ctx context.Context, input presentations.ShopCreate, opts ...repooption.TxOption) error
	UpdateShop(ctx context.Context, roleID string, input presentations.ShopUpdate, opts ...repooption.TxOption) error
	FindShopByID(ctx context.Context, roleID string) (*entity.Shop, error)
	GetAllShop(ctx context.Context, meta *common.Metadata) ([]entity.Shop, error)
}

type User interface {
	FindUserByID(ctx context.Context, userID string) (*entity.User, error)
}

type Warehouse interface {
	GetAllWarehouse(ctx context.Context, shopID string) ([]entity.Warehouse, error)
}
