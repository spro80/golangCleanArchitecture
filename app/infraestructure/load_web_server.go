package infraestructure

import (
	"github.com/spro80/golangCleanArchitecture/app/application/useCase/templateUseCase"
	"github.com/spro80/golangCleanArchitecture/app/infraestructure/controllers/templateController"
	"github.com/spro80/golangCleanArchitecture/app/interfaces/web"
	"github.com/spro80/golangCleanArchitecture/app/shared/config"
)

type LoadInterface interface {
	LoadRoutes()
	StartServer()
}

type LoadHandler struct {
	config config.ConfigInterface
	web    web.WebServerHandler
}

func NewLoad(config config.ConfigInterface, web web.WebServerHandler) *LoadHandler {
	return &LoadHandler{config: config, web: web}
}

func (ws LoadHandler) LoadRoutes() {

	//Load useCase: (Are used as dependency injection in controllers.)
	templateUseCase := templateUseCase.NewTemplateUseCase()

	//Load Controller
	templateCtrl := templateController.NewTemplateController(templateUseCase)

	//Calling initialize routes
	ws.web.InitRoutes(templateCtrl)
}
