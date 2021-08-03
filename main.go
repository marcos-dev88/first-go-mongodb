package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var mongoURI = "mongodb://127.0.0.1:27017"

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	if err != nil {
		log.Fatalf("Error: %v", client)
	}

	cntx := context.Background()

	// Connect to MongoDB
	if err = client.Connect(cntx); err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Disconnection to database when code finish
	defer client.Disconnect(cntx)

	// Create database(schema) in mongo
	firstDB := client.Database("go-mongoDB")

	// Create a collection(table) in mongo
	if err = firstDB.CreateCollection(cntx, "octopus"); err != nil {
		log.Fatalf("Error: %v", err)
	}

	octopusCollection := firstDB.Collection("octopus")

	// It'll erase the database
	defer octopusCollection.Drop(cntx)

	// Insert a data in a database
	result, err := octopusCollection.InsertOne(cntx, bson.D{
		{Key: "name", Value: "Dr. Branches"},
		{Key: "age", Value: 2},
	})

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Printf("Result -> %v", result)

	cursor, err := octopusCollection.Find(cntx, bson.M{})

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	//var octopus []bson.M

	/** That's one way to get the all data from database **/
	//if err = cursor.All(cntx, &octopus); err != nil {
	//	log.Fatalf("Error: %v", err)
	//}
	//log.Printf("octopus slice -> %v", octopus)

	defer cursor.Close(cntx)

	// That's the way to get data from database with mongo
	for cursor.Next(cntx) {
		var srbranches bson.M

		if err = cursor.Decode(&srbranches); err != nil {
			log.Fatalf("Error: %v", err)
		}

		fmt.Printf("Octopus: %v", srbranches)
	}
}
