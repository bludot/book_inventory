package order

import (
	"encoding/json"
	"github.com/bludot/tempmee/order/graph/model"
	"github.com/bludot/tempmee/order/http/http_utils"
	"github.com/bludot/tempmee/order/internal/resolvers"
	"github.com/bludot/tempmee/order/internal/services/inventory_api"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type OrderExtended struct {
	model.Order
	Books []map[string]interface{}
}

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

func FindOrdersByUserId(resolver *http_utils.HTTPResolver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		ctx := r.Context()

		userId := params["userId"]

		orders, err := resolvers.GetOrdersByUserID(ctx, resolver.OrderService, userId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(http_utils.HTTPError{
				Message: "Invalid user id",
				Status:  http.StatusInternalServerError,
			})

			return
		}

		var extendedOrders []OrderExtended

		for _, order := range orders {
			orderExtended := OrderExtended{
				Order: *order,
			}
			for _, productId := range order.Products {
				book, err := inventory_api.NewInventoryApi(resolver.Config.InventoryAPIConfig).GetBookByID(ctx, productId)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(w).Encode(http_utils.HTTPError{
						Message: "Invalid product id",
						Status:  http.StatusInternalServerError,
					})

					return
				}
				orderExtended.Books = append(orderExtended.Books, book)
			}
			extendedOrders = append(extendedOrders, orderExtended)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(extendedOrders)

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
	router.HandleFunc("/api/order/{orderId}", FindOrderById(resolver)).Methods("GET")
	router.HandleFunc("/api/order", CreateOrder(resolver)).Methods("POST")
	router.HandleFunc("/api/order/{orderId}/status", UpdateOrderStatus(resolver)).Methods("PATCH")
	router.HandleFunc("/api/order/user/{userId}", FindOrdersByUserId(resolver)).Methods("GET")

}
