package main

import (
	"CalculatorRestApi/config"
	"CalculatorRestApi/internal/logger"
	"CalculatorRestApi/internal/models"
	"CalculatorRestApi/internal/server"
	"log"
)

func main() {

	cfg, err := config.LoadConfig("config/config.json")
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
