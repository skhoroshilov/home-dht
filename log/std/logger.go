// Package std provides dhtlog.Logger interface implementation using standard
// golang log package.
package std

import (
	"log"

	dhtlog "github.com/skhoroshilov/home-dht/log"
)

// loggers is a private type that implements dhtlog.Logger interface using golang log package.
type logger struct {
}

// NewLogger creates new instance of dhtlog.Logger that writes messages using log package.
func NewLogger() dhtlog.Logger {
	stdlogger := &logger{}

	log.SetFlags(0)
	log.SetOutput(dateFormatter{})

	return dhtlog.NewLogger(stdlogger)
}

func (logger *logger) Write(level string, message string) {
	log.Printf(formatMessage(level, message))
}

func (logger *logger) Writef(level string, message string, args ...interface{}) {
	log.Printf(formatMessage(level, message), args...)
}

func formatMessage(level string, message string) string {
	return level + " " + message
}
