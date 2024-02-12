package TelegramClient

import (
	"fmt"
	tgbotapi "github.com/matterbridge/telegram-bot-api/v6"
	"kz_bot/internal/models"
	"strconv"
	"strings"
	"time"
)

func (t *Telegram) logicMix(m *tgbotapi.Message, edit bool) {
	go t.imHere(m.Chat.ID, m.Chat)
	t.accesChatTg(m) //это была начальная функция при добавлени бота в группу
	ThreadID := m.MessageThreadID
	if !m.IsTopicMessage && ThreadID != 0 {
		ThreadID = 0
	}
	ChatId := strconv.FormatInt(m.Chat.ID, 10) + fmt.Sprintf("/%d", ThreadID)
	if t.prefixCompendium(m, ChatId) {
		return
	}
	url := t.handleDownload(m)
	if m.Text == "" && m.Caption != "" {
		m.Text = m.Caption
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
			}{
				Mesid: m.MessageID,
			},
			Config: config,
			Option: models.Option{
				InClient: true,
			},
		}
		if in.Mtext == "" && config.Forward {
			t.DelMessageSecond(ChatId, strconv.Itoa(m.MessageID), 180)
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
				Config:  &models.BridgeConfig{},
			}
			t.ChanBridgeMessage <- mes
			time.Sleep(3 * time.Second)
			t.storage.ReloadDbArray()
			t.bridgeConfig = t.storage.BridgeConfigs
		}()
	}
	if tg {
		go func() {
			if len(m.Text) < 3500 { //игнорируем сообщения большой длины
				t.handlePoll(m)
				if m.Document != nil {
					url, _ = t.t.GetFileDirectURL(m.Document.FileID)
					if m.Text == "" {
						m.Text = m.Document.FileName
					}
				}
				if len(m.Photo) > 0 {
					url, _ = t.t.GetFileDirectURL(m.Photo[len(m.Photo)-1].FileID)
				}
				mes := models.BridgeMessage{
					Text:          m.Text,
					Sender:        username,
					Tip:           "tg",
					FileUrl:       url,
					Avatar:        t.GetAvatar(m.From.ID, m.From.String()),
					ChatId:        ChatId,
					MesId:         strconv.Itoa(m.MessageID),
					TimestampUnix: m.Time().Unix(),
					GuildId:       chatName,
					Config:        &bridgeConfig,
				}

				if m.ReplyToMessage != nil && m.ReplyToMessage.ForumTopicCreated == nil {
					mes.Reply = &models.BridgeMessageReply{
						Text:        m.ReplyToMessage.Text,
						UserName:    t.nameOrNick(m.ReplyToMessage.From.UserName, m.ReplyToMessage.From.FirstName),
						TimeMessage: m.ReplyToMessage.Time().Unix(),
						Avatar:      t.GetAvatar(m.ReplyToMessage.From.ID, m.ReplyToMessage.From.String()),
						FileUrl:     t.handleDownload(m.ReplyToMessage),
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
			}
		}()
	}
}
