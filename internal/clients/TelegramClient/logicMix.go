package TelegramClient

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"kz_bot/internal/models"
	"kz_bot/internal/storage/CorpsConfig/hades"
	"kz_bot/internal/storage/memory"
	"strings"
)

func (t *Telegram) logicMix(m *tgbotapi.Message) {
	okAlliance, corp := hades.HadesStorage.AllianceChatTg(m.Chat.ID)
	if okAlliance {
		t.sendToFilterHades(m, corp, 0)
	}

	okWs1, corp := hades.HadesStorage.Ws1ChatTg(m.Chat.ID)
	if okWs1 {
		t.sendToFilterHades(m, corp, 1)
	}

	// тут я передаю чат айди и проверяю должен ли бот реагировать на этот чат
	ok, config := t.storage.Cache.CheckChannelConfigTG(m.Chat.ID)
	t.accesChatTg(m) //это была начальная функция при добавлени бота в группу
	if ok {
		t.sendToFilterRs(m, config)
	}
	good, relayConfig := t.storage.CorpsConfig.RelayCache.CheckChannelConfigTG(m.Chat.ID)
	if good || strings.HasPrefix(m.Text, ".") {
		t.SendToRelayChatFilter(m, relayConfig)
	}
}
func (t *Telegram) sendToFilterHades(m *tgbotapi.Message, corp models.Corporation, channelType int) {
	if m.Text != "" {
		if filterRsPl(m.Text) {
			return
		}
		mes := models.MessageHades{
			Text:        m.Text,
			Sender:      t.nameOrNick(m.From.UserName, m.From.FirstName),
			Avatar:      t.GetAvatar(m.From.ID),
			ChannelType: channelType,
			Corporation: corp.Corp,
			Command:     "text",
			Messager:    "tg",
			Tg: models.MessageHadesTg{
				MessageId: m.MessageID,
			},
		}
		t.ChanToGame <- mes
	}
}
func (t *Telegram) sendToFilterRs(m *tgbotapi.Message, config memory.CorpporationConfig) {
	name := t.nameNick(m.From.UserName, m.From.FirstName, m.Chat.ID)
	in := models.InMessage{
		Mtext:       m.Text,
		Tip:         "tg",
		Name:        name,
		NameMention: "@" + name,
		Tg: struct {
			Mesid  int
			Nameid int64
		}{
			Mesid:  m.MessageID,
			Nameid: m.From.ID,
		},
		Config: config,
		Option: models.Option{
			InClient: true,
		},
	}

	t.ChanRsMessage <- in
}
func (t *Telegram) SendToRelayChatFilter(m *tgbotapi.Message, config models.RelayConfig) {
	//username := t.nameOrNick(m.From.UserName, m.From.FirstName)
	fmt.Printf("\n\nMessage %+v\n\n", m)
	fmt.Printf("\n\nReplyToMessage %+v\n\n", m.ReplyToMessage)
	fmt.Printf("\n\nChat %+v\n\n", m.Chat)
	//if config.RelayName == "" && config.GuildName == "" {
	//	mes := models.RelayMessage{
	//		Text:   m.Text,
	//		Tip:    "tg",
	//		Author: username,
	//		Tg: models.RelayMessageTg{
	//			ChatId:        m.Chat.ID,
	//			MesId:         m.MessageID,
	//			Avatar:        t.GetAvatar(m.From.ID),
	//			GuildId:       m.Chat.ID,
	//			TimestampUnix: m.Date,
	//		},
	//	}
	//	//fmt.Printf(" logicmix.  %+v\n", mes)
	//	d.ChanRelay <- mes
	//	return
	//}
	//
	//if len(m.Attachments) > 0 {
	//	for _, attach := range m.Attachments { //вложеные файлы
	//		m.Content = m.Content + "\n" + attach.URL
	//	}
	//}
	//
	//mes := models.RelayMessage{
	//	Text:   d.replaceTextMessage(m.Content, m.GuildID),
	//	Tip:    "ds",
	//	Author: username,
	//	Ds: models.RelayMessageDs{
	//		ChatId:        m.ChannelID,
	//		MesId:         m.ID,
	//		Avatar:        m.Author.AvatarURL("128"),
	//		GuildId:       m.GuildID,
	//		TimestampUnix: m.Timestamp.Unix(),
	//	},
	//	Config: config,
	//}
	//if m.MessageReference != nil {
	//	usernameR := m.ReferencedMessage.Author.Username
	//	if m.ReferencedMessage.Member != nil && m.ReferencedMessage.Member.Nick != "" {
	//		usernameR = m.ReferencedMessage.Member.Nick
	//	}
	//	mes.Ds.Reply.UserName = usernameR
	//	mes.Ds.Reply.Text = d.replaceTextMessage(m.ReferencedMessage.Content, m.GuildID)
	//	mes.Ds.Reply.Avatar = m.ReferencedMessage.Author.AvatarURL("128")
	//	mes.Ds.Reply.TimeMessage = m.ReferencedMessage.Timestamp
	//}
	//
	//d.ChanRelay <- mes
}
