package scheduling

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IRedis"
)

type Mutex struct {
	redis IRedis.RedisFactory
	store string
}

func (this *Mutex) Create(event *Event) bool {
	if this.redis != nil {
		var _, err = this.redis.Connection(this.store).Set(event.MutexName(), "1", event.expiresAt)
		return err == nil
	}
	return true
}
func (this *Mutex) Exists(event *Event) bool {
	if this.redis != nil {
		var num, err = this.redis.Connection(this.store).Exists(event.MutexName())
		return err == nil && num > 0
	}
	return true
}

func (this *Mutex) Forget(event *Event) {
	if this.redis != nil {
		_, _ = this.redis.Connection(this.store).Del(event.MutexName())
	}
}
