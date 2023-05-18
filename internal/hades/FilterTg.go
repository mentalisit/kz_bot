package hades

import (
	"kz_bot/internal/models"
	"kz_bot/internal/storage/CorpsConfig/hades"
)

func (h *Hades) filterTg(m models.MessageHades) {
	if h.ifComands(m) {
		return
	}
	h.tgConvertToMessage(models.Message{
		Text:        m.Text,
		Sender:      m.Sender,
		Avatar:      m.Avatar,
		ChannelType: m.ChannelType,
		Corporation: m.Corporation,
		Command:     m.Command,
		Messager:    m.Messager,
	})
}

func (h *Hades) tgConvertToMessage(msg models.Message) {
	ok, corp := hades.HadesStorage.AllianceName(msg.Corporation)
	if ok && msg.Command == "text" {
		sender := "(TG)" + msg.Sender

		if corp.DsChat != "" && msg.ChannelType == 0 {
			h.cl.Ds.SendWebhookForHades(msg.Text, sender, corp.DsChat, corp.GuildId, msg.Avatar)
		}
		if corp.DsChatWS1 != "" && msg.ChannelType == 1 {
			h.cl.Ds.SendWebhookForHades(msg.Text, sender, corp.DsChatWS1, corp.GuildId, msg.Avatar)
		}
	}
	if ok {
		h.toGame <- msg
	}
}
