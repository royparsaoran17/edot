package repositories

import (
	"context"
	"product-se/internal/common"
	"product-se/internal/entity"
	"product-se/internal/presentations"
	"product-se/internal/repositories/repooption"
)

type Warehouse interface {
	CreateWarehouse(ctx context.Context, input presentations.WarehouseCreate, opts ...repooption.TxOption) error
	UpdateWarehouse(ctx context.Context, roleID string, input presentations.WarehouseUpdate, opts ...repooption.TxOption) error
	FindWarehouseByID(ctx context.Context, roleID string) (*entity.Warehouse, error)
	GetAllWarehouse(ctx context.Context, meta *common.Metadata) ([]entity.Warehouse, error)
	GetStock(ctx context.Context, productID string) ([]entity.ProductStock, error)
	CreateMoveStockWarehouse(ctx context.Context, input presentations.WarehouseCreateMoveStock, opts ...repooption.TxOption) error
}

type Product interface {
	CreateProduct(ctx context.Context, input presentations.ProductCreate, opts ...repooption.TxOption) error
	UpdateProduct(ctx context.Context, productID string, input presentations.ProductUpdate, opts ...repooption.TxOption) error
	GetAllProduct(ctx context.Context, meta *common.Metadata) ([]entity.Product, error)
	FindProductByID(ctx context.Context, productID string) (*entity.Product, error)
	GetStockMovement(ctx context.Context) ([]entity.StockMovement, error)
	GetStockMovementByProductID(ctx context.Context, productID string) ([]entity.StockMovement, error)
	GetStock(ctx context.Context, productID string) ([]entity.ProductStock, error)
	GetStockDetail(ctx context.Context, productID, warehouseID string) (*entity.ProductStock, error)
	CreateProductStock(ctx context.Context, input presentations.ProductCreateStock, opts ...repooption.TxOption) error
	UpdateProductStock(ctx context.Context, stockID string, input presentations.ProductUpdateStock, opts ...repooption.TxOption) error
}
