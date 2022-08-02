package registerUserController

import (
	"fmt"

	"github.com/spro80/golangCleanArchitecture/app/application/useCase/registerUserUseCase"
	"github.com/spro80/golangCleanArchitecture/app/domain/entities"
)

type ControllerRegisterUserInterface interface {
	HandlerRegisterUserController(u *entities.User) (*entities.User, int, error)
}

type ControllerRegisterUserHandler struct {
	useCase registerUserUseCase.UseCaseRegisterUserInterface
}

func NewRegisterUserController(useCase registerUserUseCase.UseCaseRegisterUserInterface) *ControllerRegisterUserHandler {
	return &ControllerRegisterUserHandler{useCase: useCase}
}

func (r *ControllerRegisterUserHandler) HandlerRegisterUserController(u *entities.User) (*entities.User, int, error) {
	fmt.Println("[register_user_controller] Init in HandlerRegisterUserController")
	fmt.Printf("[register_user_controller] UserName: [%v]", u.UserName)

	fmt.Println("[register_user_controller] Calling HandlerRegisterUserUseCase")
	userEntity, statusCode, err := r.useCase.HandlerRegisterUserUseCase(u)
	if err != nil {
		fmt.Printf("[register_user_controller] Error: [%s]", err.Error())
		return userEntity, statusCode, err
	}

	fmt.Printf("[register_user_controller] userEntity: [%v]", userEntity)
	fmt.Printf("[register_user_controller] userEntity: [%d]", statusCode)
	fmt.Println("[register_user_controller] HandlerRegisterUserUseCase was called successfully")
	fmt.Println("[register_user_controller] End in HandlerRegisterUserController")

	return userEntity, statusCode, nil
}
