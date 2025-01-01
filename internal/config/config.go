package config

import "github.com/psv2522/rss-aggregator/internal/database"

type ApiConfig struct {
	DB *database.Queries
}