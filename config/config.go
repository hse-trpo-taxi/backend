package config

import (
	"log"
	"os"
)

type Config struct {
	ServerPort string
	DBPath     string
}

func LoadConfig() *Config {
	config := &Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		DBPath:     getEnv("DB_PATH", "./taxi.db"),
	}
	
	log.Printf("Configuration loaded: Port=%s, DBPath=%s", config.ServerPort, config.DBPath)
	return config
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
