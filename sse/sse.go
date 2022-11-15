package sse

import (
	"errors"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IHttp"

	"sync"
)

var (
	ConnectionDontExistsErr = errors.New("connection does not exist")
)

type Sse struct {
	fdMutex     sync.Mutex
	connMutex   sync.Mutex
	connections map[uint64]IHttp.SseConnection
	count       uint64
}

func (sse *Sse) Add(connect IHttp.SseConnection) {
	sse.connMutex.Lock()
	defer sse.connMutex.Unlock()
	sse.connections[connect.Fd()] = connect
}

func (sse *Sse) GetFd() uint64 {
	sse.fdMutex.Lock()
	defer sse.fdMutex.Unlock()
	sse.count++
	var fd = sse.count
	return fd
}

func (sse *Sse) Close(fd uint64) error {
	var conn, exists = sse.connections[fd]
	if exists {
		sse.connMutex.Lock()
		defer sse.connMutex.Unlock()
		delete(sse.connections, fd)
		return conn.Close()
	}

	return ConnectionDontExistsErr
}

func (sse *Sse) Send(fd uint64, message interface{}) error {
	var conn, exists = sse.connections[fd]
	if exists {
		return conn.Send(message)
	}

	return ConnectionDontExistsErr
}
