package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Server struct {
		Port string `json:"port"`
	}
	Logger struct {
		Level  string `json:"level"`
		Format string `json:"format"`
	}
}

func LoadConfig(path string) (*Config, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	var cfg Config
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
