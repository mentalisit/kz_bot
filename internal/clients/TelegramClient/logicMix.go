package TelegramClient

import (
	"fmt"
	tgbotapi "github.com/samuelemusiani/telegram-bot-api"
	"strconv"

	"kz_bot/internal/models"
	"strings"
)

func (t *Telegram) logicMix2(m *tgbotapi.Message, edit bool) {
	t.accesChatTg(m) //это была начальная функция при добавлени бота в группу
	ThreadID := m.MessageThreadID
	if !m.IsTopicMessage && ThreadID != 0 {
		ThreadID = 0
	}
	ChatId := strconv.FormatInt(m.Chat.ID, 10) + fmt.Sprintf("/%d", ThreadID)
	url := t.handleDownload(m)
	if m.Text == "" && m.Caption != "" {
		m.Text = m.Caption
	}
	t.handlePoll(m)
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
			}{
				Mesid: m.MessageID,
			},
			Config: config,
			Option: models.Option{
				InClient: true,
			},
		}

		t.ChanRsMessage <- in
	}

	tg, bridgeConfig := t.BridgeCheckChannelConfigTg(ChatId)

	if strings.HasPrefix(m.Text, ".") {
		go func() {
			username := t.nameOrNick(m.From.UserName, m.From.FirstName)
			mes := models.BridgeMessage{
				Text:   m.Text,
				Sender: username,
				Tip:    "tg",
				ChatId: ChatId,
				MesId:  strconv.Itoa(m.MessageID),
			}
			t.ChanBridgeMessage <- mes
		}()
	}
	if tg {
		go func() {
			if m.Document != nil {
				url, _ = t.t.GetFileDirectURL(m.Document.FileID)
				if m.Text == "" {
					m.Text = m.Document.FileName
				}
			}
			username := t.nameOrNick(m.From.UserName, m.From.FirstName)
			mes := models.BridgeMessage{
				Text:          m.Text,
				Sender:        username,
				Tip:           "tg",
				FileUrl:       url,
				Avatar:        t.GetAvatar(m.From.ID),
				ChatId:        ChatId,
				MesId:         strconv.Itoa(m.MessageID),
				TimestampUnix: m.Time().Unix(),
				GuildId:       m.Chat.Title,
				Config:        &bridgeConfig,
			}

			if m.ReplyToMessage != nil && m.ReplyToMessage.Text != "" {
				mes.Reply = &models.BridgeMessageReply{
					Text:        m.ReplyToMessage.Text,
					UserName:    t.nameOrNick(m.ReplyToMessage.From.UserName, m.ReplyToMessage.From.FirstName),
					TimeMessage: m.ReplyToMessage.Time().Unix(),
					Avatar:      t.GetAvatar(m.ReplyToMessage.From.ID),
				}
			}
			if m.ForwardFrom != nil {
				forwardName := t.nameOrNick(m.ForwardFrom.UserName, m.ForwardFrom.FirstName)
				text := fmt.Sprintf("Пересланное сообщение от %s \n %s ", forwardName, mes.Text)
				mes.Text = text
			}
			if edit {
				mes.Tip = "tge"
			}

			t.ChanBridgeMessage <- mes
		}()
	}
}

