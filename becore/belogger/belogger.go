package belogger

import (
	"github.com/kataras/golog"
)

// BeLogger is the interface for the logger
type BeLogger struct {
	*golog.Logger
}

// TODO: Customize logger in the future https://github.com/kataras/golog/blob/master/_examples/customize-levels/new-level/main.go
// NewBeLogger returns a new instance of BeLogger
func NewBeLogger() *BeLogger {
	belogger := BeLogger{golog.New()}
	belogger.SetTimeFormat("2006-01-02 15:04:05")
	belogger.SetLevel("debug")

	return &belogger
}
