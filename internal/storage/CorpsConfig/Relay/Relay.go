package Relay

import (
	"github.com/sirupsen/logrus"
	"kz_bot/pkg/clientDB/postgresqlS"
)

type RelayStorage struct {
	client postgresqlS.Client
	log    *logrus.Logger
}

func NewRelayStorage(client postgresqlS.Client, log *logrus.Logger) *RelayStorage {
	s := &RelayStorage{client: client, log: log}
	return s
}
