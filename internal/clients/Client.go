package clients

import (
	"github.com/sirupsen/logrus"
	"kz_bot/internal/clients/DiscordClient"
	"kz_bot/internal/clients/TelegramClient"
	"kz_bot/internal/clients/WhatsappClient"
	"kz_bot/internal/config"
	"kz_bot/internal/models"
	"kz_bot/internal/storage"
)

type Clients struct {
	Ds     *DiscordClient.Discord
	Tg     *TelegramClient.Telegram
	Wa     *WhatsappClient.Whatsapp
	Inbox  chan models.InMessage
	ToGame chan models.Message
}

func NewClients(log *logrus.Logger, st *storage.Storage, cfg *config.ConfigBot) *Clients {
	c := &Clients{}
	//inbox channel
	c.Inbox = make(chan models.InMessage, 10)
	c.ToGame = make(chan models.Message, 10)
	var toGameForDiscord = make(chan models.Message, 10)
	var toGameForTelegram = make(chan models.Message, 10)

	c.Ds = DiscordClient.NewDiscord(c.Inbox, toGameForDiscord, log, st, cfg)

	c.Tg = TelegramClient.NewTelegram(c.Inbox, toGameForTelegram, log, st, cfg)

	//c.Wa = WhatsappClient.NewWhatsapp(c.Inbox, log, st, cfg)

	go c.HadesBridge(toGameForDiscord, toGameForTelegram)
	return c
}
