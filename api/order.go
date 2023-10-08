package api

import (
	"ecom_in_go/models/order"
	"ecom_in_go/models/orderline"
	"ecom_in_go/storage"
	"ecom_in_go/transformers"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type OrderHandler struct {
	OrderStore     storage.OrderStore
	OrderLineStore storage.OrderLineStore
	CustomerStore  storage.CustomerStore
	ProductStore   storage.ProductStore
	VariantStore   storage.VariantStore
}

func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	var cIDs, oIDs []int
	var oNums, oRefs []string

	orderIds, ok1 := r.URL.Query()["orderIds"]
	orderNumbers, ok2 := r.URL.Query()["orderNumbers"]
	customerIDs, ok3 := r.URL.Query()["customerIds"]
	orderReferences, ok4 := r.URL.Query()["orderReferences"]

	if ok1 && len(orderIds[0]) > 0 {
		ids, err := parseCommaSeparatedInt(orderIds[0])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		oIDs = ids
	}

	if ok2 && len(orderNumbers[0]) > 0 {
		ids, err := parseCommaSeparatedString(orderNumbers[0])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		oNums = ids
	}

	if ok3 && len(customerIDs[0]) > 0 {
		ids, err := parseCommaSeparatedInt(customerIDs[0])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cIDs = ids
	}

	if ok4 && len(orderReferences[0]) > 0 {
		ids, err := parseCommaSeparatedString(orderReferences[0])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		oRefs = ids
	}

	orders, err := h.OrderStore.GetOrdersByFilter(oIDs, oNums, cIDs, oRefs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	orderResponses := transformers.TransformOrders(orders) // Transform the orders into DTOs

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orderResponses) // Send the DTOs as JSON
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	// Extracting order ID from the URL
	vars := mux.Vars(r)
	orderID := vars["id"]

	order, err := h.OrderStore.GetOrderByID(orderID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

// prase comma separated string into list of strings
func parseCommaSeparatedString(str string) ([]string, error) {
	var list []string
	for _, id := range strings.Split(str, ",") {
		list = append(list, id)
	}
	return list, nil
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {

	// - Extract body
	// - Parse into json
	// - Validate
	// - Check if Order exists
	// - Check if Customer exists, else create
	// - Check or create product
	// - Check or create variant
	// - Create order
	// - Create order lines

	// 1. Extract body
	var orderRequest order.CreateOrderRequest
	err := json.NewDecoder(r.Body).Decode(&orderRequest)

	// if orderExists(orderRequest.OrderNumber) {
	//     return nil, errors.New("order already exists")
	// }

	// 2. Check or create customer
	customer, err := h.CustomerStore.GetOrCreateCustomer(orderRequest.Customer)
	if err != nil {

		println("Error creating order: %v\n", err)
	}

	// 4. Create the order
	order := &order.Order{
		OrderNumber: orderRequest.OrderNumber,
		CustomerID:  customer.ID,
		// Populate other fields
	}
	if err := h.OrderStore.GetOrCreateOrder(order).Error; err != nil {
		// fatal error
		println("Error creating order: %v\n", err)
		return
	}

	// 5. Create order lines
	var orderlines []orderline.OrderLineRequest
	for _, ol := range orderRequest.OrderLines {
		_, err := h.ProductStore.GetOrCreateProductAndVariant(ol.Product, ol.Product.Variant)
		if err != nil {
			println("Error creating order: %v\n", err)
			return
		}
		orderLine := &orderline.OrderLineRequest{
			OrderID:   order.ID,
			ProductID: ol.ProductID,
			VariantID: ol.VariantID,
			Product:   ol.Product,
			Quantity:  ol.Quantity,
		}
		orderlines = append(orderlines, *orderLine)
	}
	_, err = h.OrderLineStore.CreateOrderLines(orderlines, order.ID)
	if err != nil {
		return
	}

}

// parse comma separated string into list of ints'
func parseCommaSeparatedInt(str string) ([]int, error) {
	var list []int
	for _, id := range strings.Split(str, ",") {
		if i, err := strconv.Atoi(id); err == nil {
			list = append(list, i)
		} else {
			return nil, err
		}
	}
	return list, nil
}
