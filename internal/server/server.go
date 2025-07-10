package server

import (
	"CalculatorRestApi/internal/calc"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RunEchoServer() {
	e := echo.New()

	e.POST("/json", postSumHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

// Структура для запроса json
type Input struct {
	Numbers []int `json:"numbers"`
}

// Структура для ответа json
type Result struct {
	Result int `json:"result"`
}

func postSumHandler(c echo.Context) error {
	var inp Input

	if err := c.Bind(&inp); err != nil {
		return c.String(http.StatusBadRequest, "Ошибка запроса")
	}

	if len(inp.Numbers) == 0 {
		return c.String(http.StatusBadRequest, "Нет вводных данных")
	}

	// !!!На таске 4 заменить логер
	c.Logger().Infof("Получены числа: %v", inp.Numbers)

	result := calc.SumNumbers(inp.Numbers)

	return c.JSON(http.StatusOK, Result{Result: result})
}
