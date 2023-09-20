package http_utils

import (
	"github.com/bludot/tempmee/order/config"
	"github.com/bludot/tempmee/order/internal/services/inventory_api"
	"github.com/bludot/tempmee/order/internal/services/jwt"
	"github.com/bludot/tempmee/order/internal/services/order"
)

type HTTPResolver struct {
	Config       config.Config
	OrderService order.OrderServiceImpl
	InventoryApi inventory_api.InventoryApiImpl
	JWTService   jwt.JWTServiceImpl
}
