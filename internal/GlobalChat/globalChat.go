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
	storage *storage.Storage
	client  *clients.Clients
	log     *logrus.Logger
	inbox   chan models.InGlobalMessage
	in      models.InGlobalMessage
}

func NewChat(storage *storage.Storage, client *clients.Clients, log *logrus.Logger) *Chat {
	c := &Chat{storage: storage, client: client, log: log, inbox: client.GlobalChat}
	c.loadInbox()
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
		for _, global := range *memory.G {
			if global.DsChannel != "" && global.DsChannel != c.in.Ds.ChatId {
				username := fmt.Sprintf("%s ([%s]%s)", c.in.Name, tip, c.in.Config.CorpName)
				if c.in.Ds.Reply.Reply.Text != "" {
					reply := c.in.Ds.Reply
					reply.GuildId = global.GuildId
					reply.ChatId = global.DsChannel
					c.client.Ds.SendWebhookReply(reply)
				} else {
					c.client.Ds.SendWebhook(c.in.Content, username, global.DsChannel, global.GuildId, c.in.Ds.Avatar)
				}
			}
		}
	}
}
