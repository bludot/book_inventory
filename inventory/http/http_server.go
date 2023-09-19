package http

import (
	"fmt"
	"github.com/bludot/tempmee/inventory/config"
	"github.com/bludot/tempmee/inventory/http/handlers"
	userHanlders "github.com/bludot/tempmee/inventory/http/http_routes/user"
	"github.com/bludot/tempmee/inventory/http/http_utils"
	"github.com/bludot/tempmee/inventory/internal/services/jwt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func SetupHttpServer(cfg config.Config) *mux.Router {

	r := mux.NewRouter()

	//database := db.NewDatabase(cfg.DBConfig)
	jwtService := jwt.NewJWTService(cfg.JWTConfig, map[string]interface{}{})
	httpResolver := &http_utils.HTTPResolver{
		Config:     cfg,
		JWTService: jwtService,
	}

	r.HandleFunc("/healthcheck", handlers.HealthCheckHandler()).Methods("GET")
	r.HandleFunc("/api/user/register", userHanlders.RegisterUserHandler(httpResolver)).Methods("POST")
	r.HandleFunc("/api/user/signin", userHanlders.SignInUserHandler(httpResolver)).Methods("POST")

	return r
}

func StartHttpServer() error {
	cfg := config.LoadConfigOrPanic()
	router := SetupHttpServer(cfg)

	log.Printf("connect to http://localhost:%d/", cfg.AppConfig.Port)

	return http.ListenAndServe(fmt.Sprintf(":%d", cfg.AppConfig.Port), router)
}
