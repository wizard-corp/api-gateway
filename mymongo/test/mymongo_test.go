package mymongo

import (
	"context"
	"testing"
	"time"

	"github.com/wizard-corp/api-gateway/mymongo"
)

func TestMongoInterface(t *testing.T) {
	server := &mymongo.MongoServer{
		Host:     "127.0.0.1",
		Port:     27017,
		User:     "mongo",
		Password: "mongo",
		Database: "test",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mymongo.NewMongoClient(server)
	if err != nil {
		t.Error(mymongo.DB_CONNECTION_FAILED)
	}

	// Access database and collection
	db := client.Database(server.Database)
	col := db.Collection("person")

	// Sample data for insertion
	data := map[string]interface{}{
		"name": "John Doe",
		"age":  30,
		"city": "New York",
	}

	// Insert data with proper error handling
	_, err = col.InsertOne(ctx, data)
	if err != nil {
		t.Error(mymongo.DB_INSERT_FAILED)
	}

	// Define filter for finding documents
	filter := map[string]interface{}{"name": "John Doe"}

	// Find documents with proper error handling and cursor closure
	cursor, err := col.Find(ctx, filter)
	if err != nil {
		t.Error(mymongo.DB_FIND_FAILED)
	}
	defer cursor.Close(ctx)

	// Iterate through results (optional)
	var foundDocument bool
	for cursor.Next(ctx) {
		var result map[string]interface{}
		err := cursor.Decode(&result)
		if err != nil {
			t.Error(mymongo.DB_DECODE_FAILED)
		}
		foundDocument = true
	}

	if !foundDocument {
		// Indicate no document found using an error (consider a custom error type)
		t.Error(mymongo.DB_COLLECTION_NOT_FOUND)
	}
} // Success
