package routes

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spro80/golangCleanArchitecture/app/infraestructure/web/middlewares"
	"github.com/spro80/golangCleanArchitecture/app/infraestructure/web/models/request_models"
	"github.com/spro80/golangCleanArchitecture/app/interfaces/input/controllers/registerUserController"
	"github.com/spro80/golangCleanArchitecture/app/shared/utils/response"
	"net/http"
)

type RegisterUserRouteInterface interface {
	HandlerRegisterUserRoute(serverContext echo.Context) error
}

type RegisterUserRouteHandler struct {
	controller     registerUserController.ControllerRegisterUserInterface
	responseStruct response.ResponseInterface
}

/*
type Response struct {
	Status            int    `json:"status"`
	StatusDescription string `json:"statusDescription"`
	Data              string `json:"data"`
	Error             string `json:"error"`
}
*/

func NewRegisterUserRoute(e *echo.Echo, controller registerUserController.ControllerRegisterUserInterface, responseStruct response.ResponseInterface) *RegisterUserRouteHandler {
	h := &RegisterUserRouteHandler{controller: controller, responseStruct: responseStruct}
	e.POST("/registerUser", h.HandlerRegisterUserRoute, middlewares.ContextMiddleWare)
	return h
}

func (r RegisterUserRouteHandler) HandlerRegisterUserRoute(serverContext echo.Context) error {

	fmt.Println("[register_user_route] Init in HandlerRegisterUserRoute")
	var response response.ResponseStruct

	var errBind error
	userRequest := new(request_models.User)
	if errBind = serverContext.Bind(userRequest); errBind != nil {
		response = r.responseStruct.HandlerCreateResponseSuccess(http.StatusBadRequest, "Register User was called with error", "", errBind.Error())
		return errBind
	}

	fmt.Printf("[register_user_route]  userRequest: [%v] ", userRequest)

	user, statusCode, errCtrl := r.controller.HandlerRegisterUserController(serverContext.Get("traceContext").(context.Context), userRequest)
	fmt.Printf("\n [register_user_route] user: [%v]", user)
	fmt.Printf("\n [register_user_route] statusCode: [%d]", statusCode)

	if errCtrl != nil {
		fmt.Printf("\n [register_user_route] Error: [%s]", errCtrl.Error())
	}
	if user == nil {
		response = r.responseStruct.HandlerCreateResponseSuccess(statusCode, "User already exist", "", "")
		return serverContext.JSON(http.StatusOK, response)
	}

	fmt.Println("\n [register_user_route] End in HandlerRegisterUserRoute")

	//response = r.response.CreateResponse(statusCode, "Register User was called successfully", "", "")
	response = r.responseStruct.HandlerCreateResponseSuccess(statusCode, "Register User was called successfully", "", "")
	return serverContext.JSON(http.StatusOK, response)
}

/*
func (r Response) CreateResponse(status int, statusDes string, data string, errorDes string) Response {
	response := Response{
		Status:            status,
		StatusDescription: statusDes,
		Data:              data,
		Error:             errorDes,
	}
	return response
}*/
