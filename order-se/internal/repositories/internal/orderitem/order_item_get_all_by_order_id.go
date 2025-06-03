package orderitem

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/pkg/errors"
	"order-se/internal/consts"
	"order-se/internal/entity"
)

func (r orderItem) GetAllOrderItemByOrderID(ctx context.Context, orderID string) ([]entity.OrderItem, error) {
	query := `
	SELECT jsonb_agg(
		jsonb_build_object(
			'id', c.id,
			'order_id', c.order_id,
			'product_id', c.product_id,
			'product', (
				SELECT json_build_object(
					'id', p.id,
					'name', p.name,
					'description', p.description,
					'price', p.price,
					'unit', p.unit,
					'sku', p.sku,
					'category', p.category,
					'created_at', p.created_at::timestamptz,
					'updated_at', p.updated_at::timestamptz,
					'deleted_at', p.deleted_at::timestamptz
				)
				FROM products p
				WHERE p.id = c.product_id
			),
			'warehouse_id', c.warehouse_id,
			'warehouse', (
				SELECT json_build_object(
					'id', w.id,
					'name', w.name,
					'shop_id', w.shop_id,
					'is_active', w.is_active,
					'created_at', w.created_at::timestamptz,
					'updated_at', w.updated_at::timestamptz,
					'deleted_at', w.deleted_at::timestamptz
				)
				FROM warehouses w
				WHERE w.id = c.warehouse_id
			),
			'quantity', c.quantity,
			'price', c.price,
			'created_at', c.created_at::timestamptz,
			'updated_at', c.updated_at::timestamptz,
			'deleted_at', c.deleted_at::timestamptz
		)
	) AS result
	FROM order_items c
	WHERE c.order_id = $1
		AND c.deleted_at IS NULL;
	`
	
	var b []byte
	err := r.db.QueryRow(ctx, &b, query, orderID)
	if err != nil {
		sqlErr := r.db.ParseSQLError(err)
		switch sqlErr {
		case sql.ErrNoRows:
			return nil, consts.ErrUserNotFound
		default:
			return nil, errors.Wrap(err, "failed to fetch order item from db")
		}
	}

	var orders []entity.OrderItem
	if err := json.Unmarshal(b, &orders); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal byte to order items")
	}

	return orders, nil
}
