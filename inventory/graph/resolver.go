package graph

import (
	"github.com/bludot/tempmee/inventory/config"
	"github.com/bludot/tempmee/inventory/internal/services/jwt"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Config     config.Config
	JWTService jwt.JWTServiceImpl
}
