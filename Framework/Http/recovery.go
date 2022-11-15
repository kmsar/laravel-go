package Http

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IHttp"
	"github.com/kmsar/laravel-go/Framework/Contracts/IPipeline"
	"github.com/kmsar/laravel-go/Framework/Support/Exceptions"
)

func (this *Router) recovery(request *Request, next IPipeline.Pipe) (result interface{}) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			if res := this.errHandler(panicValue, request); res != nil { //The response returned by the exception handler takes precedence
				HandleResponse(res, request)
			} else {
				HandleResponse(panicValue, request) //If the exception handler does not define a response, respond directly to the value of panic
			}
			result = nil
		}
	}()
	return next(request)
}

func (this *Router) errHandler(err interface{}, request IHttp.IHttpRequest) (result interface{}) {
	var httpException Exception
	switch rawErr := err.(type) {
	case Exception:
		httpException = rawErr
	default:
		httpException = Exception{
			Exception: Exceptions.ResolveException(err),
			Request:   request,
		}
	}

	// 调用容器内的异常处理器
	return this.app.StaticCall(exceptionHandler, httpException)[0]
}
