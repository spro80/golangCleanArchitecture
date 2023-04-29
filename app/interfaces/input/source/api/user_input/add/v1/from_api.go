package source_user_input_add_v1

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	controllers_add_user_controller "github.com/spro80/golangCleanArchitecture/app/interfaces/input/controllers/user_add_controller"
	"github.com/spro80/golangCleanArchitecture/app/interfaces/input/source"
	user_input_add_v1_request "github.com/spro80/golangCleanArchitecture/app/interfaces/input/source/api/user_input/add/v1/request"
	shared_utils_response "github.com/spro80/golangCleanArchitecture/app/shared/utils/response"
	"net/http"
)

type fromApi struct {
	userAddController controllers_add_user_controller.UserAddControllerInterface
	response          shared_utils_response.ResponseInterface
}

func NewFromApi(userAddController controllers_add_user_controller.UserAddControllerInterface, response shared_utils_response.ResponseInterface) source.FromApiInterface {
	return &fromApi{userAddController, response}
}

func (i *fromApi) FromApi(serverContext echo.Context) error {
	fmt.Println("Init in FromApi")
	//var response response_models.ResponseModelSuccess

	var errBind error
	//userRequest := new(request_models.User)
	//userRequest := new(user_input_add_v1_request.UserAddRequest)
	var userRequest = user_input_add_v1_request.UserAddRequest{}
	if errBind = serverContext.Bind(&userRequest); errBind != nil {
		//response := i.response.HandlerCreateResponseSuccess(http.StatusBadRequest, "Register User was called with error", "", errBind.Error())
		return errBind
	}
	//fmt.Printf("userRequest: [%v] ", userRequest)

	user, statusCode, errCtrl := i.userAddController.HandlerUserAddController(serverContext.Get("traceContext").(context.Context), &userRequest)
	fmt.Printf("\n [user_input_add] user: [%v]", user)
	fmt.Printf("\n [user_input_add] statusCode: [%d]", statusCode)

	if errCtrl != nil {
		fmt.Printf("\n [user_input_add] Error: [%s]", errCtrl.Error())
	}

	if user == nil {
		fmt.Printf("\n [user_input_add] user: [%v]", user)
		response := i.response.HandlerCreateResponseSuccess(statusCode, "User already exist", "", "")
		return serverContext.JSON(http.StatusOK, response)
	}

	fmt.Println("\n [user_input_add] End in HandlerRegisterUserRoute")
	response := i.response.HandlerCreateResponseSuccess(statusCode, "User was created successfully", "", "")
	return serverContext.JSON(http.StatusOK, response)
}
