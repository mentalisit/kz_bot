package relay

import (
	"kz_bot/internal/models"
	"time"
)

func (r *Relay) removeIfTimeDay() {
	for {
		<-time.After(1 * time.Hour)
		if len(r.messages) > 0 {
			var mem []models.RelayMessageMemory
			for _, memory := range r.messages {
				if time.Now().Unix()-memory.Timestamp < 86400 {
					mem = append(mem, memory)
				}
			}
			r.messages = mem
		}
	}
}
func (r *Relay) RemoveMessage(MesId string) {
	if len(r.messages) > 0 {
		var mem []models.RelayMessageMemory
		for _, memory := range r.messages {
			if r.ifMessageId(memory, MesId) {
				for _, s := range memory.MessageDs {
					go r.client.Ds.DeleteMessage(s.ChatId, s.MessageId)
				}
			} else {
				mem = append(mem, memory)
			}
		}
		r.messages = mem
	}
}
func (r *Relay) ifMessageId(memory models.RelayMessageMemory, MesId string) bool {
	for _, s := range memory.MessageDs {
		if s.MessageId == MesId {
			return true
		}
	}
	return false
}
