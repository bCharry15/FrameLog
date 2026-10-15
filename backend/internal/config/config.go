package config

import (
	"errors"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	TMDBReadAccessToken string
}

func Load() (Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return Config{}, errors.New("could not load .env configuration")
		}
	}

	token := strings.TrimSpace(os.Getenv("TMDB_READ_ACCESS_TOKEN"))

	if token == "" {
		return Config{}, errors.New("TMDB_READ_ACCESS_TOKEN is required")
	}

	return Config{
		TMDBReadAccessToken: token,
	}, nil
}
