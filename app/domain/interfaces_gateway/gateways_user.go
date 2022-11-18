package interfaces

import (
	user_entities_interface "github.com/spro80/golangCleanArchitecture/app/domain/entity/user/interfaces"
)

type RepositoryGatewayInterface interface {
	FindUserByRut(rut string) (user_entities_interface.UserEntityInterface, error)
}
