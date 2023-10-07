package variant

type VariantResponse struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	ProductID int    `gorm:"foreignKey:ID" json:"-"`
	Name      string `json:"name"`
	Stock     int    `json:"stock"`
	UpdatedAt string `json:"updatedAt"`
	CreatedAt string `json:"createdAt"`
}
