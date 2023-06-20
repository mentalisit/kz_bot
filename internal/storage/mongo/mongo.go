package mongo

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"kz_bot/internal/config"
)

type DB struct {
	s   *mongo.Client
	log *logrus.Logger
}

func InitMongoDB(log *logrus.Logger) *DB {
	uri := config.Instance.Mongo
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Panic(err)
	}

	d := &DB{
		s:   client,
		log: log,
	}
	return d
}
