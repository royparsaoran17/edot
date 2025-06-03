package shop

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"shop-se/internal/common"
	"shop-se/internal/entity"
	"shop-se/pkg/logger"
	"strings"
)

func (r shop) GetAllShop(ctx context.Context, meta *common.Metadata) ([]entity.Shop, error) {
	params, err := common.ParamFromMetadata(meta, &r)
	if err != nil {
		return nil, errors.Wrap(err, "parse params from meta")
	}

	query := `
		SELECT 
			id, 
			name, 
    		owner_id, 
			created_at::timestamptz,
			updated_at::timestamptz, 
			deleted_at::timestamptz
		FROM shops 
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

	shops := make([]entity.Shop, 0)

	err = r.db.Query(ctx, &shops, query, params.Limit, params.Offset, params.DateFrom, params.DateEnd)
	if err != nil {
		logger.Info(err.Error())
		return nil, errors.Wrap(err, "failed to get all shops from database")
	}

	return shops, nil
}
