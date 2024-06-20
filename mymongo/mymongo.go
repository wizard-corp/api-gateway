package mymongo

import (
	"context"
	"errors"
	"fmt"
	"log"
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

func NewConnection(server *MongoServer) (*MongoDB, error) {
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

	db := client.Database(server.Database)

	return &MongoDB{Client: client, Database: db}, nil
}

func CloseConnection(mdb *MongoDB) {
	if mdb.Client == nil {
		return
	}

	err := mdb.Client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}

func TestInfrastructure(server *MongoServer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Establish connection with proper error handling
	mdb, err := NewConnection(server)
	if err != nil {
		return fmt.Errorf("error in connection: %w", err)
	}
	defer CloseConnection(mdb) // Ensure connection closure (even on successful test)

	// Access database and collection
	col := mdb.Database.Collection("person")

	// Sample data for insertion
	data := map[string]interface{}{
		"name": "John Doe",
		"age":  30,
		"city": "New York",
	}

	// Insert data with proper error handling
	_, err = col.InsertOne(ctx, data)
	if err != nil {
		return fmt.Errorf("error inserting data: %w", err)
	}

	// Define filter for finding documents
	filter := map[string]interface{}{"name": "John Doe"}

	// Find documents with proper error handling and cursor closure
	cursor, err := col.Find(ctx, filter)
	if err != nil {
		return fmt.Errorf("error finding documents: %w", err)
	}
	defer cursor.Close(ctx)

	// Iterate through results (optional)
	var foundDocument bool
	for cursor.Next(ctx) {
		var result map[string]interface{}
		err := cursor.Decode(&result)
		if err != nil {
			return fmt.Errorf("error decoding result: %w", err)
		}
		foundDocument = true
	}

	if !foundDocument {
		// Indicate no document found using an error (consider a custom error type)
		return errors.New("document not found")
	}

	return nil // Success
}
