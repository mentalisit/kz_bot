package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"kz_bot/config"
	"kz_bot/internal/bot"
	"kz_bot/internal/clients/discordClient"
	"kz_bot/internal/clients/telegramClient"
	"kz_bot/internal/dbase/dbaseMysql"
)

func main() {
	fmt.Println("ЗАПУСК БОТА")
	log := logrus.New()
	cfg := config.InitConfig()
	log.Println("загрузка")

	db := &dbaseMysql.Db{}
	tg := &telegramClient.Telegram{}
	ds := &discordClient.Ds{}
	//wa := &watsappClient.Watsapp{}

	//go wa.InitWA()

	db.DbConnection()

	tg.InitTG(cfg.TokenT, db.Db)
	ds.InitDS(cfg.TokenD, db.Db)

	db.ReadBotCorpConfig()
	go bot.NewBot(tg, ds, db).InitBot()

	<-make(chan struct{})
	//return
}
