package gate

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IAuth"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IExeption"
)

type Exception struct {
	IExeption.Exception

	User IAuth.Authorizable
	Ability   string
	Arguments []interface{}
}
