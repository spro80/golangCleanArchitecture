package registerUserController

import (
	"fmt"

	"github.com/spro80/golangCleanArchitecture/app/application/useCase/registerUserUseCase"
)

type RegisterUserControllerInterface interface {
	HandlerRegisterUserController() error
}

type RegisterUserControllerHandler struct {
	useCase registerUserUseCase.RegisterUserUseCaseInterface
}

func NewRegisterUserController(useCase registerUserUseCase.RegisterUserUseCaseInterface) *RegisterUserControllerHandler {
	return &RegisterUserControllerHandler{useCase: useCase}
}

func (r *RegisterUserControllerHandler) HandlerRegisterUserController() error {
	fmt.Println("[register_user_controller] Init in HandlerRegisterUserController")
	fmt.Println("[register_user_controller] Calling HandlerRegisterUserUseCase")

	err := r.useCase.HandlerRegisterUserUseCase()
	if err != nil {
		fmt.Printf("[register_user_controller] Error: [%s]", err.Error())
	}

	fmt.Println("[register_user_controller] HandlerRegisterUserUseCase was called successfully")
	fmt.Println("[register_user_controller] End in HandlerRegisterUserController")

	return nil
}
