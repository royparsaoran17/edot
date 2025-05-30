package dto

import (
	"product-se/internal/consts"
	"product-se/internal/entity"
	"product-se/internal/presentations"
)

func ProductToResponse(src entity.Product) presentations.ProductDetail {
	x := presentations.ProductDetail{
            ID: src.ID,
            Name: src.Name,
            Description: src.Description,
            Price: src.Price,
            Unit: src.Unit,
            Sku: src.Sku,
            Category: src.Category,
            IsActive: src.IsActive,
            CreatedAt: src.CreatedAt,
            UpdatedAt: src.UpdatedAt,
            DeletedAt: src.DeletedAt,
	}

	if !src.CreatedAt.IsZero() {
		x.CreatedAt = src.CreatedAt.Format(consts.LayoutDateTimeFormat)
	}
	//
	//if !src.UpdatedAt.IsZero() {
	//	x.UpdatedAt = src.UpdatedAt.Format(consts.LayoutDateTimeFormat)
	//}
	//
	//if !src.DeletedAt.IsZero() {
	//	x.DeletedAt = src.DeletedAt.Format(consts.LayoutDateTimeFormat)
	//}

	return x
}

func ProductsToResponse(inputs []entity.Product) []presentations.ProductDetail {
	var (
		result = []presentations.ProductDetail{}
	)

	for i := 0; i < len(inputs); i++ {
		result = append(result, ProductToResponse(inputs[i]))
	}

	return result
}