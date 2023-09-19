package handlers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/bludot/tempmee/order/config"
	"github.com/bludot/tempmee/order/graph"
	"github.com/bludot/tempmee/order/graph/generated"
	"github.com/bludot/tempmee/order/internal/services/jwt"

	"github.com/bludot/tempmee/order/internal/directives"
	"net/http"
)

func BuildGraphQLHandler(conf config.Config) http.Handler {
	//database := db.NewDatabase(conf.DBConfig)
	//userRepository := user2.NewUserRepository(database)
	//userService := user.NewUserService(userRepository)
	jwtService := jwt.NewJWTService(conf.JWTConfig, map[string]interface{}{})
	resolvers := &graph.Resolver{
		Config: conf,
		//UserService: userService,
		JWTService: jwtService,
	}

	cfg := generated.Config{Resolvers: resolvers, Directives: directives.GetDirectives()}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(cfg))

	return srv
}
