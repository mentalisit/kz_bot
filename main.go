package main

import (
	"fmt"
	"kz_bot/config"
	"kz_bot/internal/bot"
	discord "kz_bot/internal/clients/ds"
	telega "kz_bot/internal/clients/tg"
	db_Mysql "kz_bot/internal/dbase/dbaseMysql"
	"time"
)

func main() {
	fmt.Println("ЗАПУСК БОТА")
	cfg := config.InitConfig()

	tg := &telega.Telegram{}
	ds := &discord.Ds{}
	db := &db_Mysql.Db{}

	go func() {
		tg.InitTG(cfg.TokenT)
	}()

	go func() {
		ds.InitDS(cfg.TokenD)
	}()

	go func() {
		db.DbConnection()
	}()

	time.Sleep(time.Second * 5)
	bot.NewBot(*tg, *ds, *db)

	db.ReadBotCorpConfig()

	fmt.Println("35", tg.BotName())
	fmt.Println("36", ds.BotName())

	<-make(chan struct{})
	return
}
