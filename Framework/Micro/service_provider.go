package Micro

import (
	"fmt"
	"github.com/kmsar/laravel-go/Framework/Contracts/IConfig"
	"github.com/kmsar/laravel-go/Framework/Logs"

	"github.com/kmsar/laravel-go/Framework/Contracts/IFoundation"
	"go-micro.dev/v4"
	"runtime/debug"
)

type ServiceProvider struct {
	app             IFoundation.IApplication
	ServiceRegister func(service micro.Service) error
}

func (s *ServiceProvider) Boot(application IFoundation.IApplication) {
	//TODO implement me
	panic("implement me")
}

func (s *ServiceProvider) Register(application IFoundation.IApplication) {
	s.app = application
	application.NamedSingleton("micro", func(config IConfig.Config) micro.Service {
		var (
			microConfig = config.Get("micro").(Config)
			service     = micro.NewService()
			options     = append(microConfig.CustomOptions, micro.HandleSignal(microConfig.Signal))
		)
		if microConfig.Registry != nil {
			options = append(options, micro.Registry(microConfig.Registry))
		}

		if microConfig.Auth != nil {
			options = append(options, micro.Auth(microConfig.Auth))
		}

		if microConfig.Broker != nil {
			options = append(options, micro.Broker(microConfig.Broker))
		}

		if microConfig.Cmd != nil {
			options = append(options, micro.Cmd(microConfig.Cmd))
		}

		if microConfig.Config != nil {
			options = append(options, micro.Config(microConfig.Config))
		}

		if microConfig.Client != nil {
			options = append(options, micro.Client(microConfig.Client))
		}

		if microConfig.Server != nil {
			options = append(options, micro.Server(microConfig.Server))
		}

		if microConfig.Store != nil {
			options = append(options, micro.Store(microConfig.Store))
		}

		if microConfig.Client != nil {
			options = append(options, micro.Client(microConfig.Client))
		}

		if microConfig.Runtime != nil {
			options = append(options, micro.Runtime(microConfig.Runtime))
		}

		if microConfig.Transport != nil {
			options = append(options, micro.Transport(microConfig.Transport))
		}

		if microConfig.Profile != nil {
			options = append(options, micro.Profile(microConfig.Profile))
		}
		if microConfig.Context != nil {
			options = append(options, micro.Context(microConfig.Context))
		}

		for _, handler := range microConfig.BeforeStart {
			options = append(options, micro.BeforeStart(handler))
		}
		for _, handler := range microConfig.BeforeStop {
			options = append(options, micro.BeforeStop(handler))
		}

		service.Init(options...)

		return service
	})
}

func (s *ServiceProvider) Start() error {
	defer func() {
		if err, ok := recover().(error); ok && err != nil {
			Logs.WithError(err).Error("micro.ServiceProvider.Start: micro service start failed")
			debug.PrintStack()
			s.app.Stop()
		}
	}()
	return s.app.Call(func(service micro.Service) error {
		fmt.Println(service, s.app)
		if err := s.ServiceRegister(service); err != nil {
			defer s.app.Stop()
			return err
		}

		return service.Run()
	})[0].(error)
}

func (s *ServiceProvider) Stop() {

}
