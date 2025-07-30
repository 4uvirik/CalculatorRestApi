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

func TestSumHandler(t *testing.T) {

	type testCase struct {
		name   string
		body   models.Request
		code   int
		result any
	}

	tests := []testCase{
		{
			name: "Правильный запрос",
			body: models.Request{
				Numbers: []int{1, 3, 5},
				Token:   "user123",
			},
			code:   http.StatusOK,
			result: models.Response{Result: 9},
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
			bodyBytes, _ := json.Marshal(tt.body)

			req := httptest.NewRequest(http.MethodPost, "/calculate/sum", bytes.NewReader(bodyBytes))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, rec)

			logger := logrus.New()
			store := models.NewSafeStore()

			err := SumHandler(c, logger, store)

			assert.NoError(t, err)
			assert.Equal(t, tt.code, rec.Code)

			responseBody, _ := io.ReadAll(rec.Body)

			if tt.code == http.StatusOK {
				var actual models.Response
				_ = json.Unmarshal(responseBody, &actual)
				assert.Equal(t, tt.result, actual)
			} else {
				var actual models.ErrorResponse
				_ = json.Unmarshal(responseBody, &actual)
				assert.Equal(t, tt.result, actual)
			}
		})
	}
}
