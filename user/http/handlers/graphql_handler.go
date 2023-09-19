package handlers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/bludot/tempmee/user/config"
	"github.com/bludot/tempmee/user/graph"
	"github.com/bludot/tempmee/user/graph/generated"
	"github.com/bludot/tempmee/user/internal/db"
	"github.com/bludot/tempmee/user/internal/services/jwt"

	user2 "github.com/bludot/tempmee/user/internal/db/repositories/user"
	"github.com/bludot/tempmee/user/internal/directives"
	"github.com/bludot/tempmee/user/internal/services/user"
	"net/http"
)

func BuildGraphQLHandler(conf config.Config) http.Handler {
	database := db.NewDatabase(conf.DBConfig)
	userRepository := user2.NewUserRepository(database)
	userService := user.NewUserService(userRepository)
	jwtService := jwt.NewJWTService(conf.JWTConfig, map[string]interface{}{})
	resolvers := &graph.Resolver{
		Config:      conf,
		UserService: userService,
		JWTService:  jwtService,
	}

	cfg := generated.Config{Resolvers: resolvers, Directives: directives.GetDirectives()}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(cfg))

	return srv
}
