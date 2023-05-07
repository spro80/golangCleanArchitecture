package user_get_all_controller

import (
	"context"
	"fmt"
	"github.com/spro80/golangCleanArchitecture/app/application/useCase/user_get_all_use_case"
	"github.com/spro80/golangCleanArchitecture/app/interfaces/gateways/parser"
	user_input_get_all_v1_response "github.com/spro80/golangCleanArchitecture/app/interfaces/input/source/api/user_input/get/v1/response"
)

type UserGetAllControllerInterface interface {
	HandlerUserGetAllController(ctx context.Context, userId string) ([]user_input_get_all_v1_response.UserResponse, int, error)
}

type UserGetAllControllerHandler struct {
	useCase                    user_get_all_use_case.UserGetAllUseCaseInterface
	parserUserEntityToResponse parser.UserEntityToUserResponseInterface
}

func NewUserGetAllController(useCase user_get_all_use_case.UserGetAllUseCaseInterface, parserUserEntityToResponse parser.UserEntityToUserResponseInterface) *UserGetAllControllerHandler {
	return &UserGetAllControllerHandler{useCase, parserUserEntityToResponse}
}

func (r *UserGetAllControllerHandler) HandlerUserGetAllController(ctx context.Context, userId string) ([]user_input_get_all_v1_response.UserResponse, int, error) {
	fmt.Printf("\n [user_get_all_controller] Init in HandlerUserGetAllController | userId: [%s]", userId)

	fmt.Println("\n [user_get_all_controller] Before to call HandlerUserGetAllUseCase")
	usersEntity, statusCode, err := r.useCase.HandlerUserGetAllUseCase(ctx, userId)
	if err != nil {
		fmt.Printf("\n [user_get_all_controller] | Error from user get all use case with error: [%s]", err.Error())
		//return nil, statusCode, err
	}

	userResponse := r.parserUserEntityToResponse.UserEntityToUserResponseHandler(usersEntity)
	fmt.Printf("\n [user_get_all_controller] usersEntity: [%v]", usersEntity)
	fmt.Printf("\n [user_get_all_controller] User get all use case was called succesfully")

	return userResponse, statusCode, nil
}
