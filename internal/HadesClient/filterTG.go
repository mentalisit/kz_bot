package HadesClient

import (
	"kz_bot/internal/models"
)

func (h *Hades) filterTg(m models.MessageHades) {
	if h.ifComands(m) {
		return
	}
	in := models.MessageHadesClient{
		Text:        m.Text,
		Sender:      m.Sender,
		Avatar:      m.Avatar,
		ChannelType: m.ChannelType,
		Corporation: m.Corporation,
		Command:     m.Command,
		Messager:    m.Messager,
	}
	corp := h.getConfig(in.Corporation)
	if corp.Corp != "" && in.Command == "text" {
		sender := "(TG)" + in.Sender

		if corp.DsChat != "" && in.ChannelType == 0 {
			h.cl.Ds.SendWebhookForHades(in.Text, sender, corp.DsChat, corp.GuildId, in.Avatar)
		}
		if corp.DsChatWS1 != "" && in.ChannelType == 1 {
			h.cl.Ds.SendWebhookForHades(in.Text, sender, corp.DsChatWS1, corp.GuildId, in.Avatar)
		}
		h.toGame <- in
	}
}
