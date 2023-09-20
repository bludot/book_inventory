package http

import (
	"fmt"
	"github.com/bludot/tempmee/order/config"
	"github.com/bludot/tempmee/order/http/handlers"
	order3 "github.com/bludot/tempmee/order/http/http_routes/order"
	"github.com/bludot/tempmee/order/http/http_utils"
	"github.com/bludot/tempmee/order/internal/db"
	"github.com/bludot/tempmee/order/internal/db/repositories/order"
	"github.com/bludot/tempmee/order/internal/services/inventory_api"
	"github.com/bludot/tempmee/order/internal/services/jwt"
	order2 "github.com/bludot/tempmee/order/internal/services/order"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func SetupHttpServer(cfg config.Config) *mux.Router {

	r := mux.NewRouter()

	database := db.NewDatabase(cfg.DBConfig)
	orderRepository := order.NewOrderRepository(database)
	orderService := order2.NewOrderService(orderRepository)
	inventoryApi := inventory_api.NewInventoryApi(cfg.InventoryAPIConfig)
	jwtService := jwt.NewJWTService(cfg.JWTConfig, map[string]interface{}{})
	httpResolver := &http_utils.HTTPResolver{
		Config:       cfg,
		OrderService: orderService,
		InventoryApi: inventoryApi,
		JWTService:   jwtService,
	}

	r.HandleFunc("/healthcheck", handlers.HealthCheckHandler()).Methods("GET")
	order3.RegisterOrderRoutes(httpResolver, r)

	return r
}

func StartHttpServer() error {
	cfg := config.LoadConfigOrPanic()
	router := SetupHttpServer(cfg)

	log.Printf("connect to http://localhost:%d/", cfg.AppConfig.Port)

	return http.ListenAndServe(fmt.Sprintf(":%d", cfg.AppConfig.Port), router)
}
