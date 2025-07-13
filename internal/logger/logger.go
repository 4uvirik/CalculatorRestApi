package logger

import (
	"CalculatorRestApi/config"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

func Init(cfg *config.LoggerConfig) *logrus.Logger {

	log := logrus.New()

	// Формат логирвоания
	switch strings.ToLower(cfg.Format) {
	case "json":
		log.SetFormatter(&logrus.JSONFormatter{})
	default:
		log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	}

	// Уровень логирования
	level, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	log.SetLevel(level)

	// Все логи отправляем в терминал
	log.SetOutput(os.Stdout)

	return log
}
