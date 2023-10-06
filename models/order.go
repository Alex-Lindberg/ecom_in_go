package models

import "time"

// CREATE TABLE orders (
//     id SERIAL PRIMARY KEY,
//     order_reference VARCHAR(255) NOT NULL,
//     customer_id INTEGER REFERENCES customers(id),
//     updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
//     created_at TIMESTAMP NOT NULL DEFAULT NOW()
// );

type Order struct {
	ID             int       `json:"id"`
	OrderID        string    `json:"orderID"`
	OrderReference string    `json:"orderReference"`
	CustomerID     int       `json:"customerID"`
	UpdatedAt      time.Time `json:"updatedAt"`
	CreatedAt      time.Time `json:"createdAt"`
}
