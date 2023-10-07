package order

import (
	"ecom_in_go/models/customer"
	"ecom_in_go/models/orderline"
)

type OrderResponse struct {
	ID          int                           `json:"id"`
	OrderNumber string                        `json:"orderNumber"`
	Customer    customer.CustomerResponse     `gorm:"foreignKey:ID" json:"customer"`
	OrderLines  []orderline.OrderLineResponse `json:"orderLines"`
	UpdatedAt   string                        `json:"updatedAt"`
	CreatedAt   string                        `json:"createdAt"`
}
