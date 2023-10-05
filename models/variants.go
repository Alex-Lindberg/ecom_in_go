package models

import "time"

type Variant struct {
	ID        int       `json:"id"`
	ProductID int       `json:"productId"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}
