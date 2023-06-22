package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"kz_bot/internal/config"
	"kz_bot/internal/models"
	"kz_bot/internal/storage/mongo"
	"kz_bot/pkg/clientDB/postgresqlS"
)

var HadesStorage *Hades

type Hades struct {
	client postgresqlS.Client
	log    *logrus.Logger
}

func NewHades(log *logrus.Logger) *Hades {
	h := &Hades{log: log}
	client, err := postgresqlS.NewClient(context.Background(), log, 5, config.Instance)
	if err != nil {
		return nil
	}
	h.client = client
	HadesStorage = h
	h.ReadCorporation()
	return h
}
func (h *Hades) ReadCorporation() {

	read :=
		`SELECT * FROM kzbot.bridge`
	results, err := h.client.Query(context.Background(), read)
	if err != nil {
		h.log.Println("Ошибка чтения крнфигурации корпораций hades", err)
	}
	fmt.Printf("hades: ")
	m := mongo.InitMongoDB(h.log)
	for results.Next() {
		var t models.Corporation
		err = results.Scan(&t.Id, &t.Corp, &t.DsChat, &t.DsChatWS1, &t.DsChatWS2,
			&t.GuildId, &t.TgChat, &t.TgChatWS1)

		m.InsertCorporation(models.CorporationHadesClient{
			Corp:      t.Corp,
			DsChat:    t.DsChat,
			DsChatWS1: t.DsChatWS1,
			DsChatWS2: t.DsChatWS2,
			GuildId:   t.GuildId,
			TgChat:    t.TgChat,
			TgChatWS1: t.TgChatWS1,
		})
	}
	fmt.Println()
	//go h.reloadConsoleClient(clients)

}
