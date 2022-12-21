package repository

import (
	"context"

	"github.com/spro80/golangCleanArchitecture/app/infraestructure/persistence/mongo/client"
	"github.com/spro80/golangCleanArchitecture/app/infraestructure/persistence/mongo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryInterface interface {
	Save()
}

type UserRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository(dbClient *client.MongoClient) *UserRepository {
	collection := dbClient.GetCollection("user")
	return &UserRepository{collection}
}

func (r *UserRepository) Save(user *models.UserModel) (*models.UserModel, error) {
	//log.Info("Saving into database")
	id, err := r.collection.InsertOne(context.Background(), user)
	if err != nil {
		return &models.UserModel{}, err
	}
	user.ID = id.InsertedID.(primitive.ObjectID)
	return user, nil
}
