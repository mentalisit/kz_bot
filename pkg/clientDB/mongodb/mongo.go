package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"kz_bot/internal/config"
)

func NewMongoClient() (client *mongo.Client, err error) {
	uri := config.Instance.Mongo
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	return client, err
}
