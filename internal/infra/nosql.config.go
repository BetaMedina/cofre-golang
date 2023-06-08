package infra

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client

func init() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(GetEnvs("MONGO_CONNECTION_STRING")))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	Client = client
}

func GetConnection(collection string) *mongo.Collection {
	return (*mongo.Collection)(Client.Database(GetEnvs("DATABASE")).Collection(collection))
}
