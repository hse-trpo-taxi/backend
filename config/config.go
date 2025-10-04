// Package config provides configuration management for the taxi service backend.
// It handles loading configuration from environment variables with sensible defaults.
package config

import (
	"fmt"
	"log"
	"os"
)

// Config holds the configuration settings for the taxi service.
// It includes server and database connection parameters.
type Config struct {
	// ServerPort specifies the port on which the HTTP server will listen
	ServerPort string
	// DatabaseDSN contains the PostgreSQL connection string
	DatabaseDSN string
}

// LoadConfig creates and returns a new Config instance with values loaded from environment variables.
// If environment variables are not set, it uses sensible defaults.
// SERVER_PORT defaults to "8080" and DATABASE_URL is constructed from individual database parameters.
func LoadConfig() *Config {
	config := &Config{
		ServerPort:  getEnv("SERVER_PORT", "8080"),
		DatabaseDSN: getEnv("DATABASE_URL", getDefaultPostgresURL()),
	}

	log.Printf("Configuration loaded: Port=%s", config.ServerPort)
	return config
}

// getEnv retrieves an environment variable value or returns a default value if not set.
// This is a helper function to simplify configuration loading with fallbacks.
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// getDefaultPostgresURL constructs a PostgreSQL connection string from individual environment variables.
// It uses the following environment variables with their defaults:
// - DB_HOST (default: "localhost")
// - DB_PORT (default: "5432")
// - DB_USER (default: "postgres")
// - DB_PASSWORD (default: "postgres")
// - DB_NAME (default: "taxi")
// - DB_SSLMODE (default: "disable")
func getDefaultPostgresURL() string {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "postgres")
	dbname := getEnv("DB_NAME", "taxi")
	sslmode := getEnv("DB_SSLMODE", "disable")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)
}
