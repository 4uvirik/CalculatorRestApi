package handlers

import (
	"CalculatorRestApi/internal/calc"
	"CalculatorRestApi/internal/models"
	"fmt"
	"github.com/labstack/echo/v4"
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

	// Пробую использовать новый логер
	logger.Infof("Получены числа: %v", req.Numbers)

	result := calc.MultiplyNumbers(req.Numbers)

	// Генерация ключа на основе чисел
	key := fmt.Sprint(req.Numbers)

	// Сохранение в памяти ключ: значение
	store.Save(key, result)

	return c.JSON(http.StatusOK, models.Response{Result: result})
}
