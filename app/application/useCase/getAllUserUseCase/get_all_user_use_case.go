package getAllUserUseCase

import (
	"context"
	"fmt"
	user_entities_interface "github.com/spro80/golangCleanArchitecture/app/domain/entity/user_entity/interfaces"
	"github.com/spro80/golangCleanArchitecture/app/domain/interfaces_gateway"
	"net/http"
)

type UseCaseGetAllUserInterface interface {
	HandlerGetAllUserUseCase(ctx context.Context) ([]user_entities_interface.UserEntityInterface, int, error)
}

type UseCaseGetAllUserHandler struct {
	userGateway interfaces_gateway.RepositoryGatewayInterface
}

func NewGetAllUserUseCase(userGateway interfaces_gateway.RepositoryGatewayInterface) *UseCaseGetAllUserHandler {
	return &UseCaseGetAllUserHandler{userGateway}
}

func (ru *UseCaseGetAllUserHandler) HandlerGetAllUserUseCase(ctx context.Context) ([]user_entities_interface.UserEntityInterface, int, error) {
	fmt.Printf("\n [get_all_user_use_case] Init in HandlerGetAllUserUseCase")

	userResponse, err := ru.userGateway.FindAllUsers(ctx)
	if err != nil {
		fmt.Printf("\n [get_all_user_use_case] Error with message:[%s]", err.Error())
		//TODO: create error from database
		return nil, http.StatusBadRequest, err
	}

	fmt.Printf("\n [get_all_user_use_case] userResponse:[%v]", userResponse)
	return userResponse, http.StatusOK, nil
}
