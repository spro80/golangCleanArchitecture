package parser

import (
	user_entity_interface "github.com/spro80/golangCleanArchitecture/app/domain/entity/user_entity/interfaces"
	user_input_get_all_v1_response "github.com/spro80/golangCleanArchitecture/app/interfaces/input/source/api/user_input/get/v1/response"
)

type UserEntityToUserResponseInterface interface {
	UserEntityToUserResponseHandler(userInterface []user_entity_interface.UserEntityInterface) []user_input_get_all_v1_response.UserResponse
}

type userEntityToUserResponse struct {
}

func NewUserEntityToUserResponse() *userEntityToUserResponse {
	return &userEntityToUserResponse{}
}

func (r *userEntityToUserResponse) UserEntityToUserResponseHandler(userInterface []user_entity_interface.UserEntityInterface) []user_input_get_all_v1_response.UserResponse {

	userResponse := user_input_get_all_v1_response.UserResponse{}
	var usersResponse []user_input_get_all_v1_response.UserResponse

	for k, _ := range userInterface {
		userResponse.UserId = userInterface[k].GetUserId()
		userResponse.Rut = userInterface[k].GetRut()
		userResponse.FirstName = userInterface[k].GetFirstName()
		userResponse.LastName = userInterface[k].GetLastName()
		userResponse.Email = userInterface[k].GetEmail()
		userResponse.UserName = userInterface[k].GetUserName()
		userResponse.Password = userInterface[k].GetPassword()
		userResponse.Valid = userInterface[k].GetValid()

		usersResponse = append(usersResponse, userResponse)
	}

	return usersResponse
}
