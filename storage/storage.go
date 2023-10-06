package storage

import "ecom_in_go/models"

type ProductStore interface {
	GetProducts() ([]models.Product, error)
	GetProductByID(productID int) (*models.Product, error)
}

type VariantStore interface {
	GetVariantsByProductID(productID int) ([]models.Variant, error)
	// GetVariantsByProductIDs(productIDs []int) (map[int][]models.Variant, error)
}

type CustomerStore interface {
	GetCustomers() ([]models.Customer, error)
	GetCustomerByID(customerID int) (*models.Customer, error)
}

type OrderStore interface {
	GetOrdersByCustomerID(customerID int) ([]models.Order, error)
	GetOrdersByOrderReference(orderReference int) (*models.Order, error)
	GetOrdersByFilter(orderIDs []int, orderNumbers []string, customerIDs []int, orderReferences []string) ([]models.Order, error)
	GetOrderByID(orderID string) (*models.Order, error)
}

type OrderLineStore interface {
	GetOrderLinesByOrderID(orderID string) ([]models.OrderLine, error)
	GetOrderLinesByOrderIDs(orderIDs []string) (map[string][]models.OrderLine, error)
}
