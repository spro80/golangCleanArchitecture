package user_update_use_case

import (
	"context"
	"fmt"
	user_entities_interface "github.com/spro80/golangCleanArchitecture/app/domain/entity/user_entity/interfaces"
	"github.com/spro80/golangCleanArchitecture/app/domain/interfaces_gateway"
	"net/http"
)

type UserUpdateUseCaseInterface interface {
	HandlerUserUpdateUseCase(ctx context.Context, userEntity user_entities_interface.UserEntityInterface) (user_entities_interface.UserEntityInterface, int, error)
}

type UserUpdateUseCaseHandler struct {
	userGateway interfaces_gateway.RepositoryGatewayInterface
}

func NewUserUpdateUseCase(userGateway interfaces_gateway.RepositoryGatewayInterface) *UserUpdateUseCaseHandler {
	return &UserUpdateUseCaseHandler{userGateway}
}

func (ue *UserUpdateUseCaseHandler) HandlerUserUpdateUseCase(ctx context.Context, userEntity user_entities_interface.UserEntityInterface) (user_entities_interface.UserEntityInterface, int, error) {
	fmt.Printf("\n [user_edit_use_case] Init in HandlerUserAddUseCase")

	userResponse, err := ue.userGateway.FindUserByRut(ctx, userEntity.GetRut())
	if err != nil {
		fmt.Printf("\n [user_edit_use_case] Error in FindUserByRut:[%s] with error:[%s]", userResponse.GetRut(), err.Error())
		//TODO: create error from database
		return nil, http.StatusInternalServerError, err
	}

	fmt.Printf("\n [user_update_use_case] userResponse:[%v]", userResponse)
	if userResponse.GetRut() != "" {
		fmt.Println("\n [user_update_use_case] User exist in DB. Calling to update.")

		userSave, errGateway := ue.userGateway.UserUpdate(ctx, userEntity)
		if errGateway != nil {
			fmt.Printf("\n [user_update_use_case] Error in Save User | RUT:[%s]  | Messagge Error:[%s]", userResponse.GetRut(), errGateway.Error())
			//TODO: create error from database
			return nil, http.StatusInternalServerError, err
		}

		fmt.Println("\n [user_update_use_case] User was updated successfully")
		return userSave, http.StatusCreated, nil

	} else {
		fmt.Println("\n [user_update_use_case] User do not exist in DB. Is not possible to update")

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
