package HadesClient

import (
	"fmt"
)

func (h *Hades) logicAlliance() {
	sender := "(🎮)" + h.in.Sender
	config := h.getConfig(h.in.Corporation)
	if h.in.Command == "text" {
		if h.in.Corporation == "UKR Spase" {
			h.numToRole()
		}

		if h.in.MessageId > h.getChatIdAlliance() {
			h.idMessage[h.in.Corporation] = h.in.MessageId
			err := h.storage.HadesClient.UpdateCorpMesId(h.in.Corporation, h.in.MessageId)
			if err != nil {
				h.log.Println("logicAlliance() " + err.Error())
				return
			}
			if config.DsChat != "" {
				if h.ifEditMessage(h.in, config) {
					return
				}
				go h.cl.Ds.SendWebhookForHades(h.in.Text, sender, config.DsChat, config.GuildId, h.in.Avatar)
			}
			if config.TgChat != 0 {
				go h.cl.Tg.SendChannel(config.TgChat, sender+"\n"+h.in.Text)
			}
			fmt.Printf("Alliance %s Name %s: %s\n", h.in.Corporation, h.in.Sender, h.in.Text)

		} else if h.in.MessageId == h.getChatIdAlliance() {
			if config.DsChat != "" {
				h.ifEditMessage(h.in, config)
			}
		}
	} else if h.in.Command != "text" {
		if h.in.Command == "ответ ds" {
			mesid := h.cl.Ds.SendWebhookForHades(h.in.Text, sender, config.DsChat, config.GuildId, h.in.Avatar)
			go h.cl.Ds.DeleteMesageSecond(config.DsChat, mesid, 180)
		}
		if h.in.Command == "ответ tg" {
			go h.cl.Tg.SendChannelDelSecond(config.TgChat, h.in.Text, 180)
		}
		if h.in.Command == "access" {
			h.CheckMember(h.in.Sender, h.in.Corporation, h.in.MessageId)
		}
		if h.in.Command == "rang" {
			h.CheckMemberRang(h.in.Sender, h.in.Corporation, h.in.MessageId)
		}
	}
}
