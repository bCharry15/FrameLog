package catalog

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"framelog/backend/internal/tmdb"
)

func writeError(
	w http.ResponseWriter,
	status int,
	code string,
	message string,
) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(map[string]any{
		"error": map[string]string{
			"code":    code,
			"message": message,
		},
	}); err != nil {
		log.Printf("Failed to encode error response: %v", err)
	}
}

func (h *Handler) GetMovie(w http.ResponseWriter, r *http.Request) {
	movieID, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || movieID <= 0 {
		writeError(
			w,
			http.StatusBadRequest,
			"INVALID_MOVIE_ID",
			"Movie ID must be a positive integer.",
		)
		return
	}

	movie, err := h.client.Details(r.Context(), movieID)

	if errors.Is(err, tmdb.ErrMovieNotFound) {
		writeError(
			w,
			http.StatusNotFound,
			"MOVIE_NOT_FOUND",
			"Movie not found.",
		)
		return
	}

	if err != nil {
		log.Printf("Failed to fetch movie %d: %v", movieID, err)

		writeError(
			w,
			http.StatusBadGateway,
			"CATALOG_UNAVAILABLE",
			"Could not retrieve movie details.",
		)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if err := json.NewEncoder(w).Encode(movie); err != nil {
		log.Printf("Failed to encode movie details: %v", err)
	}
}
