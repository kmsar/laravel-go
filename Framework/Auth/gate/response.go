package gate

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IHttp"
	"net/http"
)

type Response struct {
	Allowed bool        `json:"allowed"`
	Message string      `json:"message"`
	Code    interface{} `json:"code"`
}

func (this *Response) Status() int {
	if this.Allowed {
		return http.StatusOK
	}

	return http.StatusUnauthorized
}

func (this *Response) Response(ctx IHttp.HttpContext) error {
	return ctx.JSON(this.Status(), this)
}
