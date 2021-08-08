package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Repository interface {
	GetAllOctopus() ([]bson.M, error)
	CreateOctopus(name string, age int) error
}

type repository struct {
	mongodb MongoDB
}

func NewRepository(mongodb MongoDB) *repository {
	return &repository{mongodb: mongodb}
}

func (r repository) GetAllOctopus() ([]bson.M, error) {
	_, table, cntx, err := r.mongodb.GetConn()

	cursor, err := table.Find(cntx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Fatalf("Error to close context -> %v", err)
		}
	}(cursor, cntx)

	//That's the way to get data from database with mongo
	var octopus []bson.M

	for cursor.Next(cntx) {
		srbranches := bson.M{}
		if err = cursor.Decode(&srbranches); err != nil {
			return nil, err
		}
		octopus = append(octopus, srbranches)
	}

	/** That's one way to get the all data from database **/
	//if err = cursor.All(cntx, &octopus); err != nil {
	//	log.Fatalf("Error: %v", err)
	//}
	//log.Printf("octopus slice -> %v", octopus)

	return octopus, nil
}


func (r repository) CreateOctopus(name string, age int) error {

	_, table, cntx, err := r.mongodb.GetConn()

	if err != nil {
		return err
	}

	_, err = table.InsertOne(cntx, bson.D {
		{Key: "name", Value: name},
		{Key: "age", Value: age},
	})

	if err != nil {
		return err
	}

	return nil
}
