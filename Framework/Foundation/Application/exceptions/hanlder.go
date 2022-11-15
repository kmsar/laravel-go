package exceptions

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IExeption"
	"github.com/kmsar/laravel-go/Framework/Logs"
	"github.com/kmsar/laravel-go/Framework/Support/Utils"

	"reflect"
)

type DefaultExceptionHandler struct {
	dontReportExceptions []reflect.Type
}

func NewDefaultHandler(dontReportExceptions []IExeption.Exception) DefaultExceptionHandler {
	return DefaultExceptionHandler{Utils.ConvertToTypes(dontReportExceptions)}
}

func (handler DefaultExceptionHandler) Handle(exception IExeption.Exception) (result interface{}) {
	Logs.WithException(exception).
		WithField("exception", reflect.TypeOf(exception).String()).
		Error("DefaultExceptionHandler")

	return
}

func (handler DefaultExceptionHandler) Report(exception IExeption.Exception) {
}

func (handler DefaultExceptionHandler) ShouldReport(exception IExeption.Exception) bool {
	return !Utils.IsInstanceIn(exception, handler.dontReportExceptions...)
}
