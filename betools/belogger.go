package betools

import (
	"go.uber.org/zap"
)

// Logger structure
type Logger struct {
	*zap.SugaredLogger
}

// Logger returns a new logger
func NewLogger() *Logger {
	logger, _ := zap.NewProduction()
	return &Logger{logger.Sugar()}
}
