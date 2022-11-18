package registerUserController

import (
	"context"
	"fmt"
	"github.com/spro80/golangCleanArchitecture/app/application/useCase/registerUserUseCase"
	user_entities_interface "github.com/spro80/golangCleanArchitecture/app/domain/entity/user/interfaces"
	"github.com/spro80/golangCleanArchitecture/app/infraestructure/web/models/request_models"
)

type ControllerRegisterUserInterface interface {
	HandlerRegisterUserController(ctx context.Context, requestUser *request_models.User) (user_entities_interface.UserEntityInterface, int, error)
}

type ControllerRegisterUserHandler struct {
	useCase registerUserUseCase.UseCaseRegisterUserInterface
}

func NewRegisterUserController(useCase registerUserUseCase.UseCaseRegisterUserInterface) *ControllerRegisterUserHandler {
	return &ControllerRegisterUserHandler{useCase}
}

//func (r *ControllerRegisterUserHandler) HandlerRegisterUserController(u *user.User) (*user.User, int, error) {
func (r *ControllerRegisterUserHandler) HandlerRegisterUserController(ctx context.Context, requestUser *request_models.User) (user_entities_interface.UserEntityInterface, int, error) {
	fmt.Println("[register_user_controller] Init in HandlerRegisterUserController")
	fmt.Printf("[register_user_controller] [%s]", ctx)
	fmt.Printf("[register_user_controller] [%v]", requestUser)
	//fmt.Printf("[register_user_controller] UserName: [%v]", u.UserName)
	/*
		userInterface := user.NewUserEntity()

		fmt.Println("[register_user_controller] Calling HandlerRegisterUserUseCase")
		userEntity, statusCode, err := r.useCase.HandlerRegisterUserUseCase(ctx, userInterface)
		if err != nil {
			fmt.Printf("[register_user_controller] Error: [%s]", err.Error())
			return nil, statusCode, err
		}
	*/
	//fmt.Printf("[register_user_controller] userEntity: [%v]", userEntity)
	//fmt.Printf("[register_user_controller] userEntity: [%d]", statusCode)
	fmt.Println("[register_user_controller] HandlerRegisterUserUseCase was called successfully")
	fmt.Println("[register_user_controller] End in HandlerRegisterUserController")

	return nil, 200, nil
	//return userEntity, statusCode, nil
}
