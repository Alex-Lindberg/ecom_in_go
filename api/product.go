package api

import (
	"ecom_in_go/storage"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductHandler struct {
	ProductStore storage.ProductStore
}

func (h *ProductHandler) GetProducts(writer http.ResponseWriter, request *http.Request) {
	products, err := h.ProductStore.GetProducts()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(products)
}

func (h *ProductHandler) GetProduct(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	println(vars["id"])
	productID, err := strconv.Atoi(vars["id"])
	product, err := h.ProductStore.GetProductByID(productID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(product)
}
