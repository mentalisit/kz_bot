package clients

import (
	"github.com/mentalisit/logger"
	"kz_bot/internal/clients/DiscordClient"
	"kz_bot/internal/clients/TelegramClient"
	"kz_bot/internal/config"
	"kz_bot/internal/storage"
)

type Clients struct {
	Ds      *DiscordClient.Discord
	Tg      *TelegramClient.Telegram
	storage *storage.Storage
}

func NewClients(log *logger.Logger, st *storage.Storage, cfg *config.ConfigBot) *Clients {
	c := &Clients{storage: st}

	c.Ds = DiscordClient.NewDiscord(log, st, cfg)

	c.Tg = TelegramClient.NewTelegram(log, st, cfg)

	return c
}
func (c *Clients) DeleteMessageTimer() {
	if config.Instance.BotMode != "dev" {
		m := c.storage.Temp.TimerDeleteMessage()
		if len(m) > 0 {
			for _, timer := range m {
				if timer.Dsmesid != "" {
					go c.Ds.DeleteMesageSecond(timer.Dschatid, timer.Dsmesid, timer.Timed)
				}
				if timer.Tgmesid != "" {
					go c.Tg.DelMessageSecond(timer.Tgchatid, timer.Tgmesid, timer.Timed)
				}
			}
		}
	}
}
