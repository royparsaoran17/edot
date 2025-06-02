package stockreservation

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"order-se/internal/consts"
	"order-se/internal/entity"
)

func (r stockReservation) FindStockReservationByID(ctx context.Context, stockReservationID string) (*entity.StockReservation, error) {
	query := `
		SELECT 
			id, 
			order_id, 
			product_id, 
			warehouse_id, 
			quantity, 
			price, 
			reserved_at, 
			expires_at, 
			created_at::timestamptz,
			updated_at::timestamptz, 
			deleted_at::timestamptz
		FROM stock_reservations 
		WHERE id = $1
		  AND deleted_at is null
`

	var stockReservation entity.StockReservation

	err := r.db.QueryRow(ctx, &stockReservation, query, stockReservationID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, consts.ErrDataNotFound
		default:
			return nil, errors.Wrap(err, "failed to fetch row from db")
		}
	}

	return &stockReservation, nil
}
