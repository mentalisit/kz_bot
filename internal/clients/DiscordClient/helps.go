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
	if c.MesidDsHelp == "" {
		c.MesidDsHelp = d.hhelp1(c.DsChannel)
		return c.MesidDsHelp
	} else {
		messages, err := d.s.ChannelMessages(c.DsChannel, 10, "", c.MesidDsHelp, "")
		if err != nil {
			go d.DeleteMessage(c.DsChannel, c.MesidDsHelp)
			c.MesidDsHelp = d.hhelp1(c.DsChannel)
			return c.MesidDsHelp
		}
		if len(messages) > 2 {
			go d.DeleteMessage(c.DsChannel, c.MesidDsHelp)
			c.MesidDsHelp = d.hhelp1(c.DsChannel)
		}
	}
	return c.MesidDsHelp
}

func (d *Discord) hhelp1(chatid string) string {
	m := d.SendEmbedText(chatid, d.getLang(chatid, "spravka"),
		fmt.Sprintf("%s \n\n%s", d.getLang(chatid, "botUdalyaet"), d.getLang(chatid, "hhelpText")))
	return m.ID
}
