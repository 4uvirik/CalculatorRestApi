package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfig(t *testing.T) {

	type testCase struct {
		name           string
		fileContents   string
		file           bool
		expectError    bool
		expectedConfig *Config
	}

	tests := []testCase{
		{
			name:         "OK",
			fileContents: `{"server": {"port": ":8080"}, "logger": {"level": "DEBUG", "format": "json"}}`,
			file:         true,
			expectError:  false,
			expectedConfig: &Config{
				Server: ServerConfig{Port: ":8080"},
				Logger: LoggerConfig{Level: "DEBUG", Format: "json"},
			},
		},
		{
			name:        "Нет файла",
			file:        false,
			expectError: true,
		},
		{
			name:         "Не правильный json",
			fileContents: `{"port": ":8080"`, // нет закрывающей скобки
			file:         true,
			expectError:  true,
		},
		{
			name:         "Пустой файл",
			fileContents: ``,
			file:         true,
			expectError:  true,
		},
		{
			name:         "пустое значение логер конфига",
			fileContents: `{"server": {"port": ":8080"}}`,
			file:         true,
			expectError:  false,
			expectedConfig: &Config{
				Server: ServerConfig{Port: ":8080"},
				Logger: LoggerConfig{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var path string // переменная в которую будет записан путь к файлу

			if tt.file {
				tmpFile, err := os.CreateTemp("", "config-*.json")
				assert.NoError(t, err)
				defer os.Remove(tmpFile.Name())
				path = tmpFile.Name()

				_, err = tmpFile.WriteString(tt.fileContents)
				assert.NoError(t, err)
				tmpFile.Close()
			} else {
				path = filepath.Join(os.TempDir(), "notexist.json")
			}

			cfg, err := LoadConfig(path)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedConfig, cfg)
			}
		})
	}
}
