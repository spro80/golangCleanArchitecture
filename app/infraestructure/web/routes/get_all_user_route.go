package routes

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spro80/golangCleanArchitecture/app/infraestructure/web/middlewares"
	"github.com/spro80/golangCleanArchitecture/app/infraestructure/web/models/request_models"
	"github.com/spro80/golangCleanArchitecture/app/interfaces/input/controllers/getAllUserController"
	"net/http"
)

type GetAllUserRouteInterface interface {
	HandlerGetAllUserRoute(serverContext echo.Context) error
}

type GetAllUserRouteHandler struct {
	controller getAllUserController.ControllerGetAllUserHandler
	response   Response
}

type ResponseGetAllUser struct {
	Status            int    `json:"status"`
	StatusDescription string `json:"statusDescription"`
	Data              string `json:"data"`
	Error             string `json:"error"`
}

func NewGetAllUserRoute(e *echo.Echo, controller getAllUserController.ControllerGetAllUserHandler) *GetAllUserRouteHandler {
	h := &GetAllUserRouteHandler{controller: controller}
	e.POST("/getAllUser", h.HandlerGetAllUserRoute, middlewares.ContextMiddleWare)
	return h
}

func (r GetAllUserRouteHandler) HandlerGetAllUserRoute(serverContext echo.Context) error {

	fmt.Println("[get_all_user_route] Init in HandlerGetAllUserRoute")
	var response Response

	var errBind error
	userRequest := new(request_models.User)
	if errBind = serverContext.Bind(userRequest); errBind != nil {
		response = r.response.CreateResponse(http.StatusBadRequest, "Get All User was called with error", "", errBind.Error())
		return errBind
	}

	user, statusCode, errCtrl := r.controller.HandlerGetAllUserController(serverContext.Get("traceContext").(context.Context), userRequest)
	fmt.Printf("\n [get_all_user_route] user: [%v]", user)
	fmt.Printf("\n [get_all_user_route] statusCode: [%d]", statusCode)

	if errCtrl != nil {
		fmt.Printf("\n [get_all_user_route] Error: [%s]", errCtrl.Error())
	}
	fmt.Println("\n [get_all_user_route] End in HandlerGetAllUserRoute")

	response = r.response.CreateGetAllResponse(statusCode, "Get All User was called successfully", "", "")
	return serverContext.JSON(http.StatusOK, response)
}

func (r Response) CreateGetAllResponse(status int, statusDes string, data string, errorDes string) Response {
	response := Response{
		Status:            status,
		StatusDescription: statusDes,
		Data:              data,
		Error:             errorDes,
	}
	return response
}
