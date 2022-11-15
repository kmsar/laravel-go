package ILogger

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IExeption"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
)

type Logger interface {

	// WithFields adding data
	WithFields(fields Support.Fields) Logger

	// WithField Add data by given key value.
	WithField(key string, value interface{}) Logger

	// WithError add error.
	WithError(err error) Logger

	// WithException Delegate exception management to a custom exception handler.
	WithException(exception IExeption.Exception) Logger

	// Info Adds a log record at the INFO level.
	Info(msg string)

	// Warn Adds a log record at the WARNING level.
	Warn(msg string)

	// Debug Adds a log record at the DEBUG level.
	Debug(msg string)

	// Error Adds a log record at the ERROR level.
	Error(msg string)

	// Fatal Adds a log record at the FATAL level.
	Fatal(msg string)
}
