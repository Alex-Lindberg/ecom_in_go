package api

import (
	"ecom_in_go/storage"
	"encoding/json"
	"fmt"
	"net/http"
)

type OrderLineHandler struct {
	OrderLineStore storage.OrderLineStore
}

func (h *OrderLineHandler) GetOrderLines(w http.ResponseWriter, r *http.Request) {
	var oID string

	orderID, ok1 := r.URL.Query()["orderId"]
	fmt.Println(orderID)
	if ok1 && len(orderID[0]) > 0 {
		oID = orderID[0]
	}
	println(oID)
	orderLines, err := h.OrderLineStore.GetOrderLinesByOrderID(oID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orderLines)
}
