package handlers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/bludot/tempmee/inventory/config"
	"github.com/bludot/tempmee/inventory/graph"
	"github.com/bludot/tempmee/inventory/graph/generated"
	"github.com/bludot/tempmee/inventory/internal/db"
	"github.com/bludot/tempmee/inventory/internal/db/repositories/book"
	book2 "github.com/bludot/tempmee/inventory/internal/services/book"
	"github.com/bludot/tempmee/inventory/internal/services/jwt"

	"github.com/bludot/tempmee/inventory/internal/directives"
	"net/http"
)

func BuildGraphQLHandler(conf config.Config) http.Handler {
	database := db.NewDatabase(conf.DBConfig)
	bookRepo := book.NewBookRepository(database)
	bookService := book2.NewBookService(bookRepo)
	jwtService := jwt.NewJWTService(conf.JWTConfig, map[string]interface{}{})
	resolvers := &graph.Resolver{
		Config:      conf,
		JWTService:  jwtService,
		BookService: bookService,
	}

	cfg := generated.Config{Resolvers: resolvers, Directives: directives.GetDirectives()}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(cfg))

	return srv
}
