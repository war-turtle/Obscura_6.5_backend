package db

import (
	"context"
	"log"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// variable to access collections
var Collection *mongo.Collection

// Connect function to connect to mongodb instance
func Connect(uri string) {
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}

	Collection = client.Database("obscura").Collection("users")

	log.Println("Connected")
}