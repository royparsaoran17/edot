package stockreservation

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"order-se/internal/common"
	"order-se/internal/entity"
	"strings"
)

func (r stockReservation) GetAllStockReservation(ctx context.Context, meta *common.Metadata) ([]entity.StockReservation, error) {
	params, err := common.ParamFromMetadata(meta, &r)
	if err != nil {
		return nil, errors.Wrap(err, "parse params from meta")
	}

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
			WHERE 1=1
			AND deleted_at is null
			AND created_at >= GREATEST($3::date, '-infinity'::date)
			AND created_at <= LEAST($4::date, 'infinity'::date)
			ORDER BY created_at DESC
			LIMIT $1 OFFSET $2
`
	query = strings.Replace(
		query,
		"ORDER BY created_at DESC",
		fmt.Sprintf("ORDER BY %s %s", params.OrderBy, params.OrderDirection),
		1,
	)

	if params.SearchBy != "" {
		query = strings.Replace(
			query,
			"1=1",
			fmt.Sprintf("lower(%s) like '%s'", params.SearchBy, params.Search),
			1,
		)
	}

	stockReservations := make([]entity.StockReservation, 0)

	err = r.db.Query(ctx, &stockReservations, query, params.Limit, params.Offset, params.DateFrom, params.DateEnd)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all stockReservations from database")
	}

	return stockReservations, nil
}
