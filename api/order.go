package api

import (
	"ecom_in_go/storage"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type OrderHandler struct {
	OrderStore storage.OrderStore
}

func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	var cIDs, oIDs []int
	var oNums, oRefs []string

	orderIds, ok1 := r.URL.Query()["orderIds"]
	orderNumbers, ok2 := r.URL.Query()["orderNumbers"]
	customerIDs, ok3 := r.URL.Query()["customerIds"]
	orderReferences, ok4 := r.URL.Query()["orderReferences"]

	if ok1 && len(orderIds[0]) > 0 {
		for _, id := range strings.Split(orderIds[0], ",") {
			if i, err := strconv.Atoi(id); err == nil {
				oIDs = append(oIDs, i)
			}
		}
	}

	if ok2 && len(orderNumbers[0]) > 0 {
		for _, str := range strings.Split(orderNumbers[0], ",") {
			oNums = append(oNums, str)
		}
	}

	if ok3 && len(customerIDs[0]) > 0 {
		for _, id := range strings.Split(customerIDs[0], ",") {
			if i, err := strconv.Atoi(id); err == nil {
				cIDs = append(cIDs, i)
			}
		}
	}

	if ok4 && len(orderReferences[0]) > 0 {
		for _, str := range strings.Split(orderReferences[0], ",") {
			oRefs = append(oRefs, str)
		}
	}

	orders, err := h.OrderStore.GetOrdersByFilter(oIDs, oNums, cIDs, oRefs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
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
