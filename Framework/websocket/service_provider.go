package websocket

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IConfig"
	"github.com/kmsar/laravel-go/Framework/Contracts/IFoundation"
	"github.com/kmsar/laravel-go/Framework/Contracts/IWebsocket"
	"sync"
)

type ServiceProvider struct {
}

func (s ServiceProvider) Register(application IFoundation.IApplication) {
	application.NamedSingleton("websocket", func(config IConfig.Config) IWebsocket.WebSocket {
		var wsConfig = config.Get("websocket").(Config)

		upgrader = wsConfig.Upgrader

		return &WebSocket{
			connMutex:   sync.Mutex{},
			fdMutex:     sync.Mutex{},
			connections: map[uint64]IWebsocket.WebSocketConnection{},
			count:       0,
		}
	})
}

func (s ServiceProvider) Start() error {
	return nil
}

func (s ServiceProvider) Stop() {
}
