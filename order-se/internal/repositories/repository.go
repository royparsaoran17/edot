package repositories

import (
	"context"
	"database/sql"
	"order-se/internal/repositories/internal/order"
	"order-se/internal/repositories/internal/orderitem"
	"order-se/internal/repositories/internal/payment"
	"order-se/internal/repositories/internal/product"
	"order-se/internal/repositories/internal/stockreservation"
	"order-se/internal/repositories/internal/user"
	"order-se/internal/repositories/internal/warehouse"
	"order-se/pkg/databasex"
)

type Repository struct {
	Order            Order
	OrderItem        OrderItem
	Payment          Payment
	StockReservation StockReservation
	User             User
	Warehouse        Warehouse
	Product          Product
	db               databasex.Adapter
}

func NewRepository(db databasex.Adapter) *Repository {
	return &Repository{
		Order:            order.NewOrder(db),
		Payment:          payment.NewPayment(db),
		OrderItem:        orderitem.NewOrderItem(db),
		StockReservation: stockreservation.NewStockReservation(db),
		User:             user.NewUser(db),
		Product:          product.NewProduct(db),
		Warehouse:        warehouse.NewWarehouse(db),
	}
}

func (r Repository) BeginTx(ctx context.Context, options *sql.TxOptions) (*sql.Tx, error) {
	return r.db.BeginTx(ctx, options)
}
