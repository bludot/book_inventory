package http_utils

import (
	"github.com/bludot/tempmee/order/config"
	"github.com/bludot/tempmee/order/internal/services/jwt"
	"github.com/bludot/tempmee/order/internal/services/order"
)

type HTTPResolver struct {
	Config       config.Config
	OrderService order.OrderServiceImpl
	JWTService   jwt.JWTServiceImpl
}
