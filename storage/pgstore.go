package storage

import (
	"ecom_in_go/models"
	"strings"

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

func (store *PGStore) GetOrdersByCustomerID(customerID int) ([]models.Order, error) {
	var orders []models.Order
	if err := store.DB.Where("customer_id = ?", customerID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (store *PGStore) GetOrdersByOrderNumber(orderNumber int) (*models.Order, error) {
	var order models.Order
	if err := store.DB.Where("order_number = ?", orderNumber).Find(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (store *PGStore) GetOrderByID(orderID int) (*models.Order, error) {
	var order models.Order
	if err := store.DB.Where("id = ?", orderID).Find(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (store *PGStore) GetOrdersByFilter(customerIDs, orderNumbers []int) ([]models.Order, error) {
	var orders []models.Order
	query := store.DB

	if len(customerIDs) > 0 {
		placeholders := strings.Trim(strings.Repeat(",?", len(customerIDs))[1:], " ")
		query = query.Where("customer_id IN ("+placeholders+")", customerIDs)
	} else if len(orderNumbers) > 0 {
		placeholders := strings.Trim(strings.Repeat(",?", len(orderNumbers))[1:], " ")
		query = query.Where("order_number IN ("+placeholders+")", orderNumbers)
	}

	if err := query.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
