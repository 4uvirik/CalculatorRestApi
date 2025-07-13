package server

import (
	"CalculatorRestApi/internal/calc"
	"CalculatorRestApi/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RunEchoServer() {
	e := echo.New()

	e.POST("/calculate/sum", sumHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

func sumHandler(c echo.Context) error {
	var req models.SumRequest

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

	// !!!На таске 4 заменить логер
	c.Logger().Infof("Получены числа: %v", req.Numbers)

	result := calc.SumNumbers(req.Numbers)

	return c.JSON(http.StatusOK, models.SumResponse{Result: result})
}
