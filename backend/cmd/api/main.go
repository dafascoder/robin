package main

import (
	"backend/internal/datastore"
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

	db, err := datastore.NewDatabase(context.Background())

	redis, _ := datastore.InitRedis()
	if err != nil {
		logger.LogFatal().Msgf("Failed to connect to redis: %v", err)
	}

	if err != nil {
		logger.LogFatal().Msgf("Failed to connect to datastore: %v", err)
	}

	defer db.Close()

	queries := model.New(db.GetDatabaseInstance())

	resendMailClient := mail.NewMailClient()

	authRepo := repositories.NewAuthRepository(queries)
	userRepo := repositories.NewUserRepository(queries)

	authService := services.NewAuthServices(authRepo, resendMailClient, redis)
	userService := services.NewUserServices(userRepo)

	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)

	router := routes.NewRoutes(authHandler, userHandler).RegisterRoutes()

	app := server.NewServer(router)

	// Channel to listen for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// Start the server in a goroutine
	go func() {
		if err := app.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.LogError().Msgf("Failed to start server: %v", err)
		}
	}()

	logger.LogInfo().Msgf("Server started on port %s", os.Getenv("PORT"))

	// Wait for the interrupt signal
	<-quit
	logging.Logger.LogInfo().Msg("Shutting down server...")

	// Create a context with a timeout for the shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		cancel()
		db.Close()
	}()

	// Attempt to gracefully shutdown the server
	if err := app.Shutdown(ctx); err != nil {
		logger.LogFatal().Msgf("Failed to shutdown server: %v", err)
	}

}
