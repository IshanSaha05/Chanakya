package mongodb

import (
	"context"

	"github.com/IshanSaha05/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDb struct {
	context    context.Context
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

// Function to create a mongo object containing context, client, database and collection.
func (object *MongoDb) GetMongoClient() error {
	object.context = context.Background()
	client, err := mongo.Connect(object.context, options.Client().ApplyURI(config.MongoSite))

	if err != nil {
		return err
	}

	object.client = client
	object.database = nil
	object.collection = nil

	return nil
}

func (object *MongoDb) SetMongoDatabase(databaseName string) {
	object.database = object.client.Database(databaseName)
}

func (object *MongoDb) SetMongoCollection(collectionName string) error {
	allCollectionNames, err := object.database.ListCollectionNames(object.context, bson.D{})
	if err != nil {
		return err
	}

	for _, name := range allCollectionNames {
		if name == collectionName {
			object.collection = object.database.Collection(collectionName)
			return nil
		}
	}

	err = object.database.CreateCollection(object.context, collectionName)

	if err != nil {
		return err
	}

	object.collection = object.database.Collection(collectionName)
	return nil
}

func (object *MongoDb) InsertIntoDB(datas []interface{}) error {
	for _, data := range datas {
		_, err := object.collection.InsertOne(object.context, data)

		if err != nil {
			return err
		}
	}

	return nil
}
