package server

import (
	"CalculatorRestApi/build/calc"
	"fmt"
	"io"
	"log"
	"net/http"
)

func RunServer() {

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
}
