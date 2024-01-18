package TelegramClient

import (
	"fmt"
	tgbotapi "github.com/samuelemusiani/telegram-bot-api"
	"strconv"

	"kz_bot/internal/models"
	"strings"
)

func (t *Telegram) logicMix(m *tgbotapi.Message, edit bool) {
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

	username := t.nameOrNick(m.From.UserName, m.From.FirstName)
	chatName := m.Chat.Title
	if m.IsTopicMessage && m.ReplyToMessage != nil && m.ReplyToMessage.ForumTopicCreated != nil {
		chatName = fmt.Sprintf("%s/%s", chatName, m.ReplyToMessage.ForumTopicCreated.Name)
	}

	if strings.HasPrefix(m.Text, ".") {
		go func() {
			mes := models.BridgeMessage{
				Text:    m.Text,
				Sender:  username,
				Tip:     "tg",
				ChatId:  ChatId,
				MesId:   strconv.Itoa(m.MessageID),
				GuildId: chatName,
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
