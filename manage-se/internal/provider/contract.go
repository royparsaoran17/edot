package provider

import (
	"context"
	"manage-se/internal/common"
	"manage-se/internal/presentations"
	"manage-se/internal/provider/order"
	"manage-se/internal/provider/user"
)

type User interface {
	Login(ctx context.Context, input presentations.Login) (*user.UserDetailToken, error)
	Verify(ctx context.Context, input presentations.Verify) (*user.UserDetail, error)

	CreateUser(ctx context.Context, input presentations.UserCreate) (*user.UserDetail, error)
	GetListUsers(ctx context.Context, meta *common.Metadata) ([]user.User, error)

	GetListRoles(ctx context.Context) ([]user.Role, error)
	CreateRole(ctx context.Context, input presentations.RoleCreate) (*user.Role, error)
}

type Order interface {
	CreateOrderPayment(ctx context.Context, orderID string, input presentations.PaymentCreate) (*order.Order, error)
	CreateOrder(ctx context.Context, input presentations.OrderCreate) (*order.Order, error)
	GetOrderByID(ctx context.Context, orderID string) (*order.OrderDetail, error)
	GetListOrders(ctx context.Context, userID string, meta *common.Metadata) ([]order.Order, error)
}
