package controller

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://pranshujain0221:abc%40123@richpanel-assingment.rplhf7o.mongodb.net/?retryWrites=true&w=majority"

const dbName = "Microservices"
const collName = "Orders"

// MOST IMPORTANT
var collection *mongo.Collection

func Connect() error {
	fmt.Println("Connecting to MongoDB...")
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Printf("Error connecting to MongoDB: %s\n", err)
		return err
	}
	collection = client.Database(dbName).Collection(collName)
	fmt.Println("Connected to MongoDB!")
	return nil
}
