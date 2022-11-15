package Limiter

import (
	"go.uber.org/ratelimit"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IHttp"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IPipeline"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IRateLimit"
)

func Middleware(rate int, opts ...ratelimit.Option) interface{} {
	return func(request IHttp.IHttpRequest, pipe IPipeline.Pipe, limiter IRateLimit.RateLimiter) interface{} {
		limiter.Limiter("request", func() IRateLimit.Limiter {
			return ratelimit.New(rate, opts...)
		}).Take()

		return pipe(request)
	}
}
