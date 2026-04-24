package config

import (
	"os"
)

// Config holds all service configuration
type Config struct {
	Port   string
	ApiKey string
	IsDev  bool
}

// Load loads configuration from environment variables
func Load() *Config {
	return &Config{
		Port:   getEnv("PORT", "9090"),
		ApiKey: getEnv("API_KEY", "dev-secret-key"),
		IsDev:  getEnv("ENV", "development") == "development",
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
