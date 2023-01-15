package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func ConnectDB() (*mongo.Client, error) {
	log.Println("Connect to MongoDB")

	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://dail:123@cluster0.8zjtjcp.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func Collection() *mongo.Collection {
	client, _ := ConnectDB()
	collection := client.Database("users").Collection("userinfo")

	return collection
}
