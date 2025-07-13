package models

// Структура для запроса json
type SumRequest struct {
	Numbers []int `json:"numbers"`
}

// Структура для ответа json
type SumResponse struct {
	Result int `json:"result"`
}
