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
func (r *Relay) RemoveMessage() {
	if len(r.messages) > 0 {
		var mem []models.RelayMessageMemory
		if r.in.Ds != nil {
			for _, memory := range r.messages {
				if r.ifMessageIdDs(memory, r.in.Ds.MesId) {
					for _, s := range memory.MessageDs {
						go r.client.Ds.DeleteMessage(s.ChatId, s.MessageId)
					}
					for _, s := range memory.MessageTg {
						go r.client.Tg.DelMessage(s.ChatId, s.MessageId)
					}
				} else {
					mem = append(mem, memory)
				}
			}
		}
		if r.in.Tg != nil {
			for _, memory := range r.messages {
				if r.ifMessageIdTg(memory, r.in.Tg.MesId) {
					for _, s := range memory.MessageTg {
						go r.client.Tg.DelMessage(s.ChatId, s.MessageId)
					}
					for _, s := range memory.MessageDs {
						go r.client.Ds.DeleteMessage(s.ChatId, s.MessageId)
					}
				} else {
					mem = append(mem, memory)
				}
			}
		}
		r.messages = mem
	}

}
func (r *Relay) ifMessageIdDs(memory models.RelayMessageMemory, MesId string) bool {
	for _, s := range memory.MessageDs {
		if s.MessageId == MesId {
			return true
		}
	}
	return false
}
func (r *Relay) ifMessageIdTg(memory models.RelayMessageMemory, MesId int) bool {
	for _, s := range memory.MessageTg {
		if s.MessageId == MesId {
			return true
		}
	}
	return false
}
