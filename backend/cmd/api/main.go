package main

import (
	"backend/internal/database"
	"backend/internal/handlers"
	logging "backend/internal/logger"
	"backend/internal/mail"
	model "backend/internal/models"
	"backend/internal/repositories"
	"backend/internal/routes"
	"backend/internal/server"
	"backend/internal/services"
	"context"
	"errors"
	"github.com/joho/godotenv"
	"log"

	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file")
	}

	logger := logging.NewLogger()

	db, err := database.New()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	queries := model.New(db)

	resendMailClient := mail.NewMailClient()

	authRepo := repositories.NewAuthRepository(db, queries)

	authService := services.NewAuthServices(authRepo, resendMailClient)

	authHandler := handlers.NewAuthHandler(authService)

	router := routes.NewRoutes(authHandler).RegisterRoutes()

	app := server.NewServer(router)

	// Start the server in a goroutine
	go func() {
		if err := app.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.LogError().Msgf("Failed to start server: %v", err)
		}
	}()

	// Channel to listen for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// Wait for the interrupt signal
	<-quit

	// Create a context with a timeout for the shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Attempt to gracefully shutdown the server
	if err := app.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to shutdown server: %v", err)
	}

}
