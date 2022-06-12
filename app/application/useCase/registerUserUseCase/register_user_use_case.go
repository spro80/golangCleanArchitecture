package registerUserUseCase

import "fmt"

type RegisterUserUseCaseInterface interface {
	HandlerRegisterUserUseCase() error
}

type RegisterUserUseCaseHandler struct {
}

func NewRegisterUserUseCase() *RegisterUserUseCaseHandler {
	return &RegisterUserUseCaseHandler{}
}

func (r *RegisterUserUseCaseHandler) HandlerRegisterUserUseCase() error {
	fmt.Println("[register_user_use_case] Init in HandlerRegisterUserUseCase")

	fmt.Println("[register_user_use_case] End in HandlerRegisterUserUseCase")
	return nil
}
