package models

import "time"

type OrderLine struct {
	ID        int       `json:"id"`
	OrderID   string    `json:"orderID"`
	ProductID int       `json:"productID"`
	VariantID int       `json:"variantID"`
	Price     int       `json:"price"`
	Quantity  int       `json:"quantity"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}
