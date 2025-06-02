package warehouse

import (
	"context"
	"github.com/pkg/errors"
	"shop-se/internal/entity"
)

func (r warehouse) GetAllWarehouse(ctx context.Context, shopID string) ([]entity.Warehouse, error) {

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
			WHERE shop_id = $1
			AND deleted_at is null
`

	warehouses := make([]entity.Warehouse, 0)

	err := r.db.Query(ctx, &warehouses, query, shopID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all warehouses from database")
	}

	return warehouses, nil
}
