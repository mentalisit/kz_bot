package clients

import (
	"github.com/sirupsen/logrus"
	"kz_bot/internal/clients/DiscordClient"
	"kz_bot/internal/clients/TelegramClient"
	//"kz_bot/internal/clients/WhatsappClient"
	"kz_bot/internal/config"
	"kz_bot/internal/storage"
	"time"
)

type Clients struct {
	Ds *DiscordClient.Discord
	Tg *TelegramClient.Telegram
	//Wa      *WhatsappClient.Whatsapp
	storage *storage.Storage
}

func NewClients(log *logrus.Logger, st *storage.Storage, cfg *config.ConfigBot) *Clients {
	c := &Clients{storage: st}

	c.Ds = DiscordClient.NewDiscord(log, st, cfg)

	c.Tg = TelegramClient.NewTelegram(log, st, cfg)

	//c.Wa = WhatsappClient.NewWhatsapp(c.Inbox, log, st, cfg)

	go c.deleteMessageTimer()
	return c
}
func (c *Clients) deleteMessageTimer() {
	for {
		<-time.After(1 * time.Minute)
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
