package order

import (
	"context"
	"github.com/pkg/errors"
	"manage-se/internal/common"
	"manage-se/internal/presentations"
	"manage-se/internal/provider"
	"manage-se/internal/provider/order"
)

type service struct {
	provider *provider.Provider
}

func NewService(provider *provider.Provider) Order {
	return &service{provider: provider}
}

func (s *service) GetAllOrder(ctx context.Context, userID string, meta *common.Metadata) ([]order.Order, error) {
	orders, err := s.provider.Order.GetListOrders(ctx, userID, meta)
	if err != nil {
		return nil, errors.Wrap(err, "getting all orders ")
	}

	return orders, nil
}

func (s *service) GetOrderByID(ctx context.Context, orderID string) (*order.OrderDetail, error) {
	orderDetail, err := s.provider.Order.GetOrderByID(ctx, orderID)
	if err != nil {
		return nil, errors.Wrap(err, "getting  order details")
	}

	return orderDetail, nil
}

func (s *service) CreateOrder(ctx context.Context, input presentations.OrderCreate) (*order.Order, error) {
	if err := input.Validate(); err != nil {
		return nil, errors.Wrap(err, "validation error")
	}

	orderDetail, err := s.provider.Order.CreateOrder(ctx, input)
	if err != nil {
		return nil, errors.Wrap(err, "getting  order details")
	}

	return orderDetail, nil
}

func (s *service) CreateOrderPayment(ctx context.Context, orderID string, input presentations.PaymentCreate) (*order.Order, error) {
	if err := input.Validate(); err != nil {
		return nil, errors.Wrap(err, "validation(s) error")
	}

	orderDetail, err := s.provider.Order.CreateOrderPayment(ctx, orderID, input)
	if err != nil {
		return nil, errors.Wrap(err, "getting  order details")
	}

	return orderDetail, nil
}
