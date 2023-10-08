package orderline

import (
	"ecom_in_go/models/product"
)

type OrderLineResponse struct {
	ID        int                              `gorm:"primaryKey" json:"id"`
	OrderID   int                              `json:"-"`
	VariantID int                              `json:"variantId"`
	ProductID int                              `json:"productId"`
	Product   product.OrderLineProductResponse `gorm:"foreignKey:ID" json:"product"`
	Price     int                              `json:"price"`
	Quantity  int                              `json:"quantity"`
	UpdatedAt string                           `json:"updatedAt"`
	CreatedAt string                           `json:"createdAt"`
}

type OrderLineRequest struct {
	OrderID   int                             `json:"-"`
	VariantID int                             `json:"variantId"`
	ProductID int                             `json:"productId"`
	Product   product.OrderLineProductRequest `gorm:"foreignKey:ID" json:"product"`
	Price     int                             `json:"price"`
	Quantity  int                             `json:"quantity"`
}
