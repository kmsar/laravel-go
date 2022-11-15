package sse

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IHttp"
)

func Default() interface{} {
	return New(&DefaultController{})
}

type DefaultController struct {
}

func (d *DefaultController) OnConnect(request IHttp.IHttpRequest, fd uint64) error {
	return nil
}

func (d *DefaultController) OnClose(fd uint64) {
}
