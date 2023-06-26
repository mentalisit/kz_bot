package mongo

import (
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"kz_bot/pkg/clientDB/mongodb"
)

type DB struct {
	s   *mongo.Client
	log *logrus.Logger
}

func InitMongoDB(log *logrus.Logger) *DB {
	client, err := mongodb.NewMongoClient()
	if err != nil {
		log.Println("InitMongoDB() " + err.Error())
		return nil
	}

	d := &DB{
		s:   client,
		log: log,
	}
	return d
}
