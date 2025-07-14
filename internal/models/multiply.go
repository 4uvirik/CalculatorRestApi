package models

// Структура для запроса json
type MultiplyRequest struct {
	Numbers []int `json:"numbers"`
}

// Структура для ответа json
type MultiplyResponse struct {
	Result int `json:"result"`
}
