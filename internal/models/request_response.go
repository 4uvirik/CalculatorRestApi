package models

// Структура для запроса json
type Request struct {
	Token   string `json:"token"`
	Numbers []int  `json:"numbers"`
}

// Структура для ответа json
type Response struct {
	Result int `json:"result"`
}
