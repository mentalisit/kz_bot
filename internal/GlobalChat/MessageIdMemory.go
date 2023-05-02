package GlobalChat

import (
	"kz_bot/internal/models"
	"time"
)

func (c *Chat) removeIfTimeDay() {
	for {
		<-time.After(1 * time.Hour)
		if len(c.GlobalChatMemoryMessageId) > 0 {
			var mem []models.MessageMemory
			for _, memory := range c.GlobalChatMemoryMessageId {
				if time.Now().Unix()-memory.Timestamp < 86400 {
					mem = append(mem, memory)
				}
			}
			c.GlobalChatMemoryMessageId = mem
		}
	}
}
func (c *Chat) RemoveMessage(MesId string) {
	if len(c.GlobalChatMemoryMessageId) > 0 {
		var mem []models.MessageMemory
		for _, memory := range c.GlobalChatMemoryMessageId {
			if c.ifMessageId(memory, MesId) {
				for _, s := range memory.Message {
					go c.client.Ds.DeleteMessage(s.ChatId, s.MessageId)
				}
			} else {
				mem = append(mem, memory)
			}
		}
		c.GlobalChatMemoryMessageId = mem
	}
}
func (c *Chat) ifMessageId(memory models.MessageMemory, MesId string) bool {
	for _, s := range memory.Message {
		if s.MessageId == MesId {
			return true
		}
	}
	return false
}
