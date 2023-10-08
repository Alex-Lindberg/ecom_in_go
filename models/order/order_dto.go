package order

import (
	"ecom_in_go/models/customer"
	"ecom_in_go/models/orderline"
)

type OrderResponse struct {
	ID             int                           `json:"id"`
	OrderNumber    string                        `json:"orderNumber"`
	OrderReference string                        `json:"orderReference"`
	Customer       customer.CustomerResponse     `gorm:"foreignKey:ID" json:"customer"`
	OrderLines     []orderline.OrderLineResponse `json:"orderLines"`
	UpdatedAt      string                        `json:"updatedAt"`
	CreatedAt      string                        `json:"createdAt"`
}

type CreateOrderRequest struct {
	OrderNumber    string                       `json:"orderNumber"`
	OrderReference string                       `json:"orderReference"`
	Customer       customer.CustomerRequest     `gorm:"foreignKey:ID" json:"customer"`
	OrderLines     []orderline.OrderLineRequest `json:"orderLines"`
}
