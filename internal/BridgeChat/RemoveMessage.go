package BridgeChat

import (
	"kz_bot/internal/models"
	"time"
)

func (b *Bridge) removeIfTimeDay() {
	for {
		<-time.After(1 * time.Hour)
		if len(b.messages) > 0 {
			var mem []models.BridgeTempMemory
			for _, memory := range b.messages {
				if time.Now().Unix()-memory.Timestamp < 86400 {
					mem = append(mem, memory)
				}
			}
			b.messages = mem
		}
	}
}
func (b *Bridge) RemoveMessage() {
	if len(b.messages) > 0 {
		var mem []models.BridgeTempMemory
		for _, memory := range b.messages {
			if b.ifMessageIdDs(memory, b.in.Ds.MesId) {
				for _, s := range memory.MessageDs {
					go b.client.Ds.DeleteMessage(s.ChatId, s.MessageId)
				}
				for _, s := range memory.MessageTg {
					go b.client.Tg.DelMessage(s.ChatId, s.MessageId)
				}
			} else {
				mem = append(mem, memory)
			}
		}
		b.messages = mem
	}
}

func (b *Bridge) ifMessageIdDs(memory models.BridgeTempMemory, MesId string) bool {
	for _, s := range memory.MessageDs {
		if s.MessageId == MesId {
			return true
		}
	}
	return false
}
func (b *Bridge) ifMessageIdTg(memory models.BridgeTempMemory, MesId int) bool {
	for _, s := range memory.MessageTg {
		if s.MessageId == MesId {
			return true
		}
	}
	return false
}
