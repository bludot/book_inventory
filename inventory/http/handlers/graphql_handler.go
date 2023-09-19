package handlers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/bludot/tempmee/inventory/config"
	"github.com/bludot/tempmee/inventory/graph"
	"github.com/bludot/tempmee/inventory/graph/generated"
	"github.com/bludot/tempmee/inventory/internal/services/jwt"

	"github.com/bludot/tempmee/inventory/internal/directives"
	"net/http"
)

func BuildGraphQLHandler(conf config.Config) http.Handler {
	//database := db.NewDatabase(conf.DBConfig)
	jwtService := jwt.NewJWTService(conf.JWTConfig, map[string]interface{}{})
	resolvers := &graph.Resolver{
		Config:     conf,
		JWTService: jwtService,
	}

	cfg := generated.Config{Resolvers: resolvers, Directives: directives.GetDirectives()}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(cfg))

	return srv
}
