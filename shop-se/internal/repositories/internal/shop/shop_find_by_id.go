package shop

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"shop-se/internal/consts"
	"shop-se/internal/entity"
)

func (r shop) FindShopByID(ctx context.Context, shopID string) (*entity.Shop, error) {
	query := `
SELECT 
    id, 
    name, 
    owner_id, 
    created_at::timestamptz,
    updated_at::timestamptz, 
    deleted_at::timestamptz
FROM shops 
WHERE id = $1
  AND deleted_at is null
`

	var shop entity.Shop

	err := r.db.QueryRow(ctx, &shop, query, shopID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, consts.ErrShopNotFound
		default:
			return nil, errors.Wrap(err, "failed to fetch row from db")
		}
	}

	return &shop, nil
}
