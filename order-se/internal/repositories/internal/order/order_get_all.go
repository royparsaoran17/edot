package order

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"order-se/internal/common"
	"order-se/internal/entity"
	"strings"
)

func (r order) GetAllOrder(ctx context.Context, userID string, meta *common.Metadata) ([]entity.Order, error) {
	params, err := common.ParamFromMetadata(meta, &r)
	if err != nil {
		return nil, errors.Wrap(err, "parse params from meta")
	}

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
			WHERE 1=1
			AND user_id = $1
			AND deleted_at is null
			AND created_at >= GREATEST($3::date, '-infinity'::date)
			AND created_at <= LEAST($4::date, 'infinity'::date)
			ORDER BY created_at DESC
			LIMIT $2 OFFSET $3
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

	orders := make([]entity.Order, 0)

	err = r.db.Query(ctx, &orders, query, userID, params.Limit, params.Offset, params.DateFrom, params.DateEnd)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all orders from database")
	}

	return orders, nil
}
