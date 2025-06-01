package warehouse

import (
	"context"
	"product-se/internal/common"
	"product-se/internal/presentations"

	"product-se/internal/entity"
)

type Warehouse interface {
	GetAllWarehouse(ctx context.Context, meta *common.Metadata) ([]entity.Warehouse, error)
	GetWarehouseByID(ctx context.Context, warehouseID string) (*entity.WarehouseDetail, error)
	UpdateWarehouseByID(ctx context.Context, warehouseID string, input presentations.WarehouseUpdate) error
	CreateWarehouse(ctx context.Context, input presentations.WarehouseCreate) (*entity.Warehouse, error)
	AddWarehouseStock(ctx context.Context, input presentations.WarehouseStock) (*entity.ProductStock, error)
	DeductWarehouseStock(ctx context.Context, input presentations.WarehouseStock) (*entity.ProductStock, error)
	MoveWarehouseStock(ctx context.Context, input presentations.WarehouseCreateMoveStock) error
}
