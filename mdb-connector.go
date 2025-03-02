package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect to the MongoDB server using port 27017

const mongoURI = "mongodb://localhost:27017"

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Error creating MongoDB client:", err)
	}


	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()


	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	defer client.Disconnect(ctx)

	fmt.Println("Connected to MongoDB!")


	database := client.Database("testdb")
	collection := database.Collection("users")

	// How to insert a document
  
	user := bson.M{"name": "Alice", "age": 30, "city": "New York"}
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal("Error inserting document:", err)
	}
	fmt.Println("Inserted document ID:", result.InsertedID)

	// Find to find a document
	var foundUser bson.M
	err = collection.FindOne(ctx, bson.M{"name": "Alice"}).Decode(&foundUser)
	if err != nil {
		log.Fatal("Error finding document:", err)
	}
	fmt.Println("Found user:", foundUser)
}
