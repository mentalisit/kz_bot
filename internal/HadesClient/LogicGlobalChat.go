package HadesClient

import "fmt"

func (h *Hades) logicGlobalChat() {
	mID := h.getChatIdAlliance()
	if h.in.MessageId > mID {
		h.cl.Ds.SendWebhookForHades(h.in.Text, h.in.Sender, "1120333468220530728", "700238199070523412", h.in.Avatar)
		fmt.Printf("GLOBAL: %s: %s\n", h.in.Sender, h.in.Text)
	}
	h.idMessage[h.in.Corporation] = h.in.MessageId
	err := h.storage.HadesClient.UpdateCorpMesId(h.in.Corporation, h.in.MessageId)
	if err != nil {
		h.log.Println("logicGlobalChat() " + err.Error())
		return
	}
}
