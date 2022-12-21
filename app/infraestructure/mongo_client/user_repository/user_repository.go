package user_repository

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/spro80/golangCleanArchitecture/app/infraestructure/mongo_client"
	"github.com/spro80/golangCleanArchitecture/app/infraestructure/mongo_client/collections"
	"github.com/spro80/golangCleanArchitecture/app/infraestructure/mongo_client/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

type UserRepositoryInterface interface {
	FindUserByRut(userRut string) (*models.UserModel, error)
	FindAllUsers(ctx context.Context) ([]models.UserModel, error)
	SaveUser(ctx context.Context, user *models.UserModel, contextSession ...context.Context) (*models.UserModel, error)
}

type UserRepository struct {
	collection mongo_client.CollectionInterface
	dbClient   mongo_client.MongoClientInterface
}

func NewUserRepository(database mongo_client.DatabaseInterface) UserRepositoryInterface {
	collection := database.Collection(collections.User)
	return &UserRepository{collection, database.Client()}
}

func (ur *UserRepository) FindUserByRut(userRut string) (*models.UserModel, error) {
	fmt.Printf("\n [user_repository][FindByUserId] Init in FindByUserId | User Rut: [%s] ", userRut)

	ctx := context.Background()
	cur := ur.collection.FindOne(ctx, bson.M{"rut": userRut})

	var user models.UserModel
	if err := cur.Decode(&user); err != nil {
		fmt.Printf("\n [user_repository][FindByUserId] Error finding userRut: [%s] | error: [%s]", userRut, err.Error())
		//TODO: create error generic
		//return &user, nil
	}

	fmt.Printf("\n [user_repository][FindByUserId] End FindByUserRut | User Rut:[%s]", userRut)
	return &user, nil
}

func (ur *UserRepository) FindAllUsers(ctx context.Context) ([]models.UserModel, error) {
	fmt.Printf("\n [user_repository][FindAllUsers] Init in FindAllUsers")

	limitRetryRelease := 100
	/*
		minutesAgo := 100
		limitRetryRelease := 100

		timeMinutesAgo := time.Duration(-minutesAgo) * time.Minute
		dateNow := time.Now()
		dateFrom := dateNow.Add(timeMinutesAgo).Format("2006-01-02T15:04:05")
		dateTo := dateNow.Format("2006-01-02T15:04:05")
	*/
	opts := options.Find().SetLimit(int64(limitRetryRelease))
	filter := bson.M{
		"valid": bson.M{"$eq": true},
	}
	cursor, errFind := ur.collection.Find(ctx, filter, opts)
	if errFind != nil {
		log.Error("[user_repository][FindAllUsers] Error in Find with message: [%s]", errFind.Error())
		return nil, errFind
		//return nil, infrastructure_errors.New(map[string]interface{}{"RetryRelease": "ErrorFindingShippingGroupsIdToRetryRelease"}, errFind.Error(), infrastructure_errors.DatabaseException)
	}

	var userListModels []models.UserModel
	errCursor := cursor.AllCursor(ctx, &userListModels)
	if errCursor != nil {
		log.Error("[user_repository][FindAllUsers] Error in Get AllCursor with message: [%s]", errCursor.Error())
		return nil, errCursor
		//return nil, infrastructure_errors.New(map[string]interface{}{"RetryRelease": "ErrorFindingShippingGroupsIdToRetryRelease"}, errCursor.Error(), infrastructure_errors.OrderStatusAllCursorException)
	}

	/*
		if len(userList) == 0 {
			fmt.Println("There is not Users in the DB")
		}*/

	fmt.Printf("\n [user_repository][FindAllUsers] End FindAllUsers")
	return userListModels, nil
}

func (ur *UserRepository) SaveUser(ctx context.Context, user *models.UserModel, contextSession ...context.Context) (*models.UserModel, error) {

	fmt.Printf("\n [user_repository][SaveUser] Init in SaveUser | User Rut: [%s] ", user.Rut)

	if contextSession != nil {
		ctx = contextSession[0]
	}
	id, err := ur.collection.InsertOne(ctx, user)
	if err != nil {
		fmt.Printf("\n [user_repository][SaveUser] Error saving new user | user Rut: [%s]", user.Rut)
		//TODO: create error generic
		return nil, err
	}

	idWithoutObjectId := strings.Replace(id, "ObjectID(\"", "", -1)
	idWithoutObjectId = strings.Replace(idWithoutObjectId, "\")", "", -1)
	insertId, err := primitive.ObjectIDFromHex(idWithoutObjectId)
	if err != nil {
		fmt.Printf("Error : [%s] ", err.Error())
		//TODO: create error generic
		return nil, err
	}

	user.ID = insertId
	fmt.Printf("\n [user_repository][SaveUser] User was saved succesfully | user Rut: [%s] ", user.Rut)
	return user, nil
}
