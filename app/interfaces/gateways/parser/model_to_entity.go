package parser

import (
	user_entity "github.com/spro80/golangCleanArchitecture/app/domain/entity/user_entity"
	user_entity_interface "github.com/spro80/golangCleanArchitecture/app/domain/entity/user_entity/interfaces"
	"github.com/spro80/golangCleanArchitecture/app/infraestructure/mongo_client/models"
)

func UserModelToEntity(userModel *models.UserModel) user_entity_interface.UserEntityInterface {

	userEntity := user_entity.NewUserEntity()
	userEntity.SetRut(userModel.Rut)
	userEntity.SetUserName(userModel.UserName)
	userEntity.SetPassword(userModel.Password)
	userEntity.SetEmail(userModel.Email)
	userEntity.SetFirstName(userModel.FirstName)
	userEntity.SetLastName(userModel.LastName)
	userEntity.SetValid(userModel.Valid)

	return userEntity
}
