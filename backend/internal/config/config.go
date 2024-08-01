package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	DBHost          string
	DBPort          int
	DBName          string
	DBUser          string
	DBPass          string
	Port            int
	DatabaseSSL     bool
	ResendAPI       string
	DatabaseUrl     string
	RedisUrl        string
	AuthToken       string
	RefreshToken    string
	TokenExpiration struct {
		DurationString string
		Duration       time.Duration
	}
}

// Token Expiration

var Env = initConfig()

const (
	fallbackDBPort = 5432
)

func initConfig() Config {
	tokexpirationStr := getEnvOrError("TOKEN_EXPIRATION")
	duration, expErr := time.ParseDuration(tokexpirationStr)
	if expErr != nil {
		panic(fmt.Sprint("Invalid token expiration"))
	}
	
	return Config{
		DBHost:       getEnvOrError("DB_HOST"),
		DBPort:       getEnvAsInt("DB_PORT", fallbackDBPort),
		DBName:       getEnvOrError("DB_NAME"),
		DBUser:       getEnvOrError("DB_USERNAME"),
		DBPass:       getEnvOrError("DB_PASSWORD"),
		DatabaseSSL:  getEnvAsBool("DB_SSLMODE", false),
		Port:         getEnvAsInt("PORT", 8080),
		ResendAPI:    getEnvOrError("RESEND_API"),
		DatabaseUrl:  getEnvOrError("DATABASE_URL"),
		AuthToken:    getEnvOrError("AUTH_TOKEN_SECRET"),
		RefreshToken: getEnvOrError("REFRESH_TOKEN_SECRET"),
		RedisUrl:     getEnvOrError("REDIS_URL"),
		TokenExpiration: struct {
			DurationString string
			Duration       time.Duration
		}{
			DurationString: tokexpirationStr,
			Duration:       duration,
		},
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
