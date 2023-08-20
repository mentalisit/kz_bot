package TelegramClient

import (
	"fmt"
	tgbotapi "github.com/matterbridge/telegram-bot-api/v6"
	"strconv"

	//tgbotapi "github.com/matterbridge/telegram-bot-api/v6"

	"kz_bot/internal/models"
	"strings"
)

func (t *Telegram) logicMix(m *tgbotapi.Message) {
	t.accesChatTg(m) //это была начальная функция при добавлени бота в группу

	ThreadID := m.MessageThreadID

	if m.MessageThreadID != 0 && m.MessageID-m.MessageThreadID < 10 {
		ThreadID = 0
	}
	ChatId := strconv.FormatInt(m.Chat.ID, 10) + fmt.Sprintf("/%d", ThreadID)

	fmt.Printf("   TG IN MessageID %d ThreadID %d ChatID %d Text %s \n",
		m.MessageID, ThreadID, m.Chat.ID, m.Text)

	// RsClient
	ok, config := t.CheckChannelConfigTG(ChatId)
	if ok {
		t.sendToFilterRs(m, config)
	}
	// client hs
	corpAlliance := t.getCorpHadesAlliance(ChatId)
	if corpAlliance.Corp != "" {
		t.sendToFilterHades(m, corpAlliance, 0)
		return
	}
	corpWs1 := t.getCorpHadesWs1(ChatId)
	if corpWs1.Corp != "" {
		t.sendToFilterHades(m, corpWs1, 1)
		return
	}

	tg, bridgeConfig := t.BridgeCheckChannelConfigTg(ChatId)
	if tg || strings.HasPrefix(m.Text, ".") {
		go t.SendToBridgeChatFilter(m, bridgeConfig)
	}
}
func (t *Telegram) sendToFilterHades(m *tgbotapi.Message, corp models.CorporationHadesClient, channelType int) {
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
func (t *Telegram) sendToFilterRs(m *tgbotapi.Message, config models.CorporationConfig) {
	name := t.nameNick(m.From.UserName, m.From.FirstName, config.TgChannel)
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

func (t *Telegram) SendToBridgeChatFilter(m *tgbotapi.Message, config models.BridgeConfig) {
	username := t.nameOrNick(m.From.UserName, m.From.FirstName)
	if len(m.Photo) != 0 {
		for _, ph := range m.Photo {
			url := t.GetPic(ph.FileID)
			m.Text = m.Text + url + " "
		}
		m.Text = m.Text + m.Caption
	}

	mes := models.BridgeMessage{
		Text:   m.Text,
		Sender: username,
		Tip:    "tg",
		Tg: models.BridgeMessageTg{
			ChatId:        strconv.FormatInt(m.Chat.ID, 10) + fmt.Sprintf("/%d", m.MessageThreadID),
			MesId:         m.MessageID,
			Avatar:        t.GetAvatar(m.From.ID),
			TimestampUnix: m.Time().Unix(),
			GroupName:     m.Chat.Title,
		},
	}
	if config.HostRelay == "" {

		//fmt.Printf(" logicmix.  %+v\n", mes)
		t.ChanBridgeMessage <- mes
		return
	}
	mes.Config = config

	if m.ReplyToMessage != nil && m.ReplyToMessage.Text != "" {
		mes.Tg.Reply.Text = m.ReplyToMessage.Text
		mes.Tg.Reply.UserName = t.nameOrNick(m.ReplyToMessage.From.UserName, m.ReplyToMessage.From.FirstName)
		mes.Tg.Reply.TimeMessage = m.ReplyToMessage.Time().Unix()
		mes.Tg.Reply.Avatar = t.GetAvatar(m.ReplyToMessage.From.ID)
	}
	t.ChanBridgeMessage <- mes
}
