package CorpsConfig

import (
	"context"
	"github.com/sirupsen/logrus"
	"kz_bot/internal/config"
	"kz_bot/internal/storage/CorpsConfig/db"
	"kz_bot/pkg/clientDB/postgresqlS"
)

type Corps struct {
	client postgresqlS.Client
	log    *logrus.Logger
	debug  bool
	db     *db.Repository
}

func NewCorps(log *logrus.Logger, cfg *config.ConfigBot) *Corps {
	//создаем клиента инет базы данных
	client, err := postgresqlS.NewClient(context.Background(), log, 5, cfg)
	if err != nil {
		log.Fatalln("Ошибка подключения к облачной ДБ ", err)
	}

	//инициализируем репо
	repo := db.NewRepository(client, log)

	return &Corps{
		client: client,
		log:    log,
		debug:  cfg.IsDebug,
		db:     repo,
	}
}
