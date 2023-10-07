package order

import (
	"ecom_in_go/models/customer"
	"ecom_in_go/models/orderline"
	"time"
)

type Order struct {
	ID             int                   `gorm:"primaryKey" json:"id"`
	OrderNumber    string                `json:"orderNumber"`
	OrderReference string                `json:"orderReference"`
	CustomerID     int                   `json:"-"`
	Customer       customer.Customer     `gorm:"foreignKey:CustomerID" json:"customer"`
	OrderLines     []orderline.OrderLine `gorm:"foreignKey:OrderID" json:"orderLines"`
	UpdatedAt      time.Time             `json:"updatedAt"`
	CreatedAt      time.Time             `json:"createdAt"`
}
