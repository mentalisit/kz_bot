package main

import (
	"fmt"

	"kz_bot/config"
	"kz_bot/internal/bot"
	"kz_bot/internal/clients/discordClient"
	"kz_bot/internal/clients/telegramClient"
	"kz_bot/internal/dbase/dbaseMysql"
	"kz_bot/internal/hserver"
)

func main() {
	fmt.Println("ЗАПУСК БОТА")
	cfg := config.InitConfig()

	db := &dbaseMysql.Db{}
	tg := &telegramClient.Telegram{}
	ds := &discordClient.Ds{}
	//wa := &watsappClient.Watsapp{}

	//go wa.InitWA()

	db.DbConnection()

	tg.InitTG(cfg.TokenT)
	ds.InitDS(cfg.TokenD)
	//нужно проверить нужны ли тут горутины

	//time.Sleep(time.Second * 5)
	db.ReadBotCorpConfig()
	go bot.NewBot(tg, ds, db).InitBot()

	//тест сервера
	a := db.ReadAllTable()
	h := hserver.Httpserver{}
	h.NewServer("Mentalisit", a)

	//<-make(chan struct{})
	//return
}

/*
	config:=hserver.NewConfig()
		_,err:=toml.Decode(ConfigPath,config)
	if err != nil {
		log.Println(err)
	}
		s:=hserver.New(config)
		if err := s.Start(); err!=nil{log.Fatal(err)}
*/
