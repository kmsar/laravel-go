package Container

import "github.com/kmsar/laravel-go/Framework/Contracts/IExeption"

type ContainerExceptionInterface struct {
	IExeption.Exception
}

type NotFoundExceptionInterface struct {
	IExeption.Exception
}
