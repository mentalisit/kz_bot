package BridgeChat

import (
	"fmt"
	"kz_bot/internal/models"
	"strings"
)

func (b *Bridge) logicMessage() {
	if b.checkingForIdenticalMessage() {
		return
	}
	var memory models.BridgeTempMemory
	memory.RelayName = b.in.Config.NameRelay
	if b.in.Tip == "ds" {
		memory.Timestamp = b.in.Ds.TimestampUnix
		memory.MessageDs = append(memory.MessageDs, struct {
			MessageId string
			ChatId    string
		}{MessageId: b.in.Ds.MesId, ChatId: b.in.Ds.ChatId})
		for _, d := range b.in.Config.ChannelDs {
			if d.ChannelId != b.in.Ds.ChatId {
				if d.ChannelId != "" {
					var mes string
					text := replaceTextMap(b.in.Text, d.MappingRoles)
					if b.in.Ds.Reply.Text != "" {
						mes = b.client.Ds.SendWebhookReply(text, b.GetSenderName(),
							d.ChannelId, d.GuildId, b.in.Ds.Avatar,
							b.in.Ds.Reply.Text,
							b.in.Ds.Reply.Avatar,
							b.in.Ds.Reply.UserName,
							b.in.Ds.Reply.TimeMessage)
					} else {
						texts := b.replaceTextMentionRsRole(text, d.GuildId)
						mes = b.client.Ds.SendWebhook(texts, b.GetSenderName(),
							d.ChannelId, d.GuildId,
							b.in.Ds.Avatar)
					}
					memory.MessageDs = append(memory.MessageDs, struct {
						MessageId string
						ChatId    string
					}{
						MessageId: mes,
						ChatId:    d.ChannelId,
					})

				}
			}
		}
		for _, d := range b.in.Config.ChannelTg {
			if d.ChannelId != "" {
				text := replaceTextMap(b.in.Text, d.MappingRoles)
				textTg := fmt.Sprintf("%s\n%s", b.GetSenderName(), text)
				if b.in.Ds.Reply.Text != "" {
					textTg = fmt.Sprintf("%s\n%s\nReply: %s", b.GetSenderName(), text, b.in.Ds.Reply.Text)
				}
				mesTg := b.client.Tg.SendChannel(d.ChannelId, textTg)
				memory.MessageTg = append(memory.MessageTg, struct {
					MessageId int
					ChatId    string
				}{MessageId: mesTg, ChatId: d.ChannelId})
			}
		}
	} else if b.in.Tip == "tg" {
		memory.Timestamp = b.in.Tg.TimestampUnix
		memory.MessageTg = append(memory.MessageTg, struct {
			MessageId int
			ChatId    string
		}{MessageId: b.in.Tg.MesId, ChatId: b.in.Tg.ChatId})
		for _, c := range b.in.Config.ChannelTg {
			if c.ChannelId != b.in.Tg.ChatId {
				if c.ChannelId != "" {
					text := replaceTextMap(b.in.Text, c.MappingRoles)
					textTg := fmt.Sprintf("%s\n%s", b.GetSenderName(), text)
					if b.in.Tg.Reply.Text != "" {
						textTg = fmt.Sprintf("%s\n%s\nReply: %s", b.GetSenderName(), text, b.in.Tg.Reply.Text)
					}
					mesTg := b.client.Tg.SendChannel(c.ChannelId, textTg)
					memory.MessageTg = append(memory.MessageTg, struct {
						MessageId int
						ChatId    string
					}{MessageId: mesTg, ChatId: c.ChannelId})

				}
			}
		}
		for _, d := range b.in.Config.ChannelDs {
			if d.ChannelId != "" {
				text := replaceTextMap(b.in.Text, d.MappingRoles)
				var mes string
				if b.in.Tg.Reply.Text != "" {
					if b.in.Tg.Reply.UserName == "gote1st_bot" {
						at := strings.SplitN(b.in.Tg.Reply.Text, "\n", 2)
						b.in.Tg.Reply.UserName = at[0]
						b.in.Tg.Reply.Text = at[1]
					}
					mes = b.client.Ds.SendWebhookReply(text, b.GetSenderName(),
						d.ChannelId, d.GuildId, b.in.Tg.Avatar,
						b.in.Tg.Reply.Text,
						b.in.Tg.Reply.Avatar,
						b.in.Tg.Reply.UserName,
						b.in.Tg.Reply.TimeMessage)
				} else {
					texts := b.replaceTextMentionRsRole(text, d.GuildId)
					mes = b.client.Ds.SendWebhook(texts, b.GetSenderName(),
						d.ChannelId, d.GuildId,
						b.in.Tg.Avatar)
				}
				memory.MessageDs = append(memory.MessageDs, struct {
					MessageId string
					ChatId    string
				}{
					MessageId: mes,
					ChatId:    d.ChannelId,
				})

			}
		}
	} else if b.in.Tip == "del" {
		b.RemoveMessage()
	}
	b.messages = append(b.messages, memory)
}
