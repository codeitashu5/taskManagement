package mongoClient

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"taskManagement/envornment"
)

// MongoDB MongoClient making MongoClient singleTon to be used to create database using it
var MongoDB *mongo.Client

func SetUpMongo() error {
	var err error
	uri := envornment.GetMongoURI()
	if uri == "" {
		return errors.New("error getting mongo uri")
	}

	MongoDB, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	return err
}

func ShutdownMongo() error {
	err := MongoDB.Disconnect(context.TODO())
	return err
}
