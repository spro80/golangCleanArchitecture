package infraestructure

import (
	"github.com/spro80/golangCleanArchitecture/app/application/useCase/registerUserUseCase"
	"github.com/spro80/golangCleanArchitecture/app/application/useCase/templateUseCase"
	"github.com/spro80/golangCleanArchitecture/app/infraestructure/controllers/registerUserController"
	"github.com/spro80/golangCleanArchitecture/app/infraestructure/controllers/templateController"
	"github.com/spro80/golangCleanArchitecture/app/infraestructure/web"
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
	registerUserUseCase := registerUserUseCase.NewRegisterUserUseCase()

	//Load Controller
	templateCtrl := templateController.NewTemplateController(templateUseCase)
	registerUserCtrl := registerUserController.NewRegisterUserController(registerUserUseCase)

	//Calling initialize routes
	ws.web.InitRoutes(templateCtrl, registerUserCtrl)

}
