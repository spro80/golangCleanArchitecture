package registerUserUseCase

import (
	"fmt"
	"github.com/go-playground/validator/v10"

	"github.com/spro80/golangCleanArchitecture/app/domain/entities"
)

var validate *validator.Validate

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
	fmt.Printf("[register_user_use_case] UserName: [%v]", u.UserName)
	fmt.Printf("[register_user_use_case] Password: [%v]", u.Password)

	user := entities.User{
		IdUser:    u.Rut,
		Rut:       u.Rut,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		UserName:  u.UserName,
		Password:  u.Password,
	}
	fmt.Println(user)

	validate = validator.New()
	err := validate.Struct(user)
	if err != nil {
		fmt.Println("[register_user_use_case] Error in validation in creation of Entity")
		return user, err
	}

	/*gatewayUser := gateways.NewOrderGateway(&models.UserModel{}, &repository.UserRepository{})
	user, err := gatewayUser.SaveUser(&user)
	if err != nil {
		fmt.Println("Error")
		return err
	}*/

	fmt.Println("[register_user_use_case] End in HandlerRegisterUserUseCase")
	return user, nil
}
