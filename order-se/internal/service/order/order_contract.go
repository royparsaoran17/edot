package order

import (
	"context"
	"order-se/internal/common"
	"order-se/internal/presentations"

	"order-se/internal/entity"
)

type Order interface {
	GetAllOrder(ctx context.Context, userID string, meta *common.Metadata) ([]entity.Order, error)
	GetOrderByID(ctx context.Context, orderID string) (*entity.OrderDetail, error)
	CreateOrder(ctx context.Context, input presentations.Order) (*entity.Order, error)
	CreateOrderPayment(ctx context.Context, input presentations.OrderPayment) (*entity.Order, error)
}
