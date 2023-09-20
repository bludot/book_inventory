package order

import (
	"encoding/json"
	"github.com/bludot/tempmee/order/graph/model"
	"github.com/bludot/tempmee/order/http/http_utils"
	"github.com/bludot/tempmee/order/internal/resolvers"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func FindOrderById(resolver *http_utils.HTTPResolver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		ctx := r.Context()

		orderId, err := strconv.Atoi(params["orderId"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(http_utils.HTTPError{
				Message: "Invalid order id",
				Status:  http.StatusBadRequest,
			})

			return
		}
		order, err := resolver.OrderService.FindOrderById(ctx, orderId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(http_utils.HTTPError{
				Message: "Invalid order id",
				Status:  http.StatusInternalServerError,
			})

			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(order)
	}
}

func CreateOrder(resolver *http_utils.HTTPResolver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var order model.CreateOrderInput
		json.NewDecoder(r.Body).Decode(&order)

		ctx := r.Context()

		createdOrder, err := resolvers.CreateOrder(ctx, resolver.OrderService, order)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(http_utils.HTTPError{
				Message: "Invalid order id",
				Status:  http.StatusInternalServerError,
			})

			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(createdOrder)
	}
}

func UpdateOrderStatus(resolver *http_utils.HTTPResolver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var order model.UpdateOrderStatusInput
		json.NewDecoder(r.Body).Decode(&order)

		ctx := r.Context()

		updatedOrder, err := resolvers.UpdateOrderStatus(ctx, resolver.OrderService, params["orderId"], order.Status)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(http_utils.HTTPError{
				Message: "Invalid order id",
				Status:  http.StatusInternalServerError,
			})

			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updatedOrder)
	}
}

func RegisterOrderRoutes(resolver *http_utils.HTTPResolver, router *mux.Router) {
	router.HandleFunc("/order/{orderId}", FindOrderById(resolver)).Methods("GET")
	router.HandleFunc("/order", CreateOrder(resolver)).Methods("POST")
	router.HandleFunc("/order/{orderId}/status", UpdateOrderStatus(resolver)).Methods("PATCH")

}
