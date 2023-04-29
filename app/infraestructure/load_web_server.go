package infraestructure

import (
	"fmt"
	use_case_user_add "github.com/spro80/golangCleanArchitecture/app/application/useCase/user_add_use_case"
	"github.com/spro80/golangCleanArchitecture/app/application/useCase/user_delete_use_case"
	"github.com/spro80/golangCleanArchitecture/app/application/useCase/user_update_use_case"
	"github.com/spro80/golangCleanArchitecture/app/infraestructure/mongo_client"
	"github.com/spro80/golangCleanArchitecture/app/infraestructure/mongo_client/user_repository"
	"github.com/spro80/golangCleanArchitecture/app/infraestructure/web"
	"github.com/spro80/golangCleanArchitecture/app/interfaces/gateways/user_gateway"
	controllers_add_user_controller "github.com/spro80/golangCleanArchitecture/app/interfaces/input/controllers/user_add_controller"
	"github.com/spro80/golangCleanArchitecture/app/interfaces/input/controllers/user_delete_controller"
	"github.com/spro80/golangCleanArchitecture/app/interfaces/input/controllers/user_update_controller"
	source_user_input_add_v1 "github.com/spro80/golangCleanArchitecture/app/interfaces/input/source/api/user_input/add/v1"
	source_user_input_delete_v1 "github.com/spro80/golangCleanArchitecture/app/interfaces/input/source/api/user_input/delete/v1"
	source_user_input_update_v1 "github.com/spro80/golangCleanArchitecture/app/interfaces/input/source/api/user_input/update/v1"
	"github.com/spro80/golangCleanArchitecture/app/shared/config"
	shared_utils_response "github.com/spro80/golangCleanArchitecture/app/shared/utils/response"
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
	sharedUtilsResponse := shared_utils_response.NewResponse()

	//Load repository
	userRepository := user_repository.NewUserRepository(clientDatabase)

	//Load gateways
	userGateway := user_gateway.NewRepositoryGateway(userRepository)

	// Load UseCase
	// User UseCase
	userAddUseCase := use_case_user_add.NewUserAddUseCase(userGateway)
	userUpdateUseCase := user_update_use_case.NewUserUpdateUseCase(userGateway)
	userDeleteUseCase := user_delete_use_case.NewUserDeleteUseCase(userGateway)
	//getAllUserUseCase := getAllUserUseCase.NewGetAllUserUseCase(userGateway)
	//deleteUserUseCase := deleteUserUseCase2.NewDeleteUserUseCase(userGateway)

	//Load Controller
	userAddController := controllers_add_user_controller.NewUserAddController(userAddUseCase)
	userUpdateController := user_update_controller.NewUserUpdateController(userUpdateUseCase)
	userDeleteController := user_delete_controller.NewUserDeleteController(userDeleteUseCase)
	//templateCtrl := templateController.NewTemplateController(templateUseCase)
	//getAllUserCtrl := getAllUserController.NewGetAllUserController(getAllUserUseCase)
	//deleteUserCtrl := deleteUserController.NewDeleteUserController(deleteUserUseCase)

	//Load Input
	userAddInput := source_user_input_add_v1.NewFromApi(userAddController, sharedUtilsResponse)
	userUpdateInput := source_user_input_update_v1.NewFromApi(userUpdateController, sharedUtilsResponse)
	userDeleteInput := source_user_input_delete_v1.NewFromApi(userDeleteController, sharedUtilsResponse)

	//Calling initialize routes
	//ws.web.InitRoutes(responseStruct, templateCtrl, getAllUserCtrl, userAddController, deleteUserCtrl)
	//ws.web.InitRoutes(responseStruct, templateCtrl, userAddController)
	ws.web.InitRoutes(userAddInput, userUpdateInput, userDeleteInput)
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
