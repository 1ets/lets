package types

import (
	"github.com/1ets/lets"
	"github.com/1ets/lets/websocket"
)

// Default gRPC configuration
const (
	SERVER_WS_PORT = "5050"
	SERVER_WS_MODE = "debug"
)

// Interface for accessable method
type IWebSocketServer interface {
	GetPort() string
	GetMode() string
	GetHandler() func(*websocket.WebSocketHandle)
}

// Serve information
type WebSocketServer struct {
	Port   string
	Mode   string
	Router func(*websocket.WebSocketHandle)
}

// Get Port
func (hs *WebSocketServer) GetPort() string {
	if hs.Port == "" {
		lets.LogW("Config: SERVER_WS_PORT is not set, using default configuration.")

		return SERVER_WS_PORT
	}

	return hs.Port
}

// Get Mode
func (hs *WebSocketServer) GetMode() string {
	if hs.Mode == "" {
		lets.LogW("Config: SERVER_WS_MODE is not set, using default configuration.")

		return SERVER_WS_MODE
	}

	return hs.Mode
}

// Get Router
func (hs *WebSocketServer) GetHandler() func(*websocket.WebSocketHandle) {
	return hs.Router
}
