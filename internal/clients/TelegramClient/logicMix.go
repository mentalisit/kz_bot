package TelegramClient

import (
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
	username := t.nameOrNick(m.From.UserName, m.From.FirstName)
	if config.RelayName == "" && config.GuildName == "" {
		mes := models.RelayMessage{
			Text:   m.Text,
			Tip:    "tg",
			Author: username,
			Tg: &models.RelayMessageTg{
				ChatId:        m.Chat.ID,
				MesId:         m.MessageID,
				Avatar:        t.GetAvatar(m.From.ID),
				GuildId:       m.Chat.ID,
				TimestampUnix: m.Time().Unix(),
			},
		}
		//fmt.Printf(" logicmix.  %+v\n", mes)
		t.ChanRelay <- mes
		return
	}

	//if len(m.Attachments) > 0 {
	//	for _, attach := range m.Attachments { //вложеные файлы
	//		m.Content = m.Content + "\n" + attach.URL
	//	}
	//}

	mes := models.RelayMessage{
		Text:   m.Text,
		Tip:    "tg",
		Author: username,
		Tg: &models.RelayMessageTg{
			ChatId:        m.Chat.ID,
			MesId:         m.MessageID,
			Avatar:        t.GetAvatar(m.From.ID),
			GuildId:       m.Chat.ID,
			TimestampUnix: m.Time().Unix(),
		},
		Config: &config,
	}
	if m.ReplyToMessage != nil && m.ReplyToMessage.Text != "" {
		mes.Tg.Reply.Text = m.ReplyToMessage.Text
		mes.Tg.Reply.UserName = t.nameOrNick(m.ReplyToMessage.From.UserName, m.ReplyToMessage.From.FirstName)
		mes.Tg.Reply.TimeMessage = m.ReplyToMessage.Time().Unix()
		mes.Tg.Reply.Avatar = t.GetAvatar(m.ReplyToMessage.From.ID)
	}
	t.ChanRelay <- mes
}
