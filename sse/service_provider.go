package sse

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IFoundation"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IHttp"
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

func (s ServiceProvider) Stop() {
}
