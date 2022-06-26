package registerUserUseCase

import (
	"fmt"

	"github.com/spro80/golangCleanArchitecture/app/domain/entities"
)

type RegisterUserUseCaseInterface interface {
	HandlerRegisterUserUseCase(u entities.User) (entities.User, error)
}

type RegisterUserUseCaseHandler struct {
}

func NewRegisterUserUseCase() *RegisterUserUseCaseHandler {
	return &RegisterUserUseCaseHandler{}
}

func (r *RegisterUserUseCaseHandler) HandlerRegisterUserUseCase(u entities.User) (entities.User, error) {
	fmt.Println("[register_user_use_case] Init in HandlerRegisterUserUseCase")

	fmt.Printf("[register_user_use_case] Rut: [%v]", u.Rut)
	fmt.Printf("[register_user_use_case] FirstName: [%v]", u.FirstName)
	fmt.Printf("[register_user_use_case] LastName: [%v]", u.LastName)
	fmt.Printf("[register_user_use_case] Email: [%v]", u.Email)
	fmt.Printf("[register_user_use_case] Password: [%v]", u.Password)

	user := entities.User{
		IdUser:    "14515778",
		Rut:       "145157781",
		FirstName: "michael",
		LastName:  "clark",
		Email:     "mclark@gmail.com",
		UserName:  "mclark",
		Password:  "123456",
	}
	fmt.Println(user)

	/*gatewayUser := gateways.NewOrderGateway(&models.UserModel{}, &repository.UserRepository{})
	user, err := gatewayUser.SaveUser(&user)
	if err != nil {
		fmt.Println("Error")
		return err
	}*/

	fmt.Println("[register_user_use_case] End in HandlerRegisterUserUseCase")
	return user, nil
}
