package clients

import (
	"github.com/sirupsen/logrus"
	"kz_bot/config"
	"kz_bot/internal/clients/discordClient"
	"kz_bot/internal/clients/telegramClient"
	"kz_bot/internal/clients/watsappClient"
	"kz_bot/internal/dbase"
)

type Client struct {
	Tg telegramClient.Tg
	Ds discordClient.Ds
	Wa watsappClient.Wa
}

func NewClient(cfg config.ConfigBot, db dbase.Db, log *logrus.Logger) *Client {
	telegram := telegramClient.Telegram{}
	telegram.InitTG(cfg.TokenT, db, log, cfg.Debug)

	ds := discordClient.Discord{}
	ds.InitDS(cfg.TokenD, db, log, cfg.Debug)

	wa := watsappClient.Watsapp{}
	//wa.InitWA(db)

	return &Client{Tg: &telegram, Ds: &ds, Wa: &wa}
}
