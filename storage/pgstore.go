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
	if err := store.DB.Preload("Variants").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (store *PGStore) GetProductByID(productID int) (*product.Product, error) {
	var product product.Product
	if err := store.DB.Preload("Variants").Where("id = ?", productID).Find(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
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

func (store *PGStore) GetOrCreateProductAndVariant(newProduct product.OrderLineProductRequest, variant variant.VariantRequest) (product.Product, error) {
	var existingProduct product.Product
	if err := store.DB.Preload("Variants").Where("name = ?", newProduct.Name).Find(&existingProduct).Error; err != nil {
		if err := store.DB.Create(&newProduct).Error; err != nil {
			return existingProduct, err
		} //
		return existingProduct, nil
	}
	return existingProduct, nil
}

func (store *PGStore) GetOrCreateCustomer(newCustomer customer.CustomerRequest) (customer.Customer, error) {
	var existingCustomer customer.Customer
	if err := store.DB.Where("email = ?", newCustomer.Email).Find(&existingCustomer).Error; err != nil {
		if err := store.DB.Create(&newCustomer).Error; err != nil {
			return existingCustomer, err
		}
		return existingCustomer, nil
	}
	return existingCustomer, nil
}

func (store *PGStore) GetOrCreateOrder(newOrder *order.Order) error {
	var existingOrder order.Order
	if err := store.DB.Where("order_number = ?", newOrder.OrderNumber).Find(&existingOrder).Error; err != nil {
		if err := store.DB.Create(&newOrder).Error; err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (store *PGStore) CreateOrderLines(orderLines []orderline.OrderLine, orderID int) ([]orderline.OrderLine, error) {
	if err := store.DB.CreateInBatches(orderLines, 100).Error; err != nil {
		store.DB.Rollback()
		return nil, err
	}
	store.DB.Commit()

	return orderLines, nil
}
