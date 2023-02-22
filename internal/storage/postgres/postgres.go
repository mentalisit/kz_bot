package postgres

import (
	"context"
	"github.com/sirupsen/logrus"
	"kz_bot/internal/config"
	"kz_bot/pkg/clientDB/postgresLocal"
)

type Db struct {
	db    postgresLocal.Client
	log   *logrus.Logger
	debug bool
}

func NewDb(log *logrus.Logger, cfg *config.ConfigBot) *Db {
	db, err := postgresLocal.NewClient(context.Background(), log, 5, cfg)
	if err != nil {
		log.Fatalln("Ошибка подключения к local ДБ ", err)
	}
	return &Db{
		db:    db,
		log:   log,
		debug: cfg.IsDebug,
	}
}
