package source_user_input_delete_v1

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spro80/golangCleanArchitecture/app/interfaces/input/controllers/user_delete_controller"
	"github.com/spro80/golangCleanArchitecture/app/interfaces/input/source"
	user_input_delete_v1_request "github.com/spro80/golangCleanArchitecture/app/interfaces/input/source/api/user_input/delete/v1/request"
	shared_utils_response "github.com/spro80/golangCleanArchitecture/app/shared/utils/response"
	"net/http"
)

type fromApi struct {
	userDeleteController user_delete_controller.UserDeleteControllerInterface
	response             shared_utils_response.ResponseInterface
}

func NewFromApi(userDeleteController user_delete_controller.UserDeleteControllerInterface, response shared_utils_response.ResponseInterface) source.FromApiInterface {
	return &fromApi{userDeleteController, response}
}

func (i *fromApi) FromApi(serverContext echo.Context) error {
	fmt.Println("Init in FromApi")

	//var response user_input_delete_v1_request

	var errBind error
	//userRequest := new(request_models.User)

	var userRequest = user_input_delete_v1_request.UserDeleteRequest{}

	if errBind = serverContext.Bind(&userRequest); errBind != nil {
		fmt.Println("Error in error: ", errBind.Error())
		//response = r.response.CreateResponse(http.StatusBadRequest, "Get All User was called with error", "", errBind.Error())
		//response = r.responseStruct.HandlerCreateResponseSuccess(http.StatusBadRequest, "Register User was called with error", "", errBind.Error())
		return errBind
	}

	fmt.Printf("[user_delete_route] userRequest: [%v]", userRequest)

	countDelete, statusCode, errCtrl := i.userDeleteController.HandlerUserDeleteController(serverContext.Get("traceContext").(context.Context), userRequest)
	if errCtrl != nil {
		fmt.Printf("\n [delete_user_route] Error: [%s]", errCtrl.Error())
	}
	fmt.Printf("\n [delete_user_route] countDelete: [%v]", countDelete)
	fmt.Printf("\n [delete_user_route] statusCode: [%d]", statusCode)

	if countDelete == 0 {
		fmt.Printf("\n [delete_user_route] The User don´t was deleted. RUT: [%s]", userRequest.Rut)
		//response = CreateDeleteResponse(statusCode, userRequest.Rut, "The User don´t was deleted", "")
		//return serverContext.JSON(http.StatusOK, response)
	}

	//response = CreateDeleteResponse(statusCode, userRequest.Rut, "The User was deleted successfully", "")
	//fmt.Println("\n [delete_user_route] End in HandlerDeleteUserRoute")
	//return serverContext.JSON(http.StatusOK, response)

	/*fmt.Println("\n [user_input_add] End in HandlerRegisterUserRoute")
	response := i.response.HandlerCreateResponseSuccess(statusCode, "User was created successfully", "", "")
	return serverContext.JSON(http.StatusOK, response)
	*/

	response := i.response.HandlerCreateResponseSuccess(statusCode, "User was deleted successfully", "", "")
	return serverContext.JSON(http.StatusOK, response)

}
