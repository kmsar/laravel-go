package Limiter

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IRateLimit"
	"sync"
)

type RateLimiter struct {
	limiters sync.Map
}

func (rate *RateLimiter) Limiter(name string, limiter func() IRateLimit.Limiter) IRateLimit.Limiter {
	if limit, exists := rate.limiters.Load(name); exists {
		return limit.(IRateLimit.Limiter)
	}
	limit := limiter()
	rate.limiters.Store(name, limit)
	return limit
}
