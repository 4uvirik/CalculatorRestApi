package server

import (
	"CalculatorRestApi/build/calc"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func RunEchoServer() {
	e := echo.New()

	e.POST("/json", handlePost)

	e.Logger.Fatal(e.Start(":8080"))
}

// Структура для POST запроса json
type Message struct {
	Text string `json:"text"`
}

func handlePost(c echo.Context) error {
	var msg Message

	if err := c.Bind(&msg); err != nil {
		return c.String(http.StatusBadRequest, "Ошибка запроса")
	}

	fmt.Println("Получены числа:", msg.Text)

	calc, err := calc.ResultCalc(msg.Text)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "Сумма чисел равна:\n"+strconv.Itoa(calc))
}

/*func RunServer() {

	registerRoutes()
	fmt.Println("Сервер запущен")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Ошибка сервера: %v", err)
	}
}

func handlePost(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Введите числа через пробел, я посчитаю их сумму\n")
	switch req.Method {
	case http.MethodPost:
		body, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "Ошибка чтения тела запроса", http.StatusBadRequest)
			return
		}
		calc, err := calc.ResultCalc(string(body))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Сумма чисел равна: %d\n", calc)
	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

func registerRoutes() {
	http.HandleFunc("/", handlePost)
}*/
