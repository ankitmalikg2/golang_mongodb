package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// You will be using this Trainer type later in the program
type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
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
	filter := bson.D{{"name", "Ash"}}
	// update := bson.D{
	// 	{"$inc", bson.D{
	// 		{"age", 1},
	// 	}},
	// }

	update2 := bson.D{
		{"$set", bson.D{
			{"age", 8},
		}},
	}
	fmt.Println(update2)
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	fmt.Println("Connected to MongoDB!")
}
