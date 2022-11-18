package getAllUserController

import (
	"context"
	"fmt"
	"github.com/spro80/golangCleanArchitecture/app/application/useCase/registerUserUseCase"
	"github.com/spro80/golangCleanArchitecture/app/domain/entity/user_entity"
	user_entities_interface "github.com/spro80/golangCleanArchitecture/app/domain/entity/user_entity/interfaces"
	"github.com/spro80/golangCleanArchitecture/app/infraestructure/web/models/request_models"
)

type ControllerGetAllUserInterface interface {
	HandlerGetAllUserController(ctx context.Context, requestUser *request_models.User) (user_entities_interface.UserEntityInterface, int, error)
}

type ControllerGetAllUserHandler struct {
	useCase registerUserUseCase.UseCaseRegisterUserInterface
}

func NewGetAllUserController(useCase registerUserUseCase.UseCaseRegisterUserInterface) *ControllerGetAllUserHandler {
	return &ControllerGetAllUserHandler{useCase}
}

//func (r *ControllerRegisterUserHandler) HandlerRegisterUserController(u *user.User) (*user.User, int, error) {
func (r *ControllerGetAllUserHandler) HandlerGetAllUserController(ctx context.Context, requestUser *request_models.User) (user_entities_interface.UserEntityInterface, int, error) {
	fmt.Println("\n [get_all_user_controller] Init in HandlerGetAllUserController")

	userEntityData := user_entity.NewUserEntity()
	userEntityData.SetRut(requestUser.Rut)
	userEntityData.SetFirstName(requestUser.FirstName)
	userEntityData.SetLastName(requestUser.LastName)
	userEntityData.SetEmail(requestUser.Email)
	userEntityData.SetUserName(requestUser.UserName)
	userEntityData.SetPassword(requestUser.Password)

	userEntity, statusCode, err := r.useCase.HandlerRegisterUserUseCase(ctx, userEntityData)
	if err != nil {
		fmt.Printf("\n [get_all_user_controller] User Rut: [%s] | Message Error from UseCase: [%s]", requestUser.Rut, err.Error())
		return nil, statusCode, err
	}

	fmt.Printf("\n [get_all_user_controller] userEntity: [%v]", userEntity)
	fmt.Printf("\n [get_all_user_controller] Resgister UseCase was called succesfully | user RUT: [%s]", requestUser.Rut)

	return userEntity, statusCode, nil
}
