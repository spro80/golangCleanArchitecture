package user_entity

import "github.com/spro80/golangCleanArchitecture/app/domain/entity/user_entity/interfaces"

type User struct {
	IdUser    string `json:"idUser" validate:"required"`
	Rut       string `json:"rut" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	UserName  string `json:"userName" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Valid     bool   `json:"valid" validate:"required"`
	Profile   user_entity_interface.ProfileEntityInterface
}

type Profile struct {
	ProfileId       int    `json:"profileId" validate:"required"`
	ProfileStatus   bool   `json:"profileStatus" validate:"required"`
	ProfileDateInit string `json:"profileDateInit" validate:"required"`
	ProfileDateEnd  string `json:"profileDateEnd" validate:"required"`
	ProfileAllTime  bool   `json:"profileAllTime" validate:"required"`
}
