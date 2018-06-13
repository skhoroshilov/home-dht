// Package log implements a simple logging package. It defines interface Logger
// that contains methods for logging messages using log levels. It also hase
// a predefined private type logger that implements Logger interface and writes
// messages to log using Writer interface.
//
// To create new logger backend, implement Writer interface and pass it to NewLogger
// function.
package log

import (
	"os"
)

// Logger provides methods for logging messages using predefined log levels.
// The Fatal functions call os.Exit(1) after writing the log message.
type Logger interface {
	Debug(message string)
	Debugf(message string, args ...interface{})

	Info(message string)
	Infof(message string, args ...interface{})

	Error(message string)
	Errorf(message string, args ...interface{})

	Fatal(message string)
	Fatalf(message string, args ...interface{})
}

// Writer provides methods to write log message to some output.
type Writer interface {
	Write(level string, message string)
	Writef(level string, message string, args ...interface{})
}

type logLevel string

const (
	levelDebug = "DEBUG"
	levelInfo  = "INFO"
	levelError = "ERROR"
	levelFatal = "FATAL"
)

// logger is a simple Logger implementation.
type logger struct {
	writer Writer
}

// NewLogger creates new instance of logger type using specified Writer.
func NewLogger(writer Writer) Logger {
	return &logger{
		writer: writer,
	}
}

func (logger *logger) Debug(message string) {
	logger.writer.Write(levelDebug, message)
}

func (logger *logger) Debugf(message string, args ...interface{}) {
	logger.writer.Writef(levelDebug, message, args...)
}

func (logger *logger) Info(message string) {
	logger.writer.Write(levelInfo, message)
}

func (logger *logger) Infof(message string, args ...interface{}) {
	logger.writer.Writef(levelInfo, message, args...)
}

func (logger *logger) Error(message string) {
	logger.writer.Write(levelError, message)
}

func (logger *logger) Errorf(message string, args ...interface{}) {
	logger.writer.Writef(levelError, message, args...)
}

func (logger *logger) Fatal(message string) {
	logger.writer.Write(levelFatal, message)
	os.Exit(1)
}

func (logger *logger) Fatalf(message string, args ...interface{}) {
	logger.writer.Writef(levelFatal, message, args...)
	os.Exit(1)
}
