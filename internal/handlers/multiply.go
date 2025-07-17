package handlers

import (
	"CalculatorRestApi/internal/calc"
	"CalculatorRestApi/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	"net/http"
)

func MultiplyHandler(c echo.Context, logger *logrus.Logger, store *models.SafeStore) error {
	var req models.Request

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: "Ошибка запроса",
		})
	}

	if len(req.Numbers) == 0 {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: "Нет полученных данных",
		})
	}

	if len(req.Token) == 0 {
		log.Warn("Token not found")
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: "Токен не обнаружен",
		})
	}

	// Пробую использовать новый логер
	logger.Infof("Numbers received: %v", req.Numbers)

	result := calc.MultiplyNumbers(req.Numbers)

	// Генерация ключа на основе чисел
	key := req.Token

	// Сохранение в памяти ключ: значение
	store.Save(key, result)
	logger.Infof("Stored result: %d for token: %s", result, key)

	return c.JSON(http.StatusOK, models.Response{Result: result})
}
