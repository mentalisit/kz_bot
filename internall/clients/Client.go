package clients

import (
	"github.com/sirupsen/logrus"
	"kz_bot/config"
	"kz_bot/internall/clients/DiscordClient"
	"kz_bot/internall/clients/TelegramClient"
	"kz_bot/internall/clients/WhatsappClient"
	"kz_bot/internall/models"
	"kz_bot/internall/storage"
)

type Clients struct {
	Ds    *DiscordClient.Discord
	Tg    *TelegramClient.Telegram
	Wa    *WhatsappClient.Whatsapp
	Inbox chan models.InMessage
}

func NewClients(log *logrus.Logger, st *storage.Storage, cfg config.ConfigBot) *Clients {
	//inbox channel
	var inbox = make(chan models.InMessage, 10)

	ds := DiscordClient.NewDiscord(inbox, log, st, cfg)

	tg := TelegramClient.NewTelegram(inbox, log, st, cfg)

	wa := WhatsappClient.NewWhatsapp(inbox, log, st, cfg)

	return &Clients{
		Ds:    ds,
		Tg:    tg,
		Wa:    wa,
		Inbox: inbox,
	}
}
func (c Clients) name() {

}
