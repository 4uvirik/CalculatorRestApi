package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type ServerConfig struct {
	Port string `json:"port"`
}
type LoggerConfig struct {
	Level  string `json:"level"`
	Format string `json:"format"`
}

type Config struct {
	Server ServerConfig `json:"server"`
	Logger LoggerConfig `json:"logger"`
}

func LoadConfig(path string) (*Config, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("file config not open %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	var cfg Config
	if err := decoder.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("decode config failed %w", err)
	}
	return &cfg, nil
}
