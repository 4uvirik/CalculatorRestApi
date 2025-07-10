package main

import (
	"CalculatorRestApi/config"
	"CalculatorRestApi/internal/server"
	"fmt"
	"log"
)

func main() {

	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatalf("config not loaded ")
	}

	fmt.Println("Порт:", cfg.Server)
	fmt.Println("Уровень логирования:", cfg.Logger.Level, "Формат лоигрвоания", cfg.Logger.Format)

	server.RunEchoServer()

}
