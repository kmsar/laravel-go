package Logs

import (
	"github.com/apex/log"
	"github.com/kmsar/laravel-go/Framework/Contracts/IExeption"
	"github.com/kmsar/laravel-go/Framework/Contracts/ILogger"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
)

var (
	Debug = false
)

type ApexLogger struct {
	Entry *log.Entry
}

func (this *ApexLogger) WithFields(m Support.Fields) ILogger.Logger {
	if this == nil || this.Entry == nil {
		this = &ApexLogger{
			Entry: log.WithFields(log.Fields(m)),
		}
	}

	this.Entry = this.Entry.WithFields(log.Fields(m))

	return this
}

func (this *ApexLogger) WithField(key string, value interface{}) ILogger.Logger {
	if this == nil || this.Entry == nil {
		this = &ApexLogger{
			Entry: log.WithField(key, value),
		}
	}

	this.Entry = this.Entry.WithField(key, value)

	return this
}

func (this *ApexLogger) WithError(err error) ILogger.Logger {
	if this == nil || this.Entry == nil {
		this = &ApexLogger{
			Entry: log.WithError(err),
		}
	}

	this.Entry = this.Entry.WithError(err)

	return this
}

func (this *ApexLogger) WithException(err IExeption.Exception) ILogger.Logger {
	if this == nil || this.Entry == nil {
		this = &ApexLogger{
			Entry: log.WithError(err).WithFields(log.Fields(err.Fields())),
		}
	}

	this.Entry = this.Entry.WithError(err).WithFields(log.Fields(err.Fields()))

	return this
}

func (this ApexLogger) Info(msg string) {
	this.Entry.Info(msg)
}

func (this ApexLogger) Warn(msg string) {
	this.Entry.Warn(msg)
}

func (this ApexLogger) Debug(msg string) {
	if Debug {
		this.Entry.Debug(msg)
	}
}

func (this ApexLogger) Error(msg string) {
	this.Entry.Error(msg)
}

func (this ApexLogger) Fatal(msg string) {
	this.Entry.Fatal(msg)
}
