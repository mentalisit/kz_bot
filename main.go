package main

import (
	"fmt"
	"kz_bot/config"
	"kz_bot/internal/bot"
	"kz_bot/internal/clients/discordClient"
	"kz_bot/internal/clients/telegramClient"
	"kz_bot/internal/dbase/dbaseMysql"
	"time"
)

func main() {
	fmt.Println("ЗАПУСК БОТА")
	cfg := config.InitConfig()

	tg := &telegramClient.Telegram{}
	ds := &discordClient.Ds{}
	db := &dbaseMysql.Db{}

	go tg.InitTG(cfg.TokenT)
	go ds.InitDS(cfg.TokenD)
	go db.DbConnection()

	time.Sleep(time.Second * 5)
	db.ReadBotCorpConfig()

	bot.NewBot(tg, ds, db).InitBot()

	<-make(chan struct{})
	return
}
