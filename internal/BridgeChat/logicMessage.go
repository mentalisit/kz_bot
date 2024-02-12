package BridgeChat

import (
	"fmt"
	"kz_bot/internal/models"
	"strings"
	"sync"
	"time"
)

func (b *Bridge) logicMessage() {
	if b.checkingForIdenticalMessage() {
		return
	}
	if b.in.Tip == "delDs" {
		b.RemoveMessage()
		return
	}
	if b.in.Tip == "dse" {
		b.EditMessageDS()
		return
	}
	if b.in.Tip == "tge" {
		//b.EditMessageTG()// нужно исправить
		return
	}
	var memory models.BridgeTempMemory
	memory.RelayName = b.in.Config.NameRelay
	if b.ifTipDs(&memory) {
	} else if b.ifTipTg(&memory) {
	}
	go func() {
		time.Sleep(20 * time.Second)
		if b.in.Text != "" {
			b.log.InfoStruct("зависло ", b.in)
			b.log.InfoStruct("Reply", b.in.Reply)
		}
	}()
	memory.Wg.Wait()
	b.in = models.BridgeMessage{}
	b.messages = append(b.messages, memory)
}

func (b *Bridge) ifTipDs(memory *models.BridgeTempMemory) (ok bool) {
	if b.in.Tip == "ds" {
		ok = true
		memory.Wg.Add(1)
		memory.Timestamp = b.in.TimestampUnix
		memory.MessageDs = append(memory.MessageDs, models.MessageDs{
			MessageId: b.in.MesId,
			ChatId:    b.in.ChatId,
		})

		// Создаем WaitGroup для ожидания завершения всех горутин
		var wg sync.WaitGroup
		// Создаем канал для получения результатов (ID сообщений)
		resultChannelDs := make(chan models.MessageDs, 10)

		for _, d := range b.in.Config.ChannelDs {
			if d.ChannelId != b.in.ChatId {
				if d.ChannelId != "" {
					texts := b.replaceTextMentionRsRole(replaceTextMap(b.in.Text, d.MappingRoles), d.GuildId)
					wg.Add(1)
					if b.in.Reply != nil && b.in.Reply.Text != "" {
						go b.client.Ds.SendWebhookReplyAsync(texts, b.GetSenderName(), d.ChannelId, d.GuildId, b.in.Avatar, b.in.Reply, resultChannelDs, &wg)
					} else if b.in.FileUrl != "" {
						go b.client.Ds.SendFileAsync(texts, b.GetSenderName(), d.ChannelId, d.GuildId, b.in.FileUrl, b.in.Avatar, resultChannelDs, &wg)
					} else {
						go b.client.Ds.SendWebhookAsync(texts, b.GetSenderName(), d.ChannelId, d.GuildId, b.in.Avatar, resultChannelDs, &wg)
					}
				}
			}
		}

		// Создаем канал для получения результатов (ID сообщений)
		resultChannelTg := make(chan models.MessageTg, 10)
		for _, d := range b.in.Config.ChannelTg {
			if d.ChannelId != "" {
				text := replaceTextMap(b.in.Text, d.MappingRoles)
				textTg := fmt.Sprintf("%s\n%s", b.GetSenderName(), text)
				if b.in.Reply != nil && b.in.Reply.Text != "" {
					textTg = fmt.Sprintf("%s\n%s\nReply: %s", b.GetSenderName(), text, b.in.Reply.Text)
				}
				wg.Add(1)
				if b.in.FileUrl != "" {
					go b.client.Tg.SendFileFromURLAsync(d.ChannelId, textTg, b.in.FileUrl, resultChannelTg, &wg)
				} else {
					go b.client.Tg.SendChannelAsync(d.ChannelId, textTg, resultChannelTg, &wg)
				}
			}
		}
		go func() {
			wg.Wait()
			close(resultChannelDs)
			close(resultChannelTg)
			for value := range resultChannelDs {
				memory.MessageDs = append(memory.MessageDs, value)
			}
			for value := range resultChannelTg {
				memory.MessageTg = append(memory.MessageTg, value)
			}
			memory.Wg.Done()
		}()

	}
	return ok
}
func (b *Bridge) ifTipTg(memory *models.BridgeTempMemory) (ok bool) {
	if b.in.Tip == "tg" {
		ok = true
		memory.Wg.Add(1)
		memory.Timestamp = b.in.TimestampUnix
		memory.MessageTg = append(memory.MessageTg, struct {
			MessageId string
			ChatId    string
		}{MessageId: b.in.MesId, ChatId: b.in.ChatId})

		// Создаем WaitGroup для ожидания завершения всех горутин
		var wg sync.WaitGroup
		// Создаем канал для получения результатов (ID сообщений)
		resultChannelTg := make(chan models.MessageTg, 10)

		for _, c := range b.in.Config.ChannelTg {
			if c.ChannelId != b.in.ChatId {
				if c.ChannelId != "" {
					wg.Add(1)
					text := replaceTextMap(b.in.Text, c.MappingRoles)
					textTg := fmt.Sprintf("%s\n%s", b.GetSenderName(), text)
					if b.in.Reply != nil && (b.in.Reply.Text != "" || b.in.Reply.FileUrl != "") {
						if b.in.Reply.Text != "" {
							textTg = fmt.Sprintf("%s\n%s\nReply: %s", b.GetSenderName(), text, b.in.Reply.Text)
						} else if b.in.Reply.FileUrl != "" {
							go b.client.Tg.SendFileFromURLAsync(c.ChannelId, textTg, b.in.Reply.FileUrl, resultChannelTg, &wg)
						}
					} else if b.in.FileUrl != "" {
						go b.client.Tg.SendFileFromURLAsync(c.ChannelId, textTg, b.in.FileUrl, resultChannelTg, &wg)
					} else {
						go b.client.Tg.SendChannelAsync(c.ChannelId, textTg, resultChannelTg, &wg)
					}
				}
			}
		}

		// DS
		var chatids []string
		for _, d := range b.in.Config.ChannelDs {
			if d.ChannelId != "" {
				chatids = append(chatids, d.ChannelId)
			}
		}
		lenChannels := len(chatids)
		resultChannelDs := make(chan models.MessageDs, lenChannels)
		if b.in.Reply != nil && b.in.Reply.Text != "" && b.in.Reply.UserName == "gote1st_bot" {
			at := strings.SplitN(b.in.Reply.Text, "\n", 2)
			b.in.Reply.UserName = at[0]
			b.in.Reply.Text = at[1]
		}
		if lenChannels > 0 {
			wg.Add(lenChannels)
			b.client.Ds.SendBridgeAsync(b.in.Text, b.GetSenderName(), chatids, "", b.in.FileUrl, b.in.Avatar, b.in.Reply, resultChannelDs, &wg)
		}

		go func() {
			wg.Wait()
			close(resultChannelTg)
			close(resultChannelDs)
			for value := range resultChannelTg {
				memory.MessageTg = append(memory.MessageTg, value)
			}
			for value := range resultChannelDs {
				memory.MessageDs = append(memory.MessageDs, value)
			}
			memory.Wg.Done()
		}()
	}
	return ok
}
