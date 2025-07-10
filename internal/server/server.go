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

// Структура для POST запроса json
type Message struct {
	Text string `json:"text"`
}

// Структура для ответа json
type Result struct {
	Result int `json:"result"`
}

func postSumHandler(c echo.Context) error {
	var msg Message

	if err := c.Bind(&msg); err != nil {
		return c.String(http.StatusBadRequest, "Ошибка запроса")
	}

	c.Logger().Infof("Получены числа: %s", msg.Text)

	result, err := calc.SumForPostString(msg.Text)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, Result{Result: result})
}
