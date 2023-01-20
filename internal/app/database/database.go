package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func ConnectDB() *mongo.Client {
	log.Print("Connect to MongoDB")

	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://dail:123@cluster0.8zjtjcp.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed: %v", err)
	}

	log.Print("Connected OK")
	return client
}

func Collection() *mongo.Collection {
	client := ConnectDB()
	collection := client.Database("users").Collection("userinfo")

	return collection
}
