package repositories

import (
	"context"
	"order-se/internal/common"
	"order-se/internal/entity"
	"order-se/internal/presentations"
	"order-se/internal/repositories/repooption"
)

type Order interface {
	CreateOrder(ctx context.Context, input presentations.OrderCreate, opts ...repooption.TxOption) error
	UpdateOrder(ctx context.Context, roleID string, input presentations.OrderUpdate, opts ...repooption.TxOption) error
	FindOrderByID(ctx context.Context, roleID string) (*entity.Order, error)
	GetAllOrder(ctx context.Context, userID string, meta *common.Metadata) ([]entity.Order, error)
}

type Product interface {
	FindProductByID(ctx context.Context, productID string) (*entity.Product, error)
	GetStockDetail(ctx context.Context, productID, warehouseID string) (*entity.ProductStock, error)
}

type User interface {
	FindUserByID(ctx context.Context, userID string) (*entity.User, error)
}

type Warehouse interface {
	GetAllWarehouse(ctx context.Context, orderID string) ([]entity.Warehouse, error)
}

type Payment interface {
	CreatePayment(ctx context.Context, input presentations.PaymentCreate, opts ...repooption.TxOption) error
	FindPaymentByID(ctx context.Context, roleID string) (*entity.Payment, error)
	GetAllPayment(ctx context.Context, meta *common.Metadata) ([]entity.Payment, error)
}

type StockReservation interface {
	CreateStockReservation(ctx context.Context, input presentations.StockReservationCreate, opts ...repooption.TxOption) error
	FindStockReservationByID(ctx context.Context, roleID string) (*entity.StockReservation, error)
	GetAllStockReservation(ctx context.Context, meta *common.Metadata) ([]entity.StockReservation, error)
}

type OrderItem interface {
	CreateOrderItem(ctx context.Context, input presentations.OrderItemCreate, opts ...repooption.TxOption) error
	FindOrderItemByID(ctx context.Context, roleID string) (*entity.OrderItem, error)
	GetAllOrderItemByOrderID(ctx context.Context, orderID string) ([]entity.OrderItem, error)
}
