package BridgeChat

import (
	"fmt"
	"kz_bot/internal/models"
	"regexp"
	"strings"
)

var messageTextAuthor [2]string

// проверка на повторное сообщение
func (b *Bridge) checkingForIdenticalMessage() bool {
	if messageTextAuthor[0] == b.in.Text && messageTextAuthor[1] == b.in.Sender {
		b.delIncomingMessage()
		return true
	}
	messageTextAuthor[0] = b.in.Text
	messageTextAuthor[1] = b.in.Sender
	return false
}

// удаление входящего сообщения
func (b *Bridge) delIncomingMessage() {
	if b.in.Tip == "ds" {
		go b.client.Ds.DeleteMessage(b.in.Ds.ChatId, b.in.Ds.MesId)
	} else if b.in.Tip == "tg" {
		go b.client.Tg.DelMessage(b.in.Tg.ChatId, b.in.Tg.MesId)
	}
}

// GetSenderName конконтенация имени
func (b *Bridge) GetSenderName() string {
	AliasName := ""
	if b.in.Tip == "ds" {
		for _, d := range b.in.Config.ChannelDs {
			if d.ChannelId == b.in.Ds.ChatId {
				AliasName = d.AliasName
			}
		}
	} else if b.in.Tip == "tg" {
		for _, d := range b.in.Config.ChannelTg {
			if d.ChannelId == b.in.Tg.ChatId {
				AliasName = d.AliasName
			}
		}
	}
	return fmt.Sprintf("%s ([%s]%s)", b.in.Sender, strings.ToUpper(b.in.Tip), AliasName)
}
func (b *Bridge) GuildName() string {
	if b.in.Tip == "ds" {
		return b.client.Ds.GuildChatName(b.in.Ds.ChatId, b.in.Ds.GuildId)
	}
	if b.in.Tip == "tg" {
		return b.in.Tg.GroupName
	}
	return ""
}

func (b *Bridge) ifTipDelSend(text string) {
	if b.in.Tip == "ds" {
		go b.client.Ds.SendChannelDelSecond(b.in.Ds.ChatId, "```"+text+"```", 30)
		go b.client.Ds.DeleteMessage(b.in.Ds.ChatId, b.in.Ds.MesId)
	} else if b.in.Tip == "tg" {
		go b.client.Tg.SendChannelDelSecond(b.in.Tg.ChatId, text, 30)
		go b.client.Tg.DelMessage(b.in.Tg.ChatId, b.in.Tg.MesId)
	}
}

func (b *Bridge) ifChannelTip(relay *models.BridgeConfig) {
	if b.in.Tip == "ds" {
		relay.ChannelDs = append(relay.ChannelDs, models.BridgeConfigDs{
			ChannelId:       b.in.Ds.ChatId,
			GuildId:         b.in.Ds.GuildId,
			CorpChannelName: b.GuildName(),
			AliasName:       "",
			MappingRoles:    map[string]string{},
		})
	}
	if b.in.Tip == "tg" {
		relay.ChannelTg = append(relay.ChannelTg, models.BridgeConfigTg{
			ChannelId:       b.in.Tg.ChatId,
			CorpChannelName: b.GuildName(),
			AliasName:       "",
			MappingRoles:    map[string]string{},
		})
	}
}

func (b *Bridge) replaceTextMentionRsRole(input, guildId string) string {
	//for _, s := range b.in.Config.Role {
	//	if strings.HasPrefix(b.in.Text, s) {
	//
	//	}
	//}

	re := regexp.MustCompile(`@&rs([4-9]|1[0-2])`)
	output := re.ReplaceAllStringFunc(input, func(s string) string {
		return b.client.Ds.TextToRoleRsPing(s[2:], guildId)
	})
	return output
}
