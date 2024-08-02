package routes

import (
	"backend/internal/handlers"
	"net/http"
)

type Routes struct {
	AuthHandler *handlers.AuthHandler
}

const apiVersion = "/v1/api"

func NewRoutes(authHandler *handlers.AuthHandler) *Routes {
	return &Routes{
		AuthHandler: authHandler,
	}

}

func (r *Routes) RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Auth Routes
	mux.HandleFunc("POST /login", r.AuthHandler.HandleSignIn)
	mux.HandleFunc("POST /register", r.AuthHandler.HandleSignUp)
	mux.HandleFunc("POST /logout", r.AuthHandler.HandleSignOut)
	mux.HandleFunc("POST /verify", r.AuthHandler.HandleVerifyAccount)

	v1 := http.NewServeMux()
	v1.Handle("/v1/api/", http.StripPrefix("/v1/api", mux))

	return v1
}
