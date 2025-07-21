package handlers

import (
	"CalculatorRestApi/internal/calc"
	"CalculatorRestApi/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

// @Summary      Умножение чисел
// @Description  Умножает числа полученные через POST запрос и возвращает результат
// @Tags		 math
// @Accept		 json
// @Produce 	 json
// @Param 		 multiply body models.Request true "данные для умножения"
// @Success		 200 {object} models.Response
// @Failure 	 400 {object} models.ErrorResponse
// @Router 		 /calculate/multiply [post]
func MultiplyHandler(c echo.Context, logger *logrus.Logger, store *models.SafeStore) error {
	var req models.Request

	if err := c.Bind(&req); err != nil {
		logger.Warn("JSON binding error")
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: "Ошибка запроса",
		})
	}

	if len(req.Numbers) == 0 {
		logger.Warn("Empty list of numbers")
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: "Нет полученных данных",
		})
	}

	if len(req.Token) == 0 {
		logger.Warn("Token not found")
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: "Токен не обнаружен",
		})
	}

	// Пробую использовать новый логер
	logger.WithField("numbers", req.Numbers).Info("Numbers received")

	result := calc.MultiplyNumbers(req.Numbers)

	// Генерация ключа на основе чисел
	key := req.Token

	// Сохранение в памяти ключ: значение
	store.Save(key, result)
	logger.WithFields(logrus.Fields{
		"result": result,
		"token":  key,
	}).Info("Stored result")

	return c.JSON(http.StatusOK, models.Response{Result: result})
}
