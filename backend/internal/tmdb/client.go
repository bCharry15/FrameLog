package tmdb

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Movie struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
	Overview    string  `json:"overview"`
	PosterPath  *string `json:"poster_path"`
	ReleaseDate string  `json:"release_date"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount   int     `json:"vote_count"`
}

type MoviePage struct {
	Page         int     `json:"page"`
	Results      []Movie `json:"results"`
	TotalPages   int     `json:"total_pages"`
	TotalResults int     `json:"total_results"`
}

type Client struct {
	token      string
	httpClient *http.Client
}

func NewClient(token string) *Client {
	return &Client{
		token: token,
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *Client) Popular(ctx context.Context) (MoviePage, error) {
	const endpoint = "https://api.themoviedb.org/3/movie/popular?language=es-CO&region=CO&page=1"

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		endpoint,
		nil,
	)
	if err != nil {
		return MoviePage{}, fmt.Errorf("create TMDB request: %w", err)
	}

	request.Header.Set("Authorization", "Bearer "+c.token)
	request.Header.Set("Accept", "application/json")

	response, err := c.httpClient.Do(request)
	if err != nil {
		return MoviePage{}, fmt.Errorf("request TMDB movies: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return MoviePage{}, fmt.Errorf(
			"TMDB returned HTTP %d",
			response.StatusCode,
		)
	}

	var page MoviePage

	if err := json.NewDecoder(
		io.LimitReader(response.Body, 4*1024*1024),
	).Decode(&page); err != nil {
		return MoviePage{}, fmt.Errorf("decode TMDB movies: %w", err)
	}

	if page.Results == nil {
		page.Results = []Movie{}
	}

	return page, nil
}
