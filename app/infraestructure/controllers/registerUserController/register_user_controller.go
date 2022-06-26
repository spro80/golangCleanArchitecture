package registerUserController

import (
	"fmt"

	"github.com/spro80/golangCleanArchitecture/app/application/useCase/registerUserUseCase"
	"github.com/spro80/golangCleanArchitecture/app/domain/entities"
)

type RegisterUserControllerInterface interface {
	HandlerRegisterUserController(u entities.User) error
}

type RegisterUserControllerHandler struct {
	useCase registerUserUseCase.RegisterUserUseCaseInterface
}

func NewRegisterUserController(useCase registerUserUseCase.RegisterUserUseCaseInterface) *RegisterUserControllerHandler {
	return &RegisterUserControllerHandler{useCase: useCase}
}

func (r *RegisterUserControllerHandler) HandlerRegisterUserController(u entities.User) error {
	fmt.Println("[register_user_controller] Init in HandlerRegisterUserController")
	fmt.Println("[register_user_controller] Calling HandlerRegisterUserUseCase")

	//fmt.Printf("[register_user_controller] UserName: [%v]", u.Username)
	//fmt.Printf("[register_user_controller] UserName: [%v]", u.Username)

	userEntity, err := r.useCase.HandlerRegisterUserUseCase(u)
	if err != nil {
		fmt.Printf("[register_user_controller] Error: [%s]", err.Error())
	}

	fmt.Println("[register_user_controller] userEntity: [%v]", userEntity)
	fmt.Println("[register_user_controller] HandlerRegisterUserUseCase was called successfully")
	fmt.Println("[register_user_controller] End in HandlerRegisterUserController")

	return nil
}
