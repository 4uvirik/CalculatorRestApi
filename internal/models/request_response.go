package models

// Request - структура запроса
type Request struct {
	Token   string `json:"token"`
	Numbers []int  `json:"numbers"`
}

// Response - структура успешного ответа
type Response struct {
	Result int `json:"result"`
}
