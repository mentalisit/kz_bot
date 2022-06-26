package main

import (
	"fmt"
	"kz_bot/config"
	"kz_bot/internal/bot"
	"kz_bot/internal/clients/discordClient"
	"kz_bot/internal/clients/telegramClient"
	"kz_bot/internal/dbase/dbaseMysql"
	"kz_bot/internal/logger"
)

func main() {
	fmt.Println("ЗАПУСК БОТА")
	cfg := config.InitConfig()
	log := logger.NewLoggerTG(cfg.LogToken, cfg.LogChatId)
	log.Println("🚀  загрузка  🚀")

	db := &dbaseMysql.Db{}
	tg := &telegramClient.Telegram{}
	ds := &discordClient.Ds{}
	//wa := &watsappClient.Watsapp{}

	//go wa.InitWA()

	db.InitDB(log, cfg)

	tg.InitTG(cfg.TokenT, db.Db, log)
	ds.InitDS(cfg.TokenD, db.Db, log)

	db.ReadBotCorpConfig()

	go bot.NewBot(tg, ds, db, log).InitBot()

	<-make(chan struct{})
	//return
}
