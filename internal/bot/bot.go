package bot

import (
	"database/sql"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"kz_bot/internal/clients/ds"
	"kz_bot/internal/models"
)

type Bot struct {
	Tg tgbotapi.BotAPI
	Ds *ds.Ds
	Db *sql.DB
}

func NewBot(tg tgbotapi.BotAPI, ds ds.Ds) *Bot {
	return &Bot{Tg: tg, Ds: &ds}
}
func (b *Bot) SendIF() {
	name := b.Tg.Self.UserName
	//b.Ds.Send("","")

	fmt.Println(" tg", name)
	fmt.Println(" ds", b.Ds.NameBot())
}
func (b *Bot) InitBot() {
	in := <-models.ChTg
	fmt.Println(in.Tip)
}

func LogicRs() {

}
