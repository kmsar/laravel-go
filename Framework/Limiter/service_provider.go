package Limiter

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IFoundation"
	"github.com/kmsar/laravel-go/Framework/Contracts/IRateLimit"
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
func (this *ServiceProvider) Boot(application IFoundation.IApplication) {
	//TODO implement me
	panic("implement me")
}
func (s ServiceProvider) Stop() {
}
