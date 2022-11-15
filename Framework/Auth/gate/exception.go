package gate

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IAuth"
	"github.com/kmsar/laravel-go/Framework/Contracts/IExeption"
)

type Exception struct {
	IExeption.Exception

	User      IAuth.Authorizable
	Ability   string
	Arguments []interface{}
}
