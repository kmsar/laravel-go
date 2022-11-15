package IHttp

type Sse interface {

	// Add  a connection.
	Add(connect SseConnection)

	// GetFd get new fd.
	GetFd() uint64

	// Close  fd.
	Close(fd uint64) error

	// Send  a message to the connection with the specified fd.
	Send(fd uint64, message interface{}) error
}

type SseConnection interface {

	// Fd Get connection ID.
	Fd() uint64

	// Send  a message to this connection.
	Send(message interface{}) error

	// Close  connection ID.
	Close() error
}

type SseController interface {

	// OnConnect Can handle some authentication and other operations when connecting.
	OnConnect(request IHttpRequest, fd uint64) error

	// OnClose Handling close events.
	OnClose(fd uint64)
}
