package product

import (
	"ecom_in_go/models/variant"
	"time"
)

type Product struct {
	ID        int               `gorm:"primaryKey" json:"id"`
	Name      string            `json:"name"`
	Variants  []variant.Variant `json:"variant"`
	UpdatedAt time.Time         `json:"updatedAt"`
	CreatedAt time.Time         `json:"createdAt"`
}
