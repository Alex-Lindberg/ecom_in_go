package variant

import "time"

type Variant struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	ProductID int       `json:"-"`
	Name      string    `json:"name"`
	Stock     int       `json:"stock"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}
