package controllers_add_user_controller

import (
	"context"
	"fmt"
	use_case_user_add "github.com/spro80/golangCleanArchitecture/app/application/useCase/user_add_use_case"
	"github.com/spro80/golangCleanArchitecture/app/domain/entity/user_entity"
	user_entities_interface "github.com/spro80/golangCleanArchitecture/app/domain/entity/user_entity/interfaces"
	user_input_add_v1_request "github.com/spro80/golangCleanArchitecture/app/interfaces/input/source/api/user_input/add/v1/request"
)

type UserAddControllerInterface interface {
	HandlerUserAddController(ctx context.Context, requestUser *user_input_add_v1_request.UserAddRequest) (user_entities_interface.UserEntityInterface, int, error)
}

type UserAddControllerHandler struct {
	useCase use_case_user_add.UserAddUseCaseInterface
}

func NewUserAddController(useCase use_case_user_add.UserAddUseCaseInterface) UserAddControllerInterface {
	return &UserAddControllerHandler{useCase}
}

func (r *UserAddControllerHandler) HandlerUserAddController(ctx context.Context, requestUser *user_input_add_v1_request.UserAddRequest) (user_entities_interface.UserEntityInterface, int, error) {
	fmt.Println("\n [add_user_controller] Init in HandlerRegisterUserController")

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

	userEntity, statusCode, err := r.useCase.HandlerUserAddUseCase(ctx, userEntityData)
	if err != nil {
		fmt.Printf("\n [add_user_controller] User Rut: [%s] | Message Error from UseCase: [%s]", requestUser.Rut, err.Error())
		return nil, statusCode, err
	}

	fmt.Printf("\n [add_user_controller] userEntity: [%v]", userEntity)
	fmt.Printf("\n [add_user_controller] End of HandlerRegisterUserController | user RUT: [%s]", requestUser.Rut)

	return userEntity, statusCode, nil
}

func (r *UserAddControllerHandler) createUserEntity(requestUser *user_input_add_v1_request.UserAddRequest) user_entities_interface.UserEntityInterface {

	userEntityData := user_entity.NewUserEntity()
	userEntityData.SetUserId(requestUser.UserId)
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
