package web

import (
	"fmt"
	"github.com/spro80/golangCleanArchitecture/app/interfaces/input/source"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spro80/golangCleanArchitecture/app/infraestructure/web/routes"
	"github.com/spro80/golangCleanArchitecture/app/shared/config"
)

type WebServerInterface interface {
	StartServer() error
	InitRoutes()
}

type WebServerHandler struct {
	echoServer *echo.Echo
	config     config.ConfigInterface
}

func NewWebServer(config config.ConfigInterface) *WebServerHandler {
	echoServer := echo.New()
	echoServer.HideBanner = true

	return &WebServerHandler{echoServer: echoServer, config: config}
}

func (ws WebServerHandler) InitRoutes(
	userAddInputV1FromApi source.FromApiInterface,
	userDeleteInputV1FromApi source.FromApiInterface) {

	fmt.Println("[server] Init in InitRoutes")
	//config := config.NewConfig()

	routes.NewHealthCheckRoute(ws.echoServer, ws.config)
	routes.NewUserAddRoute(ws.echoServer, userAddInputV1FromApi, userDeleteInputV1FromApi)
	//routes.NewGetAllUserRoute(ws.echoServer, getAllUserCtrlInterface, responseInterface)
	//routes.NewDeleteUserRoute(ws.echoServer, deleteUserCtrlInterface, responseInterface)
	fmt.Println("[server] InitRoutes called successfully")
}

func (ws WebServerHandler) Start() error {
	fmt.Println("[server] Init in method Start")

	/*configServer, err := ws.config.Handler()
	if err != nil {
		fmt.Println("err in get handler")
	}
	fmt.Println(configServer)
	fmt.Println(reflect.TypeOf(configServer))
	*/

	port, _ := ws.config.GetPort()
	//url, _ := ws.config.GetUrl()
	//portAddr := ":" + strconv.Itoa(port)
	portAddr := ":" + port

	server := &http.Server{
		//Addr:         fmt.Sprintf(":%s", "9090"),
		Addr:         portAddr,
		ReadTimeout:  3 * time.Minute,
		WriteTimeout: 3 * time.Minute,
	}
	ws.echoServer.Logger.Fatal(ws.echoServer.StartServer(server))

	return nil
}
