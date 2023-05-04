package GlobalChat

import "C"
import (
	"fmt"
	"github.com/sirupsen/logrus"
	"kz_bot/internal/clients"
	"kz_bot/internal/models"
	"kz_bot/internal/storage"
	"kz_bot/internal/storage/memory"
	"regexp"
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
		c.in.Content = filterMessageLinks(c.in.Content)
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
					mes = c.client.Ds.SendWebhookReply(c.in.Content, username,
						global.DsChannel, global.GuildId, c.in.Ds.Avatar,
						c.in.Ds.Reply.Text,
						c.in.Ds.Reply.Avatar,
						c.in.Ds.Reply.UserName,
						c.in.Ds.Reply.TimeMessage)
				} else {
					texts := c.replaceTextMentionRsRole(c.in.Content, global.GuildId)
					mes = c.client.Ds.SendWebhook(texts, username,
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
func filterMessageLinks(input string) string {
	// Регулярное выражение для поиска ссылок
	re := regexp.MustCompile(`(https?://[^\s]+)`)
	// Список разрешенных ссылок
	allowedLinks := []string{
		"https://t.me/",
		"https://discord.com/channels/",
		"https://discord.gg/",
		"https://userxinos.github.io/",
	}
	// Запрещенная ссылка
	forbiddenLink := "запрещенная ссылка"
	// Заменяем все ссылки, кроме разрешенных, на запрещенную ссылку
	output := re.ReplaceAllStringFunc(input, func(link string) string {
		for _, allowedLink := range allowedLinks {
			if strings.HasPrefix(link, allowedLink) {
				return link
			}
		}
		return forbiddenLink
	})
	return output
}
func (c *Chat) replaceTextMentionRsRole(input, guildId string) string {
	re := regexp.MustCompile(`@&rs([4-9]|1[0-2])`)
	output := re.ReplaceAllStringFunc(input, func(s string) string {
		return c.client.Ds.TextToRoleRsPing(s[2:], guildId)
	})
	return output
}
