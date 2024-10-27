package handlers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var taskCollection *mongo.Collection

// ConnectDB initializes MongoDB connection
func ConnectDB() {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	taskCollection = client.Database("tasks_db").Collection("tasks")
	fmt.Println("Connected to MongoDB!")
}

// GetTaskCollection returns the task collection
func GetTaskCollection() *mongo.Collection {
	return taskCollection
}
