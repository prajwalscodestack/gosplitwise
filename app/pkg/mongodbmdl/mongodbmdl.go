// mongo_wrapper.go
package mongodbmdl

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB represents a MongoDB connection.
type MongoDB struct {
	client   *mongo.Client
	database *mongo.Database
	ctx      context.Context
}

var client *mongo.Client

func Init() {
	client = initConnection()
}
func initConnection() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	clnt, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic("error while connecting to mongodb:" + err.Error())
	}

	// Check if the connection is successful
	err = clnt.Ping(context.TODO(), nil)
	if err != nil {
		panic("error while connecting to mongodb:" + err.Error())
	}
	log.Println("CONNECTED TO MONGODB")
	return clnt
}

// NewMongoDB creates a new MongoDB connection.
func NewMongoDB(dbName string, ctx context.Context) (*MongoDB, error) {
	return &MongoDB{
		client:   client,
		database: client.Database(dbName),
		ctx:      ctx,
	}, nil
}

// Close closes the MongoDB connection.
func (db *MongoDB) Close() error {
	err := db.client.Disconnect(db.ctx)
	if err != nil {
		return err
	}
	return nil
}

// Collection returns a collection from the MongoDB database.
func (db *MongoDB) Collection(collectionName string) *mongo.Collection {
	return db.database.Collection(collectionName)
}

// SaveDocument saves a document in the specified collection.
func (db *MongoDB) SaveDocument(collectionName string, document interface{}) (*mongo.InsertOneResult, error) {
	collection := db.Collection(collectionName)
	result, err := collection.InsertOne(db.ctx, document)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FetchDocument fetches a single document from the specified collection.
func (db *MongoDB) FetchDocument(collectionName string, filter interface{}) *mongo.SingleResult {
	collection := db.Collection(collectionName)
	result := collection.FindOne(db.ctx, filter)
	return result
}

// UpdateDocument updates a document in the specified collection.
func (db *MongoDB) UpdateDocument(collectionName string, filter, update interface{}) (*mongo.UpdateResult, error) {
	collection := db.Collection(collectionName)
	result, err := collection.UpdateOne(db.ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return result, nil
}
