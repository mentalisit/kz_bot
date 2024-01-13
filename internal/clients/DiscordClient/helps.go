package DiscordClient

import (
	"fmt"
	"kz_bot/internal/models"
)

func (d *Discord) Help(Channel string) {
	mId := d.hhelp1(Channel)
	d.DeleteMesageSecond(Channel, mId, 184)
}

func (d *Discord) HelpChannelUpdate(c models.CorporationConfig) string {
	return d.hhelp1(c.DsChannel)
}

func (d *Discord) hhelp1(chatid string) string {
	m := d.SendEmbedText(chatid, d.getLang(chatid, "spravka"),
		fmt.Sprintf("%s \n\n%s", d.getLang(chatid, "botUdalyaet"), d.getLang(chatid, "hhelpText")))
	return m.ID
}

//func (d *Discord) restoredb(c *models.CorporationConfig) {
//	list := ReadCorps()
//	for _, config := range list {
//		if config.DsChannel == c.DsChannel {
//			c.CorpName = config.CorpName
//			c.DsChannel = config.DsChannel
//			c.TgChannel = config.TgChannel
//			c.WaChannel = config.WaChannel
//			c.Country = config.Country
//			c.DelMesComplite = config.DelMesComplite
//			c.Primer = config.Primer
//			c.Guildid = config.Guildid
//		}
//	}
//}
