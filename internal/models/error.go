package models

// Структура для ошибок json
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
