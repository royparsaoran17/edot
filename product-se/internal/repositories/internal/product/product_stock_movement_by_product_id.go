package product

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/pkg/errors"
	"product-se/internal/consts"
	"product-se/internal/entity"
)

func (r product) GetStockMovementByProductID(ctx context.Context, productID string) ([]entity.StockMovement, error) {
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
            'from_warehouse_id', s.from_warehouse_id,
            'from_warehouse',(
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
                    WHERE s.from_warehouse_id, = fw.id
                    AND fw.deleted_at is null
            ),
            'to_warehouse_id', s.to_warehouse_id,
            'to_warehouse',(
                SELECT
					json_build_object(
						'id', tw.id,
						'name', tw.name,
						'is_active', tw.is_active,
						'created_at', tw.created_at::timestamptz,
						'updated_at', tw.updated_at::timestamptz,
						'deleted_at', tw.deleted_at::timestamptz
					)
                FROM warehouses tw
                    WHERE s.to_warehouse, = tw.id
                    AND tw.deleted_at is null
            ),
            'notes', s.notes,
            'quantity', s.quantity,
            'created_at', p.created_at::timestamptz,
            'updated_at', p.updated_at::timestamptz,
            'deleted_at', p.deleted_at::timestamptz
        )
    FROM
        stock_movements s
    WHERE s.product_id = $1
        AND s.deleted_at is null;`

	var b []byte
	err := r.db.QueryRow(ctx, &b, query, productID)
	if err != nil {
		sqlErr := r.db.ParseSQLError(err)
		switch sqlErr {
		case sql.ErrNoRows:
			return nil, consts.ErrStockNotFound
		default:
			return nil, errors.Wrap(err, "failed to fetch stock from db")
		}
	}

	var stock []entity.StockMovement
	if err := json.Unmarshal(b, &stock); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal byte to user")
	}

	return stock, nil
}
