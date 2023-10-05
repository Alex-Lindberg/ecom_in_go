package api

import (
	"ecom_in_go/storage"
	"encoding/json"
	"net/http"
)

type CustomerHandler struct {
	CustomerStore storage.CustomerStore
}

func (handler *CustomerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := handler.CustomerStore.GetCustomers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
