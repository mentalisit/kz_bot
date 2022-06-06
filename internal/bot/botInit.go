package bot

import (
	"fmt"
	"kz_bot/internal/clients/discordClient"
	"kz_bot/internal/clients/telegramClient"
	"kz_bot/internal/dbase/dbaseMysql"
	"kz_bot/internal/models"
)

type Bot struct {
	Tg telegramClient.TelegramInterface
	Ds discordClient.DiscordInterface
	Db *dbaseMysql.Db
}

func NewBot(tg telegramClient.TelegramInterface, ds discordClient.DiscordInterface, db *dbaseMysql.Db) *Bot {
	return &Bot{Tg: tg, Ds: ds, Db: db}
}
func (b *Bot) InitBot() {
	for {
		select {
		case in := <-models.ChTg:
			b.LogicRs(in)
		case in := <-models.ChDs:
			b.LogicRs(in)

		}
	}
}

func (b *Bot) LogicRs(in models.InMessage) {
	fmt.Println(in.Name, "пишет")
	fmt.Println(b.Ds.BotName(), b.Tg.BotName())

}
