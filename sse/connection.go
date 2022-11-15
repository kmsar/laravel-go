package sse

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IHttp"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Exceptions"
)

type Connection struct {
	fd        uint64
	msgPipe   chan interface{}
	closePipe chan bool
}

func NewConnection(pipe chan interface{}, closePipe chan bool, fd uint64) IHttp.SseConnection {
	return &Connection{
		fd:        fd,
		msgPipe:   pipe,
		closePipe: closePipe,
	}
}

func (conn *Connection) Fd() uint64 {
	return conn.fd
}

func (conn *Connection) Close() error {
	conn.closePipe <- true
	return nil
}

func (conn *Connection) Send(msg interface{}) (err error) {
	defer func() {
		if v := recover(); v != nil {
			err = Exception{Exceptions.WithRecover(v, Support.Fields{
				"msg": msg,
			})}
		}
	}()
	conn.msgPipe <- msg
	return
}
