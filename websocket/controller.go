package websocket

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IHttp"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IWebsocket"
)

func Default(handler func(frame IWebsocket.WebSocketFrame)) interface{} {
	return New(&DefaultController{Handler: handler})
}

type DefaultController struct {
	Handler func(frame IWebsocket.WebSocketFrame)
}

func (d *DefaultController) OnConnect(request IHttp.IHttpRequest, fd uint64) error {
	return nil
}

func (d *DefaultController) OnMessage(frame IWebsocket.WebSocketFrame) {
	d.Handler(frame)
}

func (d *DefaultController) OnClose(fd uint64) {
}
