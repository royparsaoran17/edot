package user

import (
	"auth-se/internal/common"
	"auth-se/internal/entity"
	"auth-se/pkg/logger"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

func (c user) GetAllUser(ctx context.Context, meta *common.Metadata) ([]entity.User, error) {
	params, err := common.ParamFromMetadata(meta, &c)
	if err != nil {
		return nil, errors.Wrap(err, "parse params from meta")
	}

	query := `
		SELECT 
			id, 
			name, 
			email, 
			phone, 
			role_id,
			created_at::timestamptz,
			updated_at::timestamptz, 
			deleted_at::timestamptz
		FROM users 
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

	roles := make([]entity.User, 0)

	err = c.db.Fetch(ctx, &roles, query, params.Limit, params.Offset, params.DateFrom, params.DateEnd)
	if err != nil {
		logger.Info(err.Error())
		return nil, errors.Wrap(err, "failed to get all roles from database")
	}

	query = `select count(id)
		from users
		where  1=1
  		and created_at >= GREATEST($1::date, '-infinity'::date)
  		and created_at <= LEAST($2::date, 'infinity'::date)`

	if params.SearchBy != "" {
		query = strings.Replace(
			query,
			"1=1",
			fmt.Sprintf("lower(%s) like '%s'", params.SearchBy, params.Search),
			1,
		)
	}

	var count int
	err = c.db.FetchRow(ctx, &count, query, params.DateFrom, params.DateEnd)
	if err != nil {
		logger.Info(err.Error())
		return nil, errors.Wrap(err, "fetch count")
	}

	meta.Total = count
	return roles, nil
}
