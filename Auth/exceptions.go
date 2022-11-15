package Auth

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IExeption"
)

type Exception struct {
	IExeption.Exception
}

type GuardException struct {
	IExeption.Exception
}

type UserProviderException struct {
	IExeption.Exception
}
