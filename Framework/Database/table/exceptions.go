package table

import "github.com/kmsar/laravel-go/Framework/Contracts/IExeption"

type CreateException struct {
	IExeption.Exception
}

type InsertException struct {
	IExeption.Exception
}

type UpdateException struct {
	IExeption.Exception
}

type DeleteException struct {
	IExeption.Exception
}

type SelectException struct {
	IExeption.Exception
}

type NotFoundException struct {
	IExeption.Exception
}
