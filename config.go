package main

import (
	"os"
)

// Config holds the application's configuration.
type Config struct {
	APIEndpoint string
	// Add other configuration parameters here, e.g., APIKey, Timeout, etc.
}

// LoadConfig loads the configuration from environment variables.
func LoadConfig() *Config {
	cfg := &Config{
		APIEndpoint: os.Getenv("API_ENDPOINT"),
		// Load other configuration parameters from environment variables
	}
	return cfg
}
