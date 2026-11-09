// Package config contains configuration logic for the application
package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Config represents a configuration for the application
type Config struct {
	// BaseURL represents the base URL of the API
	BaseURL string `json:"base_url"`
	// Timeout represents the timeout for API requests
	Timeout time.Duration `json:"timeout"`
	// DarkMode represents whether the application is in dark mode
	DarkMode bool `json:"dark_mode"`
}

// defaultConfig returns the default configuration
func defaultConfig() *Config {
	return &Config{
		DarkMode: false,
		BaseURL:  "https://devapi.wodify.com/v1",
	}
}

// path returns the path to the configuration file
func path() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("reading config directory: %v", err)
	}
	return fmt.Sprintf("%s/Wodify/config.json", dir), nil
}

// Read reads the application's configuration
func Read() (*Config, error) {
	filePath, err := path()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(filePath)
	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("reading config file: %v", err)
	}
	if os.IsNotExist(err) {
		return defaultConfig(), nil
	}

	config := &Config{}
	err = json.Unmarshal(data, config)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling config file: %v", err)
	}

	return config, nil
}

// Write writes the configuration to the configuration file.
func (c *Config) Write() error {
	filePath, err := path()
	if err != nil {
		return err
	}
	err = os.MkdirAll(filepath.Dir(filePath), 0o755)
	if err != nil {
		return fmt.Errorf("creating config directory: %v", err)
	}

	data, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return fmt.Errorf("marshalling config: %v", err)
	}
	err = os.WriteFile(filePath, data, 0o644)
	if err != nil {
		return fmt.Errorf("writing config file: %v", err)
	}

	return nil
}
