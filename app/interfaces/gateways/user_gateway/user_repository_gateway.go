package user_gateway

import (
	"context"
	"fmt"
	"github.com/spro80/golangCleanArchitecture/app/domain/entity/user_entity"
	user_entities_interface "github.com/spro80/golangCleanArchitecture/app/domain/entity/user_entity/interfaces"
	"github.com/spro80/golangCleanArchitecture/app/domain/interfaces_gateway"
	"github.com/spro80/golangCleanArchitecture/app/infraestructure/mongo_client/user_repository"
	"github.com/spro80/golangCleanArchitecture/app/interfaces/gateways/parser"
)

type RepositoryGateway struct {
	userRepository user_repository.UserRepositoryInterface
}

func NewRepositoryGateway(repository user_repository.UserRepositoryInterface) interfaces_gateway.RepositoryGatewayInterface {
	return &RepositoryGateway{repository}
}

func (g *RepositoryGateway) FindUserByRut(userRut string) (user_entities_interface.UserEntityInterface, error) {
	fmt.Printf("\n [user_gateway] Init in FindUserByRut | User Rut: [%s] ", userRut)

	userModel, err := g.userRepository.FindUserByRut(userRut)
	if err != nil {
		fmt.Printf("\n [user_gateway]: Error in called to Repository FindByUserId | User Rut: [%s] | Error with message: [%s] ", userRut, err.Error())
		//TODO: create error generic
		return nil, err
	}

	fmt.Printf("\n [user_gateway]: Repository FindByUserId was called succesfully | User Rut: [%s] ", userRut)
	return parser.UserModelToEntity(userModel), nil
}

func (g *RepositoryGateway) FindAllUsers(ctx context.Context) ([]user_entities_interface.UserEntityInterface, error) {
	fmt.Printf("\n [user_gateway][FindAllUsers] Init in FindAllUsers")

	userModel, err := g.userRepository.FindAllUsers(ctx)
	if err != nil {
		fmt.Printf("\n [user_gateway][FindAllUsers] Error in called to Repository FindAllUsers | Error with message: [%s] ", err.Error())
		//TODO: create error generic
		return nil, err
	}

	countUsers := len(userModel)
	if len(userModel) > 0 {
		fmt.Printf("Existen [%d] usuarios en en la base de datos.", countUsers)
	} else {
		fmt.Println("No se encontraron usuarios en la base de datos.")
	}

	var users []user_entities_interface.UserEntityInterface

	fmt.Println("antes de Iterar por el ciclo forrrr......")
	for pos, value := range userModel {
		fmt.Println("Iterando por el ciclo forrrr......")
		fmt.Println(pos, value)
		user := user_entity.NewUserEntity()
		user.SetRut(userModel[pos].Rut)
		user.SetUserName(userModel[pos].UserName)
		user.SetPassword(userModel[pos].Password)
		user.SetEmail(userModel[pos].Email)
		user.SetFirstName(userModel[pos].FirstName)
		user.SetLastName(userModel[pos].LastName)
		user.SetValid(userModel[pos].Valid)
		users = append(users, user)
	}

	fmt.Printf("\n [user_gateway][FindAllUsers] Repository FindAllUsers was called succesfully")
	//return parser.UserModelToEntity(userModel), nil
	return users, nil
}

func (g *RepositoryGateway) SaveUser(ctx context.Context, userEntity user_entities_interface.UserEntityInterface) (user_entities_interface.UserEntityInterface, error) {

	fmt.Printf("\n [user_gateway] Init in Save User | User Rut: [%s] ", userEntity.GetRut())
	userModel := parser.UserEntityToModel(userEntity)

	user, err := g.userRepository.SaveUser(ctx, userModel)
	if err != nil {
		fmt.Printf("\n [user_gateway] Error in called to repository SaveUser | userRut: [%s]  | Error with message: [%s]", userEntity.GetRut(), err.Error())
		//TODO: create error generic
		return nil, err
	}

	return parser.UserModelToEntity(user), nil
}
