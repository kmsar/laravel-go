package IRateLimit

import "time"

type Limiter interface {

	// Take Get next pass time.
	Take() time.Time
}

type RateLimiter interface {

	// Limiter Get the current limiter by the given name.
	Limiter(name string, limiter func() Limiter) Limiter
}
