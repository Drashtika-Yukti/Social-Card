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
		Port:   getEnv("PORT", "7860"),
		ApiKey: getEnv("API_KEY", "drashtika-social-secret"),
		IsDev:  getEnv("ENV", "development") == "development",
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
