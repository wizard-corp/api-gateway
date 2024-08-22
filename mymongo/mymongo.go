package mymongo

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoServer struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

var (
	mongoClient *mongo.Client
	once        sync.Once
)

type Database interface {
	FindOne(context.Context, interface{}) SingleResult
	InsertOne(context.Context, interface{}) (interface{}, error)
	InsertMany(context.Context, []interface{}) ([]interface{}, error)
	DeleteOne(context.Context, interface{}) (int64, error)
	Find(context.Context, interface{}, ...*options.FindOptions) (Cursor, error)
	CountDocuments(context.Context, interface{}, ...*options.CountOptions) (int64, error)
	Aggregate(context.Context, interface{}) (Cursor, error)
	UpdateOne(context.Context, interface{}, interface{}, ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	UpdateMany(context.Context, interface{}, interface{}, ...*options.UpdateOptions) (*mongo.UpdateResult, error)
}

func NewMongoClient(server *MongoServer) (*mongo.Client, error) {
	once.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		mongodbURI := fmt.Sprintf("mongodb://%s:%s@%s:%d", server.User, server.Password, server.Host, server.Port)

		if server.User == "" || server.Password == "" {
			mongodbURI = fmt.Sprintf("mongodb://%s:%d", server.Host, server.Port)
		}

		client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbURI))
		if err != nil {
			return nil, err
		}

		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			return nil, err
		}

		mongoClient = client
	})

	return mongoClient, nil
}

func CloseMongoClient() error {
	if mongoClient == nil {
		return nil // Client not initialized, nothing to close
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := mongoClient.Disconnect(ctx)
	if err != nil {
		return err
	}

	mongoClient = nil
	return nil
}
