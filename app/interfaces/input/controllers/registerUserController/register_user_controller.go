package registerUserController

import (
	"context"
	"fmt"
	"github.com/spro80/golangCleanArchitecture/app/application/useCase/registerUserUseCase"
	"github.com/spro80/golangCleanArchitecture/app/domain/entity/user_entity"
	user_entities_interface "github.com/spro80/golangCleanArchitecture/app/domain/entity/user_entity/interfaces"
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
	fmt.Println("\n [register_user_controller] Init in HandlerRegisterUserController")

	/*
		userEntityData := user_entity.NewUserEntity()
		userEntityData.SetRut(requestUser.Rut)
		userEntityData.SetFirstName(requestUser.FirstName)
		userEntityData.SetLastName(requestUser.LastName)
		userEntityData.SetEmail(requestUser.Email)
		userEntityData.SetUserName(requestUser.UserName)
		userEntityData.SetPassword(requestUser.Password)
		userEntityData.SetValid(requestUser.Valid)
	*/
	userEntityData := r.createUserEntity(requestUser)

	userEntity, statusCode, err := r.useCase.HandlerRegisterUserUseCase(ctx, userEntityData)
	if err != nil {
		fmt.Printf("\n [register_user_controller] User Rut: [%s] | Message Error from UseCase: [%s]", requestUser.Rut, err.Error())
		return nil, statusCode, err
	}

	fmt.Printf("\n [register_user_controller] userEntity: [%v]", userEntity)
	fmt.Printf("\n [register_user_controller] End of HandlerRegisterUserController | user RUT: [%s]", requestUser.Rut)

	return userEntity, statusCode, nil
}

func (r *ControllerRegisterUserHandler) createUserEntity(requestUser *request_models.User) user_entities_interface.UserEntityInterface {

	userEntityData := user_entity.NewUserEntity()
	userEntityData.SetRut(requestUser.Rut)
	userEntityData.SetFirstName(requestUser.FirstName)
	userEntityData.SetLastName(requestUser.LastName)
	userEntityData.SetEmail(requestUser.Email)
	userEntityData.SetUserName(requestUser.UserName)
	userEntityData.SetPassword(requestUser.Password)
	userEntityData.SetValid(requestUser.Valid)

	profile := user_entity.NewProfileEntity()
	profile.SetProfileId(requestUser.Profile.ProfileId)
	profile.SetProfileStatus(requestUser.Profile.ProfileStatus)
	profile.SetProfileDateInit(requestUser.Profile.ProfileDateInit)
	profile.SetProfileDateEnd(requestUser.Profile.ProfileDateEnd)
	profile.SetProfileAllTime(requestUser.Profile.ProfileAllTime)

	userEntityData.SetProfile(profile)
	fmt.Printf("\n [register_user_controller] userEntityData:[%v]", &userEntityData)

	return userEntityData
}
