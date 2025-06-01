package warehouse

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"product-se/internal/common"
	"product-se/internal/consts"
	"product-se/internal/entity"
	"product-se/internal/presentations"
	"product-se/internal/repositories"
)

type service struct {
	repo *repositories.Repository
}

func NewService(repo *repositories.Repository) Warehouse {
	return &service{repo: repo}
}

func (s *service) GetAllWarehouse(ctx context.Context, meta *common.Metadata) ([]entity.Warehouse, error) {
	warehouses, err := s.repo.Warehouse.GetAllWarehouse(ctx, meta)
	if err != nil {
		return nil, errors.Wrap(err, "getting all warehouses on ")
	}

	return warehouses, nil
}

func (s *service) GetWarehouseByID(ctx context.Context, warehouseID string) (*entity.WarehouseDetail, error) {
	warehouses, err := s.repo.Warehouse.FindWarehouseByID(ctx, warehouseID)
	if err != nil {
		return nil, errors.Wrapf(err, "getting warehouse id %s", warehouseID)
	}

	products, err := s.repo.Warehouse.GetStock(ctx, warehouseID)
	if err != nil {
		return nil, errors.Wrapf(err, "getting warehouse id %s", warehouseID)
	}

	return &entity.WarehouseDetail{
		ID:        warehouses.ID,
		Name:      warehouses.Name,
		IsActive:  warehouses.IsActive,
		CreatedAt: warehouses.CreatedAt,
		UpdatedAt: warehouses.UpdatedAt,
		DeletedAt: warehouses.DeletedAt,
		Products:  products,
	}, nil
}

func (s *service) UpdateWarehouseByID(ctx context.Context, warehouseID string, input presentations.WarehouseUpdate) error {
	if err := input.Validate(); err != nil {
		return errors.Wrap(err, "validation(s) error")
	}

	_, err := s.repo.Warehouse.FindWarehouseByID(ctx, warehouseID)
	if err != nil {
		return errors.Wrapf(err, "getting warehouse id %s", warehouseID)
	}

	if err := s.repo.Warehouse.UpdateWarehouse(ctx, warehouseID, input); err != nil {
		return errors.Wrap(err, "updating warehouse")
	}

	return nil
}

func (s *service) CreateWarehouse(ctx context.Context, input presentations.WarehouseCreate) (*entity.Warehouse, error) {
	if err := input.Validate(); err != nil {
		return nil, errors.Wrap(err, "validation(s) error")
	}

	warehouseID := uuid.NewString()
	input.ID = warehouseID
	err := s.repo.Warehouse.CreateWarehouse(ctx, input)
	if err != nil {
		return nil, errors.Wrap(err, "creating warehouse")

	}

	warehouse, err := s.repo.Warehouse.FindWarehouseByID(ctx, warehouseID)
	if err != nil {
		return nil, errors.Wrapf(err, "getting warehouse id %s", warehouseID)
	}

	return warehouse, nil
}

func (s *service) AddWarehouseStock(ctx context.Context, input presentations.WarehouseStock) (*entity.ProductStock, error) {
	warehouse, err := s.repo.Warehouse.FindWarehouseByID(ctx, input.ID)
	if err != nil {
		return nil, errors.Wrapf(err, "getting warehouse id %s", input.ID)
	}

	product, err := s.repo.Product.FindProductByID(ctx, input.ProductID)
	if err != nil {
		return nil, errors.Wrapf(err, "getting warehouse id %s", input.ID)
	}

	productWarehouseStock, err := s.repo.Product.GetStockDetail(ctx, product.ID, warehouse.ID)
	if err != nil {
		return nil, errors.Wrapf(err, "getting warehouse id %s", input.ID)
	}

	if err := s.repo.Product.UpdateProductStock(ctx, input.ID, presentations.ProductUpdateStock{
		Quantity: productWarehouseStock.Quantity + input.Quantity,
	}); err != nil {
		return nil, errors.Wrap(err, "updating warehouse")
	}

	updatedProductWarehouseStock, err := s.repo.Product.GetStockDetail(ctx, product.ID, warehouse.ID)
	if err != nil {
		return nil, errors.Wrapf(err, "getting warehouse id %s", input.ID)
	}

	return updatedProductWarehouseStock, nil
}

func (s *service) DeductWarehouseStock(ctx context.Context, input presentations.WarehouseStock) (*entity.ProductStock, error) {
	warehouse, err := s.repo.Warehouse.FindWarehouseByID(ctx, input.ID)
	if err != nil {
		return nil, errors.Wrapf(err, "getting warehouse id %s", input.ID)
	}

	product, err := s.repo.Product.FindProductByID(ctx, input.ProductID)
	if err != nil {
		return nil, errors.Wrapf(err, "getting warehouse id %s", input.ID)
	}

	productWarehouseStock, err := s.repo.Product.GetStockDetail(ctx, product.ID, warehouse.ID)
	if err != nil {
		return nil, errors.Wrapf(err, "getting warehouse id %s", input.ID)
	}

	if productWarehouseStock.Quantity < input.Quantity {
		return nil, consts.ErrWarehouseStockEmpty
	}

	if err := s.repo.Product.UpdateProductStock(ctx, input.ID, presentations.ProductUpdateStock{
		Quantity: productWarehouseStock.Quantity - input.Quantity,
	}); err != nil {
		return nil, errors.Wrap(err, "updating warehouse")
	}

	updatedProductWarehouseStock, err := s.repo.Product.GetStockDetail(ctx, product.ID, warehouse.ID)
	if err != nil {
		return nil, errors.Wrapf(err, "getting warehouse id %s", input.ID)
	}

	return updatedProductWarehouseStock, nil
}
