package hades

import (
	"kz_bot/internal/models"
	"kz_bot/internal/storage/CorpsConfig/hades"
)

func (h *Hades) filterGame(msg models.Message) {
	ok, corp := hades.HadesStorage.AllianceName(msg.Corporation)
	if msg.Corporation == "UKR Spase" {
		msg = numToRole(msg)
	}
	sender := "(🎮)" + msg.Sender
	if ok && msg.Command == "text" {
		if msg.ChannelType == 0 && corp.DsChat != "" {
			if h.ifEditMessage(msg, corp) {
				return
			}
			h.cl.Ds.SendWebhookForHades(msg.Text, sender, corp.DsChat, corp.GuildId, msg.Avatar)
		}
		if msg.ChannelType == 1 && corp.DsChatWS1 != "" {
			h.cl.Ds.SendWebhookForHades(msg.Text, sender, corp.DsChatWS1, corp.GuildId, msg.Avatar)
		}
		if msg.ChannelType == 2 && corp.DsChatWS2 != "" {
			h.cl.Ds.SendWebhookForHades(msg.Text, sender, corp.DsChatWS2, corp.GuildId, msg.Avatar)
		}

		text := "(🎮)" + msg.Sender + ": " + msg.Text
		if msg.ChannelType == 0 && corp.TgChat != 0 {
			if h.ifEditMessage(msg, corp) {
				return
			}
			h.cl.Tg.SendChannel(corp.TgChat, text)
		}
		if msg.ChannelType == 1 && corp.TgChatWS1 != 0 {
			h.cl.Tg.SendChannel(corp.TgChatWS1, text)
		}
	} else if ok && msg.Command != "text" {
		if msg.Command == "ответ ds" {
			mesid := h.cl.Ds.SendWebhookForHades(msg.Text, sender, corp.DsChat, corp.GuildId, msg.Avatar)
			h.cl.Ds.DeleteMesageSecond(corp.DsChat, mesid, 180)
		}
		if msg.Command == "ответ tg" {
			h.cl.Tg.SendChannelDelSecond(corp.TgChat, msg.Text, 180)
		}
	}
}