//func (t *Telegram) logicMix(m *tgbotapi.Message) {
//	t.accesChatTg(m) //это была начальная функция при добавлени бота в группу
//	if m.Text == "" && m.Caption != "" {
//		m.Text = m.Caption
//		//len(m.Photo)
//	}
//	ThreadID := m.MessageThreadID
//	if !m.IsTopicMessage && ThreadID != 0 {
//		ThreadID = 0
//	}
//
//	ttt := fmt.Sprintf("   tg new message chat %s OT %s text %s", m.Chat.Title, m.From, m.Text)
//	if m.ReplyToMessage != nil {
//		ttt += fmt.Sprintf(" ответ %s на сообщение %s", m.ReplyToMessage.From, m.ReplyToMessage.Text)
//	}
//	if len(m.Photo) > 0 {
//		ttt += fmt.Sprintf(" вложено фото %+v", m.Photo[len(m.Photo)-1])
//	}
//	fmt.Printf(ttt+" ThreadID %d\n", ThreadID)
//	fmt.Printf("   tg message %+v \n\n", m)
//
//	if m.ReplyToMessage != nil {
//		fmt.Printf("   tg messageReplyToMessage %+v \n\n", m.ReplyToMessage)
//		if m.ReplyToMessage.ForumTopicCreated != nil {
//			fmt.Printf("   tg messageReplyToMessage.ForumTopicCreated %+v \n\n", m.ReplyToMessage.ForumTopicCreated)
//		}
//	}
//	if m.ForwardFromChat != nil {
//		fmt.Printf("   tg messageForwardFromChat %+v \n\n", m.ForwardFromChat)
//	}
//	if m.ForwardFrom != nil {
//		fmt.Printf("   tg messageForwardFrom %+v \n\n", m.ForwardFrom)
//	}
//	if m.Document != nil {
//		fmt.Printf("   tg messageDocument %+v \n\n", m.Document)
//	}
//
//	fmt.Printf("\n\n")
//
//	ChatId := strconv.FormatInt(m.Chat.ID, 10) + fmt.Sprintf("/%d", ThreadID)
//
//	// RsClient
//	ok, config := t.CheckChannelConfigTG(ChatId)
//	if ok {
//		t.sendToFilterRs(m, config)
//	}
//	// client hs
//	//corpAlliance := t.getCorpHadesAlliance(ChatId)
//	//if corpAlliance.Corp != "" {
//	//	t.sendToFilterHades(m, corpAlliance, 0)
//	//	return
//	//}
//	//corpWs1 := t.getCorpHadesWs1(ChatId)
//	//if corpWs1.Corp != "" {
//	//	t.sendToFilterHades(m, corpWs1, 1)
//	//	return
//	//}
//
//	tg, bridgeConfig := t.BridgeCheckChannelConfigTg(ChatId)
//	if tg || strings.HasPrefix(m.Text, ".") {
//		go t.SendToBridgeChatFilter(m, bridgeConfig)
//	}
//}
//func (t *Telegram) sendToFilterHades(m *tgbotapi.Message, corp models.CorporationHadesClient, channelType int) {
//	if m.Text != "" {
//		if filterRsPl(m.Text) {
//			return
//		}
//		mes := models.MessageHades{
//			Text:        m.Text,
//			Sender:      t.nameOrNick(m.From.UserName, m.From.FirstName),
//			Avatar:      t.GetAvatar(m.From.ID),
//			ChannelType: channelType,
//			Corporation: corp.Corp,
//			Command:     "text",
//			Messager:    "tg",
//			MessageId:   strconv.Itoa(m.MessageID),
//		}
//		t.ChanToGame <- mes
//	}
//}
//func (t *Telegram) sendToFilterRs(m *tgbotapi.Message, config models.CorporationConfig) {
//	name := t.nameNick(m.From.UserName, m.From.FirstName, config.TgChannel)
//	in := models.InMessage{
//		Mtext:       m.Text,
//		Tip:         "tg",
//		Name:        name,
//		NameMention: "@" + name,
//		Tg: struct {
//			Mesid int
//			//Nameid int64
//		}{
//			Mesid: m.MessageID,
//			//Nameid: m.From.ID,
//		},
//		Config: config,
//		Option: models.Option{
//			InClient: true,
//		},
//	}
//
//	t.ChanRsMessage <- in
//}

//func (t *Telegram) SendToBridgeChatFilter(m *tgbotapi.Message, config models.BridgeConfig) {
//	username := t.nameOrNick(m.From.UserName, m.From.FirstName)
//	if len(m.Photo) != 0 {
//		for _, ph := range m.Photo {
//			url := t.GetPic(ph.FileID)
//			m.Text = m.Text + url + " "
//		}
//		m.Text = m.Text + m.Caption
//	}
//
//	mes := models.BridgeMessage{
//		Text:   m.Text,
//		Sender: username,
//		Tip:    "tg",
//		Tg: models.BridgeMessageTg{
//			ChatId:        strconv.FormatInt(m.Chat.ID, 10) + fmt.Sprintf("/%d", m.MessageThreadID),
//			MesId:         m.MessageID,
//			Avatar:        t.GetAvatar(m.From.ID),
//			TimestampUnix: m.Time().Unix(),
//			GroupName:     m.Chat.Title,
//		},
//	}
//	if config.HostRelay == "" {
//
//		//fmt.Printf(" logicmix.  %+v\n", mes)
//		t.ChanBridgeMessage <- mes
//		return
//	}
//	mes.Config = config
//
//	if m.ReplyToMessage != nil && m.ReplyToMessage.Text != "" {
//		mes.Tg.Reply.Text = m.ReplyToMessage.Text
//		mes.Tg.Reply.UserName = t.nameOrNick(m.ReplyToMessage.From.UserName, m.ReplyToMessage.From.FirstName)
//		mes.Tg.Reply.TimeMessage = m.ReplyToMessage.Time().Unix()
//		mes.Tg.Reply.Avatar = t.GetAvatar(m.ReplyToMessage.From.ID)
//	}
//	t.ChanBridgeMessage <- mes
//}
