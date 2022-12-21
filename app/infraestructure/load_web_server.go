package infraestructure

import (
	"fmt"
	"github.com/spro80/golangCleanArchitecture/app/application/useCase/getAllUserUseCase"
	"github.com/spro80/golangCleanArchitecture/app/application/useCase/registerUserUseCase"
	"github.com/spro80/golangCleanArchitecture/app/application/useCase/templateUseCase"
	"github.com/spro80/golangCleanArchitecture/app/infraestructure/mongo_client"
	"github.com/spro80/golangCleanArchitecture/app/infraestructure/mongo_client/user_repository"
	"github.com/spro80/golangCleanArchitecture/app/infraestructure/web"
	"github.com/spro80/golangCleanArchitecture/app/interfaces/gateways/user_gateway"
	"github.com/spro80/golangCleanArchitecture/app/interfaces/input/controllers/getAllUserController"
	"github.com/spro80/golangCleanArchitecture/app/interfaces/input/controllers/registerUserController"
	"github.com/spro80/golangCleanArchitecture/app/interfaces/input/controllers/templateController"
	"github.com/spro80/golangCleanArchitecture/app/shared/config"
	"github.com/spro80/golangCleanArchitecture/app/shared/utils/response"
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

	clientMongo, err := LoadDatabase()
	if err != nil {
		fmt.Printf("[Dependencies]: Error initializing. database %s", err.Error())
		//return err
	}

	clientDatabase, err := clientMongo.Database("administration")
	if err != nil {
		fmt.Printf("[Dependencies]: Error in get name of database %s", err.Error())
		//return err
	}

	//Load shared
	responseStruct := response.NewResponse()
	//Load repository
	userRepository := user_repository.NewUserRepository(clientDatabase)

	//Load gateways
	userGateway := user_gateway.NewRepositoryGateway(userRepository)

	//Load useCase: (Are used as dependency injection in controllers.)
	templateUseCase := templateUseCase.NewTemplateUseCase()
	getAllUserUseCase := getAllUserUseCase.NewGetAllUserUseCase(userGateway)
	registerUserUseCase := registerUserUseCase.NewRegisterUserUseCase(userGateway)

	//Load Controller
	templateCtrl := templateController.NewTemplateController(templateUseCase)
	getAllUserCtrl := getAllUserController.NewGetAllUserController(getAllUserUseCase)
	registerUserCtrl := registerUserController.NewRegisterUserController(registerUserUseCase)

	//Calling initialize routes
	ws.web.InitRoutes(responseStruct, templateCtrl, getAllUserCtrl, registerUserCtrl)

}

func LoadDatabase() (mongo_client.MongoClientInterface, error) {
	mongoClient, err := mongo_client.NewClient("mongodb://127.0.0.1:27017")
	if err != nil {
		fmt.Printf("[Dependencies]: Error instantiating the MongoClient")
		//return nil, infrastructure_errors.New(map[string]interface{}{"UrlClient": config.GetString("database-ox.url")}, err.Error(), infrastructure_errors.DatabaseException)
	}
	err = mongoClient.Connect()
	if err != nil {
		fmt.Printf("[Dependencies]: Error connecting to database | Error: [%s]", err.Error())
		//return nil, infrastructure_errors.New(map[string]interface{}{"database": config.GetString("database-ox.name")}, err.Error(), infrastructure_errors.DatabaseException)
	}
	return mongoClient, nil
}
