package websocket

import "github.com/gorilla/websocket"

type Config struct {
	Upgrader websocket.Upgrader
}
