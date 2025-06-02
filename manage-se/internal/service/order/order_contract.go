package order

import (
	"context"
	"manage-se/internal/common"
	"manage-se/internal/presentations"
	"manage-se/internal/provider/order"
)

type Order interface {
	GetAllOrder(ctx context.Context, userID string, meta *common.Metadata) ([]order.Order, error)
	GetOrderByID(ctx context.Context, orderID string) (*order.OrderDetail, error)
	CreateOrder(ctx context.Context, input presentations.OrderCreate) (*order.Order, error)
	CreateOrderPayment(ctx context.Context, orderID string, input presentations.PaymentCreate) (*order.Order, error)
}
