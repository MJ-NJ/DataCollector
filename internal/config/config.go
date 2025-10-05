package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config holds application configuration
type Config struct {
	Target    string            `json:"target"`
	Depth     int               `json:"depth"`
	FileTypes []string          `json:"filetypes"`
	Patterns  map[string]string `json:"patterns,omitempty"`
	Output    string            `json:"output,omitempty"`
}

// Load reads configuration from a JSON file
func Load(filepath string) (*Config, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config JSON: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return &cfg, nil
}

// Validate checks if configuration is valid
func (c *Config) Validate() error {
	if c.Target == "" {
		return fmt.Errorf("target URL cannot be empty")
	}

	if c.Depth < 0 {
		return fmt.Errorf("depth cannot be negative")
	}

	if c.Depth > 10 {
		return fmt.Errorf("depth cannot exceed 10 (prevents infinite crawling)")
	}

	return nil
}

// Save writes configuration to a JSON file
func (c *Config) Save(filepath string) error {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(filepath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

// Default returns default configuration
func Default() *Config {
	return &Config{
		Target:    "https://example.com",
		Depth:     2,
		FileTypes: []string{"pdf", "jpg", "png"},
		Output:    "./downloads",
	}
}
