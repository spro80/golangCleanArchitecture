package routes

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spro80/golangCleanArchitecture/app/domain/entity/user_entity"
	"github.com/spro80/golangCleanArchitecture/app/infraestructure/web/middlewares"
	"github.com/spro80/golangCleanArchitecture/app/infraestructure/web/models/request_models"
	"github.com/spro80/golangCleanArchitecture/app/interfaces/input/controllers/getAllUserController"
	"github.com/spro80/golangCleanArchitecture/app/shared/utils/response"
	"net/http"
)

type GetAllUserRouteInterface interface {
	HandlerGetAllUserRoute(serverContext echo.Context) error
}

type GetAllUserRouteHandler struct {
	controller     getAllUserController.ControllerGetAllUserInterface
	responseStruct response.ResponseInterface
}

type ResponseGetAllUser struct {
	Status            int    `json:"status"`
	StatusDescription string `json:"statusDescription"`
	Data              []user_entity.User
	Error             string `json:"error"`
}

func NewGetAllUserRoute(e *echo.Echo, controller getAllUserController.ControllerGetAllUserInterface, responseStruct response.ResponseInterface) *GetAllUserRouteHandler {
	h := &GetAllUserRouteHandler{controller: controller, responseStruct: responseStruct}
	e.GET("/getAllUser", h.HandlerGetAllUserRoute, middlewares.ContextMiddleWare)
	return h
}

func (r GetAllUserRouteHandler) HandlerGetAllUserRoute(serverContext echo.Context) error {

	fmt.Println("[get_all_user_route] Init in HandlerGetAllUserRoute")
	var response ResponseGetAllUser

	var errBind error
	userRequest := new(request_models.User)
	if errBind = serverContext.Bind(userRequest); errBind != nil {
		//response = r.response.CreateResponse(http.StatusBadRequest, "Get All User was called with error", "", errBind.Error())
		//response = r.responseStruct.HandlerCreateResponseSuccess(http.StatusBadRequest, "Register User was called with error", "", errBind.Error())
		return errBind
	}

	usersEntity, statusCode, errCtrl := r.controller.HandlerGetAllUserController(serverContext.Get("traceContext").(context.Context), userRequest)
	if errCtrl != nil {
		fmt.Printf("\n [get_all_user_route] Error: [%s]", errCtrl.Error())
	}
	fmt.Printf("\n [get_all_user_route] usersEntity: [%v]", usersEntity)
	fmt.Printf("\n [get_all_user_route] statusCode: [%d]", statusCode)

	var usersList []user_entity.User
	var user user_entity.User

	for pos, _ := range usersEntity {
		user.Rut = usersEntity[pos].GetRut()
		user.UserName = usersEntity[pos].GetUserName()
		user.Password = usersEntity[pos].GetPassword()
		user.Email = usersEntity[pos].GetFirstName()
		user.FirstName = usersEntity[pos].GetFirstName()
		user.LastName = usersEntity[pos].GetLastName()
		user.Valid = usersEntity[pos].GetValid()
		usersList = append(usersList, user)
	}

	response = CreateGetAllResponse(statusCode, "Get All User was called successfully", usersList, "")
	fmt.Println("\n [get_all_user_route] End in HandlerGetAllUserRoute")
	return serverContext.JSON(http.StatusOK, response)
}

//func CreateGetAllResponse(status int, statusDes string, data []interfaces.UserEntityInterface, errorDes string) ResponseGetAllUser {
func CreateGetAllResponse(status int, statusDes string, data []user_entity.User, errorDes string) ResponseGetAllUser {
	response := ResponseGetAllUser{
		Status:            status,
		StatusDescription: statusDes,
		Data:              data,
		Error:             errorDes,
	}
	return response
}
