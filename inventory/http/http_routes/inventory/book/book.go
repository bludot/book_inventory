package book

import (
	"github.com/bludot/tempmee/inventory/http/http_utils"
	"github.com/gorilla/mux"
	"net/http"
)

func findAllBooks(resolver *http_utils.HTTPResolver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func findBookById(resolver *http_utils.HTTPResolver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func findBookByTitle(resolver *http_utils.HTTPResolver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func findBooksByAuthor(resolver *http_utils.HTTPResolver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func RegisterBookRoutes(resolver *http_utils.HTTPResolver, router *mux.Router) {
	router.HandleFunc("/api/inventory/books", findAllBooks(resolver)).Methods("GET")
	router.HandleFunc("/api/inventory/books/{id}", findBookById(resolver)).Methods("GET")
	router.HandleFunc("/api/inventory/books/{title}", findBookByTitle(resolver)).Methods("GET")
	router.HandleFunc("/api/inventory/books/{author}", findBooksByAuthor(resolver)).Methods("GET")
}
