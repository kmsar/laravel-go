package gate

import (
	"fmt"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IAuth"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IDatabase"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Utils"
	"reflect"
)

type Factory struct {
	abilities map[string]IAuth.GateChecker

	policies map[string]IAuth.Policy

	beforeHooks []IAuth.GateHook
	afterHooks  []IAuth.GateHook
}

var factory *Factory

func Check(user IAuth.Authorizable, ability string, arguments ...interface{}) bool {
	return factory.Check(user, ability, arguments...)
}

func GetFactory() IAuth.GateFactory {
	if factory == nil {
		factory = &Factory{
			abilities:   map[string]IAuth.GateChecker{},
			policies:    map[string]IAuth.Policy{},
			beforeHooks: make([]IAuth.GateHook, 0),
			afterHooks:  make([]IAuth.GateHook, 0),
		}
	}
	return factory
}

func (this *Factory) Check(user IAuth.Authorizable, ability string, arguments ...interface{}) bool {
	this.runBeforeHooks(user, ability, arguments...)
	defer this.runAfterHooks(user, ability, arguments...)

	checker, exists := this.get(ability, arguments...)

	if exists {
		return checker(user, arguments...)
	}
	return false
}

func (this *Factory) Has(ability string) bool {
	_, exists := this.abilities[ability]
	return exists
}

func (this *Factory) runBeforeHooks(user IAuth.Authorizable, ability string, arguments ...interface{}) {
	for _, hook := range this.beforeHooks {
		hook(user, ability, arguments...)
	}
}
func (this *Factory) runAfterHooks(user IAuth.Authorizable, ability string, arguments ...interface{}) {
	for _, hook := range this.afterHooks {
		hook(user, ability, arguments...)
	}
}

func (this *Factory) get(ability string, arguments ...interface{}) (IAuth.GateChecker, bool) {
	checker, exists := this.abilities[ability]

	if exists {
		return checker, exists
	}

	if len(arguments) > 0 {
		var classname string
		if class, isClass := arguments[0].(Support.Class); isClass {
			classname = class.ClassName()
		} else if model, isModel := arguments[0].(IDatabase.Model); isModel {
			classname = model.GetClass().ClassName()
		} else {
			classname = Utils.GetTypeKey(reflect.TypeOf(arguments[0]))
		}
		if this.policies[classname] != nil {
			checker, exists = this.policies[classname][ability]
		}
	}

	return checker, exists
}

func (this *Factory) Define(ability string, callback IAuth.GateChecker) IAuth.GateFactory {
	this.abilities[ability] = callback
	return this
}

func (this *Factory) Policy(class Support.Class, policy IAuth.Policy) IAuth.GateFactory {
	this.policies[class.ClassName()] = policy
	return this
}

func (this *Factory) Before(callable IAuth.GateHook) IAuth.GateFactory {
	this.beforeHooks = append(this.beforeHooks, callable)
	return this
}

func (this *Factory) After(callable IAuth.GateHook) IAuth.GateFactory {
	this.afterHooks = append(this.afterHooks, callable)
	return this
}

func (this *Factory) Abilities() []string {
	var abilities []string

	for ability, _ := range this.abilities {
		abilities = append(abilities, ability)
	}

	for name, policy := range this.policies {
		for ability, _ := range policy {
			abilities = append(abilities, fmt.Sprintf("%s@%s", name, ability))
		}
	}

	return abilities
}
