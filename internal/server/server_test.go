package server

import (
	"CalculatorRestApi/config"
	"CalculatorRestApi/docs"
	"CalculatorRestApi/internal/models"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetupRouter(t *testing.T) {

	type testCase struct {
		name           string
		expectedRoutes []string
		expectedHost   string
		cfg            *config.Config
	}

	tests := []testCase{
		{
			name: "router OK",
			expectedRoutes: []string{
				"/swagger/*",
				"/calculate/sum",
				"/calculate/multiply",
			},
			expectedHost: "localhost:8080",
			cfg:          &config.Config{Server: config.ServerConfig{Port: ":8080"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			logger := logrus.New()
			store := models.NewSafeStore()
			e := SetupRouter(tt.cfg, logger, store)

			routes := make([]string, 0, len(e.Routes()))
			for _, route := range e.Routes() {
				routes = append(routes, route.Path)
			}
			for _, expected := range tt.expectedRoutes {
				assert.Contains(t, routes, expected)
			}

			assert.Equal(t, tt.expectedHost, docs.SwaggerInfo.Host)

		})
	}

}
