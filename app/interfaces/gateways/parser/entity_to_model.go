package parser

import (
	user_entities_interface "github.com/spro80/golangCleanArchitecture/app/domain/entity/user_entity/interfaces"
	"github.com/spro80/golangCleanArchitecture/app/infraestructure/mongo_client/models"
)

func UserEntityToModel(user user_entities_interface.UserEntityInterface) *models.UserModel {

	userModel := models.UserModel{}
	userModel.Rut = user.GetRut()
	userModel.UserName = user.GetUserName()
	userModel.Password = user.GetPassword()
	userModel.Email = user.GetEmail()
	userModel.FirstName = user.GetFirstName()
	userModel.LastName = user.GetLastName()
	userModel.Valid = user.GetValid()

	return &userModel
}
