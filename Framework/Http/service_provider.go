package Http

import (
	"fmt"
	"github.com/kmsar/laravel-go/Framework/Contracts/IConfig"
	"github.com/kmsar/laravel-go/Framework/Contracts/IEvent"
	"github.com/kmsar/laravel-go/Framework/Contracts/IFoundation"
	"github.com/kmsar/laravel-go/Framework/Contracts/IRouting"
	"github.com/kmsar/laravel-go/Framework/Logs"
	"github.com/kmsar/laravel-go/Framework/Support/Str"
	"net/http"
)

type ServiceProvider struct {
	app IFoundation.IApplication

	RouteCollectors []interface{}
}

func (this *ServiceProvider) Stop() {
	this.app.Call(func(dispatcher IEvent.EventDispatcher, router IRouting.Router) {
		if err := router.Close(); err != nil {
			Logs.WithError(err).Info("Router closes error")
		}
		dispatcher.Dispatch(&ServeClosed{})
	})
}

func (this *ServiceProvider) Start() error {
	for _, collector := range this.RouteCollectors {
		this.app.Call(collector)
	}

	err := this.app.Call(func(router IRouting.Router, config IConfig.Config) error {
		httpConfig := config.Get("http").(Config)
		return router.Start(
			Str.StringOr(
				httpConfig.Address,
				fmt.Sprintf("%s:%s", httpConfig.Host, Str.StringOr(httpConfig.Port, "8000")),
			),
		)
	})[0].(error)

	if err != nil && err != http.ErrServerClosed {
		Logs.WithError(err).Error("http service can't start")
		this.app.Stop()
		return err
	}

	return nil
}

func (this *ServiceProvider) Register(app IFoundation.IApplication) {
	this.app = app

	app.NamedSingleton("Router", func() IRouting.Router {
		return New(this.app)
	})
}
