package Container

import "github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IExeption"

type ContainerExceptionInterface struct {
	IExeption.Exception
}

type NotFoundExceptionInterface struct {
	IExeption.Exception
}
