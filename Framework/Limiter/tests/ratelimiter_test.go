package tests

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IRateLimit"
	"github.com/kmsar/laravel-go/Framework/Limiter"
	"go.uber.org/ratelimit"
	"log"
	"testing"
	"time"
)

func TestLimiter(t *testing.T) {
	rate := Limiter.RateLimiter{}
	prev := time.Now()

	for i := 0; i < 20; i++ {
		now := rate.Limiter("testing", func() IRateLimit.Limiter {
			return ratelimit.New(10, ratelimit.Per(time.Second)) // per second
		}).Take()
		log.Default().Println(i, now.Sub(prev))
		prev = now
	}
}

func BenchmarkLimiter(b *testing.B) {
	rate := Limiter.RateLimiter{}
	per := b.N / 2
	if per <= 0 {
		per = 1
	}

	for i := 0; i < b.N; i++ {
		rate.Limiter("testing", func() IRateLimit.Limiter {
			return ratelimit.New(per, ratelimit.Per(time.Second)) // per second
		}).Take()
	}
}
