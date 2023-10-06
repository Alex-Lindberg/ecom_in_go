package api

import (
	"ecom_in_go/storage"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func (handler *CustomerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	println(vars["id"])
	customerID, err := strconv.Atoi(vars["id"])
	customer, err := handler.CustomerStore.GetCustomerByID(customerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}
