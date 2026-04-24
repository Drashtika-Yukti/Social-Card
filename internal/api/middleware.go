package api

import (
	"log"
	"net/http"
	"time"

	"golang.org/x/time/rate"
	"social-forge/internal/config"
)

// Middleware is a function that wraps an http.Handler
type Middleware func(http.Handler) http.Handler

// AuthMiddleware checks for a valid API Key in the X-API-KEY header
func AuthMiddleware(cfg *config.Config) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get("X-API-KEY")
			if apiKey != cfg.ApiKey {
				http.Error(w, "Unauthorized: Invalid or missing API Key", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// LoggerMiddleware logs basic info about every request
func LoggerMiddleware() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
		})
	}
}

// RateLimitMiddleware applies a simple global rate limit
func RateLimitMiddleware(limit rate.Limit, burst int) Middleware {
	limiter := rate.NewLimiter(limit, burst)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !limiter.Allow() {
				http.Error(w, "Too Many Requests: Rate limit exceeded", http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// Chain applies a sequence of middlewares to a handler
func Chain(h http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}
