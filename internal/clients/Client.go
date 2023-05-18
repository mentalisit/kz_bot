package clients

import (
	"github.com/sirupsen/logrus"
	"kz_bot/internal/clients/DiscordClient"
	"kz_bot/internal/clients/TelegramClient"
	"kz_bot/internal/clients/WhatsappClient"
	"kz_bot/internal/config"
	"kz_bot/internal/storage"
)

type Clients struct {
	Ds *DiscordClient.Discord
	Tg *TelegramClient.Telegram
	Wa *WhatsappClient.Whatsapp
}

func NewClients(log *logrus.Logger, st *storage.Storage, cfg *config.ConfigBot) *Clients {
	c := &Clients{}

	c.Ds = DiscordClient.NewDiscord(log, st, cfg)

	c.Tg = TelegramClient.NewTelegram(log, st, cfg)

	//c.Wa = WhatsappClient.NewWhatsapp(c.Inbox, log, st, cfg)

	return c
}
