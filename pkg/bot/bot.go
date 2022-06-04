package bot

import (
	"database/sql"
	"fmt"
	"github.com/bwmarrin/discordgo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"kz_bot/pkg/models"
)

type Bot struct {
	Tg tgbotapi.BotAPI
	Ds *discordgo.Session
	Db *sql.DB
}

func NewBot(tg tgbotapi.BotAPI, ds *discordgo.Session) *Bot {
	return &Bot{Tg: tg, Ds: ds}
}
func (b *Bot) SendIF() {
	name := b.Tg.Self.UserName
	id := b.Ds.State.User.Username

	fmt.Println(" tg", name)
	fmt.Println(" ds", id)
}
func (b *Bot) Run() {

}

func LogicRs(in models.InMessage) {
	fmt.Println(in)

}
