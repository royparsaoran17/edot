package orderitem

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"order-se/internal/consts"
	"order-se/internal/entity"
)

func (r orderItem) FindOrderItemByID(ctx context.Context, orderItemID string) (*entity.OrderItem, error) {
	query := `
		SELECT 
			id, 
			order_id, 
			product_id, 
			warehouse_id, 
			quantity, 
			price, 
			created_at::timestamptz,
			updated_at::timestamptz, 
			deleted_at::timestamptz
		FROM order_items 
		WHERE id = $1
		  AND deleted_at is null
`

	var orderItem entity.OrderItem

	err := r.db.QueryRow(ctx, &orderItem, query, orderItemID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, consts.ErrDataNotFound
		default:
			return nil, errors.Wrap(err, "failed to fetch row from db")
		}
	}

	return &orderItem, nil
}
