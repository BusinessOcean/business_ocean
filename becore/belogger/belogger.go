package belogger

import (
	"github.com/kataras/golog"
)

// BeLogger is the interface for the logger

type BeLogger struct {
	*golog.Logger
}

// NewBeLogger returns a new instance of BeLogger
func NewBeLogger() BeLogger {
	logger := golog.New()

	// errorAttrs := golog.Levels[golog.ErrorLevel]

	// // Change a log level's text.
	// customColorCode := 156
	// errorAttrs.SetText("custom text", customColorCode)

	// // Get (rich) text per log level.
	// enableColors := true
	// errorRichText := errorAttrs.Text(enableColors)
	return BeLogger{Logger: logger}
}
