package orderline

import (
	"ecom_in_go/models/product"
	"time"
)

type OrderLine struct {
	ID        int             `gorm:"primaryKey" json:"id"`
	OrderID   int             `json:"-"`
	ProductID int             `json:"productId"`
	VariantID int             `json:"variantId"`
	Product   product.Product `gorm:"foreignKey:ProductID" json:"product"`
	Price     int             `json:"price"`
	Quantity  int             `json:"quantity"`
	UpdatedAt time.Time       `json:"updatedAt"`
	CreatedAt time.Time       `json:"createdAt"`
}
