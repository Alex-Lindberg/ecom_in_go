package transformers

import "ecom_in_go/models/variant"

func ToVariantResponse(v variant.Variant) variant.VariantResponse {
	return variant.VariantResponse{
		ID:        v.ID,
		Name:      v.Name,
		Stock:     v.Stock,
		UpdatedAt: v.UpdatedAt.String(),
		CreatedAt: v.CreatedAt.String(),
	}
}
