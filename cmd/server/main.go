package main

import (
	"fmt"
	"log"
	"net/http"

	"social-forge/internal/api"
	"social-forge/internal/config"
	_ "embed"
)

//go:embed public/index.html
var indexHTML []byte

func main() {
	cfg := config.Load()

	// Initialize handlers
	mux := http.NewServeMux()

	// Protected endpoints
	protectedHandler := api.Chain(http.HandlerFunc(api.GenerateCardHandler), api.LoggerMiddleware(), api.RateLimitMiddleware(1, 3), api.AuthMiddleware(cfg))
	mux.Handle("/generate", protectedHandler)

	// Public endpoints
	mux.HandleFunc("/health", api.HealthCheckHandler)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.Write(indexHTML)
	})

	fmt.Printf("🚀 Drashtika SocialForge starting on :%s (Mode: %s)\n", cfg.Port, getMode(cfg))
	fmt.Printf("🔑 Configured API Key: %s\n", cfg.ApiKey)

	log.Fatal(http.ListenAndServe(":"+cfg.Port, mux))
}

func getMode(cfg *config.Config) string {
	if cfg.IsDev {
		return "development"
	}
	return "production"
}
