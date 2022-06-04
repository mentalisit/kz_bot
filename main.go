package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"kz_bot/config"
	"kz_bot/pkg/bot"
	discord "kz_bot/pkg/clients/ds"
	telega "kz_bot/pkg/clients/tg"
	"time"
)

func main() {
	fmt.Println("ЗАПУСК БОТА")
	cfg := config.InitConfig()

	var tg *tgbotapi.BotAPI
	var ds *discordgo.Session

	go func() {
		tg = telega.InitTG(cfg.TokenT)
	}()

	go func() {
		ds = discord.InitDS(cfg.TokenD)
	}()

	time.Sleep(time.Second * 5)
	bot.NewBot(*tg, ds).SendIF()

	fmt.Println("35", tg.Self.UserName)
	fmt.Println("36", ds.StateEnabled)

	<-make(chan struct{})
	return
}
