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
	GetOrCreateProductAndVariant(product product.OrderLineProductRequest, variant variant.VariantRequest) (product.Product, error)
}

type VariantStore interface {
}

type CustomerStore interface {
	GetCustomers() ([]customer.Customer, error)
	GetCustomerByID(customerID int) (*customer.Customer, error)
	GetOrCreateCustomer(customer customer.CustomerRequest) (customer.Customer, error)
}

type OrderStore interface {
	GetOrdersByFilter(orderIDs []int, orderNumbers []string, customerIDs []int, orderReferences []string) ([]order.Order, error)
	GetOrderByID(orderID string) (*order.Order, error)
	GetOrCreateOrder(order *order.Order) error
}

type OrderLineStore interface {
	CreateOrderLines(orderLines []orderline.OrderLineRequest, orderID int) ([]orderline.OrderLine, error)
}
