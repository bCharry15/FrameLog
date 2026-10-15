package catalog

import (
	"encoding/json"
	"log"
	"net/http"

	"framelog/backend/internal/tmdb"
)

type Handler struct {
	client *tmdb.Client
}

func NewHandler(client *tmdb.Client) *Handler {
	return &Handler{client: client}
}

func (h *Handler) ListMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	page, err := h.client.Popular(r.Context())
	if err != nil {
		log.Printf("Failed to fetch popular movies: %v", err)

		w.WriteHeader(http.StatusBadGateway)

		if err := json.NewEncoder(w).Encode(map[string]any{
			"error": map[string]string{
				"code":    "CATALOG_UNAVAILABLE",
				"message": "Could not retrieve movies.",
			},
		}); err != nil {
			log.Printf("Failed to encode error response: %v", err)
		}

		return
	}

	if err := json.NewEncoder(w).Encode(page); err != nil {
		log.Printf("Failed to encode movie response: %v", err)
	}
}
