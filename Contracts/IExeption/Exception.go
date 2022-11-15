package IExeption

import "github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"

type Exception interface {
	error
	Support.FieldsProvider
}

type ExceptionHandler interface {

	// Handle the exception, and return the specified result.
	Handle(exception Exception) interface{}

	// ShouldReport Determine whether to report.
	ShouldReport(exception Exception) bool

	// Report  exception.
	Report(exception Exception)
}
