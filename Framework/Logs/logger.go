package Logs

import (
	"github.com/apex/log"
	"github.com/kmsar/laravel-go/Framework/Contracts/IExeption"
	"github.com/kmsar/laravel-go/Framework/Contracts/ILogger"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
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
