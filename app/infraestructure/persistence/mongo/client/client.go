package client

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type MongoClientInterface interface {
	Connect() (*mongo.Client, error)
	Disconnect() error
	SetDatabase(dbName string) *mongo.Database
	GetCollection(collectionName string) *mongo.Collection
	Ping() error
}

type MongoClient struct {
	client *mongo.Client
	Ctx    context.Context
	db     *mongo.Database
}

func NewMongoClient(uri string) (*MongoClient, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		//log.Error(err.Error())
		return nil, err
	}

	return &MongoClient{client: client}, nil
}

func (m *MongoClient) Connect(dbName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := m.client.Connect(ctx)
	if err != nil {
		//log.Error("error connecting to database %s", err.Error())
		return err
	}
	m.SetDatabase(dbName)

	//log.Info("checking database ....")
	e := m.Ping()
	if e != nil {
		return e
	}
	//log.Info("mongo database connected successfully")
	return nil
}

func (m *MongoClient) Disconnect() error {
	e := m.Ping()
	if e != nil {
		//log.Error("database is not connected cannot disconnect")
		return e
	}
	//log.Info("disconnecting mongo database %s", m.db.Name())
	e = m.client.Disconnect(m.Ctx)
	if e != nil {
		//log.Error("an error occurred trying disconnect database %s", e)
		return e
	}
	//log.Info("database is disconnect successfully")
	return nil
}

func (m *MongoClient) Ping() error {
	e := m.client.Ping(m.Ctx, readpref.Primary())

	if e != nil {
		//log.Error("error in ping to database")
		return e
	}
	//log.Info("ping successfully to database")
	return nil

}

func (m *MongoClient) SetDatabase(dbName string) *mongo.Database {
	m.db = m.client.Database(dbName)
	return m.db
}

func (m *MongoClient) GetCollection(collectionName string) *mongo.Collection {
	return m.db.Collection(collectionName)
}
