package Limiter

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IHttp"
	"github.com/kmsar/laravel-go/Framework/Contracts/IPipeline"
	"github.com/kmsar/laravel-go/Framework/Contracts/IRateLimit"
	"go.uber.org/ratelimit"
)

func Middleware(rate int, opts ...ratelimit.Option) interface{} {
	return func(request IHttp.IHttpRequest, pipe IPipeline.Pipe, limiter IRateLimit.RateLimiter) interface{} {
		limiter.Limiter("request", func() IRateLimit.Limiter {
			return ratelimit.New(rate, opts...)
		}).Take()

		return pipe(request)
	}
}
