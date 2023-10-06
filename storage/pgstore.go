package storage

import (
	"ecom_in_go/models"

	"gorm.io/gorm"
)

type PGStore struct {
	DB *gorm.DB
}

func (store *PGStore) GetProducts() ([]models.Product, error) {
	var products []models.Product
	if err := store.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (store *PGStore) GetProductByID(productID int) (*models.Product, error) {
	var product models.Product
	if err := store.DB.Where("id = ?", productID).Find(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (store *PGStore) GetVariantsByProductID(productID int) ([]models.Variant, error) {
	var variants []models.Variant
	if err := store.DB.Where("product_id = ?", productID).Find(&variants).Error; err != nil {
		return nil, err
	}
	return variants, nil
}

func (store *PGStore) GetCustomers() ([]models.Customer, error) {
	var customers []models.Customer
	if err := store.DB.Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}

func (store *PGStore) GetCustomerByID(customerID int) (*models.Customer, error) {
	var customer models.Customer
	if err := store.DB.Where("id = ?", customerID).Find(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (store *PGStore) GetOrdersByCustomerID(customerID int) ([]models.Order, error) {
	var orders []models.Order
	if err := store.DB.Where("customer_id = ?", customerID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (store *PGStore) GetOrdersByOrderReference(orderReference int) (*models.Order, error) {
	var order models.Order
	if err := store.DB.Where("order_reference = ?", orderReference).Find(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (store *PGStore) GetOrderByID(orderID string) (*models.Order, error) {
	var order models.Order
	if err := store.DB.Where("id = ?", orderID).Find(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (store *PGStore) GetOrdersByFilter(orderIDs []int, orderNumbers []string, customerIDs []int, orderReferences []string) ([]models.Order, error) {
	var orders []models.Order
	query := store.DB

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

func (store *PGStore) GetOrderLinesByOrderID(orderID string) ([]models.OrderLine, error) {
	var orderLines []models.OrderLine
	println(orderID)
	if err := store.DB.Where("order_id = ?", orderID).Find(&orderLines).Error; err != nil {
		return nil, err
	}
	return orderLines, nil
}

func (store *PGStore) GetOrderLinesByOrderIDs(orderIDs []string) (map[string][]models.OrderLine, error) {
	var orderLines []models.OrderLine
	if err := store.DB.Where("order_id IN (?)", orderIDs).Find(&orderLines).Error; err != nil {
		return nil, err
	}

	orderLinesMap := make(map[string][]models.OrderLine)
	for _, orderLine := range orderLines {
		orderLinesMap[orderLine.OrderID] = append(orderLinesMap[orderLine.OrderID], orderLine)
	}
	return orderLinesMap, nil
}
