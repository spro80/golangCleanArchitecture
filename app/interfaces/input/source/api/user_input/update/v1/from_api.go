package source_user_input_update_v1

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spro80/golangCleanArchitecture/app/interfaces/input/controllers/user_update_controller"
	"github.com/spro80/golangCleanArchitecture/app/interfaces/input/source"
	user_input_update_v1_request "github.com/spro80/golangCleanArchitecture/app/interfaces/input/source/api/user_input/update/v1/request"
	shared_utils_response "github.com/spro80/golangCleanArchitecture/app/shared/utils/response"
	"net/http"
)

type fromApi struct {
	userUpdateController user_update_controller.UserUpdateControllerInterface
	response             shared_utils_response.ResponseInterface
}

func NewFromApi(userUpdateController user_update_controller.UserUpdateControllerInterface, response shared_utils_response.ResponseInterface) source.FromApiInterface {
	return &fromApi{userUpdateController, response}
}

func (i *fromApi) FromApi(serverContext echo.Context) error {
	fmt.Println("Init in FromApi")
	//var response response_models.ResponseModelSuccess

	var errBind error
	var userRequest = user_input_update_v1_request.UserUpdateRequest{}
	if errBind = serverContext.Bind(&userRequest); errBind != nil {
		//response := i.response.HandlerCreateResponseSuccess(http.StatusBadRequest, "Register User was called with error", "", errBind.Error())
		return errBind
	}
	//fmt.Printf("userRequest: [%v] ", userRequest)

	user, statusCode, errCtrl := i.userUpdateController.HandlerUserUpdateController(serverContext.Get("traceContext").(context.Context), &userRequest)
	fmt.Printf("\n [user_update_input] user: [%v]", user)
	fmt.Printf("\n [user_update_input] statusCode: [%d]", statusCode)

	if errCtrl != nil {
		fmt.Printf("\n [user_update_input] Error: [%s]", errCtrl.Error())
	}

	if user == nil {
		fmt.Printf("\n [user_update_input] user: [%v]", user)
		response := i.response.HandlerCreateResponseSuccess(statusCode, "User already exist", "", "")
		return serverContext.JSON(http.StatusOK, response)
	}

	fmt.Println("\n [user_update_input] End in HandlerUserUpdateController")
	response := i.response.HandlerCreateResponseSuccess(statusCode, "User was updated successfully", "", "")
	return serverContext.JSON(http.StatusOK, response)
}
