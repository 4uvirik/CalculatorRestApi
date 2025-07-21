package main

import (
	"CalculatorRestApi/config"
	"CalculatorRestApi/internal/logger"
	"CalculatorRestApi/internal/models"
	"CalculatorRestApi/internal/server"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found")
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatalf("CONFIG_PATH not found. Set the correct CONFIG_PATH")
	}

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// инициализируем логгер
	log := logger.Init(&cfg.Logger)

	// инициализируем мапу для сохранения вычеслений
	store := models.NewSafeStore()

	server.RunEchoServer(cfg, log, store)

	log.Infof("Configuration loaded: %v", cfg)
	log.Infof("Server is running on port: %s", cfg.Server.Port)
}
