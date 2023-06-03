package hades

import "kz_bot/internal/models"

func (h *Hades) delSendMessageIfTip(text string, m models.MessageHades, corporation models.Corporation) {
	if m.Messager == "ds" {
		go h.cl.Ds.SendChannelDelSecond(corporation.DsChat, "```"+text+"```", 10)
		go h.cl.Ds.DeleteMesageSecond(corporation.DsChat, m.Ds.MessageId, 10)
	}
	if m.Messager == "tg" {
		go h.cl.Tg.SendChannelDelSecond(corporation.TgChat, "<em>"+text+"</em>", 10)
		go h.cl.Tg.DelMessageSecond(corporation.TgChat, m.Tg.MessageId, 10)
	}
}
