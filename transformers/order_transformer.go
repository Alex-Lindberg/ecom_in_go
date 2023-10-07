package transformers

import (
	"ecom_in_go/models/customer"
	"ecom_in_go/models/order"
	"ecom_in_go/models/orderline"
	"ecom_in_go/models/product"
	"ecom_in_go/models/variant"
	"fmt"
)

// Convert order domain object to DTO
func ToOrderResponse(domainOrder order.Order) order.OrderResponse {
	// Map the basic order details
	orderResponse := order.OrderResponse{
		ID:          domainOrder.ID,
		OrderNumber: domainOrder.OrderNumber,
		UpdatedAt:   domainOrder.UpdatedAt.String(),
		CreatedAt:   domainOrder.CreatedAt.String(),
	}

	// Map the customer detail
	orderResponse.Customer = customer.CustomerResponse{
		ID:        domainOrder.Customer.ID,
		Name:      domainOrder.Customer.Name,
		Email:     domainOrder.Customer.Email,
		UpdatedAt: domainOrder.Customer.UpdatedAt.String(),
		CreatedAt: domainOrder.Customer.CreatedAt.String(),
	}

	fmt.Println("domainOrder.OrderLines", domainOrder.OrderLines)

	// Map the order lines and included products
	orderLines := make([]orderline.OrderLineResponse, len(domainOrder.OrderLines))
	for i, ol := range domainOrder.OrderLines {

		// Filter out the variant by orderline.VariantID
		variant := variant.Variant{}
		for _, v := range ol.Product.Variants {
			fmt.Println(v.ID, ol.VariantID)
			if v.ID == ol.VariantID {
				variant = v
				break
			}
		}

		orderLines[i] = orderline.OrderLineResponse{
			ID:        ol.ID,
			ProductID: ol.ProductID,
			VariantID: ol.VariantID,
			Product: product.OrderLineProductResponse{
				ID:        ol.Product.ID,
				Name:      ol.Product.Name,
				Variant:   ToVariantResponse(variant),
				UpdatedAt: ol.Product.UpdatedAt.String(),
				CreatedAt: ol.Product.CreatedAt.String(),
			},
			Price:     ol.Price,
			Quantity:  ol.Quantity,
			UpdatedAt: ol.UpdatedAt.String(),
			CreatedAt: ol.CreatedAt.String(),
		}
	}

	orderResponse.OrderLines = orderLines

	return orderResponse
}

func TransformOrders(domainOrders []order.Order) []order.OrderResponse {
	orderResponses := make([]order.OrderResponse, len(domainOrders))

	for i, domainOrder := range domainOrders {
		orderResponses[i] = ToOrderResponse(domainOrder) // Assume you already have a function that transforms a single Order to OrderResponse
	}

	return orderResponses
}
