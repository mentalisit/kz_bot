package hades

import (
	"kz_bot/internal/models"
	"kz_bot/internal/storage/CorpsConfig/hades"
)

func (h *Hades) filterDs(m models.MessageHades) {
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
func (h *Hades) dsConvertToMessage(msg models.Message) {
	ok, corp := hades.HadesStorage.AllianceName(msg.Corporation)
	if ok && msg.Command == "text" {
		if corp.TgChat != 0 && msg.ChannelType == 0 {
			text := "(DS)" + msg.Sender + ": " + msg.Text
			h.cl.Tg.SendChannel(corp.TgChat, text)
		}
		if corp.TgChatWS1 != 0 && msg.ChannelType == 1 {
			text := "(DS)" + msg.Sender + ": " + msg.Text
			h.cl.Tg.SendChannel(corp.TgChatWS1, text)
		}
	}
	if ok {
		h.toGame <- msg
	}
}
