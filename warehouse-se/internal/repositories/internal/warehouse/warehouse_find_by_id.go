package warehouse

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"warehouse-se/internal/consts"
	"warehouse-se/internal/entity"
)

func (r warehouse) FindWarehouseByID(ctx context.Context, warehouseID string) (*entity.Warehouse, error) {
	query := `
SELECT 
    id, 
    name, 
    shop_id,
    is_active, 
    created_at::timestamptz,
    updated_at::timestamptz, 
    deleted_at::timestamptz
FROM warehouses 
WHERE id = $1
  AND deleted_at is null
`

	var warehouse entity.Warehouse

	err := r.db.QueryRow(ctx, &warehouse, query, warehouseID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, consts.ErrWarehouseNotFound
		default:
			return nil, errors.Wrap(err, "failed to fetch row from db")
		}
	}

	return &warehouse, nil
}
