package interfaces_gateway

import (
	"context"
	user_entities_interface "github.com/spro80/golangCleanArchitecture/app/domain/entity/user_entity/interfaces"
)

type RepositoryGatewayInterface interface {
	FindAllUsers(ctx context.Context) ([]user_entities_interface.UserEntityInterface, error)
	FindUserByRut(ctx context.Context, rut string) (user_entities_interface.UserEntityInterface, error)
	SaveUser(ctx context.Context, user user_entities_interface.UserEntityInterface) (user_entities_interface.UserEntityInterface, error)
	UserUpdate(ctx context.Context, user user_entities_interface.UserEntityInterface) (user_entities_interface.UserEntityInterface, error)
	DeleteUserByRut(ctx context.Context, rut string) (int64, error)
}
