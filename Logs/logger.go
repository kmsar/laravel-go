package Logs

import (
	"github.com/apex/log"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IExeption"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/ILogger"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
)

func WithFields(fields Support.Fields) ILogger.Logger {
	return &ApexLogger{Entry: log.WithFields(log.Fields(fields))}
}

func WithError(err error) ILogger.Logger {
	return &ApexLogger{Entry: log.WithError(err)}
}

func WithException(exception IExeption.Exception) ILogger.Logger {
	return &ApexLogger{Entry: log.WithError(exception).WithFields(log.Fields(exception.Fields()))}
}

func Default() ILogger.Logger {
	return &ApexLogger{Entry: log.WithFields(log.Fields(Support.Fields{}))}
}

func WithInterface(value interface{}) ILogger.Logger {
	return WithField("value", value)
}

func WithField(key string, value interface{}) ILogger.Logger {
	return &ApexLogger{Entry: log.WithField(key, value)}
}
