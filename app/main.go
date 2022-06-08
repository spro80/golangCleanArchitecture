package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spro80/golangCleanArchitecture/app/interfaces/web"
	"github.com/spro80/golangCleanArchitecture/app/shared/config"
)

func main() {
	fmt.Println("[main.go] Init.")

	//Signs Captured
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	configurations := config.NewConfig()

	fmt.Println("[main.go] Calling to New Web Server")
	ws := web.NewWebServer(configurations)

	ws.InitRoutes()

	fmt.Println("[main.go] Calling webServer Start")
	go ws.Start()

	fmt.Println("[main.go] Closing from main")
	//Graceful Shutdown process
	sig := <-quit
	fmt.Println("[main.go] signal captured from os")
	gracefulShutdown(sig)
}

func gracefulShutdown(sig os.Signal) {
	fmt.Println("[main.go] Closing gracefulShutdown")
	fmt.Println("[main.go] TODO shutdown webServer")
}
