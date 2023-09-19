package book

import (
	"encoding/json"
	"github.com/bludot/tempmee/inventory/http/http_utils"
	"github.com/bludot/tempmee/inventory/internal/resolvers"
	"github.com/gorilla/mux"
	"net/http"
)

func FindAllBooks(resolver *http_utils.HTTPResolver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		books, err := resolvers.Books(ctx, resolver.BookService)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(http_utils.HTTPError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})

			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)

	}
}

func FindBookById(resolver *http_utils.HTTPResolver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		if _, ok := params["id"]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(http_utils.HTTPError{
				Message: "Invalid request",
				Status:  http.StatusBadRequest,
			})

			return
		}

		ctx := r.Context()

		book, err := resolvers.BookByID(ctx, resolver.BookService, params["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(http_utils.HTTPError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})

			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(book)

	}
}

func FindBookByTitle(resolver *http_utils.HTTPResolver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		if _, ok := params["title"]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(http_utils.HTTPError{
				Message: "Invalid request",
				Status:  http.StatusBadRequest,
			})

			return
		}

		ctx := r.Context()

		book, err := resolvers.BookByTitle(ctx, resolver.BookService, params["title"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(http_utils.HTTPError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})

			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(book)

	}
}

func FindBooksByAuthor(resolver *http_utils.HTTPResolver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		if _, ok := params["author"]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(http_utils.HTTPError{
				Message: "Invalid request",
				Status:  http.StatusBadRequest,
			})

			return
		}

		ctx := r.Context()

		books, err := resolvers.BooksByAuthor(ctx, resolver.BookService, params["author"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(http_utils.HTTPError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})

			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)

	}
}

func RegisterBookRoutes(resolver *http_utils.HTTPResolver, router *mux.Router) {
	router.HandleFunc("/api/inventory/books", FindAllBooks(resolver)).Methods("GET")
	router.HandleFunc("/api/inventory/books/{id}", FindBookById(resolver)).Methods("GET")
	router.HandleFunc("/api/inventory/books/title/{title}", FindBookByTitle(resolver)).Methods("GET")
	router.HandleFunc("/api/inventory/books/author/{author}", FindBooksByAuthor(resolver)).Methods("GET")
}
