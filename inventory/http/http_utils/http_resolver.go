package http_utils

import (
	"github.com/bludot/tempmee/inventory/config"
	"github.com/bludot/tempmee/inventory/internal/services/book"
	"github.com/bludot/tempmee/inventory/internal/services/jwt"
)

type HTTPResolver struct {
	Config      config.Config
	JWTService  jwt.JWTServiceImpl
	BookService book.BookServiceImpl
}
