package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"kz_bot/config"
	"kz_bot/internal/bot"
	discord "kz_bot/internal/clients/ds"
	telega "kz_bot/internal/clients/tg"
	"time"
)

func main() {
	fmt.Println("ЗАПУСК БОТА")
	cfg := config.InitConfig()

	var tg *tgbotapi.BotAPI
	//var ds *discordgo.Session
	ds := &discord.Ds{}
	go func() {
		tg = telega.InitTG(cfg.TokenT)
	}()

	go func() {
		discord.InitDS(cfg.TokenD)
	}()

	time.Sleep(time.Second * 5)
	bot.NewBot(*tg, *ds).SendIF()

	fmt.Println("35", tg.Self.UserName)
	fmt.Println("36", ds.NameBot())

	<-make(chan struct{})
	return
}
