package IWebsocket

import "github.com/kmsar/laravel-go/Framework/Contracts/IHttp"

type WebSocket interface {

	// Add   connection, return fd.
	Add(connect WebSocketConnection)

	// GetFd get new fd.
	GetFd() uint64

	// Close  the specified connection.
	Close(fd uint64) error

	// Send  a message to the connection with the specified fd.
	Send(fd uint64, message interface{}) error
}

type WebSocketConnection interface {
	WebSocketSender

	// Fd get fd.
	Fd() uint64

	// Close  the connection.
	Close() error
}

type WebSocketSender interface {

	// Send  a message to this connection
	Send(message interface{}) error

	// SendBytes send a message to this connection.
	SendBytes(bytes []byte) error

	// SendBinary send binary message to this connection.
	SendBinary(bytes []byte) error
}

type WebSocketFrame interface {
	WebSocketSender

	// Connection get current connection.
	Connection() WebSocketConnection

	// Raw get the original message.
	Raw() []byte

	// RawString get string message.
	RawString() string

	// Parse json parameters.
	Parse(v interface{}) error
}

type WebSocketController interface {

	// OnConnect Can handle some authentication and other operations when connecting.
	OnConnect(request IHttp.IHttpRequest, fd uint64) error

	// OnMessage What to do when a new message arrives.
	OnMessage(frame WebSocketFrame)

	// OnClose Handling close events.
	OnClose(fd uint64)
}
