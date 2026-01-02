package logger

import (
	"github.com/cristiano-pacheco/go-starter-kit/internal/shared/modules/config"
	"github.com/cristiano-pacheco/go-starter-kit/pkg/logger"
)

type Logger interface {
	logger.Logger
}

func New(cfg config.Config) Logger {
	logLevel := logger.MustLogLevel(cfg.Log.LogLevel)
	logConfig := logger.Config{
		LogLevel: logLevel,
	}
	return logger.New(logConfig)
}
