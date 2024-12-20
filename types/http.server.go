package types

import (
	"github.com/1ets/lets"

	"github.com/gin-gonic/gin"
)

// Default gRPC configuration
const (
	SERVER_HTTP_PORT = "5000"
	SERVER_HTTP_MODE = "debug"
)

// Interface for accessable method
type IHttpServer interface {
	GetPort() string
	GetMode() string
	GetMiddleware() func(*gin.Engine)
	GetRouter() func(*gin.Engine)
	GetGzip() bool
}

// Serve information
type HttpServer struct {
	Port       string
	Middleware func(*gin.Engine)
	Router     func(*gin.Engine)
	Gzip       bool
	Mode       string
}

// Get Port
func (hs *HttpServer) GetPort() string {
	if hs.Port == "" {
		lets.LogW("Config: SERVER_HTTP_PORT is not set, using default configuration.")

		return SERVER_HTTP_PORT
	}

	return hs.Port
}

// Get Mode
func (hs *HttpServer) GetMode() string {
	if hs.Mode == "" {
		lets.LogW("Config: SERVER_HTTP_MODE is not set, using default configuration.")

		return SERVER_HTTP_MODE
	}

	return hs.Mode
}

// Get Middleware
func (hs *HttpServer) GetMiddleware() func(*gin.Engine) {
	return hs.Middleware
}

// Get Router
func (hs *HttpServer) GetRouter() func(*gin.Engine) {
	return hs.Router
}

// Get Gzip
func (hs *HttpServer) GetGzip() bool {
	return hs.Gzip
}
