package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"framelog/backend/internal/catalog"
	"framelog/backend/internal/config"
	"framelog/backend/internal/tmdb"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	response := HealthResponse{
		Status:  "ok",
		Service: "framelog-api",
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode health response: %v", err)
	}
}

func main() {
	appConfig, err := config.Load()
	if err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	log.Println("TMDB token loaded from configuration")

	mux := http.NewServeMux()

	tmdbClient := tmdb.NewClient(appConfig.TMDBReadAccessToken)
	catalogHandler := catalog.NewHandler(tmdbClient)

	mux.HandleFunc("GET /api/v1/health/live", healthHandler)
	mux.HandleFunc("GET /api/v1/movies", catalogHandler.ListMovies)

	server := &http.Server{
		Addr:              "127.0.0.1:8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Printf("FrameLog API listening on http://%s", server.Addr)

	if err := server.ListenAndServe(); err != nil &&
		err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}
}
