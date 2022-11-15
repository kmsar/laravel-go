package exceptions

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IExeption"
)

type TransactionException struct {
	IExeption.Exception
}

type RollbackException struct {
	IExeption.Exception
}

type BeginException struct {
	IExeption.Exception
}
