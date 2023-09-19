package http

import (
	"fmt"
	"github.com/bludot/tempmee/inventory/config"
	"github.com/bludot/tempmee/inventory/http/handlers"
	book3 "github.com/bludot/tempmee/inventory/http/http_routes/inventory/book"
	"github.com/bludot/tempmee/inventory/http/http_utils"
	"github.com/bludot/tempmee/inventory/internal/db"
	"github.com/bludot/tempmee/inventory/internal/db/repositories/book"
	book2 "github.com/bludot/tempmee/inventory/internal/services/book"
	"github.com/bludot/tempmee/inventory/internal/services/jwt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func SetupHttpServer(cfg config.Config) *mux.Router {

	r := mux.NewRouter()

	database := db.NewDatabase(cfg.DBConfig)
	bookRepo := book.NewBookRepository(database)
	bookService := book2.NewBookService(bookRepo)
	jwtService := jwt.NewJWTService(cfg.JWTConfig, map[string]interface{}{})
	httpResolver := &http_utils.HTTPResolver{
		Config:      cfg,
		JWTService:  jwtService,
		BookService: bookService,
	}

	r.HandleFunc("/healthcheck", handlers.HealthCheckHandler()).Methods("GET")
	book3.RegisterBookRoutes(httpResolver, r)

	return r
}

func StartHttpServer() error {
	cfg := config.LoadConfigOrPanic()
	router := SetupHttpServer(cfg)

	log.Printf("connect to http://localhost:%d/", cfg.AppConfig.Port)

	return http.ListenAndServe(fmt.Sprintf(":%d", cfg.AppConfig.Port), router)
}
