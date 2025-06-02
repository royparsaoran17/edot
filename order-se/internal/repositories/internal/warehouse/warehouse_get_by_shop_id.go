package warehouse

import (
	"context"
	"github.com/pkg/errors"
	"order-se/internal/entity"
)

func (r warehouse) GetAllWarehouse(ctx context.Context, orderID string) ([]entity.Warehouse, error) {

	query := `
		SELECT 
			id, 
			name, 
    		order_id,
    		is_active, 
			created_at::timestamptz,
			updated_at::timestamptz, 
			deleted_at::timestamptz
		FROM warehouses 
			WHERE order_id = $1
			AND deleted_at is null
`

	warehouses := make([]entity.Warehouse, 0)

	err := r.db.Query(ctx, &warehouses, query, orderID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all warehouses from database")
	}

	return warehouses, nil
}
