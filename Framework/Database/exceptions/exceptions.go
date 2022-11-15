package exceptions

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IExeption"
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
