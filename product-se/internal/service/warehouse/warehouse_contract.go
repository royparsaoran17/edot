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
}
