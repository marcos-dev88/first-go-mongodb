package main

import (
	"github.com/marcos-dev88/first-go-mongodb/database"
	"log"
)

var mongoURI = "mongodb://127.0.0.1:27017"

func main() {

	db := database.NewMongoDB("go-mongoDB", "test", mongoURI)

	dbMethods := database.NewRepository(db)

	octopus, _ := dbMethods.GetAllOctopus()

	//dbMethods.CreateOctopus("testOctopus", 1)

	log.Printf("Octopus -> %v", octopus)
}
