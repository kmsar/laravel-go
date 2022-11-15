package websocket

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/ISerialize"
	"github.com/kmsar/laravel-go/Framework/Contracts/IWebsocket"
)

type Frame struct {
	IWebsocket.WebSocketConnection
	raw        []byte
	serializer ISerialize.Serializer
}

func NewFrame(raw []byte, conn IWebsocket.WebSocketConnection, serializer ISerialize.Serializer) IWebsocket.WebSocketFrame {
	return &Frame{
		WebSocketConnection: conn,
		raw:                 raw,
		serializer:          serializer,
	}
}

func (frame *Frame) Connection() IWebsocket.WebSocketConnection {
	return frame.WebSocketConnection
}

func (frame *Frame) Raw() []byte {
	return frame.raw
}

func (frame *Frame) RawString() string {
	return string(frame.raw)
}

func (frame *Frame) Parse(v interface{}) error {
	return frame.serializer.UnSerialize(frame.RawString(), v)
}
