package frameworks

import (
	"fmt"
	"time"

	"github.com/1ets/lets"
	"github.com/1ets/lets/types"
	"github.com/1ets/lets/websocket"

	"github.com/gin-gonic/gin"
)

// WebSocket Configurations
var WebSocketConfig types.IWebSocketServer

// WebSocket service struct
type webSocketServer struct {
	server  string
	engine  *gin.Engine
	handler func(*websocket.WebSocketHandle) // Handle endpoint event
}

// Initialize service
func (ws *webSocketServer) init() {
	gin.SetMode(WebSocketConfig.GetMode())

	ws.server = fmt.Sprintf(":%s", WebSocketConfig.GetPort())
	ws.engine = gin.New()
	ws.handler = WebSocketConfig.GetHandler()

	var defaultLogFormatter = func(param gin.LogFormatterParams) string {
		var statusColor, methodColor, resetColor string
		if param.IsOutputColor() {
			statusColor = param.StatusCodeColor()
			methodColor = param.MethodColor()
			resetColor = param.ResetColor()
		}

		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}

		return fmt.Sprintf("%s[WebSocket]%s %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
			"\x1b[32m", resetColor,
			param.TimeStamp.Format("2006-01-02 15:04:05"),
			statusColor, param.StatusCode, resetColor,
			param.Latency,
			param.ClientIP,
			methodColor, param.Method, resetColor,
			param.Path,
			param.ErrorMessage,
		)
	}

	ws.engine.Use(gin.LoggerWithFormatter(defaultLogFormatter))
}

// Run service
func (ws *webSocketServer) serve() {
	ws.engine.Run(ws.server)
}

// Start WebSocket service
func WebSocket() {
	if WebSocketConfig == nil {
		return
	}

	lets.LogI("WebSocket Server Starting ...")

	var ws webSocketServer

	ws.init()

	var wsHandler websocket.WebSocketHandle = websocket.WebSocketHandle{
		Gin: ws.engine,
	}
	ws.handler(&wsHandler)
	ws.serve()
}
