package getAllUserController

import (
	"context"
	"fmt"
	"github.com/spro80/golangCleanArchitecture/app/application/useCase/getAllUserUseCase"
	user_entities_interface "github.com/spro80/golangCleanArchitecture/app/domain/entity/user_entity/interfaces"
	"github.com/spro80/golangCleanArchitecture/app/infraestructure/web/models/request_models"
)

type ControllerGetAllUserInterface interface {
	HandlerGetAllUserController(ctx context.Context, requestUser *request_models.User) ([]user_entities_interface.UserEntityInterface, int, error)
}

type ControllerGetAllUserHandler struct {
	useCase getAllUserUseCase.UseCaseGetAllUserInterface
}

func NewGetAllUserController(useCase getAllUserUseCase.UseCaseGetAllUserInterface) *ControllerGetAllUserHandler {
	return &ControllerGetAllUserHandler{useCase}
}

func (r *ControllerGetAllUserHandler) HandlerGetAllUserController(ctx context.Context, requestUser *request_models.User) ([]user_entities_interface.UserEntityInterface, int, error) {
	fmt.Println("\n [get_all_user_controller] Init in HandlerGetAllUserController")

	usersEntity, statusCode, err := r.useCase.HandlerGetAllUserUseCase(ctx)
	if err != nil {
		fmt.Printf("\n [get_all_user_controller] | Error from use case with message: [%s]", err.Error())
		return nil, statusCode, err
	}

	fmt.Printf("\n [get_all_user_controller] usersEntity: [%v]", usersEntity)
	fmt.Printf("\n [get_all_user_controller] Resgister UseCase was called succesfully | user RUT: [%s]", requestUser.Rut)

	return usersEntity, statusCode, nil
}
