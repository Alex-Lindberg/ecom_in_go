package storage

import "ecom_in_go/models"

type ProductStore interface {
	GetProducts() ([]models.Product, error)
}

type VariantStore interface {
	GetVariantsByProductID(productID int) ([]models.Variant, error)
	// GetVariantsByProductIDs(productIDs []int) (map[int][]models.Variant, error)
}

type CustomerStore interface {
	GetCustomers() ([]models.Customer, error)
}

type OrderStore interface {
	GetOrdersByCustomerID(customerID int) ([]models.Order, error)
	GetOrdersByOrderNumber(orderNumber int) (*models.Order, error)
	GetOrdersByFilter(customerIDs, orderNumbers []int) ([]models.Order, error)
	GetOrderByID(orderID int) (*models.Order, error)
}
