package api

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"sync"

	"social-forge/internal/generator"
	"social-forge/internal/templates"
)

var (
	// Simple in-memory cache for generated images
	imageCache sync.Map
)

type CardRequest struct {
	Markdown    string                   `json:"markdown"`
	Author      string                   `json:"author"`
	Title       string                   `json:"title"`
	Theme       string                   `json:"theme,omitempty"`
	AccentColor string                   `json:"accent_color,omitempty"`
	AvatarURL   string                   `json:"avatar_url,omitempty"`
	Overrides   templates.StyleOverrides `json:"style_overrides,omitempty"`
}

func GenerateCardHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CardRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// -- CACHE CHECK --
	reqBytes, _ := json.Marshal(req)
	hash := fmt.Sprintf("%x", sha256.Sum256(reqBytes))

	if cachedVal, ok := imageCache.Load(hash); ok {
		w.Header().Set("Content-Type", "image/png")
		w.Header().Set("X-Cache", "HIT") // Signals to user it came from memory
		w.Write(cachedVal.([]byte))
		return
	}

	// 1. Convert Markdown to Sanitized HTML
	htmlContent, err := generator.MarkdownToHTML(req.Markdown)
	if err != nil {
		http.Error(w, "Markdown processing failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 2. Render Template
	t, err := template.New("card").Parse(templates.SocialCard)
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var res bytes.Buffer
	err = t.Execute(&res, templates.CardData{
		Content:     htmlContent,
		Author:      req.Author,
		Title:       req.Title,
		Theme:       req.Theme,
		AccentColor: req.AccentColor,
		AvatarURL:   req.AvatarURL,
		Overrides:   req.Overrides,
	})
	if err != nil {
		http.Error(w, "Template execution failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 3. Generate PNG
	imgBuf, err := generator.Screenshot(res.String())
	if err != nil {
		http.Error(w, "Image generation failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// -- SAVE TO CACHE --
	imageCache.Store(hash, imgBuf)

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("X-Cache", "MISS")
	w.Write(imgBuf)
}

// HealthCheckHandler returns a 200 OK for infrastructure monitoring
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok", "service":"drashtika-social-forge"}`))
}
