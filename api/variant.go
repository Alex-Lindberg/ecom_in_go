package api

import (
	"ecom_in_go/storage"
	"encoding/json"
	"net/http"
	"strconv"
)

type VariantHandler struct {
	VariantStore storage.VariantStore
}

func (h *VariantHandler) GetVariantsByProductID(writer http.ResponseWriter, request *http.Request) {
	ids, ok := request.URL.Query()["productId"]
	if !ok || len(ids[0]) < 1 {
		http.Error(writer, "Url Param 'productId' is missing", http.StatusBadRequest)
		return
	}

	productID, err := strconv.Atoi(ids[0])
	variants, err := h.VariantStore.GetVariantsByProductID(productID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(variants)
}

// func (h *VariantHandler) GetVariantsByProductIDs(writer http.ResponseWriter, request *http.Request) {
// 	productIDs, err := getProductIDs(request)

// 	if err != nil {
// 		http.Error(writer, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	variants, err := h.Store.GetVariantsByProductIDs(productIDs)
// 	if err != nil {
// 		http.Error(writer, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	writer.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(writer).Encode(variants)
// }
