package registerUserUseCase

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/spro80/golangCleanArchitecture/app/domain/entities"
)

type UseCaseRegisterUserInterface interface {
	HandlerRegisterUserUseCase(u *entities.User) (*entities.User, int, error)
}

type UseCaseRegisterUserHandler struct {
}

func NewRegisterUserUseCase() *UseCaseRegisterUserHandler {
	return &UseCaseRegisterUserHandler{}
}

func (r *UseCaseRegisterUserHandler) HandlerRegisterUserUseCase(u *entities.User) (*entities.User, int, error) {
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

	err := validateUser(&user)
	if err != nil {
		fmt.Printf("[register_user_use_case] Error in validation of Entity User : [%s]", err.Error())
		return &user, http.StatusBadRequest, err
	}

	/*gatewayUser := gateways.NewOrderGateway(&models.UserModel{}, &repository.UserRepository{})
	user, err := gatewayUser.SaveUser(&user)
	if err != nil {
		fmt.Println("Error")
		return err
	}*/

	fmt.Println("[register_user_use_case] End in HandlerRegisterUserUseCase")
	return &user, http.StatusCreated, nil
}

func validateUser(user *entities.User) error {
	var validate *validator.Validate

	validate = validator.New()
	err := validate.Struct(user)
	if err != nil {
		return err
	}
	return nil
}
