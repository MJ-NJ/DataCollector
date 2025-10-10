package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
)

func ValidateConfig(config *ScrapeConfig) error {
	validate := validator.New()
	return validate.Struct(config)
}

func LoadConfig(path string) (*ScrapeConfig, error) {
	// Locate file and read file
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read config file: %v", err)
	}

	// Convert to struct
	var config *ScrapeConfig
	if err := json.Unmarshal(file, &config); err != nil {
		return nil, fmt.Errorf("unable to unmarshal data: %v", err)
	}

	if err := ValidateConfig(config); err != nil {
		return nil, fmt.Errorf("error validating config fields: %v", err)
	}
	// Return
	return config, nil
}
