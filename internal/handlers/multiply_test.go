package handlers

import (
	"CalculatorRestApi/internal/models"
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMultiplyHandler(t *testing.T) {

	type testCase struct { // структура для тест кейсов
		name   string
		body   models.Request
		code   int
		result any
	}

	tests := []testCase{ // Сами кейсы, пока делаю 3
		{
			name: "Правильный запрос",
			body: models.Request{
				Numbers: []int{1, 3, 5},
				Token:   "user123",
			},
			code:   http.StatusOK,
			result: models.Response{Result: 15},
		},
		{
			name: "Нету чисел",
			body: models.Request{
				Numbers: []int{},
				Token:   "user123",
			},
			code:   http.StatusBadRequest,
			result: models.ErrorResponse{Message: "Нет полученных данных"},
		},
		{
			name: "Нету токена",
			body: models.Request{
				Numbers: []int{1, 3, 5},
				Token:   "",
			},
			code:   http.StatusBadRequest,
			result: models.ErrorResponse{Message: "Токен не обнаружен"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bodyBytes, _ := json.Marshal(tt.body) // преобразуем model.request в слайс байт

			req := httptest.NewRequest(http.MethodPost, "/calculate/multiply", bytes.NewReader(bodyBytes))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON) // обозначаем ContentType как json
			rec := httptest.NewRecorder()                                    // сюда записывается то что вернет хэндлер

			e := echo.New()
			c := e.NewContext(req, rec) // оборачиваем запрос и переменную куда записывается в контекст что бы передать его логгеру

			logger := logrus.New() // задаем логгер ну и далее хранилище т.к. хэндлер принимает эти значения
			store := models.NewSafeStore()

			err := MultiplyHandler(c, logger, store) // запускаем сам хэндлер.
			assert.NoError(t, err)                   // тут типа нужно убедиться что ошибка не выскочет
			assert.Equal(t, tt.code, rec.Code)       // сравниваем заданный и полученный коды

			responseBody, _ := io.ReadAll(rec.Body) // читаем полученный Боди в переменную

			if tt.code == http.StatusOK { // Сверяем статусы
				var actual models.Response                // тут переменная будет типа респонс с результатом
				_ = json.Unmarshal(responseBody, &actual) // обратно преобразовываем слайс байт в структуру
				assert.Equal(t, tt.result, actual)        // сравниваем то что хотели получить и то что получили
			} else {
				var actual models.ErrorResponse // тут переменная будет типа errorResponse с сообщением ошибки
				_ = json.Unmarshal(responseBody, &actual)
				assert.Equal(t, tt.result, actual)
			}
		})
	}
}
