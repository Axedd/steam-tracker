package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Config holds all application configuration values.
type Config struct {
	// DatabaseURL is the full connection string for Postgres
	DatabaseURL string
	// PollInterval is how often the scraper should run (e.g. "30s").
	PollInterval time.Duration
	// HTTPPort is the port the HTTP API server listens on (e.g. "8080").
	HTTPPort string
}

// Load reads configuration from environment variables or a .env file and
// populates a Config struct.
func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	// Set defaults for any values not provided.
	viper.SetDefault("POLL_INTERVAL", "30s")
	viper.SetDefault("HTTP_PORT", "8080")

	// Attempt to read .env; ignore "file not found" errors.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	// Build Config from values.
	cfg := &Config{
		DatabaseURL:  viper.GetString("DATABASE_URL"),
		PollInterval: viper.GetDuration("POLL_INTERVAL"),
		HTTPPort:     viper.GetString("HTTP_PORT"),
	}

	// Validate required fields.
	if cfg.DatabaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is required")
	}

	return cfg, nil
}
