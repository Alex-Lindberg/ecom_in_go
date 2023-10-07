package storage

import (
	"ecom_in_go/models/customer"
	"ecom_in_go/models/order"
	"ecom_in_go/models/orderline"
	"ecom_in_go/models/product"
	"ecom_in_go/models/variant"

	"gorm.io/gorm"
)

type PGStore struct {
	DB *gorm.DB
}

func (store *PGStore) GetProducts() ([]product.Product, error) {
	var products []product.Product
	if err := store.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (store *PGStore) GetProductByID(productID int) (*product.Product, error) {
	var product product.Product
	if err := store.DB.Where("id = ?", productID).Find(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (store *PGStore) GetVariantsByProductID(productID int) ([]variant.Variant, error) {
	var variants []variant.Variant
	if err := store.DB.Where("product_id = ?", productID).Find(&variants).Error; err != nil {
		return nil, err
	}
	return variants, nil
}

func (store *PGStore) GetCustomers() ([]customer.Customer, error) {
	var customers []customer.Customer
	if err := store.DB.Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}

func (store *PGStore) GetCustomerByID(customerID int) (*customer.Customer, error) {
	var customer customer.Customer
	if err := store.DB.Where("id = ?", customerID).Find(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (store *PGStore) GetOrdersByCustomerID(customerID int) ([]order.Order, error) {
	var orders []order.Order
	if err := store.DB.Where("customer_id = ?", customerID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (store *PGStore) GetOrdersByOrderReference(orderReference int) (*order.Order, error) {
	var order order.Order
	if err := store.DB.Where("order_reference = ?", orderReference).Find(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (store *PGStore) GetOrderByID(orderID string) (*order.Order, error) {
	var order order.Order
	if err := store.DB.Where("id = ?", orderID).Find(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (store *PGStore) GetOrdersByFilter(orderIDs []int, orderNumbers []string, customerIDs []int, orderReferences []string) ([]order.Order, error) {
	var orders []order.Order
	query := store.DB.
		Preload("Customer").
		Preload("OrderLines").
		Preload("OrderLines.Product").
		Preload("OrderLines.Product.Variants")

	if len(orderIDs) > 0 {
		query = query.Or("id IN (?)", orderIDs)
	}
	if len(orderNumbers) > 0 {
		query = query.Or("order_number IN (?)", orderNumbers)
	}
	if len(orderReferences) > 0 {
		query = query.Or("order_reference IN (?)", orderReferences)
	}
	if len(customerIDs) > 0 {
		query = query.Or("customer_id IN (?)", customerIDs)
	}

	if err := query.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (store *PGStore) GetOrderLinesByOrderID(orderID int) ([]orderline.OrderLine, error) {
	var orderLines []orderline.OrderLine
	println(orderID)
	if err := store.DB.Where("order_id = ?", orderID).Find(&orderLines).Error; err != nil {
		return nil, err
	}
	return orderLines, nil
}

func (store *PGStore) GetOrderLinesByOrderIDs(orderIDs []int) (map[int][]orderline.OrderLine, error) {
	var orderLines []orderline.OrderLine
	if err := store.DB.Where("order_id IN (?)", orderIDs).Find(&orderLines).Error; err != nil {
		return nil, err
	}

	orderLinesMap := make(map[int][]orderline.OrderLine)
	for _, orderline := range orderLines {
		orderLinesMap[orderline.OrderID] = append(orderLinesMap[orderline.OrderID], orderline)
	}
	return orderLinesMap, nil
}
