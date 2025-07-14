package models

// Структура для запроса json
type Request struct {
	Numbers []int `json:"numbers"`
}

// Структура для ответа json
type Response struct {
	Result int `json:"result"`
}
