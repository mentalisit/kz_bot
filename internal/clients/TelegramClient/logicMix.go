package TelegramClient

import (
	"fmt"
	tgbotapi "github.com/musianisamuele/telegram-bot-api"
	"strconv"

	"kz_bot/internal/models"
	"strings"
)

func (t *Telegram) logicMix2(m *tgbotapi.Message) {
	t.accesChatTg(m) //это была начальная функция при добавлени бота в группу
	ThreadID := m.MessageThreadID
	if m.MessageThreadID != 0 && m.MessageID-m.MessageThreadID < 10 {
		ThreadID = 0
	}
	ChatId := strconv.FormatInt(m.Chat.ID, 10) + fmt.Sprintf("/%d", ThreadID)
	url := ""
	if len(m.Photo) > 0 {
		url, _ = t.t.GetFileDirectURL(m.Photo[len(m.Photo)-1].FileID)
		if m.Text == "" && m.Caption != "" {
			m.Text = m.Caption
		}
	}

	// RsClient
	ok, config := t.CheckChannelConfigTG(ChatId)
	if ok {
		name := t.nameNick(m.From.UserName, m.From.FirstName, config.TgChannel)
		in := models.InMessage{
			Mtext:       m.Text,
			Tip:         "tg",
			Name:        name,
			NameMention: "@" + name,
			Tg: struct {
				Mesid int
				//Nameid int64
			}{
				Mesid: m.MessageID,
				//Nameid: m.From.ID,
			},
			Config: config,
			Option: models.Option{
				InClient: true,
			},
		}

		t.ChanRsMessage <- in
	}

	// client hs
	corpAlliance := t.getCorpHadesAlliance(ChatId)
	if corpAlliance.Corp != "" {
		if m.Text != "" {
			if filterRsPl(m.Text) {
				return
			}
			mes := models.MessageHades{
				Text:        m.Text,
				Sender:      t.nameOrNick(m.From.UserName, m.From.FirstName),
				Avatar:      t.GetAvatar(m.From.ID),
				ChannelType: 0,
				Corporation: corpAlliance.Corp,
				Command:     "text",
				Messager:    "tg",
				MessageId:   strconv.Itoa(m.MessageID),
			}
			t.ChanToGame <- mes
		}
		return
	}
	corpWs1 := t.getCorpHadesWs1(ChatId)
	if corpWs1.Corp != "" {
		if m.Text != "" {
			if filterRsPl(m.Text) {
				return
			}
			mes := models.MessageHades{
				Text:        m.Text,
				Sender:      t.nameOrNick(m.From.UserName, m.From.FirstName),
				Avatar:      t.GetAvatar(m.From.ID),
				ChannelType: 1,
				Corporation: corpWs1.Corp,
				Command:     "text",
				Messager:    "tg",
				MessageId:   strconv.Itoa(m.MessageID),
			}
			t.ChanToGame <- mes
		}
		return
	}

	tg, bridgeConfig := t.BridgeCheckChannelConfigTg(ChatId)
	if tg || strings.HasPrefix(m.Text, ".") {
		go func() {
			username := t.nameOrNick(m.From.UserName, m.From.FirstName)
			mes := models.BridgeMessage{
				Text:    m.Text,
				Sender:  username,
				Tip:     "tg",
				FileUrl: url,
				Tg: models.BridgeMessageTg{
					ChatId:        ChatId,
					MesId:         m.MessageID,
					Avatar:        t.GetAvatar(m.From.ID),
					TimestampUnix: m.Time().Unix(),
					GroupName:     m.Chat.Title,
				},
			}
			if bridgeConfig.HostRelay == "" {
				t.ChanBridgeMessage <- mes
				return
			}
			mes.Config = bridgeConfig

			if m.ReplyToMessage != nil && m.ReplyToMessage.Text != "" {
				mes.Tg.Reply.Text = m.ReplyToMessage.Text
				mes.Tg.Reply.UserName = t.nameOrNick(m.ReplyToMessage.From.UserName, m.ReplyToMessage.From.FirstName)
				mes.Tg.Reply.TimeMessage = m.ReplyToMessage.Time().Unix()
				mes.Tg.Reply.Avatar = t.GetAvatar(m.ReplyToMessage.From.ID)
			}
			t.ChanBridgeMessage <- mes
		}()
	}
}
func (t *Telegram) logicMix(m *tgbotapi.Message) {
	t.accesChatTg(m) //это была начальная функция при добавлени бота в группу
	if m.Text == "" && m.Caption != "" {
		m.Text = m.Caption
		//len(m.Photo)
	}
	fmt.Printf("   tg message %+v \n", m)
	fmt.Printf("   tg messageFrom %+v \n", m.From)
	fmt.Printf("   tg messageChat %+v \n", m.Chat)
	fmt.Printf("   tg messageForwardFromChat %+v \n", m.ForwardFromChat)
	ThreadID := m.MessageThreadID

	if m.MessageThreadID != 0 && m.MessageID-m.MessageThreadID < 10 {
		ThreadID = 0
	}
	ChatId := strconv.FormatInt(m.Chat.ID, 10) + fmt.Sprintf("/%d", ThreadID)

	// RsClient
	ok, config := t.CheckChannelConfigTG(ChatId)
	if ok {
		t.sendToFilterRs(m, config)
	}
	// client hs
	//corpAlliance := t.getCorpHadesAlliance(ChatId)
	//if corpAlliance.Corp != "" {
	//	t.sendToFilterHades(m, corpAlliance, 0)
	//	return
	//}
	//corpWs1 := t.getCorpHadesWs1(ChatId)
	//if corpWs1.Corp != "" {
	//	t.sendToFilterHades(m, corpWs1, 1)
	//	return
	//}

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
			MessageId:   strconv.Itoa(m.MessageID),
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
			Mesid int
			//Nameid int64
		}{
			Mesid: m.MessageID,
			//Nameid: m.From.ID,
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
