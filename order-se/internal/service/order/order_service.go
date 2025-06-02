package order

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"order-se/internal/common"
	"order-se/internal/consts"
	"order-se/internal/entity"
	"order-se/internal/presentations"
	"order-se/internal/repositories"
	"order-se/internal/repositories/repooption"
	"time"
)

type service struct {
	repo *repositories.Repository
}

func NewService(repo *repositories.Repository) Order {
	return &service{repo: repo}
}

func (s *service) GetAllOrder(ctx context.Context, userID string, meta *common.Metadata) ([]entity.Order, error) {
	orders, err := s.repo.Order.GetAllOrder(ctx, userID, meta)
	if err != nil {
		return nil, errors.Wrap(err, "getting all orders on ")
	}

	return orders, nil
}

func (s *service) GetOrderByID(ctx context.Context, orderID string) (*entity.OrderDetail, error) {
	orders, err := s.repo.Order.FindOrderByID(ctx, orderID)
	if err != nil {
		return nil, errors.Wrapf(err, "getting order id %s", orderID)
	}

	items, err := s.repo.OrderItem.GetAllOrderItemByOrderID(ctx, orderID)
	if err != nil {
		return nil, errors.Wrapf(err, "getting order id %s", orderID)
	}

	return &entity.OrderDetail{
		Order: *orders,
		Items: items,
	}, nil
}

func (s *service) CreateOrder(ctx context.Context, input presentations.Order) (*entity.Order, error) {
	if err := input.Validate(); err != nil {
		return nil, errors.Wrap(err, "validation error")
	}

	orderID := uuid.NewString()
	var totalPrice float64
	itemDetails := make([]presentations.OrderItemCreate, 0, len(input.Items))

	// Pre-check stock & calculate total price
	for _, item := range input.Items {
		product, err := s.repo.Product.FindProductByID(ctx, item.ProductID)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to get product ID %s", item.ProductID)
		}

		stock, err := s.repo.Product.GetStockDetail(ctx, item.ProductID, item.WarehouseID)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to get stock for product %s in warehouse %s", item.ProductID, item.WarehouseID)
		}

		if stock.Quantity < item.Quantity {
			return nil, consts.ErrProductStockEmpty
		}

		totalPrice += product.Price * float64(item.Quantity)

		itemDetails = append(itemDetails, presentations.OrderItemCreate{
			ID:          uuid.NewString(),
			OrderID:     orderID,
			ProductID:   item.ProductID,
			WarehouseID: item.WarehouseID,
			Quantity:    item.Quantity,
			Price:       product.Price,
		})
	}

	// Start transaction
	tx, err := s.repo.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	if err != nil {
		return nil, errors.Wrap(err, "failed to begin transaction")
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}
	}()

	// Create Order
	orderData := presentations.OrderCreate{
		ID:         orderID,
		UserID:     input.UserID,
		Status:     "PENDING",
		TotalPrice: totalPrice,
	}

	if err := s.repo.Order.CreateOrder(ctx, orderData, repooption.WithTx(tx)); err != nil {
		_ = tx.Rollback()
		return nil, errors.Wrap(err, "failed to create order")
	}

	// Create Order Items
	for _, item := range itemDetails {
		if err := s.repo.OrderItem.CreateOrderItem(ctx, item, repooption.WithTx(tx)); err != nil {
			_ = tx.Rollback()
			return nil, errors.Wrapf(err, "failed to create order item for product %s", item.ProductID)
		}

		if err := s.repo.StockReservation.CreateStockReservation(ctx, presentations.StockReservationCreate{
			ID:          uuid.NewString(),
			OrderID:     item.OrderID,
			ProductID:   item.ProductID,
			WarehouseID: item.WarehouseID,
			Quantity:    item.Quantity,
			Price:       item.Price,
			ReservedAt:  time.Now(),
			ExpiresAt:   time.Now().Add(time.Hour * 24),
		}, repooption.WithTx(tx)); err != nil {
			_ = tx.Rollback()
			return nil, errors.Wrapf(err, "failed to create order item for product %s", item.ProductID)
		}
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return nil, errors.Wrap(err, "failed to commit transaction")
	}

	// Fetch and return final order
	order, err := s.repo.Order.FindOrderByID(ctx, orderID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get order by id %s", orderID)
	}

	return order, nil
}

func (s *service) CreateOrderPayment(ctx context.Context, input presentations.OrderPayment) (*entity.Order, error) {
	if err := input.Validate(); err != nil {
		return nil, errors.Wrap(err, "validation(s) error")
	}

	order, err := s.repo.Order.FindOrderByID(ctx, input.OrderID)
	if err != nil {
		return nil, errors.Wrapf(err, "getting order id %s", input.OrderID)
	}

	// Start transaction
	tx, err := s.repo.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	if err != nil {
		return nil, errors.Wrap(err, "failed to begin transaction")
	}

	err = s.repo.Payment.CreatePayment(ctx, presentations.PaymentCreate{
		ID:      uuid.NewString(),
		OrderID: input.OrderID,
		Method:  input.PaymentMethod,
		Status:  "PAID",
		Amount:  input.Amount,
		PaidAt:  time.Time{},
	}, repooption.WithTx(tx))
	if err != nil {
		_ = tx.Rollback()
		return nil, errors.Wrap(err, "creating order")

	}

	err = s.repo.Order.UpdateOrder(ctx, input.OrderID, presentations.OrderUpdate{Status: "PAID"}, repooption.WithTx(tx))
	if err != nil {
		_ = tx.Rollback()
		return nil, errors.Wrapf(err, "getting order id %s", input.OrderID)
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()
		return nil, errors.Wrap(err, "failed to commit transaction")
	}

	return order, nil
}
