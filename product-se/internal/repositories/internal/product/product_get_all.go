package product

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"product-se/internal/common"
	"product-se/internal/entity"
	"strings"
)

func (r product) GetAllProduct(ctx context.Context, meta *common.Metadata) ([]entity.Product, error) {
	params, err := common.ParamFromMetadata(meta, &r)
	if err != nil {
		return nil, errors.Wrap(err, "parse params from meta")
	}

	query := `
		SELECT 
			id,
			name,
			description,
			price,
			unit,
			sku,
			category,
			is_active,
			created_at::timestamptz,
			updated_at::timestamptz, 
			deleted_at::timestamptz
		FROM products 
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

	products := make([]entity.Warehouse, 0)

	err = r.db.Fetch(ctx, &warehouses, query, params.Limit, params.Offset, params.DateFrom, params.DateEnd)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all warehouses from database")
	}

	return products, nil
}
