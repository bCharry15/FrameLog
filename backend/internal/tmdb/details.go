package tmdb

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var ErrMovieNotFound = errors.New("movie not found")

type Genre struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type MovieDetails struct {
	Movie
	OriginalTitle string  `json:"original_title"`
	Tagline       string  `json:"tagline"`
	Runtime       *int    `json:"runtime"`
	Genres        []Genre `json:"genres"`
	BackdropPath  *string `json:"backdrop_path"`
}

func (c *Client) Details(
	ctx context.Context,
	movieID int64,
) (MovieDetails, error) {
	endpoint := fmt.Sprintf(
		"https://api.themoviedb.org/3/movie/%d?language=es-CO",
		movieID,
	)

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		endpoint,
		nil,
	)
	if err != nil {
		return MovieDetails{}, fmt.Errorf("create movie request: %w", err)
	}

	request.Header.Set("Authorization", "Bearer "+c.token)
	request.Header.Set("Accept", "application/json")

	response, err := c.httpClient.Do(request)
	if err != nil {
		return MovieDetails{}, fmt.Errorf("request movie details: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusNotFound {
		return MovieDetails{}, ErrMovieNotFound
	}

	if response.StatusCode != http.StatusOK {
		return MovieDetails{}, fmt.Errorf(
			"TMDB returned HTTP %d",
			response.StatusCode,
		)
	}

	var movie MovieDetails

	if err := json.NewDecoder(
		io.LimitReader(response.Body, 4*1024*1024),
	).Decode(&movie); err != nil {
		return MovieDetails{}, fmt.Errorf("decode movie details: %w", err)
	}

	if movie.Genres == nil {
		movie.Genres = []Genre{}
	}

	return movie, nil
}
