package sse

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IFoundation"
	"github.com/kmsar/laravel-go/Framework/Contracts/IHttp"
	"sync"
)

type ServiceProvider struct {
}

func (s ServiceProvider) Register(application IFoundation.IApplication) {
	application.NamedSingleton("sse", func() IHttp.Sse {
		return &Sse{
			fdMutex:     sync.Mutex{},
			connMutex:   sync.Mutex{},
			connections: map[uint64]IHttp.SseConnection{},
			count:       0,
		}
	})
}

func (s ServiceProvider) Start() error {
	return nil
}
func (this *ServiceProvider) Boot(application IFoundation.IApplication) {
	//TODO implement me
	panic("implement me")
}
func (s ServiceProvider) Stop() {
}
