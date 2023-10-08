package variant

type VariantResponse struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	ProductID int    `gorm:"foreignKey:ID" json:"-"`
	SKU       string `json:"sku"`
	Name      string `json:"name"`
	Stock     int    `json:"stock"`
	UpdatedAt string `json:"updatedAt"`
	CreatedAt string `json:"createdAt"`
}

type VariantRequest struct {
	Name      string `json:"name"`
	ProductID int    `gorm:"foreignKey:ID" json:"-"`
	SKU       string `json:"sku"`
	Price     int    `json:"price"`
	Stock     int    `json:"stock"`
}
