package server

import (
	"CalculatorRestApi/config"
	"CalculatorRestApi/internal/handlers"
	"CalculatorRestApi/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func RunEchoServer(cfg *config.Config, logger *logrus.Logger, store *models.SafeStore) {
	e := echo.New()

	e.POST("/calculate/sum", func(c echo.Context) error {
		return handlers.SumHandler(c, logger, store)
	})

	e.Logger.Fatal(e.Start(cfg.Server.Port))
}
