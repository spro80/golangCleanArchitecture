package registerUserUseCase

import (
	"context"
	"fmt"
	user_entities_interface "github.com/spro80/golangCleanArchitecture/app/domain/entity/user_entity/interfaces"
	"github.com/spro80/golangCleanArchitecture/app/domain/interfaces_gateway"
	"net/http"
)

type UseCaseRegisterUserInterface interface {
	HandlerRegisterUserUseCase(ctx context.Context, userEntity user_entities_interface.UserEntityInterface) (user_entities_interface.UserEntityInterface, int, error)
}

type UseCaseRegisterUserHandler struct {
	userGateway interfaces_gateway.RepositoryGatewayInterface
}

func NewRegisterUserUseCase(userGateway interfaces_gateway.RepositoryGatewayInterface) *UseCaseRegisterUserHandler {
	return &UseCaseRegisterUserHandler{userGateway}
}

func (ru *UseCaseRegisterUserHandler) HandlerRegisterUserUseCase(ctx context.Context, userEntity user_entities_interface.UserEntityInterface) (user_entities_interface.UserEntityInterface, int, error) {
	fmt.Printf("\n [register_user_use_case] Init in HandlerRegisterUserUseCase")

	userResponse, err := ru.userGateway.FindUserByRut(userEntity.GetRut())
	if err != nil {
		fmt.Printf("\n [register_user_use_case] Error in FindUserByRut:[%s] with error:[%s]", userResponse.GetRut(), err.Error())
		//TODO: create error from database
		return nil, http.StatusBadRequest, err
	}
	fmt.Printf("\n [register_user_use_case] userResponse:[%v]", userResponse)

	if userResponse.GetRut() == "" {
		userSave, errGateway := ru.userGateway.SaveUser(ctx, userEntity)
		if errGateway != nil {
			fmt.Printf("\n [register_user_use_case] Error in Save User | RUT:[%s]  | Messagge Error:[%s]", userResponse.GetRut(), errGateway.Error())
			//TODO: create error from database
			return nil, http.StatusBadRequest, err
		}
		fmt.Println("\n [register_user_use_case] User was saved successfully")
		return userSave, http.StatusCreated, nil
	} else {
		fmt.Println("\n [register_user_use_case] User was not saved. User already exists")
		return nil, http.StatusOK, nil
	}

	return nil, http.StatusOK, nil
}

/*
func validateUser(user *user_entity.User) error {
	var validate *validator.Validate

	validate = validator.New()
	err := validate.Struct(user)
	if err != nil {
		return err
	}
	return nil
}
*/
