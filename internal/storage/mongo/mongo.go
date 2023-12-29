package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"kz_bot/pkg/clientDB/mongodb"
	"kz_bot/pkg/logger"
)

type DB struct {
	s   *mongo.Client
	log *logger.Logger
}

func InitMongoDB(log *logger.Logger) *DB {
	client, err := mongodb.NewMongoClient()
	if err != nil {
		log.Error(err.Error())
		return nil
	}

	d := &DB{
		s:   client,
		log: log,
	}
	return d
}
