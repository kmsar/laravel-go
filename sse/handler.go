package sse

import (
	"fmt"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Logs"

	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IHttp"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/ISerialize"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Http"
)

func New(controller IHttp.SseController) interface{} {
	return func(request *Http.Request, serializer ISerialize.Serializer, sse IHttp.Sse) error {
		var fd = sse.GetFd()
		if err := controller.OnConnect(request, fd); err != nil {
			Logs.WithError(err).WithFields(request.Fields()).WithField("fd", fd).Debug("sse.New: OnConnect failed")
			return err
		}

		var response = request.Response()
		response.Header().Set("Content-Type", "text/event-stream")
		response.Header().Set("Cache-Control", "no-cache")
		response.Header().Set("Connection", "keep-alive")
		response.Header().Set("Access-Control-Allow-Origin", "*")

		var (
			messageChan = make(chan interface{})
			closeChan   = make(chan bool)
			conn        = NewConnection(messageChan, closeChan, fd)
		)
		sse.Add(conn)

		defer func() {
			controller.OnClose(fd)
			close(messageChan)
			close(closeChan)
			closeChan = nil
			messageChan = nil
		}()

		for {
			select {
			case message := <-messageChan:
				var _, err = fmt.Fprintf(response, "%s\n", handleMessage(message, serializer))
				if err != nil {
					Logs.WithError(err).
						WithField("message", message).WithField("fd", fd).
						Error("sse.New: response.Write failed")
					return nil
				}
				response.Flush()

			case <-closeChan:
				return nil

			// connection is closed then defer will be executed
			case <-request.Request().Context().Done():
				_ = sse.Close(fd)
				return nil
			}
		}
	}
}

func handleMessage(msg interface{}, serializer ISerialize.Serializer) []byte {
	switch v := msg.(type) {
	case []byte:
		return v
	case string:
		return []byte(v)
	default:
		return []byte(serializer.Serialize(msg))
	}
}
