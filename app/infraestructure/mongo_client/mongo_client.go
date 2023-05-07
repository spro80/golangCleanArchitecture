package mongo_client

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"reflect"
)

type DatabaseInterface interface {
	Collection(name string) CollectionInterface
	Client() MongoClientInterface
}

type CollectionInterface interface {
	Find(ctx context.Context, filter interface{}, opts *options.FindOptions) (CursorInterface, error)
	FindOne(interface{}, interface{}) SingleResultInterface
	InsertOne(ctx interface{}, param interface{}) (string, error)
	UpdateOne(ctx interface{}, param interface{}, update interface{}) (int, error)
	DeleteOne(ctx interface{}, filter interface{}) (int64, error)
	UpsertOne(ctx, filter, update interface{}) (string, int64, int64, error)
	UpdateMany(ctx interface{}, filter interface{}, update interface{}) (int, int, error)
}

type SingleResultInterface interface {
	Decode(v interface{}) error
}

type CursorInterface interface {
	DecodeCursor(v interface{}) error
	AllCursor(ctx context.Context, result interface{}) error
}

type MongoClientInterface interface {
	Database(string) (DatabaseInterface, error)
	Disconnect() error
	Connect() error
	StartSession() (mongo.Session, error)
	UseSession(ctx context.Context, fn func(sessCtx mongo.SessionContext) error) error
}

type mongoClient struct {
	cl *mongo.Client
}

type mongoDatabase struct {
	db *mongo.Database
}
type mongoCollection struct {
	coll *mongo.Collection
}

type mongoSingleResult struct {
	sr *mongo.SingleResult
}

type MongoCursor struct {
	cr *mongo.Cursor
}

type mongoSession struct {
	s mongo.Session
}

func NewClient(uri string) (MongoClientInterface, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	return &mongoClient{cl: client}, err
}

func (mc *mongoClient) UseSession(ctx context.Context, fn func(sessCtx mongo.SessionContext) error) error {
	return mc.cl.UseSession(ctx, fn)
}

func NewDatabase(dbName string, client MongoClientInterface) DatabaseInterface {
	db, _ := client.Database(dbName)
	return db
}

func (mc *mongoClient) Database(dbName string) (DatabaseInterface, error) {
	db := mc.cl.Database(dbName)

	log.Info("checking database ....")
	e := mc.cl.Ping(context.Background(), readpref.Primary())
	if e != nil {
		return nil, e
	}
	log.Info("mongo database connected successfully")
	return &mongoDatabase{db: db}, nil
}

func (mc *mongoClient) StartSession() (mongo.Session, error) {
	return mc.cl.StartSession()
}

func (m *mongoSession) StartTransaction(i interface{}) error {
	options := i.(*options.TransactionOptions)
	return m.s.StartTransaction(options)
}

func (m *mongoSession) AbortTransaction(ctx interface{}) error {
	sessionContext := (ctx).(mongo.SessionContext)
	return m.s.AbortTransaction(sessionContext)
}

func (m *mongoSession) CommitTransaction(ctx interface{}) error {
	sessionContext := (ctx).(mongo.SessionContext)
	return m.s.CommitTransaction(sessionContext)
}

func (m *mongoSession) WithTransaction(ctx context.Context, fn func(sessCtx mongo.Session) (interface{}, error), opts ...*interface{}) (interface{}, error) {

	mongoFn := interface{}(fn).(func(sessCtx mongo.SessionContext) (interface{}, error))
	options := interface{}(opts).(*options.TransactionOptions)

	return m.s.WithTransaction(ctx, mongoFn, options)
}

func (m *mongoSession) EndSession(ctx context.Context) {
	m.s.EndSession(ctx)
}

func (mc *mongoClient) Connect() error {
	// mongo client does not use context on connect method. There is a ticket
	// with a request to deprecate this functionality and another one with
	// explanation why it could be useful in synchronous requests.
	// https://jira.mongodb.org/browse/GODRIVER-1031
	// https://jira.mongodb.org/browse/GODRIVER-979
	return mc.cl.Connect(nil)
}

func (mc *mongoClient) Disconnect() error {
	return mc.cl.Disconnect(context.Background())
}

func (md *mongoDatabase) Collection(colName string) CollectionInterface {
	collection := md.db.Collection(colName)
	return &mongoCollection{coll: collection}
}

func (md *mongoDatabase) Client() MongoClientInterface {
	client := md.db.Client()
	return &mongoClient{cl: client}
}

func (mc *mongoCollection) FindOne(ctx interface{}, filter interface{}) SingleResultInterface {
	var singleResult *mongo.SingleResult
	ctxType := reflect.TypeOf(ctx).String()
	log.Info("[mongo_client][FindOne] ctxType:[%v]", ctxType)
	if ctxType == "*context.emptyCtx" ||
		ctxType == "*context.valueCtx" {
		log.Info("[mongo_client][FindOne] in if ctxType:[%v]", ctxType)
		sessionContext := (ctx).(context.Context)
		singleResult = mc.coll.FindOne(sessionContext, filter)
	} else {
		log.Info("[mongo_client][FindOne] in else ctxType:[%v]", ctxType)
		sessionContext := (ctx).(mongo.SessionContext)
		singleResult = mc.coll.FindOne(sessionContext, filter)
	}

	return &mongoSingleResult{sr: singleResult}
}

func (mc *mongoCollection) Find(ctx context.Context, filter interface{}, opts *options.FindOptions) (CursorInterface, error) {

	log.Info("[mongo_client][Find] MongoClient Init in Find")

	cursor, err := mc.coll.Find(ctx, filter, opts)
	if err != nil {
		log.Error("[mongo_client][Find] MongoClient with error : [%s]", err.Error())
		return nil, err
	}

	return &MongoCursor{cr: cursor}, nil
}

