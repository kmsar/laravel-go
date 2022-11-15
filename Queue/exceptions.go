package Queue

import "github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IExeption"

type Exception struct {
	IExeption.Exception
}

type DriverException struct {
	IExeption.Exception
}

type JobException struct {
	IExeption.Exception
}
