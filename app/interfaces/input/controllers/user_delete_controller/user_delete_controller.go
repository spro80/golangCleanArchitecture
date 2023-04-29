package user_delete_controller

import (
	"context"
	"fmt"
	"github.com/spro80/golangCleanArchitecture/app/application/useCase/user_delete_use_case"
	user_input_delete_v1_request "github.com/spro80/golangCleanArchitecture/app/interfaces/input/source/api/user_input/delete/v1/request"
)

type UserDeleteControllerInterface interface {
	HandlerUserDeleteController(ctx context.Context, requestUser user_input_delete_v1_request.UserDeleteRequest) (int64, int, error)
}

type UserDeleteControllerHandler struct {
	useCase user_delete_use_case.UserDeleteUseCaseInterface
}

func NewUserDeleteController(useCase user_delete_use_case.UserDeleteUseCaseInterface) *UserDeleteControllerHandler {
	return &UserDeleteControllerHandler{useCase}
}

func (r *UserDeleteControllerHandler) HandlerUserDeleteController(ctx context.Context, requestUser user_input_delete_v1_request.UserDeleteRequest) (int64, int, error) {
	fmt.Println("\n [user_delete_controller] Init in HandlerUserDeleteController")

	countDelete, statusCode, err := r.useCase.HandlerUserDeleteUseCase(ctx, requestUser.Rut)
	if err != nil {
		fmt.Printf("\n [user_delete_controller] | Error from use case with message: [%s]", err.Error())
		return 0, statusCode, err
	}

	fmt.Printf("\n [user_delete_controller] countDelete: [%v]", countDelete)
	fmt.Printf("\n [user_delete_controller] Delete User UseCase was called succesfully | user RUT: [%s]", requestUser.Rut)

	return countDelete, statusCode, nil
}
