package controller

import (
	"context"
	"fmt"

	"github.com/Pranshu321/orders-api.git/models"
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

func Insert(order models.Order) error {
	_, err := collection.InsertOne(context.Background(), order)
	if err != nil {
		return fmt.Errorf("error inserting order: %s", err)
	}
	return nil
}

func FindById(id uint64) (models.Order, error) {
	var order models.Order
	filter := map[string]interface{}{"_id": id}
	err := collection.FindOne(context.Background(), filter).Decode(&order)
	if err != nil {
		return models.Order{}, fmt.Errorf("No Such Order Present in DB: %s", err)
	}
	return order, nil
}

func Delete(id uint64) error {
	filter := map[string]interface{}{"_id": id}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("Order Not Exist: %s", err)
	}
	return nil
}

func FindAll() ([]models.Order, error) {
	var orders []models.Order
	cursor, err := collection.Find(context.Background(), map[string]interface{}{})
	if err != nil {
		return nil, fmt.Errorf("Error fetching orders: %s", err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var order models.Order
		cursor.Decode(&order)
		orders = append(orders, order)
	}
	return orders, nil
}

func Update(order models.Order) error {
	filter := map[string]interface{}{"_id": order.OrderID}
	update := map[string]interface{}{"$set": order}
	_,
		err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("Error updating order: %s", err)
	}
	return nil
}
