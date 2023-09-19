package http_utils

import (
	"github.com/bludot/tempmee/order/config"
	"github.com/bludot/tempmee/order/internal/services/jwt"
)

type HTTPResolver struct {
	Config     config.Config
	JWTService jwt.JWTServiceImpl
}
