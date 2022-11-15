package Http

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IExeption"
	"github.com/kmsar/laravel-go/Framework/Contracts/IHttp"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
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
