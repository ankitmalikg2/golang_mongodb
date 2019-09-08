package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	mongoOptions := options.Client().ApplyURI(":")
	mongo.Connect(context.TODO(), mongoOptions)

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("ankit_test").Collection("trainer")

	filter := bson.D{
		{"age", 10},
	}
	//setting options for limit 2
	findOptions := options.Find()
	findOptions.SetLimit(4)
	findOptions.SetSkip(1)

	cursor, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(context.TODO()) {
		// var result interface{}
		var result Trainer
		err = cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("result", result)
	}

	fmt.Println("ended")

}
