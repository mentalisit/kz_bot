package BridgeChat

func (b *Bridge) EditMessageDS() {
	if len(b.messages) > 0 {
		for _, memory := range b.messages {
			if b.ifMessageIdDs(memory, b.in.Ds.MesId) {
				for _, s := range memory.MessageDs {
					if b.in.Ds.MesId != s.MessageId {
						go b.client.Ds.EditWebhook(b.in.Text, b.in.Sender, s.ChatId, s.MessageId, "", b.in.Avatar)
					}
				}
				for _, s := range memory.MessageTg {
					go b.client.Tg.EditText(s.ChatId, s.MessageId, b.in.Text)
				}
			}
		}
	}
}
func (b *Bridge) EditMessageTG() {
	if len(b.messages) > 0 {
		for _, memory := range b.messages {
			if b.ifMessageIdTg(memory, b.in.Tg.MesId) {
				for _, s := range memory.MessageTg {
					if b.in.Tg.MesId != s.MessageId {
						go b.client.Tg.EditText(s.ChatId, s.MessageId, b.in.Text)
					}
				}
				for _, s := range memory.MessageDs {
					go b.client.Ds.EditWebhook(b.in.Text, b.in.Sender, s.ChatId, s.MessageId, "", b.in.Avatar)
				}
			}
		}
	}
}
