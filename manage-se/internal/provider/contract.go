package provider

import (
	"context"
	"manage-se/internal/common"
	"manage-se/internal/presentations"
	"manage-se/internal/provider/product"
	"manage-se/internal/provider/user"
)

type User interface {
	Login(ctx context.Context, input presentations.Login) (*user.UserDetailToken, error)
	Verify(ctx context.Context, input presentations.Verify) (*user.UserDetail, error)

	CreateUser(ctx context.Context, input presentations.UserCreate) (*user.UserDetail, error)
	GetListUsers(ctx context.Context, meta *common.Metadata) ([]user.User, error)

	GetListRoles(ctx context.Context) ([]user.Role, error)
	CreateRole(ctx context.Context, input presentations.RoleCreate) (*user.Role, error)
}

type Product interface {
	CreateProduct(ctx context.Context, input presentations.ProductCreate) (*product.Product, error)
	GetListProducts(ctx context.Context, meta *common.Metadata) ([]product.Product, error)

	CreateWarehouse(ctx context.Context, input presentations.WarehouseCreate) (*product.Warehouse, error)
	GetListWarehouses(ctx context.Context, meta *common.Metadata) ([]product.Warehouse, error)
}
