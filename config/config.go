package config

import (
	"fmt"
	"log"
	"os"
)

type Config struct {
	ServerPort   string
	DatabaseDSN  string
}

func LoadConfig() *Config {
	config := &Config{
		ServerPort:  getEnv("SERVER_PORT", "8080"),
		DatabaseDSN: getEnv("DATABASE_URL", getDefaultPostgresURL()),
	}
	
	log.Printf("Configuration loaded: Port=%s", config.ServerPort)
	return config
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

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
