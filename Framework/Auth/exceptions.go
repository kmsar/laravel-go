package Auth

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IExeption"
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
