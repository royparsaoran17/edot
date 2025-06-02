package payment

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"order-se/internal/consts"
	"order-se/internal/entity"
)

func (r payment) FindPaymentByID(ctx context.Context, paymentID string) (*entity.Payment, error) {
	query := `
		SELECT 
			id, 
			order_id, 
			method, 
			status, 
			amount, 
			paid_at, 
			created_at::timestamptz,
			updated_at::timestamptz, 
			deleted_at::timestamptz
		FROM payments 
		WHERE id = $1
		  AND deleted_at is null
`

	var payment entity.Payment

	err := r.db.QueryRow(ctx, &payment, query, paymentID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, consts.ErrDataNotFound
		default:
			return nil, errors.Wrap(err, "failed to fetch row from db")
		}
	}

	return &payment, nil
}
