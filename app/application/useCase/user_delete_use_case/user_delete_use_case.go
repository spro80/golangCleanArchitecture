package user_delete_use_case

import (
	"context"
	"fmt"
	"github.com/spro80/golangCleanArchitecture/app/domain/interfaces_gateway"
	"net/http"
)

type UserDeleteUseCaseInterface interface {
	HandlerUserDeleteUseCase(ctx context.Context, rut string) (int64, int, error)
}

type UserDeleteUseCaseHandler struct {
	userGateway interfaces_gateway.RepositoryGatewayInterface
}

func NewUserDeleteUseCase(userGateway interfaces_gateway.RepositoryGatewayInterface) *UserDeleteUseCaseHandler {
	return &UserDeleteUseCaseHandler{userGateway}
}

func (du *UserDeleteUseCaseHandler) HandlerUserDeleteUseCase(ctx context.Context, rut string) (int64, int, error) {
	fmt.Printf("\n [user_delete_use_case] Init in HandlerUserDeleteUseCase")
	fmt.Printf("\n [user_delete_use_case] Init in HandlerUserDeleteUseCase rut:[%v]", rut)

	countDelete, err := du.userGateway.DeleteUserByRut(ctx, rut)
	if err != nil {
		fmt.Printf("\n [user_delete_use_case] Error with message:[%s]", err.Error())
		//TODO: create error from database
		return 0, http.StatusBadRequest, err
	}

	fmt.Printf("\n [user_delete_use_case] countDelete:[%v]", countDelete)
	return countDelete, http.StatusOK, nil
}
