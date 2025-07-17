package main

import (
	"CalculatorRestApi/config"
	"CalculatorRestApi/internal/logger"
	"CalculatorRestApi/internal/models"
	"CalculatorRestApi/internal/server"
	"fmt"
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

	fmt.Println("Server port:", cfg.Server.Port)
	fmt.Println("Logger level:", cfg.Logger.Level, "Logger format", cfg.Logger.Format)

	server.RunEchoServer(cfg, log, store)

}
