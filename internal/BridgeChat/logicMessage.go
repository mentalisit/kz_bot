package BridgeChat

import (
	"fmt"
	"kz_bot/internal/models"
	"strings"
	"sync"
)

func (b *Bridge) logicMessage() {
	if b.checkingForIdenticalMessage() {
		return
	}
	if b.in.Tip == "del" {
		b.RemoveMessage()
		return
	}
	var memory models.BridgeTempMemory
	memory.RelayName = b.in.Config.NameRelay
	if b.ifTipDs(&memory) {
	} else if b.ifTipTg(&memory) {
	}
	memory.Wg.Wait()
	b.messages = append(b.messages, memory)
}

func (b *Bridge) ifTipDs(memory *models.BridgeTempMemory) (ok bool) {
	if b.in.Tip == "ds" {
		ok = true
		memory.Wg.Add(1)
		memory.Timestamp = b.in.Ds.TimestampUnix
		memory.MessageDs = append(memory.MessageDs, models.MessageDs{
			MessageId: b.in.Ds.MesId,
			ChatId:    b.in.Ds.ChatId,
		})

		// Создаем WaitGroup для ожидания завершения всех горутин
		var wg sync.WaitGroup
		// Создаем канал для получения результатов (ID сообщений)
		resultChannel := make(chan models.MessageDs, 10)

		for _, d := range b.in.Config.ChannelDs {
			if d.ChannelId != b.in.Ds.ChatId {
				if d.ChannelId != "" {
					texts := b.replaceTextMentionRsRole(replaceTextMap(b.in.Text, d.MappingRoles), d.GuildId)
					wg.Add(1)
					if b.in.Ds.Reply != nil && b.in.Ds.Reply.Text != "" {
						go b.client.Ds.SendWebhookReplyAsync(texts, b.GetSenderName(), d.ChannelId, d.GuildId, b.in.Ds.Avatar, b.in.Ds.Reply, resultChannel, &wg)
					} else if b.in.FileUrl != "" {
						go b.client.Ds.SendFileAsync(texts, b.GetSenderName(), d.ChannelId, d.GuildId, b.in.FileUrl, b.in.Ds.Avatar, resultChannel, &wg)
					} else {
						go b.client.Ds.SendWebhookAsync(texts, b.GetSenderName(), d.ChannelId, d.GuildId, b.in.Ds.Avatar, resultChannel, &wg)
					}
				}
			}
		}

		go func() {
			wg.Wait()
			close(resultChannel)
			for value := range resultChannel {
				memory.MessageDs = append(memory.MessageDs, value)
			}
			memory.Wg.Done()
		}()

		for _, d := range b.in.Config.ChannelTg {
			if d.ChannelId != "" {
				var mesTg int
				var err error
				text := replaceTextMap(b.in.Text, d.MappingRoles)
				textTg := fmt.Sprintf("%s\n%s", b.GetSenderName(), text)
				if b.in.Ds.Reply != nil && b.in.Ds.Reply.Text != "" {
					textTg = fmt.Sprintf("%s\n%s\nReply: %s", b.GetSenderName(), text, b.in.Ds.Reply.Text)
				}
				if b.in.FileUrl != "" {
					mesTg, err = b.client.Tg.SendFileFromURL(d.ChannelId, textTg, b.in.FileUrl)
					if err != nil {
						b.log.Error(fmt.Sprintf("error sendFile in Channel %s error %s", d.ChannelId, err.Error()))
					}
				} else {
					mesTg = b.client.Tg.SendChannel(d.ChannelId, textTg)
				}

				memory.MessageTg = append(memory.MessageTg, struct {
					MessageId int
					ChatId    string
				}{MessageId: mesTg, ChatId: d.ChannelId})
			}
		}
	}
	return ok
}
func (b *Bridge) ifTipTg(memory *models.BridgeTempMemory) (ok bool) {
	if b.in.Tip == "tg" {
		ok = true
		memory.Timestamp = b.in.Tg.TimestampUnix
		memory.MessageTg = append(memory.MessageTg, struct {
			MessageId int
			ChatId    string
		}{MessageId: b.in.Tg.MesId, ChatId: b.in.Tg.ChatId})
		for _, c := range b.in.Config.ChannelTg {
			if c.ChannelId != b.in.Tg.ChatId {
				if c.ChannelId != "" {
					var mesTg int
					var err error
					text := replaceTextMap(b.in.Text, c.MappingRoles)
					textTg := fmt.Sprintf("%s\n%s", b.GetSenderName(), text)
					if b.in.Tg.Reply != nil && b.in.Tg.Reply.Text != "" {
						textTg = fmt.Sprintf("%s\n%s\nReply: %s", b.GetSenderName(), text, b.in.Tg.Reply.Text)
					}
					if b.in.FileUrl != "" {
						mesTg, err = b.client.Tg.SendFileFromURL(c.ChannelId, textTg, b.in.FileUrl)
						if err != nil {
							b.log.Error(fmt.Sprintf("error sendFile in Channel %s error %s", c.ChannelId, err.Error()))
						}
					} else {
						mesTg = b.client.Tg.SendChannel(c.ChannelId, textTg)
					}
					memory.MessageTg = append(memory.MessageTg, struct {
						MessageId int
						ChatId    string
					}{MessageId: mesTg, ChatId: c.ChannelId})

				}
			}
		}

		// Создаем WaitGroup для ожидания завершения всех горутин
		var wg sync.WaitGroup
		// Создаем канал для получения результатов (ID сообщений)
		resultChannel := make(chan models.MessageDs, 10)
		for _, d := range b.in.Config.ChannelDs {

			if d.ChannelId != "" {
				memory.Wg.Add(1)
				texts := b.replaceTextMentionRsRole(replaceTextMap(b.in.Text, d.MappingRoles), d.GuildId)
				wg.Add(1)
				if b.in.Tg.Reply != nil && b.in.Tg.Reply.Text != "" {
					if b.in.Tg.Reply.UserName == "gote1st_bot" {
						at := strings.SplitN(b.in.Tg.Reply.Text, "\n", 2)
						b.in.Tg.Reply.UserName = at[0]
						b.in.Tg.Reply.Text = at[1]
					}
					go b.client.Ds.SendWebhookReplyAsync(texts, b.GetSenderName(), d.ChannelId, d.GuildId, b.in.Tg.Avatar, (*models.ReplyDs)(b.in.Tg.Reply), resultChannel, &wg)
				} else if b.in.FileUrl != "" {
					go b.client.Ds.SendFileAsync(texts, b.GetSenderName(), d.ChannelId, d.GuildId, b.in.FileUrl, b.in.Tg.Avatar, resultChannel, &wg)
				} else {
					go b.client.Ds.SendWebhookAsync(texts, b.GetSenderName(), d.ChannelId, d.GuildId, b.in.Tg.Avatar, resultChannel, &wg)
				}
			}

		}

		go func() {
			wg.Wait()
			close(resultChannel)
			for value := range resultChannel {
				memory.MessageDs = append(memory.MessageDs, value)
			}
			memory.Wg.Done()
		}()
	}
	return ok
}
