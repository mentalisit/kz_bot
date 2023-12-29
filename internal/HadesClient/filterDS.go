package HadesClient

//func (h *Hades) filterDs(m models.MessageHades) {
//	if h.ifComands(m) {
//		return
//	}
//	in := models.MessageHadesClient{
//		Text:        m.Text,
//		Sender:      m.Sender,
//		Avatar:      m.Avatar,
//		ChannelType: m.ChannelType,
//		Corporation: m.Corporation,
//		Command:     m.Command,
//		Messager:    m.Messager,
//	}
//	corp := h.getConfig(in.Corporation)
//	if corp.Corp != "" && in.Command == "text" {
//		text := "(DS)" + in.Sender + ": " + in.Text
//
//		if corp.TgChat != "" && in.ChannelType == 0 {
//			h.cl.Tg.SendChannel(corp.TgChat, text)
//		}
//		if corp.TgChatWS1 != "" && in.ChannelType == 1 {
//			h.cl.Tg.SendChannel(corp.TgChatWS1, text)
//		}
//		h.toGame <- in
//	}
//}
