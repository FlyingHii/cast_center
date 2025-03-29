package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Config holds the application's configuration.
type Config struct {
	APIEndpoint string `yaml:"api_endpoint"`
	APIKey      string `yaml:"api_key"`
	Timeout     int    `yaml:"timeout"`
	ConfigPath  string // Path to the config file
}

// loadConfigFromYAML loads the configuration from a YAML file.
func loadConfigFromYAML(configPath string) (*Config, error) {
	// Expand the path to handle relative paths and user home directory
	expandedPath, err := expandPath(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to expand config path: %w", err)
	}

	// Read the file
	file, err := os.ReadFile(expandedPath)
	if err != nil {
		// If the file doesn't exist, return nil and no error.
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Unmarshal the YAML data
	var cfg Config
	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %w", err)
	}

	// Store the config path
	cfg.ConfigPath = expandedPath

	return &cfg, nil
}

// expandPath expands the path, handling relative paths and the user's home directory.
func expandPath(path string) (string, error) {
	if path == "" {
		return "", nil
	}

	// Expand the user's home directory (e.g., "~/.config/app/config.yaml")
	if path[:2] == "~/" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("failed to get user home directory: %w", err)
		}
		path = filepath.Join(homeDir, path[2:])
	}

	// Resolve the path to an absolute path
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path: %w", err)
	}

	return absPath, nil
}

// LoadConfig loads the configuration from environment variables or a YAML file.
func LoadConfig(configPath string) *Config {
	// First, try to load from YAML
	cfg, err := loadConfigFromYAML(configPath)
	if err != nil {
		// Log the error, but continue with environment variables
		fmt.Printf("Error loading config from YAML: %v\n", err)
	}

	if cfg == nil {
		cfg = &Config{
			APIEndpoint: os.Getenv("API_ENDPOINT"),
			// Load other configuration parameters from environment variables
		}
		return cfg
	}

	// If YAML was loaded successfully, override environment variables if they are set.
	if envEndpoint := os.Getenv("API_ENDPOINT"); envEndpoint != "" {
		cfg.APIEndpoint = envEndpoint
	}

	return cfg
}
