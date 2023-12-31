package http

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/bludot/tempmee/user/config"
	"github.com/bludot/tempmee/user/http/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func SetupGraphQLServer(cfg config.Config) *mux.Router {

	router := mux.NewRouter()

	router.Handle("/ui/playground", playground.Handler("GraphQL playground", "/graphql")).Methods("GET")
	router.Handle("/graphql", handlers.BuildGraphQLHandler(cfg)).Methods("POST")
	router.Handle("/healthcheck", handlers.HealthCheckHandler()).Methods("GET")

	return router
}

func StartGraphQLServer() error {
	cfg := config.LoadConfigOrPanic()
	router := SetupGraphQLServer(cfg)

	log.Printf("connect to http://localhost:%d/ for GraphQL playground", cfg.AppConfig.Port)

	return http.ListenAndServe(fmt.Sprintf(":%d", cfg.AppConfig.Port), router)
}
