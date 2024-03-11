package websocket

import (
	"log"

	"github.com/1ets/lets"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Engine for controller
type WebSocketHandle struct {
	Gin     *gin.Engine
	handler IWebSocketEvent
}

// Setup Gin
func (wsh *WebSocketHandle) SetGin(c *gin.Engine) {
	wsh.Gin = c
}

// Setup Routing
func (wsh *WebSocketHandle) Handle(endPoint string, handler IWebSocketEvent) {
	wsh.handler = handler

	wsh.Gin.GET(endPoint, wsh.route)
}

// Setup Route Handling
func (wsh *WebSocketHandle) route(c *gin.Context) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer ws.Close()

	wsh.handler.OnConnect()

	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		lets.LogI("MessageType: %v", mt)
		lets.LogI("Message    : %s", message)

		wsh.handler.OnMessage()
	}

	wsh.handler.OnDisconnect()
}

// Interface for accessable method
type IWebSocketEvent interface {
	OnConnect()
	OnMessage()
	OnDisconnect()
}
