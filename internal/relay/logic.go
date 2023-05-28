package relay

import (
	"fmt"
	"kz_bot/internal/models"
	RelayDB "kz_bot/internal/storage/CorpsConfig/Relay"
	"strings"
	"time"
)

func (r *Relay) logic() {
	if strings.HasPrefix(r.in.Text, ".") {
		if r.ifCommand() {
			return
		}
	}
	if r.in.Config != nil {
		r.logicSend()
	}

}

func (r *Relay) logicSend() {
	if r.checkingForIdenticalMessage() {
		return
	}
	r.in.Text = filterMessageLinks(r.in.Text)
	tip := strings.ToUpper(r.in.Tip)
	username := fmt.Sprintf("%s ([%s]%s)", r.in.Author, tip, r.in.Config.GuildName)
	var memory models.RelayMessageMemory
	memory.RelayName = r.in.Config.RelayName
	if tip == "DS" {
		memory.Timestamp = r.in.Ds.TimestampUnix
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
	} else if tip == "TG" {
		memory.Timestamp = r.in.Tg.TimestampUnix
		memory.MessageTg = append(memory.MessageTg, struct {
			MessageId int
			ChatId    int64
		}{MessageId: r.in.Tg.MesId, ChatId: r.in.Tg.ChatId})
		for _, c := range *RelayDB.R {
			if c.RelayName == r.in.Config.RelayName && c.GuildName != r.in.Config.GuildName {
				if c.TgChannel != 0 {
					textTg := fmt.Sprintf("%s\n%s", username, r.in.Text)
					if r.in.Tg.Reply.Text != "" {
						textTg = fmt.Sprintf("%s\n%s\nReply: %s", username, r.in.Text, r.in.Tg.Reply.Text)
					}
					mesTg := r.client.Tg.SendChannel(c.TgChannel, textTg)
					memory.MessageTg = append(memory.MessageTg, struct {
						MessageId int
						ChatId    int64
					}{MessageId: mesTg, ChatId: c.TgChannel})

				}
				if c.DsChannel != "" {
					var mes string
					if r.in.Tg.Reply.Text != "" {
						if r.in.Tg.Reply.UserName == "gote1st_bot" {
							at := strings.SplitN(r.in.Tg.Reply.Text, "\n", 2)
							r.in.Tg.Reply.UserName = at[0]
							r.in.Tg.Reply.Text = at[1]
						}

						mes = r.client.Ds.SendWebhookReply(r.in.Text, username,
							c.DsChannel, c.GuildId, r.in.Tg.Avatar,
							r.in.Tg.Reply.Text,
							r.in.Tg.Reply.Avatar,
							r.in.Tg.Reply.UserName,
							time.Unix(r.in.Tg.Reply.TimeMessage, 0))
					} else {
						texts := r.replaceTextMentionRsRole(r.in.Text, c.GuildId)
						mes = r.client.Ds.SendWebhook(texts, username,
							c.DsChannel, c.GuildId,
							r.in.Tg.Avatar)
					}
					memory.MessageDs = append(memory.MessageDs, struct {
						MessageId string
						ChatId    string
					}{
						MessageId: mes,
						ChatId:    c.DsChannel,
					})

				}
			}

		}
		r.messages = append(r.messages, memory)
	} else if tip == "DEL" {
		r.RemoveMessage()
	}
}

//if d.blackListFilter(m.Author.ID) {
//d.DeleteMesageSecond(m.ChannelID, m.ID, 5)
//return
//}
//if d.ifAsksForRoleRs(m) {
//go d.DeleteMessage(m.ChannelID, m.ID)
//return
//}
//if ifPrefix(m.Content) {
//return
//}
//username := m.Author.Username
//if m.Member != nil && m.Member.Nick != "" {
//username = m.Member.Nick
//}
//if len(m.Attachments) > 0 {
//for _, attach := range m.Attachments { //вложеные файлы
//m.Content = m.Content + "\n" + attach.URL
//}
//}
