package DiscordClient

import (
	"fmt"
	"kz_bot/internal/models"
	"kz_bot/pkg/utils"
	"time"
)

func (d *Discord) Help(Channel string) {
	mId := d.hhelp1(Channel)
	d.DeleteMesageSecond(Channel, mId, 184)
}

func (d *Discord) Autohelpds() {
	tm := time.Now()
	mtime := tm.Format("15:04")
	if mtime == "12:00" {
		a := d.storage.ConfigRs.AutoHelp()
		for _, s := range a {
			if s.DsChannel != "" {
				if s.MesidDsHelp != "" {
					go d.DeleteMessage(s.DsChannel, s.MesidDsHelp)
					d.HelpChannelUpdate(s)
				} else {
					d.HelpChannelUpdate(s)
				}
			}
		}
		time.Sleep(time.Minute)
	} else if mtime == "03:00" {
		time.Sleep(1 * time.Second)
		utils.UpdateRun()
	}
}

func (d *Discord) HelpChannelUpdate(c models.CorporationConfig) {
	newMesidHelp := d.hhelp1(c.DsChannel)
	c.MesidDsHelp = newMesidHelp
	d.storage.ConfigRs.AutoHelpUpdateMesid(c)
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
