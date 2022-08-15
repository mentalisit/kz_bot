package telegramClient

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"kz_bot/internal/models"
)

func (t *Telegram) logicMixTelegram(m *tgbotapi.Message) {
	// тут я передаю чат айди и проверяю должен ли бот реагировать на этот чат
	ok, config := t.CorpConfig.CheckChannelConfigTG(m.Chat.ID)
	t.accesChatTg(m) //это была начальная функция при добавлени бота в группу
	if ok {
		in := models.InMessage{
			Mtext:       m.Text,
			Tip:         "tg",
			Name:        m.From.UserName,
			NameMention: "@" + m.From.UserName,
			Tg: struct {
				Mesid  int
				Nameid int64
			}{
				Mesid:  m.MessageID,
				Nameid: m.From.ID,
			},
			Config: config,
			Option: struct {
				Callback bool
				Edit     bool
				Update   bool
				Queue    bool
			}{
				Callback: false,
				Edit:     false,
				Update:   false,
			},
		}
		if t.debug {
			fmt.Printf("\n\nin logicMixTelegram %+v\n", in)
		}
		models.ChTg <- in
	}
}

func (t *Telegram) callback(cb *tgbotapi.CallbackQuery) {
	callback := tgbotapi.NewCallback(cb.ID, cb.Data)
	if _, err := t.t.Request(callback); err != nil {
		t.log.Println("ошибка запроса калбек телеги ", err)
	}
	ok, config := t.CorpConfig.CheckChannelConfigTG(cb.Message.Chat.ID)
	if ok {
		in := models.InMessage{
			Mtext:       cb.Data,
			Tip:         "tg",
			Name:        cb.From.UserName,
			NameMention: "@" + cb.From.UserName,
			Tg: struct {
				Mesid  int
				Nameid int64
			}{
				Mesid:  cb.Message.MessageID,
				Nameid: cb.From.ID,
			},
			Config: config,
			Option: struct {
				Callback bool
				Edit     bool
				Update   bool
				Queue    bool
			}{
				Callback: true,
				Edit:     true,
				Update:   false,
			},
		}
		if t.debug {
			fmt.Printf("\n\nin logicMixTelegramCallback %+v\n", in)
		}
		models.ChTg <- in
	}
}

func (t *Telegram) myChatMember(member *tgbotapi.ChatMemberUpdated) {
	if member.NewChatMember.Status == "member" {
		t.SendChannelDelSecond(member.Chat.ID, fmt.Sprintf("@%s мне нужны права админа для коректной работы", member.From.UserName), 60)
	} else if member.NewChatMember.Status == "administrator" {
		t.SendChannelDelSecond(member.Chat.ID, fmt.Sprintf("@%s спасибо ... я готов к работе \nподтверди активацию .add", member.From.UserName), 60)
	}
}

func (t *Telegram) chatMember(chMember *tgbotapi.ChatMemberUpdated) {
	if chMember.NewChatMember.IsMember {
		t.SendChannelDelSecond(chMember.Chat.ID,
			fmt.Sprintf("%s Добро пожаловать в наш чат ", chMember.NewChatMember.User.FirstName),
			60)
	}

}
