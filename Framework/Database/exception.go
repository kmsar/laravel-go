package Database

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
)

type ConnectionErrorCode int

const (
	DbDriverDontExist ConnectionErrorCode = iota
	DbConnectionDontExist
)

type DBConnectionException struct {
	error
	Connection string
	Code       ConnectionErrorCode
	fields     Support.Fields
}

func (this DBConnectionException) Fields() Support.Fields {
	this.fields["Code"] = this.Code
	return this.fields
}
