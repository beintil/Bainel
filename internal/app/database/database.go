package database

import (
	"Bainel/configs"
	"Bainel/pkg/error_handler/server_errors"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	Collection = collection()
	lineError  string
)

func connectDB() *mongo.Client {
	log.Print("Connect to MongoDB")

	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(configs.GetMongoURL()))
	if err != nil {
		lineError = "database, line 20: Error connecting to MongoDB"
		server_errors.ErrorFatal(err, lineError)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		lineError = "database, line 20: Error connecting to MongoDB"
		server_errors.ErrorFatal(err, lineError)
	}

	log.Print("Connected OK")
	return client
}

func collection() *mongo.Collection {
	client := connectDB()
	collection := client.Database("users").Collection("userinfo")

	return collection
}
