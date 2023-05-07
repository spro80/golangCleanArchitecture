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

func (g *RepositoryGateway) FindUserByRut(ctx context.Context, userRut string) (user_entities_interface.UserEntityInterface, error) {
	fmt.Printf("\n [user_gateway] Init in FindUserByRut | User Rut: [%s] ", userRut)

	userModel, err := g.userRepository.FindUserByRut(ctx, userRut)
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
		fmt.Printf("[user_gateway][FindAllUsers] There are [%d] users in DB.", countUsers)
	} else {
		fmt.Println("[user_gateway][FindAllUsers] There are not users in DB.")
	}

	var users []user_entities_interface.UserEntityInterface
	//profile user_entities_interface.ProfileEntityInterface
	//users = parser.UserModelToEntity(userModel)
	//parser.OrderModelToEntity(orderModel)

	fmt.Println("[user_gateway][FindAllUsers] Antes de Iterar por el ciclo forrrr......")
	for pos, value := range userModel {
		fmt.Printf("\n [user_gateway][FindAllUsers] Iterando por el ciclo forrrr pos: [%d]", pos)
		fmt.Printf("\n")
		fmt.Println(pos, value)
		user := user_entity.NewUserEntity()
		user.SetRut(userModel[pos].Rut)
		user.SetUserName(userModel[pos].UserName)
		user.SetPassword(userModel[pos].Password)
		user.SetEmail(userModel[pos].Email)
		user.SetFirstName(userModel[pos].FirstName)
		user.SetLastName(userModel[pos].LastName)
		user.SetValid(userModel[pos].Valid)

		/*profile := user_entity.NewProfileEntity()
		profile.SetProfileId(userModel[pos].Profile.ProfileId)
		profile.SetProfileStatus(userModel[pos].Profile.ProfileStatus)
		profile.SetProfileDateInit(userModel[pos].Profile.ProfileDateInit)
		profile.SetProfileDateEnd(userModel[pos].Profile.ProfileDateEnd)
		profile.SetProfileAllTime(userModel[pos].Profile.ProfileAllTime)
		*/

		fmt.Println("[user_gateway][FindAllUsers] Create user_entity.NewProfileEntity......")
		profileEntity := user_entity.NewProfileEntity()

		fmt.Printf("\n [user_gateway][FindAllUsers] profileEntity.SetProfile value ProfileId: [%d]", userModel[pos].Profile.ProfileId)
		fmt.Println("[user_gateway][FindAllUsers] Before call profileEntity.SetProfile")
		profileEntity.SetProfileId(userModel[pos].Profile.ProfileId)
		fmt.Println("[user_gateway][FindAllUsers] After call profileEntity.SetProfile")
		//profileEntity.SetProfileStatus(false)
		//profileEntity.SetProfileDateInit("20666")
		//profileEntity.SetProfileDateEnd("20666")
		//profileEntity.SetProfileAllTime(false)

		fmt.Printf("[user_gateway][FindAllUsers] Iterando por el ciclo forrrr profileEntity: [%v]", profileEntity.GetProfileId())

		user.SetProfile(profileEntity)
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

func (g *RepositoryGateway) UserUpdate(ctx context.Context, userEntity user_entities_interface.UserEntityInterface) (user_entities_interface.UserEntityInterface, error) {

	fmt.Printf("\n [user_gateway][UserUpdate] Init in User Update | User Rut: [%s] ", userEntity.GetRut())
	userModel := parser.UserEntityToModel(userEntity)

	user, err := g.userRepository.UserUpdate(ctx, userModel)
	if err != nil {
		fmt.Printf("\n [user_gateway][UserUpdate] Error in called to repository UserUpdate | userRut: [%s] | Error with message: [%s]", userEntity.GetRut(), err.Error())
		//TODO: create error generic
		return nil, err
	}

	return parser.UserModelToEntity(user), nil
}

func (g *RepositoryGateway) DeleteUserByRut(ctx context.Context, userRut string) (int64, error) {
	fmt.Printf("\n [user_gateway] Init in DeleteUserByRut | User Rut: [%s] ", userRut)

	fmt.Printf("\n [user_gateway] Calling DeleteUserByRut | User Rut: [%s] ", userRut)
	countDelete, err := g.userRepository.DeleteUserByRut(ctx, userRut)
	if err != nil {
		fmt.Printf("\n [user_gateway]: Error in called to Repository DeleteUserByRut | User Rut: [%s] | Error with message: [%s] ", userRut, err.Error())
		//TODO: create error generic
		return 0, err
	}

	fmt.Printf("\n [user_gateway]: Repository DeleteUserByRut was called succesfully | User Rut: [%s] ", userRut)
	return countDelete, nil
}
