package Limiter

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IFoundation"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IRateLimit"
	"sync"
)

type ServiceProvider struct {
}

func (s ServiceProvider) Register(application IFoundation.IApplication) {
	application.NamedSingleton("ratelimiter", func() IRateLimit.RateLimiter {
		return &RateLimiter{limiters: sync.Map{}}
	})
}

func (s ServiceProvider) Start() error {
	return nil
}

func (s ServiceProvider) Stop() {
}
