package server

import (
	"CalculatorRestApi/config"
	"CalculatorRestApi/internal/handlers"
	"CalculatorRestApi/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func RunEchoServer(cfg *config.Config, logger *logrus.Logger, store *models.SafeStore) {
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/calculate/sum", func(c echo.Context) error {
		return handlers.SumHandler(c, logger, store)
	})

	e.POST("/calculate/multiply", func(c echo.Context) error {
		return handlers.MultiplyHandler(c, logger, store)
	})

	e.Logger.Fatal(e.Start(cfg.Server.Port))
}
