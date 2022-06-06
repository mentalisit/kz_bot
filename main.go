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

	go tg.InitTG(cfg.TokenT)
	go ds.InitDS(cfg.TokenD)
	go db.DbConnection()

	time.Sleep(time.Second * 5)
	go bot.NewBot(*tg, *ds, *db).InitBot()

	db.ReadBotCorpConfig()

	fmt.Println("35", tg.BotName())
	fmt.Println("36", ds.BotName())

	<-make(chan struct{})
	return
}
