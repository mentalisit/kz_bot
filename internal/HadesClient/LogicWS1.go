package HadesClient

import "fmt"

func (h *Hades) logicWs1() {
	if h.in.Command == "text" {
		if h.in.MessageId > h.getChatIdWs1() {
			config := h.getConfig(h.in.Corporation)
			sender := "(🎮)" + h.in.Sender
			if config.DsChatWS1 != "" {
				go h.cl.Ds.SendWebhookForHades(h.in.Text, sender, config.DsChatWS1, config.GuildId, h.in.Avatar)
			}
			if config.TgChatWS1 != "" {
				go h.cl.Tg.SendChannel(config.TgChatWS1, sender+"\n"+h.in.Text)
			}
			fmt.Printf("ws1 %s %s: %s\n", h.in.Corporation, h.in.Sender, h.in.Text)
			h.idMessageWs1[h.in.SolarSystemId] = h.in.MessageId
			h.storage.HadesClient.UpdateWs1MesId(h.in.Corporation, h.in.MessageId, h.in.SolarSystemId)
		} else if h.in.MessageId == -1 {
			h.logicWs1Module()
		}
	}
}
func (h *Hades) logicWs1Module() {
	config := h.getConfig(h.in.Corporation)
	sender := "(🎮)" + h.in.Sender
	text := h.in.Text

	//mention user ws
	if h.in.Corporation == "Слава Украине!" {
		text = h.in.Text + "\n @YaroslavShelest @Yura_Odessa  @GnarlyV "
	}
	if config.DsChatWS1 != "" {
		go h.cl.Ds.SendWebhookForHades(text, sender, config.DsChatWS1, config.GuildId, h.in.Avatar)
	}
	if config.TgChatWS1 != "" {
		go h.cl.Tg.SendChannel(config.TgChatWS1, sender+"\n"+text)
	}
	if config.TgChatWS1 == "" && config.DsChatWS1 == "" {
		if config.DsChat != "" {
			go h.cl.Ds.SendWebhookForHades(text, sender, config.DsChat, config.GuildId, h.in.Avatar)
		}
		if config.TgChat != "" {
			go h.cl.Tg.SendChannel(config.TgChat, sender+"\n"+text)
		}
	}
	fmt.Printf("ws1module %s %s: %s\n", h.in.Corporation, h.in.Sender, h.in.Text)
}
func (h *Hades) getChatIdWs1() (mId int64) {
	if h.idMessageWs1[h.in.SolarSystemId] != 0 {
		mId = h.idMessageWs1[h.in.SolarSystemId]
	} else {
		if h.in.Corporation != "" {
			mId = h.storage.HadesClient.GetWs1MesId(h.in.Corporation, h.in.SolarSystemId)
		}
	}
	return mId
}
