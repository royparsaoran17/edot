package product

import (
	"context"
	"product-se/internal/common"
	"product-se/internal/presentations"

	"product-se/internal/entity"
)

type Product interface {
	GetAllProduct(ctx context.Context, meta *common.Metadata) ([]entity.Product, error)
	GetProductByID(ctx context.Context, productID string) (*entity.Product, error)
	UpdateProductByID(ctx context.Context, productID string, input presentations.ProductUpdate) error
	CreateProduct(ctx context.Context, input presentations.ProductCreate) (*entity.Product, error)
}
