package product

import (
	"ecom_in_go/models/variant"
)

type ProductResponse struct {
	ID        int                       `json:"id"`
	Name      string                    `json:"name"`
	Variant   []variant.VariantResponse `gorm:"foreignKey:ID" json:"variant"`
	UpdatedAt string                    `json:"updatedAt"`
	CreatedAt string                    `json:"createdAt"`
}

type OrderLineProductResponse struct {
	ID        int                     `json:"id"`
	Name      string                  `json:"name"`
	VariantID int                     `json:"-"`
	Variant   variant.VariantResponse `gorm:"foreignKey:ID" json:"variant"`
	UpdatedAt string                  `json:"updatedAt"`
	CreatedAt string                  `json:"createdAt"`
}

type OrderLineProductRequest struct {
	Name    string                 `json:"name"`
	Variant variant.VariantRequest `gorm:"foreignKey:ID" json:"variant"`
}
