package TelegramClient

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"kz_bot/internal/models"
)

const nickname = "Для того что бы БОТ мог Вас индентифицировать, создайте уникальный НикНей в настройках. Вы можете использовать a-z, 0-9 и символы подчеркивания. Минимальная длина - 5 символов."

func (t *Telegram) logicMix(m *tgbotapi.Message) {
	t.ifMessageForHades(m)

	// тут я передаю чат айди и проверяю должен ли бот реагировать на этот чат
	ok, config := t.storage.Cache.CheckChannelConfigTG(m.Chat.ID)
	t.accesChatTg(m) //это была начальная функция при добавлени бота в группу
	if ok {
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

		t.inbox <- in
	}
}

func (t *Telegram) callback(cb *tgbotapi.CallbackQuery) {
	callback := tgbotapi.NewCallback(cb.ID, cb.Data)
	if _, err := t.t.Request(callback); err != nil {
		t.log.Println("ошибка запроса калбек телеги ", err)
	}
	ok, config := t.storage.Cache.CheckChannelConfigTG(cb.Message.Chat.ID)
	if ok {
		name := t.nameNick(cb.From.UserName, cb.From.FirstName, cb.Message.Chat.ID)
		in := models.InMessage{
			Mtext:       cb.Data,
			Tip:         "tg",
			Name:        name,
			NameMention: "@" + name,
			Tg: struct {
				Mesid  int
				Nameid int64
			}{
				Mesid:  cb.Message.MessageID,
				Nameid: cb.From.ID,
			},
			Config: config,
			Option: models.Option{
				Reaction: true},
		}

		t.inbox <- in
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
func (t *Telegram) nameNick(UserName, FirstName string, chatid int64) (name string) {
	if UserName != "" {
		name = UserName
	} else {
		name = FirstName
		go t.SendChannelDelSecond(chatid, nickname, 60)
	}
	return name
}