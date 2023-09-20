package handlers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/bludot/tempmee/order/config"
	"github.com/bludot/tempmee/order/graph"
	"github.com/bludot/tempmee/order/graph/generated"
	"github.com/bludot/tempmee/order/internal/db"
	"github.com/bludot/tempmee/order/internal/db/repositories/order"
	"github.com/bludot/tempmee/order/internal/services/inventory_api"
	"github.com/bludot/tempmee/order/internal/services/jwt"
	order2 "github.com/bludot/tempmee/order/internal/services/order"

	"github.com/bludot/tempmee/order/internal/directives"
	"net/http"
)

func BuildGraphQLHandler(conf config.Config) http.Handler {
	database := db.NewDatabase(conf.DBConfig)
	orderRepository := order.NewOrderRepository(database)
	orderService := order2.NewOrderService(orderRepository)
	inventoryApi := inventory_api.NewInventoryApi(conf.InventoryAPIConfig)
	jwtService := jwt.NewJWTService(conf.JWTConfig, map[string]interface{}{})
	resolvers := &graph.Resolver{
		Config:       conf,
		OrderService: orderService,
		InventoryApi: inventoryApi,
		JWTService:   jwtService,
	}

	cfg := generated.Config{Resolvers: resolvers, Directives: directives.GetDirectives()}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(cfg))

	return srv
}
