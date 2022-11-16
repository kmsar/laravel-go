package gate

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IAuth"
	"github.com/kmsar/laravel-go/Framework/Contracts/IFoundation"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
)

type ServiceProvider struct {
	app      IFoundation.IApplication
	Policies map[Support.Class]IAuth.Policy
}

func (this *ServiceProvider) Boot(application IFoundation.IApplication) {
	//TODO implement me
	panic("implement me")
}

func (this *ServiceProvider) Register(application IFoundation.IApplication) {
	this.app = application
	application.NamedSingleton("gate.factory", func() IAuth.GateFactory {
		return GetFactory()
	})
	application.NamedBind("gate", func(factory IAuth.GateFactory, guard IAuth.Guard, ctx Support.Context) IAuth.Gate {
		instance, exists := ctx.Get("access.gate").(IAuth.Gate)
		if exists {
			return instance
		}
		user, _ := guard.User().(IAuth.Authorizable)
		instance = NewGate(factory, user)
		ctx.Set("access.gate", instance)
		return instance
	})
}

func (this *ServiceProvider) Start() error {
	this.app.Call(func(gateFactory IAuth.GateFactory) {
		for class, policy := range this.Policies {
			gateFactory.Policy(class, policy)
		}
	})
	return nil
}

func (this *ServiceProvider) Stop() {
}
