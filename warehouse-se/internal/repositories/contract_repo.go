package repositories

import (
	"context"
	"warehouse-se/internal/common"
	"warehouse-se/internal/entity"
	"warehouse-se/internal/presentations"
	"warehouse-se/internal/repositories/repooption"
)

type Warehouse interface {
	CreateWarehouse(ctx context.Context, input presentations.WarehouseCreate, opts ...repooption.TxOption) error
	UpdateWarehouse(ctx context.Context, warehouseID string, input presentations.WarehouseUpdate, opts ...repooption.TxOption) error
	FindWarehouseByID(ctx context.Context, warehouseID string) (*entity.Warehouse, error)
	GetAllWarehouse(ctx context.Context, meta *common.Metadata) ([]entity.Warehouse, error)
	GetStock(ctx context.Context, productID string) ([]entity.ProductStock, error)
	CreateMoveStockWarehouse(ctx context.Context, input presentations.WarehouseCreateMoveStock, opts ...repooption.TxOption) error
}

type Product interface {
	FindProductByID(ctx context.Context, productID string) (*entity.Product, error)
	GetStockDetail(ctx context.Context, productID, warehouseID string) (*entity.ProductStock, error)
	UpdateProductStock(ctx context.Context, stockID string, input presentations.ProductUpdateStock, opts ...repooption.TxOption) error
}
