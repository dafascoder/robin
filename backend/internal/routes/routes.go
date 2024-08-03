package routes

import (
	"backend/internal/handlers"
	"backend/internal/middleware"
	"net/http"
)

type Routes struct {
	AuthHandler *handlers.AuthHandler
	UserHandler *handlers.UserHandler
}

const apiVersion = "/v1/api"

func NewRoutes(authHandler *handlers.AuthHandler, userHandler *handlers.UserHandler) *Routes {
	return &Routes{
		AuthHandler: authHandler,
		UserHandler: userHandler,
	}

}

func (r *Routes) RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Auth Routes
	mux.HandleFunc("POST /login", r.AuthHandler.HandleSignIn)
	mux.HandleFunc("POST /register", r.AuthHandler.HandleSignUp)
	mux.HandleFunc("POST /logout", r.AuthHandler.HandleSignOut)
	mux.HandleFunc("POST /verify", r.AuthHandler.HandleVerifyAccount)

	// Health Check
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// User Routes
	mux.HandleFunc("GET /current/user", middleware.AuthMiddleware(r.UserHandler.HandleCurrentGetUser))
	mux.HandleFunc("POST /user", r.UserHandler.HandleCreateUser)
	mux.HandleFunc("PUT /user", r.UserHandler.HandleUpdateUser)
	mux.HandleFunc("DELETE /user", middleware.AuthMiddleware(r.UserHandler.HandleDeleteUser))

	v1 := http.NewServeMux()
	v1.Handle("/v1/api/", http.StripPrefix("/v1/api", mux))

	return v1
}
