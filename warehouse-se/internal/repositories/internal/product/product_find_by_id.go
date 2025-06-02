package product

import (
	"context"
	"database/sql"
	"warehouse-se/internal/consts"
	"warehouse-se/internal/entity"

	"github.com/pkg/errors"
)

func (r product) FindProductByID(ctx context.Context, productID string) (*entity.Product, error) {
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
		WHERE id = $1
		  AND deleted_at is null
`

	var product entity.Product

	err := r.db.QueryRow(ctx, &product, query, productID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, consts.ErrProductNotFound
		default:
			return nil, errors.Wrap(err, "failed to fetch row from db")
		}
	}

	return &product, nil
}
