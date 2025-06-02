package warehouse

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"warehouse-se/internal/common"
	"warehouse-se/internal/consts"
	"warehouse-se/internal/entity"
	"warehouse-se/internal/presentations"
	"warehouse-se/internal/repositories"
	"warehouse-se/internal/repositories/repooption"
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

func (s *service) MoveWarehouseStock(ctx context.Context, input presentations.WarehouseCreateMoveStock) error {
	_, err := s.repo.Warehouse.FindWarehouseByID(ctx, input.ToWarehouseID)
	if err != nil {
		return errors.Wrapf(err, "getting to warehouse id %s", input.ToWarehouseID)
	}

	fromWarehouse, err := s.repo.Warehouse.FindWarehouseByID(ctx, input.FromWarehouseID)
	if err != nil {
		return errors.Wrapf(err, "getting from warehouse id %s", input.FromWarehouseID)
	}

	product, err := s.repo.Product.FindProductByID(ctx, input.ProductID)
	if err != nil {
		return errors.Wrapf(err, "getting product id %s", input.ProductID)
	}

	productWarehouseStock, err := s.repo.Product.GetStockDetail(ctx, product.ID, fromWarehouse.ID)
	if err != nil {
		return errors.Wrapf(err, "getting warehouse id %s", fromWarehouse.ID)
	}

	if productWarehouseStock.Quantity < input.Quantity {
		return consts.ErrWarehouseStockEmpty
	}

	// start transaction
	tx, err := s.repo.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
	})

	if err != nil {
		return errors.Wrap(err, "begin tx")
	}

	if err := s.repo.Product.UpdateProductStock(ctx, input.FromWarehouseID, presentations.ProductUpdateStock{
		Quantity: productWarehouseStock.Quantity - input.Quantity,
	}, repooption.WithTx(tx)); err != nil {
		errRollBack := tx.Rollback()
		if errRollBack != nil {
			err = errors.Wrap(err, errRollBack.Error())
		}

		return errors.Wrap(err, "updating from warehouse")
	}

	if err := s.repo.Product.UpdateProductStock(ctx, input.ToWarehouseID, presentations.ProductUpdateStock{
		Quantity: productWarehouseStock.Quantity + input.Quantity,
	}, repooption.WithTx(tx)); err != nil {
		errRollBack := tx.Rollback()
		if errRollBack != nil {
			err = errors.Wrap(err, errRollBack.Error())
		}

		return errors.Wrap(err, "updating to warehouse")
	}

	if err := s.repo.Warehouse.CreateMoveStockWarehouse(ctx, input, repooption.WithTx(tx)); err != nil {
		errRollBack := tx.Rollback()
		if errRollBack != nil {
			err = errors.Wrap(err, errRollBack.Error())
		}

		return errors.Wrap(err, "create warehouse move")
	}

	// commit transaction
	errCommit := tx.Commit()
	if errCommit != nil {
		return errors.Wrap(errCommit, "commit transaction")
	}

	return nil
}

func (s *service) ActivateWarehouseByID(ctx context.Context, warehouseID string) error {
	warehouse, err := s.repo.Warehouse.FindWarehouseByID(ctx, warehouseID)
	if err != nil {
		return errors.Wrapf(err, "getting warehouse id %s", warehouseID)
	}

	if err := s.repo.Warehouse.UpdateWarehouse(ctx, warehouseID, presentations.WarehouseUpdate{
		Name:     warehouse.Name,
		IsActive: true,
	}); err != nil {
		return errors.Wrap(err, "updating warehouse")
	}

	return nil
}

func (s *service) InactiveWarehouseByID(ctx context.Context, warehouseID string) error {
	warehouse, err := s.repo.Warehouse.FindWarehouseByID(ctx, warehouseID)
	if err != nil {
		return errors.Wrapf(err, "getting warehouse id %s", warehouseID)
	}

	if err := s.repo.Warehouse.UpdateWarehouse(ctx, warehouseID, presentations.WarehouseUpdate{
		Name:     warehouse.Name,
		IsActive: false,
	}); err != nil {
		return errors.Wrap(err, "updating warehouse")
	}

	return nil
}
