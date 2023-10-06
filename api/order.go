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
	var cIDs []int
	var oRefs []string

	customerIDs, ok1 := r.URL.Query()["customerIds"]
	orderReferences, ok2 := r.URL.Query()["orderReferences"]
	fmt.Printf("orderReferences: %v\n", orderReferences)
	if ok1 && len(customerIDs[0]) > 0 {
		for _, id := range strings.Split(customerIDs[0], ",") {
			if i, err := strconv.Atoi(id); err == nil {
				cIDs = append(cIDs, i)
			}
		}
	}

	if ok2 && len(orderReferences[0]) > 0 {
		for _, str := range strings.Split(orderReferences[0], ",") {
			oRefs = append(oRefs, str)
		}
	}
	orders, err := h.OrderStore.GetOrdersByFilter(cIDs, oRefs)
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
