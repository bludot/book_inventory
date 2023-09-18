package graph

import (
	"github.com/bludot/tempmee/user/config"
	"github.com/bludot/tempmee/user/internal/services/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Config      config.Config
	UserService user.UserServiceImpl
}
