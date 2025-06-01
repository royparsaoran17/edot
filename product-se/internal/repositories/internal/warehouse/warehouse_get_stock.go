package warehouse

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"product-se/internal/entity"
)

func (r warehouse) GetStock(ctx context.Context, warehouseID string) ([]entity.ProductStock, error) {
	query := `
	SELECT 
        jsonb_build_object(
            'id', s.id,
            'product_id', s.product_id,
            'product',(
                SELECT
					json_build_object(
						'id', p.id,
						'name', p.name,
						'description', p.description,
						'price', p.price,
						'unit', p.unit,
						'sku', p.sku,
						'category', p.category,
						'is_active', p.is_active,
						'created_at', p.created_at::timestamptz,
						'updated_at', p.updated_at::timestamptz,
						'deleted_at', p.deleted_at::timestamptz
					)
                FROM products p
                    WHERE s.product_id = p.id
                    AND p.deleted_at is null
            ),
            'warehouse_id', s.warehouse_id,
            'warehouse',(
                SELECT
					json_build_object(
						'id', fw.id,
						'name', fw.name,
						'is_active', fw.is_active,
						'created_at', fw.created_at::timestamptz,
						'updated_at', fw.updated_at::timestamptz,
						'deleted_at', fw.deleted_at::timestamptz
					)
                FROM warehouses fw
                    WHERE s.warehouse_id, = fw.id
                    AND fw.deleted_at is null
            ),
            'quantity', s.quantity,
            'created_at', p.created_at::timestamptz,
            'updated_at', p.updated_at::timestamptz,
            'deleted_at', p.deleted_at::timestamptz
        )
    FROM
        product_stocks s
    WHERE s.warehouse_id = $1
        AND s.deleted_at is null;`

	var b []byte
	err := r.db.QueryRow(ctx, query, warehouseID).Scan(&b)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch stock from db")
	}

	var stock []entity.ProductStock
	if err := json.Unmarshal(b, &stock); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal byte to user")
	}

	return stock, nil
}
