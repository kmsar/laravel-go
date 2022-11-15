package Queue

import "github.com/kmsar/laravel-go/Framework/Contracts/IExeption"

type Exception struct {
	IExeption.Exception
}

type DriverException struct {
	IExeption.Exception
}

type JobException struct {
	IExeption.Exception
}
