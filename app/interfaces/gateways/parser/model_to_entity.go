package parser

import (
	user_entity "github.com/spro80/golangCleanArchitecture/app/domain/entity/user_entity"
	user_entity_interface "github.com/spro80/golangCleanArchitecture/app/domain/entity/user_entity/interfaces"
	"github.com/spro80/golangCleanArchitecture/app/infraestructure/mongo_client/models"
)

func UserModelToEntity(userModel *models.UserModel) user_entity_interface.UserEntityInterface {

	userEntity := user_entity.NewUserEntity()
	userEntity.SetUserId(userModel.UserId)
	userEntity.SetRut(userModel.Rut)
	userEntity.SetUserName(userModel.UserName)
	userEntity.SetPassword(userModel.Password)
	userEntity.SetEmail(userModel.Email)
	userEntity.SetFirstName(userModel.FirstName)
	userEntity.SetLastName(userModel.LastName)
	userEntity.SetValid(userModel.Valid)

	profileEntity := user_entity.NewProfileEntity()
	profileEntity.SetProfileId(666)
	profileEntity.SetProfileStatus(false)
	profileEntity.SetProfileDateInit("20666")
	profileEntity.SetProfileDateEnd("20666")
	profileEntity.SetProfileAllTime(false)

	userEntity.SetProfile(profileEntity)

	//Data profile
	//userEntity.SetProfileId(userModel.Profile.ProfileId)
	//userEntity.SetProfileStatus(userModel.Profile.ProfileStatus)
	//userEntity.SetProfileDateInit(userModel.Profile.ProfileDateInit)
	//userEntity.SetProfileDateEnd(userModel.Profile.ProfileDateEnd)
	//userEntity.SetProfileStatus(userModel.Profile.ProfileStatus)

	return userEntity
}
