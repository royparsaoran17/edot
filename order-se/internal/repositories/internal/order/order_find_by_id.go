package order

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"order-se/internal/consts"
	"order-se/internal/entity"
)

func (r order) FindOrderByID(ctx context.Context, orderID string) (*entity.Order, error) {
	query := `
		SELECT 
			id, 
			user_id, 
			status, 
			total_price, 
			created_at::timestamptz,
			updated_at::timestamptz, 
			deleted_at::timestamptz
		FROM orders 
		WHERE id = $1
		  AND deleted_at is null
`

	var order entity.Order

	err := r.db.QueryRow(ctx, &order, query, orderID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, consts.ErrOrderNotFound
		default:
			return nil, errors.Wrap(err, "failed to fetch row from db")
		}
	}

	return &order, nil
}
