package user_update_controller

import (
	"context"
	"fmt"
	"github.com/spro80/golangCleanArchitecture/app/application/useCase/user_update_use_case"
	"github.com/spro80/golangCleanArchitecture/app/domain/entity/user_entity"
	user_entities_interface "github.com/spro80/golangCleanArchitecture/app/domain/entity/user_entity/interfaces"
	user_input_update_v1_request "github.com/spro80/golangCleanArchitecture/app/interfaces/input/source/api/user_input/update/v1/request"
)

type UserUpdateControllerInterface interface {
	HandlerUserUpdateController(ctx context.Context, requestUser *user_input_update_v1_request.UserUpdateRequest) (user_entities_interface.UserEntityInterface, int, error)
}

type UserUpdateControllerHandler struct {
	useCase user_update_use_case.UserUpdateUseCaseInterface
}

func NewUserUpdateController(useCase user_update_use_case.UserUpdateUseCaseInterface) UserUpdateControllerInterface {
	return &UserUpdateControllerHandler{useCase}
}

func (uu *UserUpdateControllerHandler) HandlerUserUpdateController(ctx context.Context, requestUser *user_input_update_v1_request.UserUpdateRequest) (user_entities_interface.UserEntityInterface, int, error) {
	fmt.Println("\n [user_update_controller] Init in HandlerUserUpdateController")

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
	userEntityData := uu.createUserEntity(requestUser)

	userEntity, statusCode, err := uu.useCase.HandlerUserUpdateUseCase(ctx, userEntityData)
	if err != nil {
		fmt.Printf("\n [user_update_controller] User Rut: [%s] | Message Error from UseCase: [%s]", requestUser.Rut, err.Error())
		return nil, statusCode, err
	}

	fmt.Printf("\n [user_update_controller] userEntity: [%v]", userEntity)
	fmt.Printf("\n [user_update_controller] End of HandlerRegisterUserController | user RUT: [%s]", requestUser.Rut)

	return userEntity, statusCode, nil
}

func (r *UserUpdateControllerHandler) createUserEntity(requestUser *user_input_update_v1_request.UserUpdateRequest) user_entities_interface.UserEntityInterface {

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
	fmt.Printf("\n [user_update_controller] userEntityData:[%v]", &userEntityData)

	return userEntityData
}
