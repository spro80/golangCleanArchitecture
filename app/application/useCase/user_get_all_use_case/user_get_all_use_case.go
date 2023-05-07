package user_get_all_use_case

import (
	"fmt"
	user_entities_interface "github.com/spro80/golangCleanArchitecture/app/domain/entity/user_entity/interfaces"
	"github.com/spro80/golangCleanArchitecture/app/domain/interfaces_gateway"
	"net/http"
)

import (
	"context"
)

type UserGetAllUseCaseInterface interface {
	HandlerUserGetAllUseCase(ctx context.Context) ([]user_entities_interface.UserEntityInterface, int, error)
}

type UserGetAllUseCaseHandler struct {
	userGateway interfaces_gateway.RepositoryGatewayInterface
}

func NewUserGetAllUseCase(userGateway interfaces_gateway.RepositoryGatewayInterface) *UserGetAllUseCaseHandler {
	return &UserGetAllUseCaseHandler{userGateway}
}

func (ru *UserGetAllUseCaseHandler) HandlerUserGetAllUseCase(ctx context.Context) ([]user_entities_interface.UserEntityInterface, int, error) {
	fmt.Printf("\n [user_get_all_use_case] Init in HandlerUserGetAllUseCase")

	usersEntity, err := ru.userGateway.FindAllUsers(ctx)
	if err != nil {
		fmt.Printf("\n [user_get_all_use_case] Error in called to FindAllUsers, error:[%s]", err.Error())
		//TODO: create error from database
		return nil, http.StatusInternalServerError, err
	}

	fmt.Printf("\n [user_get_all_use_case] Use case finished successfully")
	return usersEntity, http.StatusOK, nil
}
