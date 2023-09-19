package http_utils

import (
	"github.com/bludot/tempmee/user/config"
	"github.com/bludot/tempmee/user/internal/services/jwt"
	"github.com/bludot/tempmee/user/internal/services/user"
)

type HTTPResolver struct {
	Config      config.Config
	UserService user.UserServiceImpl
	JWTService  jwt.JWTServiceImpl
}
