package server

import (
	"CalculatorRestApi/config"
	"CalculatorRestApi/internal/calc"
	"CalculatorRestApi/internal/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	"net/http"
)

func RunEchoServer(cfg *config.Config, logger *logrus.Logger, store *models.SafeStore) {
	e := echo.New()

	e.POST("/calculate/sum", func(c echo.Context) error {
		return sumHandler(c, store)
	})

	e.Logger.Fatal(e.Start(cfg.Server.Port))
}

func sumHandler(c echo.Context, store *models.SafeStore) error {
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

	// Генерация ключа на основе чисел
	key := fmt.Sprint(req.Numbers)

	// Сохранение в памяти ключ: значение
	store.Save(key, result)

	return c.JSON(http.StatusOK, models.SumResponse{Result: result})
}
