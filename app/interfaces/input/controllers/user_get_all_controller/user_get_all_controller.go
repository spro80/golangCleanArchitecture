package getAllUserController

import (
	"context"
	"fmt"
	"github.com/spro80/golangCleanArchitecture/app/application/useCase/user_get_all_use_case"
	user_entities_interface "github.com/spro80/golangCleanArchitecture/app/domain/entity/user_entity/interfaces"
	"github.com/spro80/golangCleanArchitecture/app/infraestructure/web/models/request_models"
)

type ControllerGetAllUserInterface interface {
	HandlerGetAllUserController(ctx context.Context, requestUser *request_models.User) ([]user_entities_interface.UserEntityInterface, int, error)
}

type ControllerGetAllUserHandler struct {
	useCase user_get_all_use_case.UseCaseGetAllUserInterface
}

func NewGetAllUserController(useCase user_get_all_use_case.UseCaseGetAllUserInterface) *ControllerGetAllUserHandler {
	return &ControllerGetAllUserHandler{useCase}
}

func (r *ControllerGetAllUserHandler) HandlerGetAllUserController(ctx context.Context, requestUser *request_models.User) ([]user_entities_interface.UserEntityInterface, int, error) {
	fmt.Println("\n [get_all_user_controller] Init in HandlerGetAllUserController")

	fmt.Println("\n [get_all_user_controller] Before to call HandlerGetAllUserUseCase")
	usersEntity, statusCode, err := r.useCase.HandlerGetAllUserUseCase(ctx)
	if err != nil {
		fmt.Printf("\n [get_all_user_controller] | Error from use case with message: [%s]", err.Error())
		return nil, statusCode, err
	}

	fmt.Printf("\n [get_all_user_controller] usersEntity: [%v]", usersEntity)
	fmt.Printf("\n [get_all_user_controller] Register UseCase was called succesfully")

	return usersEntity, statusCode, nil
}
