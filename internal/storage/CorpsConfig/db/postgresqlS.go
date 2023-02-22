package db

import (
	"github.com/sirupsen/logrus"
	"kz_bot/pkg/clientDB/postgresqlS"
)

type Repository struct {
	client postgresqlS.Client
	log    *logrus.Logger
}

func NewRepository(client postgresqlS.Client, loger *logrus.Logger) *Repository {
	return &Repository{
		client: client,
		log:    loger,
	}
}
