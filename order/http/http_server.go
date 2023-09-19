package http

import (
	"fmt"
	"github.com/bludot/tempmee/order/config"
	"github.com/bludot/tempmee/order/http/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func SetupHttpServer(cfg config.Config) *mux.Router {

	r := mux.NewRouter()

	//database := db.NewDatabase(cfg.DBConfig)
	//userRepository := userRepo.NewUserRepository(database)
	//userService := user.NewUserService(userRepository)
	//jwtService := jwt.NewJWTService(cfg.JWTConfig, map[string]interface{}{})
	//httpResolver := &http_utils.HTTPResolver{
	//	Config:      cfg,
	//
	//	JWTService:  jwtService,
	//}

	r.HandleFunc("/healthcheck", handlers.HealthCheckHandler()).Methods("GET")
	//r.HandleFunc("/api/user/register", userHanlders.RegisterUserHandler(httpResolver)).Methods("POST")
	//r.HandleFunc("/api/user/signin", userHanlders.SignInUserHandler(httpResolver)).Methods("POST")

	return r
}

func StartHttpServer() error {
	cfg := config.LoadConfigOrPanic()
	router := SetupHttpServer(cfg)

	log.Printf("connect to http://localhost:%d/", cfg.AppConfig.Port)

	return http.ListenAndServe(fmt.Sprintf(":%d", cfg.AppConfig.Port), router)
}
