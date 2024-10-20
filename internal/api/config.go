package api

import (
	"github.com/thesujai/aggregator/internal/api/handlers"
	"github.com/thesujai/aggregator/internal/database"
)

type Config struct {
	*handlers.Config
}

func NewConfig(db *database.Queries) *Config {
	return &Config{
		Config: &handlers.Config{
			DB: db,
		},
	}
}
