package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config holds application configuration
type Config struct {
	ModelPath        string `json:"model_path"`
	UseGPU           bool   `json:"use_gpu"`
	OutputFormat     string `json:"output_format"`
	Quality          int    `json:"quality"`
	BatchSize        int    `json:"batch_size"`
	AlphaMattingMode string `json:"alpha_matting_mode"`
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		ModelPath:        "assets/models/u2net.onnx",
		UseGPU:           false,
		OutputFormat:     "png",
		Quality:          95,
		BatchSize:        4,
		AlphaMattingMode: "auto",
	}
}

// Load reads configuration from a file
func Load(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()
	
	cfg := &Config{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		return nil, fmt.Errorf("failed to decode config: %w", err)
	}
	
	return cfg, nil
}

// Save writes configuration to a file
func (c *Config) Save(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create config file: %w", err)
	}
	defer file.Close()
	
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(c); err != nil {
		return fmt.Errorf("failed to encode config: %w", err)
	}
	
	return nil
}