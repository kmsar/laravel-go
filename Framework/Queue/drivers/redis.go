package drivers

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IRedis"
	"github.com/kmsar/laravel-go/Framework/Contracts/Queue"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"time"
)

type Redis struct {
	lifetime time.Duration
	redis    IRedis.RedisConnection
	key      string
}

func (r Redis) Push(job Queue.Job, queue ...string) error {
	//TODO implement me
	panic("implement me")
}

func (r Redis) PushOn(queue string, job Queue.Job) error {
	//TODO implement me
	panic("implement me")
}

func (r Redis) PushRaw(payload, queue string, options ...Support.Fields) error {
	//TODO implement me
	panic("implement me")
}

func (r Redis) Later(delay time.Time, job Queue.Job, queue ...string) error {
	//TODO implement me
	panic("implement me")
}

func (r Redis) LaterOn(queue string, delay time.Time, job Queue.Job) error {
	//TODO implement me
	panic("implement me")
}

func (r Redis) GetConnectionName() string {
	//TODO implement me
	panic("implement me")
}

func (r Redis) Release(job Queue.Job, delay ...int) error {
	//TODO implement me
	panic("implement me")
}

func (r Redis) Listen(queue ...string) chan Queue.Msg {
	//TODO implement me
	panic("implement me")
}

func (r Redis) Stop() {
	//TODO implement me
	panic("implement me")
}
