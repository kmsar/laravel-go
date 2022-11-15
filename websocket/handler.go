package websocket

import (
	"github.com/gorilla/websocket"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IExeption"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/ISerialize"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IWebsocket"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Http"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Logs"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Exceptions"
)

var (
	upgrader = websocket.Upgrader{}
)

func New(controller IWebsocket.WebSocketController) interface{} {
	return func(request *Http.Request, serializer ISerialize.Serializer, socket IWebsocket.WebSocket, handler IExeption.ExceptionHandler) error {
		var ws, err = upgrader.Upgrade(request.Context.Response(), request.Request(), nil)

		if err != nil {
			Logs.WithError(err).Error("websocket.New: Upgrade failed")
			return err
		}

		var fd = socket.GetFd()

		if err = controller.OnConnect(request, fd); err != nil {
			Logs.WithError(err).Error("websocket.New: OnConnect failed")
			return err
		}

		var conn = NewConnection(ws, fd)
		socket.Add(conn)

		defer func() {
			controller.OnClose(fd)
			if closeErr := socket.Close(conn.Fd()); closeErr != nil {
				Logs.WithError(closeErr).Error("websocket.New: Connection close failed")
			}
		}()

		for {
			// Read
			var msgType, msg, readErr = ws.ReadMessage()
			if readErr != nil {
				Logs.WithError(readErr).Error("websocket.New: Failed to read message")
				return readErr
			}

			switch msgType {
			case websocket.TextMessage, websocket.BinaryMessage:
				go handleMessage(NewFrame(msg, conn, serializer), controller, handler)
			case websocket.CloseMessage:
				return nil
			}
		}
	}
}

func handleMessage(frame IWebsocket.WebSocketFrame, controller IWebsocket.WebSocketController, handler IExeption.ExceptionHandler) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			handler.Handle(Exception{
				Exception: Exceptions.WithRecover(panicValue, Support.Fields{
					"msg": frame.RawString(),
					"fd":  frame.Connection().Fd(),
				}),
			})
		}
	}()
	controller.OnMessage(frame)
}
