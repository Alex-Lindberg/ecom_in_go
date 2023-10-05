package api

import (
	"ecom_in_go/storage"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type OrderHandler struct {
	OrderStore storage.OrderStore
}

func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	var cIDs, oNums []int

	customerIDs, ok1 := r.URL.Query()["customerIds"]
	orderNumbers, ok2 := r.URL.Query()["orderNumbers"]

	fmt.Printf("customerIDs: %v, ok1: %v\n", customerIDs, ok1)
	fmt.Printf("orderNumbers: %v, ok2: %v\n", orderNumbers, ok2)

	if ok1 && ok2 {
		http.Error(w, "Cannot filter by both customer IDs and order numbers at the same time", http.StatusBadRequest)
		return
	}

	if ok1 && len(customerIDs[0]) > 0 {
		for _, id := range strings.Split(customerIDs[0], ",") {
			if i, err := strconv.Atoi(id); err == nil {
				cIDs = append(cIDs, i)
			}
		}
	}

	if ok2 && len(orderNumbers[0]) > 0 {
		for _, num := range strings.Split(orderNumbers[0], ",") {
			if i, err := strconv.Atoi(num); err == nil {
				oNums = append(oNums, i)
			}
		}
	}

	orders, err := h.OrderStore.GetOrdersByFilter(cIDs, oNums)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	// Extracting order ID from the URL
	println("Extracting order ID from the URL")
	vars := mux.Vars(r)
	println(vars["id"])
	orderID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	order, err := h.OrderStore.GetOrderByID(orderID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}
