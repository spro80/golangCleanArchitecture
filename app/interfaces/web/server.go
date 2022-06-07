package web

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

var echoServer *echo.Echo

type ServerInterface interface {
	StartServer() error
}

type ServerHandler struct {
}

func NewServer() *ServerHandler {
	return &ServerHandler{}
}

func (s ServerHandler) Start(port int) {
	echoServer = echo.New()

	server := &http.Server{
		Addr:         string(port),
		ReadTimeout:  3 * time.Minute,
		WriteTimeout: 3 * time.Minute,
	}
	//log.Info("App listen in port %s", port)
	echoServer.Logger.Fatal(echoServer.StartServer(server))

}
