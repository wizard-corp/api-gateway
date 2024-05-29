package mymongo

import (
	"context"
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

func TestInfrastructure(server *MongoServer) string {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mdb, err := NewConnection(server)
	if err != nil {
		return fmt.Errorf("error in connection: %w", err).Error()
	}
	// Close the connection
	defer CloseConnection(mdb)

	// Access database and collection using methods on MongoDB
	col := mdb.Database.Collection("person")

	// Sample data for insertion
	data := map[string]interface{}{
		"name": "John Doe",
		"age":  30,
		"city": "New York",
	}

	// Insert data
	_, err = col.InsertOne(ctx, data)
	if err != nil {
		return fmt.Errorf("error inserting data: %w", err).Error()
	}

	// Define a filter for finding documents
	filter := map[string]interface{}{"name": "John Doe"}

	// Find documents
	cursor, err := col.Find(ctx, filter)
	if err != nil {
		return fmt.Errorf("error finding documents: %w", err).Error()
	}
	defer cursor.Close(ctx) // Ensure cursor is closed

	// Iterate through results and print them (optional)
	for cursor.Next(ctx) {
		var result map[string]interface{}
		err := cursor.Decode(&result)
		if err != nil {
			return fmt.Errorf("error decoding result: %w", err).Error()
		}
		fmt.Println(result)
	}

	return "SUCCESS"
}
