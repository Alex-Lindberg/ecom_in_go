package api

import (
	"ecom_in_go/storage"
	"encoding/json"
	"net/http"
	"strconv"
)

type OrderLineHandler struct {
	OrderLineStore storage.OrderLineStore
}

func (h *OrderLineHandler) GetOrderLines(w http.ResponseWriter, r *http.Request) {
	var oID int

	orderID, ok1 := r.URL.Query()["orderId"]

	if ok1 && len(orderID[0]) > 0 {
		// convert to int using atoi
		if i, err := strconv.Atoi(orderID[0]); err == nil {
			oID = i
		}
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
