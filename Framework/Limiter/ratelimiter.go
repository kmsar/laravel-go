package Limiter

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IRateLimit"
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
