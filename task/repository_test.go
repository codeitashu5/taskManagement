package task

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"taskManagement/envornment"
	"testing"
)

func newMongoClient() *mongo.Client {
	uri := envornment.GetMongoURI()
	if uri == "" {
		log.Fatal(errors.New("MongoDB URI is empty"))
	}

	mongoTestClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(envornment.GetMongoURI()))
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Print("test client connection successful")
	err = mongoTestClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Print("test client connection successful")

	return mongoTestClient
}

func TestMongoOperations(t *testing.T) {
	mongoClient := newMongoClient()
	defer func(mongoClient *mongo.Client, ctx context.Context) {
		err := mongoClient.Disconnect(ctx)
		if err != nil {
			log.Printf("error while disconnecting mongo : %v", err)
		}
	}(mongoClient, context.Background())

	// dummy data for the user and employee

}
