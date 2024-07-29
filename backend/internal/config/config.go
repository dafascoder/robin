package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	DBHost    string
	DBPort    int
	DBName    string
	DBUser    string
	DBPass    string
	Port      int
	DBSSL     bool
	RESENDAPI string
}

var Env = initConfig()

const (
	fallbackDBPort = 5432
)

func initConfig() Config {
	return Config{
		DBHost:    getEnvOrError("DB_HOST"),
		DBPort:    getEnvAsInt("DB_PORT", fallbackDBPort),
		DBName:    getEnvOrError("DB_NAME"),
		DBUser:    getEnvOrError("DB_USERNAME"),
		DBPass:    getEnvOrError("DB_PASSWORD"),
		DBSSL:     getEnvAsBool("DB_SSLMODE", false),
		Port:      getEnvAsInt("PORT", 8080),
		RESENDAPI: getEnvOrError("RESEND_API"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvOrError(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	panic(fmt.Sprintf("Environment variable %s is not set", key))

}

func getEnvAsInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.Atoi(value)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}

func getEnvAsBool(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		b, err := strconv.ParseBool(value)
		if err != nil {
			return fallback
		}

		return b
	}

	return fallback
}
