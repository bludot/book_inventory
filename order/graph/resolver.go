package graph

import (
	"github.com/bludot/tempmee/order/config"
	"github.com/bludot/tempmee/order/internal/services/inventory_api"
	"github.com/bludot/tempmee/order/internal/services/jwt"
	"github.com/bludot/tempmee/order/internal/services/order"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Config       config.Config
	OrderService order.OrderServiceImpl
	InventoryApi inventory_api.InventoryApiImpl
	JWTService   jwt.JWTServiceImpl
}
