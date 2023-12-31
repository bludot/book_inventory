package user

import (
	"encoding/json"
	httpUtils "github.com/bludot/tempmee/user/http/http_utils"
	"github.com/bludot/tempmee/user/internal/resolvers"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type RegisterUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterUserHandler(resolver *httpUtils.HTTPResolver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		body, err := io.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var registerUserRequest RegisterUserRequest
		err = json.Unmarshal(body, &registerUserRequest)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		user, err := resolvers.Register(ctx, resolver.UserService, registerUserRequest.Name, registerUserRequest.Email, registerUserRequest.Password)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(httpUtils.HTTPError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})

			return
		}

		// return json
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)

	}
}

type SignInUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignInUserHandler(resolver *httpUtils.HTTPResolver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		body, err := io.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var signInUserRequest SignInUserRequest
		err = json.Unmarshal(body, &signInUserRequest)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		user, err := resolvers.SignIn(ctx, resolver.UserService, resolver.JWTService, signInUserRequest.Email, signInUserRequest.Password)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(httpUtils.HTTPError{
				Message: err.Error(),
				Status:  http.StatusUnauthorized,
			})

			return
		}

		// return json
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)

	}
}

func RegisterUserRoutes(resolver *httpUtils.HTTPResolver, router *mux.Router) {
	router.HandleFunc("/api/user/register", RegisterUserHandler(resolver)).Methods("POST")
	router.HandleFunc("/api/user/signin", SignInUserHandler(resolver)).Methods("POST")
}
