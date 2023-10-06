package models

import "time"

type Order struct {
	ID             int       `json:"id"`
	OrderID        string    `json:"orderID"`
	OrderReference string    `json:"orderReference"`
	CustomerID     int       `json:"customerID"`
	UpdatedAt      time.Time `json:"updatedAt"`
	CreatedAt      time.Time `json:"createdAt"`
}
