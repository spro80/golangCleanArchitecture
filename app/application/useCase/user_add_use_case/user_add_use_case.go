package use_case_user_add

import (
	"context"
	"fmt"
	user_entities_interface "github.com/spro80/golangCleanArchitecture/app/domain/entity/user_entity/interfaces"
	"github.com/spro80/golangCleanArchitecture/app/domain/interfaces_gateway"
	"net/http"
)

type UserAddUseCaseInterface interface {
	HandlerUserAddUseCase(ctx context.Context, userEntity user_entities_interface.UserEntityInterface) (user_entities_interface.UserEntityInterface, int, error)
}

type UserAddUseCaseHandler struct {
	userGateway interfaces_gateway.RepositoryGatewayInterface
}

func NewUserAddUseCase(userGateway interfaces_gateway.RepositoryGatewayInterface) *UserAddUseCaseHandler {
	return &UserAddUseCaseHandler{userGateway}
}

func (ru *UserAddUseCaseHandler) HandlerUserAddUseCase(ctx context.Context, userEntity user_entities_interface.UserEntityInterface) (user_entities_interface.UserEntityInterface, int, error) {
	fmt.Printf("\n [user_add_use_case] Init in HandlerUserAddUseCase")

	userResponse, err := ru.userGateway.FindUserByRut(userEntity.GetRut())
	if err != nil {
		fmt.Printf("\n [user_add_use_case] Error in FindUserByRut:[%s] with error:[%s]", userResponse.GetRut(), err.Error())
		//TODO: create error from database
		return nil, http.StatusInternalServerError, err
	}

	fmt.Printf("\n [user_add_use_case] userResponse:[%v]", userResponse)
	if userResponse.GetRut() != "" {
		fmt.Println("\n [user_add_use_case] User already exists")
		return nil, http.StatusOK, nil
	} else {
		userSave, errGateway := ru.userGateway.SaveUser(ctx, userEntity)
		if errGateway != nil {
			fmt.Printf("\n [user_add_use_case] Error in Save User | RUT:[%s]  | Messagge Error:[%s]", userResponse.GetRut(), errGateway.Error())
			//TODO: create error from database
			return nil, http.StatusInternalServerError, err
		}

		fmt.Println("\n [user_add_use_case] User was saved successfully")
		return userSave, http.StatusCreated, nil
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
