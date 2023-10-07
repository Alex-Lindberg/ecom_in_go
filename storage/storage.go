package storage

import (
	"ecom_in_go/models/customer"
	"ecom_in_go/models/order"
	"ecom_in_go/models/orderline"
	"ecom_in_go/models/product"
	"ecom_in_go/models/variant"
)

type ProductStore interface {
	GetProducts() ([]product.Product, error)
	GetProductByID(productID int) (*product.Product, error)
}

type VariantStore interface {
	GetVariantsByProductID(productID int) ([]variant.Variant, error)
	// GetVariantsByProductIDs(productIDs []int) (map[int][]variant.Variant, error)
}

type CustomerStore interface {
	GetCustomers() ([]customer.Customer, error)
	GetCustomerByID(customerID int) (*customer.Customer, error)
}

type OrderStore interface {
	GetOrdersByCustomerID(customerID int) ([]order.Order, error)
	GetOrdersByOrderReference(orderReference int) (*order.Order, error)
	GetOrdersByFilter(orderIDs []int, orderNumbers []string, customerIDs []int, orderReferences []string) ([]order.Order, error)
	GetOrderByID(orderID string) (*order.Order, error)
}

type OrderLineStore interface {
	GetOrderLinesByOrderID(orderID int) ([]orderline.OrderLine, error)
	GetOrderLinesByOrderIDs(orderIDs []int) (map[int][]orderline.OrderLine, error)
}
