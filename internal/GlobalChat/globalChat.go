package GlobalChat

import "C"
import (
	"fmt"
	"github.com/sirupsen/logrus"
	"kz_bot/internal/clients"
	"kz_bot/internal/models"
	"kz_bot/internal/storage"
	"kz_bot/internal/storage/memory"
	"strings"
)

type Chat struct {
	storage                   *storage.Storage
	client                    *clients.Clients
	log                       *logrus.Logger
	inbox                     chan models.InGlobalMessage
	in                        models.InGlobalMessage
	GlobalChatMemoryMessageId []models.MessageMemory
}

func NewChat(storage *storage.Storage, client *clients.Clients, log *logrus.Logger) *Chat {
	c := &Chat{storage: storage, client: client, log: log, inbox: client.GlobalChat}
	go c.loadInbox()
	go c.removeIfTimeDay()
	return c
}
func (c *Chat) loadInbox() {
	for {
		//ПОЛУЧЕНИЕ СООБЩЕНИЙ
		select {
		case in := <-c.inbox:
			fmt.Printf("in %+v\n", in)
			c.in = in
			c.logic()
		}
	}
}
func (c *Chat) logic() {
	if c.in.Tip == "ds" {
		tip := strings.ToUpper(c.in.Tip)
		var mem models.MessageMemory
		mem.Timestamp = c.in.Ds.TimestampUnix
		mem.Message = append(mem.Message, struct {
			MessageId string
			ChatId    string
		}{
			MessageId: c.in.Ds.MesId,
			ChatId:    c.in.Ds.ChatId,
		})
		for _, global := range *memory.G {
			if global.DsChannel != "" && global.DsChannel != c.in.Ds.ChatId {
				username := fmt.Sprintf("%s ([%s]%s)", c.in.Name, tip, c.in.Config.CorpName)
				var mes string
				if c.in.Ds.Reply.Text != "" {
					c.in.Ds.GuildId = global.GuildId
					c.in.Ds.ChatId = global.DsChannel
					mes = c.client.Ds.SendWebhookReply(c.in)

				} else {
					mes = c.client.Ds.SendWebhook(c.in.Content, username,
						global.DsChannel, global.GuildId,
						c.in.Ds.Avatar)
				}
				mem.Message = append(mem.Message, struct {
					MessageId string
					ChatId    string
				}{
					MessageId: mes,
					ChatId:    global.DsChannel,
				})
			}
		}
		c.GlobalChatMemoryMessageId = append(c.GlobalChatMemoryMessageId, mem)
	} else if c.in.Tip == "del" {
		c.RemoveMessage(c.in.Content)
	}
}
