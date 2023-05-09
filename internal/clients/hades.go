package clients

import (
	"kz_bot/internal/models"
	"kz_bot/internal/storage/CorpsConfig/hades"
	"time"
)

func (c *Clients) HadesBridge(ds chan models.Message, tg chan models.Message) {
	for {
		select {
		case in := <-ds:
			c.filterDs(in)
		case in := <-tg:
			c.filterTg(in)
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}
func (c *Clients) filterDs(msg models.Message) {
	ok, corp := hades.HadesStorage.AllianceName(msg.Corporation)
	if ok && msg.Command == "text" {
		if corp.TgChat != 0 {
			text := "(DS)" + msg.Sender + ": " + msg.Text
			c.Tg.SendChannel(corp.TgChat, text)
		}
	}
	if ok {
		c.ToGame <- msg
	}
}
func (c *Clients) filterTg(msg models.Message) {
	ok, corp := hades.HadesStorage.AllianceName(msg.Corporation)
	if ok && msg.Command == "text" {
		if corp.DsChat != "" && msg.ChannelType == 0 {
			sender := "(TG)" + msg.Sender
			c.Ds.SendWebhookForHades(msg.Text, sender, corp.DsChat, corp.GuildId, msg.Avatar)
		}
		if corp.DsChatWS1 != "" && msg.ChannelType == 1 {
			sender := "(TG)" + msg.Sender
			c.Ds.SendWebhookForHades(msg.Text, sender, corp.DsChatWS1, corp.GuildId, msg.Avatar)
		}
	}
	if ok {
		c.ToGame <- msg
	}
}
