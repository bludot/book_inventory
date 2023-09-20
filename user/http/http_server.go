package http

import (
	"fmt"
	"github.com/bludot/tempmee/user/config"
	"github.com/bludot/tempmee/user/http/handlers"
	userHanlders "github.com/bludot/tempmee/user/http/http_routes/user"
	"github.com/bludot/tempmee/user/http/http_utils"
	"github.com/bludot/tempmee/user/internal/db"
	userRepo "github.com/bludot/tempmee/user/internal/db/repositories/user"
	"github.com/bludot/tempmee/user/internal/services/jwt"
	"github.com/bludot/tempmee/user/internal/services/user"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func SetupHttpServer(cfg config.Config) *mux.Router {

	r := mux.NewRouter()

	database := db.NewDatabase(cfg.DBConfig)
	userRepository := userRepo.NewUserRepository(database)
	userService := user.NewUserService(userRepository)
	jwtService := jwt.NewJWTService(cfg.JWTConfig, map[string]interface{}{})
	httpResolver := &http_utils.HTTPResolver{
		Config:      cfg,
		UserService: userService,
		JWTService:  jwtService,
	}

	r.HandleFunc("/healthcheck", handlers.HealthCheckHandler()).Methods("GET")
	userHanlders.RegisterUserRoutes(httpResolver, r)

	return r
}

func StartHttpServer() error {
	cfg := config.LoadConfigOrPanic()
	router := SetupHttpServer(cfg)

	log.Printf("connect to http://localhost:%d/", cfg.AppConfig.Port)

	return http.ListenAndServe(fmt.Sprintf(":%d", cfg.AppConfig.Port), router)
}
