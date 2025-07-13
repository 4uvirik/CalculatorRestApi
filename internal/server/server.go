package server

import (
	"CalculatorRestApi/config"
	"CalculatorRestApi/internal/calc"
	"CalculatorRestApi/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	"net/http"
)

func RunEchoServer(cfg *config.Config, logger *logrus.Logger) {
	e := echo.New()

	e.POST("/calculate/sum", sumHandler)

	e.Logger.Fatal(e.Start(cfg.Server.Port))
}

func sumHandler(c echo.Context) error {
	var req models.SumRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    400,
			Message: "Ошибка запроса",
		})
	}

	if len(req.Numbers) == 0 {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    400,
			Message: "Нет полученных данных",
		})
	}

	// Пробую использовать новый логер
	log.Infof("Получены числа: %v", req.Numbers)

	result := calc.SumNumbers(req.Numbers)

	return c.JSON(http.StatusOK, models.SumResponse{Result: result})
}