func (mc *mongoCollection) InsertOne(ctx interface{}, document interface{}) (string, error) {

	var id *mongo.InsertOneResult
	var err error
	ctxType := reflect.TypeOf(ctx).String()
	if ctxType == "*context.emptyCtx" ||
		ctxType == "*context.valueCtx" || ctxType == "*context.cancelCtx" {
		fmt.Printf("\n [mongo_client] In if: [%s] ", reflect.TypeOf(ctx).String())
		sessionContext := (ctx).(context.Context)
		id, err = mc.coll.InsertOne(sessionContext, document)
	} else {
		sessionContext := (ctx).(mongo.SessionContext)
		id, err = mc.coll.InsertOne(sessionContext, document)
	}
	if err != nil {
		return "", err
	}
	resInserted := id.InsertedID.(primitive.ObjectID).String()
	return resInserted, err

}

func (mc *mongoCollection) DeleteOne(ctx interface{}, filter interface{}) (int64, error) {
	var count *mongo.DeleteResult
	var err error

	ctxType := reflect.TypeOf(ctx).String()

	log.Info("\n [mongo_client] ctxType: %s", ctxType)
	fmt.Printf("\n [mongo_client] ctxType: [%s] ", ctxType)
	if ctxType == "*context.emptyCtx" || ctxType == "*context.valueCtx" || ctxType == "*context.cancelCtx" {
		fmt.Printf("\n [mongo_client] In if ctxType: [%s] ", reflect.TypeOf(ctx).String())
		sessionContext := (ctx).(context.Context)
		count, err = mc.coll.DeleteOne(sessionContext, filter)
		fmt.Printf("\n [mongo_client] In if count: [%v] ", count)
		fmt.Printf("\n [mongo_client] In if count.DeletedCount: [%v] ", count.DeletedCount)
		if err != nil {
			fmt.Printf("\n [mongo_client] In if error: [%s] ", err.Error())
		}

	} else {
		fmt.Printf("\n [mongo_client] In else ctxType: [%s] ", ctxType)
		sessionContext := (ctx).(mongo.SessionContext)
		count, err = mc.coll.DeleteOne(sessionContext, filter)
	}
	return count.DeletedCount, err
}

func (mc *mongoCollection) UpsertOne(ctx, filter, update interface{}) (string, int64, int64, error) {

	log.Info("type context %s", reflect.TypeOf(ctx).String())
	var updateResult *mongo.UpdateResult
	var err error

	var sessionContext context.Context
	opts := options.Update().SetUpsert(true)
	if reflect.TypeOf(ctx).String() == "*context.emptyCtx" {
		sessionContext = (ctx).(context.Context)
	} else {
		sessionContext = (ctx).(mongo.SessionContext)
	}

	updateResult, err = mc.coll.UpdateOne(sessionContext, filter, update, opts)
	if err != nil {
		log.Error("Error in called to UpdateOne: [%s]", err.Error())
		return "", 0, 0, err
	}

	matchedCount := updateResult.MatchedCount
	modifiedCount := updateResult.ModifiedCount
	upsertedCount := updateResult.UpsertedCount
	upsertedId := updateResult.UpsertedID

	var resUpserted string
	if upsertedId != nil {
		resUpserted = upsertedId.(primitive.ObjectID).String()
	}

	if matchedCount != 0 {
		log.Info("Matched and replaced an existing document in UpdateOne: [%d]", matchedCount)
	}
	if upsertedCount != 0 {
		log.Info("The new document was inserted successfully with ID: [%v]", upsertedId)
	}

	return resUpserted, modifiedCount, upsertedCount, nil
}

func (sr *mongoSingleResult) Decode(v interface{}) error {
	return sr.sr.Decode(v)
}

func (cr *MongoCursor) DecodeCursor(v interface{}) error {
	return cr.cr.Decode(v)
}

func (cr *MongoCursor) AllCursor(ctx context.Context, result interface{}) error {
	return cr.cr.All(ctx, result)
}

func (mc *mongoCollection) UpdateOne(ctx interface{}, filter interface{}, update interface{}) (int, error) {
	var count *mongo.UpdateResult
	var err error
	opts := options.Update().SetUpsert(false)

	ctxType := reflect.TypeOf(ctx).String()
	if ctxType == "*context.emptyCtx" ||
		ctxType == "*context.valueCtx" || ctxType == "*context.cancelCtx" {
		sessionContext := (ctx).(context.Context)
		count, err = mc.coll.UpdateOne(sessionContext, filter, update, opts)
	} else {
		sessionContext := (ctx).(mongo.SessionContext)
		count, err = mc.coll.UpdateOne(sessionContext, filter, update, opts)
	}
	return int(count.ModifiedCount), err
}

func (mc *mongoCollection) UpdateMany(ctx interface{}, filter interface{}, update interface{}) (int, int, error) {
	var updateResult *mongo.UpdateResult
	var err error

	opts := options.Update().SetUpsert(false)

	if reflect.TypeOf(ctx).String() == "*context.emptyCtx" || reflect.TypeOf(ctx).String() == "*context.valueCtx" || reflect.TypeOf(ctx).String() == "*context.cancelCtx" {
		sessionContext := (ctx).(context.Context)
		updateResult, err = mc.coll.UpdateMany(sessionContext, filter, update, opts)
	} else {
		sessionContext := (ctx).(mongo.SessionContext)
		updateResult, err = mc.coll.UpdateMany(sessionContext, filter, update, opts)
	}
	if err != nil {
		log.Error("Error in Update Many, err: %s", err.Error())
		return 0, 0, err
	}
	matchedCount := updateResult.MatchedCount
	modifiedCount := updateResult.ModifiedCount

	return int(matchedCount), int(modifiedCount), err
}
