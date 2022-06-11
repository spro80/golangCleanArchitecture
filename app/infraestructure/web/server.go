package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	controllers "github.com/spro80/golangCleanArchitecture/app/infraestructure/controllers/templateController"
	"github.com/spro80/golangCleanArchitecture/app/interfaces/web/routes"
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

func (ws WebServerHandler) InitRoutes(templateCtrl controllers.TemplateControllerInterface) {
	fmt.Println("[server.go] Init in InitRoutes")
	//config := config.NewConfig()

	routes.NewHealthCheckRoute(ws.echoServer, ws.config)
	routes.NewTemplateRoute(ws.echoServer, templateCtrl)

	fmt.Println("[server.go] InitRoutes called successfully")
}

func (ws WebServerHandler) Start() error {
	fmt.Println("[server.go] Init in method Start")

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
