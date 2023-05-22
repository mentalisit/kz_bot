package relay

import (
	"fmt"
	"kz_bot/internal/models"
	RelayDB "kz_bot/internal/storage/CorpsConfig/Relay"
	"strings"
)

func (r *Relay) logic() {
	if r.checkingForIdenticalMessage() {
		return
	}
	r.in.Text = filterMessageLinks(r.in.Text)
	tip := strings.ToUpper(r.in.Tip)
	username := fmt.Sprintf("%s ([%s]%s)", r.in.Author, tip, r.in.Config.RelayAlias)
	if tip == "DS" {
		var memory models.RelayMessageMemory
		memory.Timestamp = r.in.Ds.TimestampUnix
		memory.RelayName = r.in.Config.RelayName
		memory.MessageDs = append(memory.MessageDs, struct {
			MessageId string
			ChatId    string
		}{MessageId: r.in.Ds.MesId, ChatId: r.in.Ds.ChatId})
		for _, c := range *RelayDB.R {
			if c.RelayName == r.in.Config.RelayName && c.GuildName != r.in.Config.GuildName {
				if c.DsChannel != "" {
					var mes string
					if r.in.Ds.Reply.Text != "" {
						mes = r.client.Ds.SendWebhookReply(r.in.Text, username,
							c.DsChannel, c.GuildId, r.in.Ds.Avatar,
							r.in.Ds.Reply.Text,
							r.in.Ds.Reply.Avatar,
							r.in.Ds.Reply.UserName,
							r.in.Ds.Reply.TimeMessage)
					} else {
						texts := r.replaceTextMentionRsRole(r.in.Text, c.GuildId)
						mes = r.client.Ds.SendWebhook(texts, username,
							c.DsChannel, c.GuildId,
							r.in.Ds.Avatar)
					}
					memory.MessageDs = append(memory.MessageDs, struct {
						MessageId string
						ChatId    string
					}{
						MessageId: mes,
						ChatId:    c.DsChannel,
					})

				}
				if c.TgChannel != 0 {
					textTg := fmt.Sprintf("%s\n%s", username, r.in.Text)
					if r.in.Ds.Reply.Text != "" {
						textTg = fmt.Sprintf("%s\n%s\nReply: %s", username, r.in.Text, r.in.Ds.Reply.Text)
					}
					mesTg := r.client.Tg.SendChannel(c.TgChannel, textTg)
					memory.MessageTg = append(memory.MessageTg, struct {
						MessageId int
						ChatId    int64
					}{MessageId: mesTg, ChatId: c.TgChannel})
				}
			}

		}
		r.messages = append(r.messages, memory)
	} else if tip == "DEL" {
		r.RemoveMessage(r.in.Text)
	}
}
