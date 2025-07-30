package server

import (
	"CalculatorRestApi/config"
	"CalculatorRestApi/internal/handlers"
	"CalculatorRestApi/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "CalculatorRestApi/docs"
	docs "CalculatorRestApi/docs"
)

func SetupRouter(cfg *config.Config, logger *logrus.Logger, store *models.SafeStore) *echo.Echo {
	e := echo.New()

	docs.SwaggerInfo.Host = "localhost" + cfg.Server.Port

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/calculate/sum", func(c echo.Context) error {
		return handlers.SumHandler(c, logger, store)
	})

	e.POST("/calculate/multiply", func(c echo.Context) error {
		return handlers.MultiplyHandler(c, logger, store)
	})

	return e
}

func RunEchoServer(cfg *config.Config, logger *logrus.Logger, store *models.SafeStore) {

	e := SetupRouter(cfg, logger, store)
	e.Logger.Fatal(e.Start(cfg.Server.Port))
}
