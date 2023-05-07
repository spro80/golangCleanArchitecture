package source_user_input_get_all_v1

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spro80/golangCleanArchitecture/app/interfaces/input/controllers/user_get_all_controller"
	"github.com/spro80/golangCleanArchitecture/app/interfaces/input/source"
	user_input_get_all_v1_request "github.com/spro80/golangCleanArchitecture/app/interfaces/input/source/api/user_input/get/v1/request"
	user_input_get_all_v1_response "github.com/spro80/golangCleanArchitecture/app/interfaces/input/source/api/user_input/get/v1/response"
	shared_utils_response "github.com/spro80/golangCleanArchitecture/app/shared/utils/response"
	"net/http"
)

type fromApi struct {
	userGetAllController user_get_all_controller.UserGetAllControllerInterface
	response             shared_utils_response.ResponseInterface
}

func NewFromApi(userGetAllController user_get_all_controller.UserGetAllControllerInterface, response shared_utils_response.ResponseInterface /*, apiResponse input_source_api.ApiResponseInterface[user_input_get_all_v1_response.UsersResponse]*/) source.FromApiInterface {
	return &fromApi{userGetAllController, response}
}

func (i *fromApi) FromApi(serverContext echo.Context) error {
	fmt.Println("Init in FromApi")
	//var response response_models.ResponseModelSuccess
	//r.URL.Query().Get("userId")

	//userId := serverContext.Get("userId")
	userId := serverContext.QueryParam("userId")

	fmt.Printf("Init in FromApi userId: [%s]", userId)

	var errBind error
	var userRequest = user_input_get_all_v1_request.UserGetAllRequest{}
	if errBind = serverContext.Bind(&userRequest); errBind != nil {
		description := fmt.Sprintf("[user_input_get_all] Error: [%s]", errBind.Error())
		response := i.response.HandlerCreateResponseSuccess(http.StatusBadRequest, description, "", errBind.Error())
		return serverContext.JSON(http.StatusInternalServerError, response)
	}

	fmt.Printf("Init in FromApi userRequest: [%s]", userRequest.Rut)
	usersResponse, statusCode, errCtrl := i.userGetAllController.HandlerUserGetAllController(serverContext.Get("traceContext").(context.Context), userId)
	fmt.Printf("\n [user_input_get_all] user: [%v] | statusCode: [%d]", usersResponse, statusCode)

	if errCtrl != nil {
		description := fmt.Sprintf("[user_input_get_all] Error: [%s]", errCtrl.Error())
		fmt.Printf("\n [%s]", description)
		response := i.response.HandlerCreateResponseSuccess(statusCode, description, usersResponse, errCtrl.Error())
		return serverContext.JSON(http.StatusOK, response)

	}

	if len(usersResponse) == 0 {
		description := "There are not users in database"
		fmt.Printf("\n [user_input_get_all] %s", description)
		response := i.response.HandlerCreateResponseSuccess(statusCode, description, usersResponse, "")
		return serverContext.JSON(http.StatusOK, response)
	}

	/*dataResponse := map[string]interface{}{
		"data": usersResponse,
	}*/
	//return serverContext.JSON(http.StatusOK, dataResponse)

	description := fmt.Sprintf("There are %d users in DB", len(usersResponse))
	response := i.response.HandlerCreateResponseSuccess(statusCode, description, usersResponse, "")
	fmt.Println("\n [user_input_get_all] End in HandlerRegisterUserRoute")

	return serverContext.JSON(http.StatusOK, response)

}

type responseEnd struct {
	Data        []user_input_get_all_v1_response.UserResponse
	Description string
}
