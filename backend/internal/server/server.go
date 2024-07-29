package server

import (
	"backend/internal/config"
	"backend/internal/middleware"
	"fmt"
	"log"
	"net/http"
	"time"
)

func NewServer(router *http.ServeMux) *http.Server {

	log.Printf("Server is running on port %d", config.Env.Port)

	middlewareChain := middleware.Chain(middleware.LoggingMiddleware)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Env.Port),
		Handler:      middlewareChain(router),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
