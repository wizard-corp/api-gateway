package mymongo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type MongoDB struct {
	Cli *mongo.Client
	Mdb *mongo.Database
}

func NewMongoDBClient(config *MongoConfig) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongodbURI := fmt.Sprintf("mongodb://%s:%s@%s:%d", config.User, config.Password, config.Host, config.Port)
	if config.User == "" || config.Password == "" {
		mongodbURI = fmt.Sprintf("mongodb://%s:%d", config.Host, config.Port)
	}

	cli, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbURI))
	if err != nil {
		return nil, errors.New(DB_CONNECTION_FAILED)
	}
	err = cli.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, errors.New(DB_CONNECTION_FAILED)
	}

	db := cli.Database(config.Database)
	return &MongoDB{Cli: cli, Mdb: db}, nil
}

func (mongoDB *MongoDB) Close() error {
	if mongoDB.Cli == nil {
		return nil // Client not initialized, nothing to close
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := mongoDB.Cli.Disconnect(ctx); err != nil {
		return errors.New(DB_DISCONNECT_FAILED)
	}

	return nil
}
