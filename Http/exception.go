package Http

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IExeption"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IHttp"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
)

type Exception struct {
	IExeption.Exception
	Request IHttp.IHttpRequest
}

func (this Exception) Fields() Support.Fields {
	return Support.Fields{
		"method": this.Request.Request().Method,
		"path":   this.Request.Path(),
		"query":  this.Request.QueryParams(),
		"fields": this.Request.Fields(),
	}
}
