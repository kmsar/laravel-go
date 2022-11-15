package Http

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IHttp"
)

type RequestBefore struct {
	request IHttp.IHttpRequest
}

func (this *RequestBefore) Event() string {
	return "REQUEST_BEFORE"
}

func (this *RequestBefore) Sync() bool {
	return true
}
func (this *RequestBefore) Request() IHttp.IHttpRequest {
	return this.request
}

type RequestAfter struct {
	request IHttp.IHttpRequest
}

func (this *RequestAfter) Event() string {
	return "REQUEST_AFTER"
}

func (this *RequestAfter) Request() IHttp.IHttpRequest {
	return this.request
}

type ResponseBefore struct {
	request IHttp.IHttpRequest
}

func (this *ResponseBefore) Event() string {
	return "RESPONSE_BEFORE"
}

func (this *ResponseBefore) Request() IHttp.IHttpRequest {
	return this.request
}

func (this *ResponseBefore) Sync() bool {
	return true
}
