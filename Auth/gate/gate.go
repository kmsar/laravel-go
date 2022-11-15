package gate

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IAuth"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IHttp"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Exceptions"
)

type Gate struct {
	factory IAuth.GateFactory
	user    IAuth.Authorizable
}

func NewGate(factory IAuth.GateFactory, user IAuth.Authorizable) IAuth.Gate {
	return &Gate{
		factory: factory,
		user:    user,
	}
}

func (gate *Gate) Allows(ability string, arguments ...interface{}) bool {
	return gate.factory.Check(gate.user, ability, arguments...)
}

func (gate *Gate) Denies(ability string, arguments ...interface{}) bool {
	return !gate.factory.Check(gate.user, ability, arguments...)
}

func (gate *Gate) Check(abilities []string, arguments ...interface{}) bool {
	for _, ability := range abilities {
		if !gate.factory.Check(gate.user, ability, arguments...) {
			return false
		}
	}
	return true
}

func (gate *Gate) Any(abilities []string, arguments ...interface{}) bool {
	for _, ability := range abilities {
		if gate.factory.Check(gate.user, ability, arguments...) {
			return true
		}
	}
	return false
}

func (gate *Gate) Authorize(ability string, arguments ...interface{}) {
	if gate.Denies(ability, arguments...) {
		panic(Exception{
			Exception: Exceptions.New("no operating authority", nil),
			User:      gate.user,
			Ability:   ability,
			Arguments: arguments,
		})
	}
}

func (gate *Gate) Inspect(ability string, arguments ...interface{}) IHttp.IHttpResponse {
	if gate.Allows(ability, arguments...) {
		return &Response{
			Allowed: true,
			Message: "ok",
			Code:    1,
		}
	}
	return &Response{
		Allowed: false,
		Message: "no operating authority",
		Code:    0,
	}
}

func (gate *Gate) ForUser(user IAuth.Authorizable) IAuth.Gate {
	gate.user = user
	return gate
}
