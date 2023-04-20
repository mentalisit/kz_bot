package DiscordClient

import (
	"context"
	"fmt"
	"kz_bot/pkg/utils"
	"time"
)

func (d *Discord) Help(Channel string) {
	mId := d.hhelp1(Channel)
	d.DeleteMesageSecond(Channel, mId, 180)
}

func (d *Discord) Autohelpds() {
	tm := time.Now()
	mtime := tm.Format("15:04")
	if mtime == "12:00" {
		a := d.storage.CorpsConfig.AutoHelp()
		for _, s := range a {
			if s.DsChannel != "" {
				if s.MesidDsHelp != "" {
					go d.DeleteMessage(s.DsChannel, s.MesidDsHelp)
					d.HelpChannelUpdate(s.DsChannel)
				} else {
					d.HelpChannelUpdate(s.DsChannel)
				}
			}

		}
		time.Sleep(time.Minute)
	} else if mtime == "03:00" {
		time.Sleep(1 * time.Second)
		utils.UpdateRun()
	}
}

func (d *Discord) HelpChannelUpdate(dschannel string) {
	newMesidHelp := d.hhelp1(dschannel)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	d.storage.CorpsConfig.AutoHelpUpdateMesid(ctx, newMesidHelp, dschannel)
}

func (d *Discord) hhelp1(chatid string) string {
	m := d.SendEmbedText(chatid, d.getLang(chatid, "spravka"),
		fmt.Sprintf("%s \n\n%s", d.getLang(chatid, "botUdalyaet"), d.getLang(chatid, "hhelpText")))
	return m.ID
}
