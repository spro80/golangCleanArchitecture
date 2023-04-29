package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spro80/golangCleanArchitecture/app/domain/entity/user_entity"
	user_entities_interface "github.com/spro80/golangCleanArchitecture/app/domain/entity/user_entity/interfaces"
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
	Data              map[string]interface{}
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
	fmt.Printf("\n [get_all_user_route] len usersEntity: [%v]", len(usersEntity))
	fmt.Printf("\n [get_all_user_route] statusCode: [%d]", statusCode)

	//var usersList []user_entity.User
	var usersList []user_entities_interface.UserEntityInterface

	var user user_entity.User

	//var profile user_entity.Profile

	for pos, _ := range usersEntity {
		user.Rut = usersEntity[pos].GetRut()
		/*user.UserName = usersEntity[pos].GetUserName()
		user.Password = usersEntity[pos].GetPassword()
		user.Email = usersEntity[pos].GetFirstName()
		user.FirstName = usersEntity[pos].GetFirstName()
		user.LastName = usersEntity[pos].GetLastName()
		user.Valid = usersEntity[pos].GetValid()
		*/
		user2 := user_entity.NewUserEntity()
		user2.SetRut("14515771")

		profile := user_entity.NewProfileEntity()
		profile.SetProfileId(333)

		/*fmt.Printf("\n [get_all_user_route] before get profileId **************")
		profile.ProfileId = usersEntity[pos].GetProfile().GetProfileId()
		fmt.Printf("\n [get_all_user_route] after get profileId **************")
		profile.ProfileStatus = true
		fmt.Printf("\n [get_all_user_route] after ProfileStatus **************")
		profile.ProfileDateInit = "2023"
		profile.ProfileDateEnd = "2023"
		profile.ProfileAllTime = true
		*/

		fmt.Printf("\n [get_all_user_route] before user.Profile.SetProfileId /////")
		//profile.Profile.SetProfileId(222)

		//user.Profile.SetProfileId(111)

		user2.SetProfile(profile)

		//user.Profile = user_entity.Profile{}
		//user.Profile = new
		fmt.Printf("\n [get_all_user_route] after user.Profile.SetProfileId /////")
		/*
			fmt.Printf("\n [get_all_user_route] before1 **************")
			fmt.Printf("\n [get_all_user_route] profile.ProfileId: [%d] **************", profile.ProfileId)
			fmt.Printf("\n [get_all_user_route] user:[%v]", usersEntity[pos].GetRut())
			fmt.Printf("\n [get_all_user_route] user:[%v]", usersEntity[pos].GetValid())
			fmt.Printf("\n [get_all_user_route] GetProfileId:[%v]", usersEntity[pos].GetProfile())
		*/
		//fmt.Printf("\n [get_all_user_route] user:[%v]", usersEntity[pos].GetProfile().GetProfileId())

		fmt.Printf("\n [get_all_user_route] before2 **************")

		//profile.ProfileStatus = usersEntity[pos].GetProfile().GetProfileStatus()
		//profile.ProfileDateInit = usersEntity[pos].GetProfile().GetProfileDateInit()
		//profile.ProfileDateEnd = usersEntity[pos].GetProfile().GetProfileDateEnd()
		//profile.ProfileAllTime = usersEntity[pos].GetProfile().GetProfileAllTime()

		/*
			user.Profile.SetProfileId(usersEntity[pos].GetProfileId())
			user.Profile.SetProfileStatus(usersEntity[pos].GetProfileStatus())
			user.Profile.SetProfileDateInit(usersEntity[pos].GetProfileDateInit())
			user.Profile.SetProfileDateEnd(usersEntity[pos].GetProfileDateEnd())
			user.Profile.SetProfileAllTime(usersEntity[pos].GetProfileAllTime())
		*/

		usersList = append(usersList, user2)

	}

	fmt.Printf("usersList::::: [%v]", usersList[0].GetRut())
	fmt.Printf("usersList::::: [%v]", usersList[0].GetUserName())
	fmt.Printf("usersList::::: [%v]", usersList[0].GetProfile().GetProfileId())

	var mapResponse map[string]interface{}
	responseByte, err := json.Marshal(usersList)
	if err != nil {
		fmt.Printf("Error in Marshal")
		//return nil, err
	}
	err = json.Unmarshal(responseByte, &mapResponse)
	if err != nil {
		fmt.Printf("Error in UnMarshal")
		//return nil, err
	}
	fmt.Printf("mapResponse::::: [%v]", mapResponse)

	response = CreateGetAllResponse(statusCode, "Get All User was called successfully", mapResponse, "")
	fmt.Println("\n [get_all_user_route] End in HandlerGetAllUserRoute")
	return serverContext.JSON(http.StatusOK, response)
}

//func CreateGetAllResponse(status int, statusDes string, data []interfaces.UserEntityInterface, errorDes string) ResponseGetAllUser {
//func CreateGetAllResponse(status int, statusDes string, data []user_entity.User, errorDes string) ResponseGetAllUser {
//func CreateGetAllResponse(status int, statusDes string, data []user_entities_interface.UserEntityInterface, errorDes string) ResponseGetAllUser {
func CreateGetAllResponse(status int, statusDes string, mapResponse map[string]interface{}, errorDes string) ResponseGetAllUser {
	response := ResponseGetAllUser{
		Status:            status,
		StatusDescription: statusDes,
		Data:              mapResponse,
		Error:             errorDes,
	}
	return response
}
