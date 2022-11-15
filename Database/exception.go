package Database

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
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
