package TelegramClient

import (
	"fmt"
	tgbotapi "github.com/matterbridge/telegram-bot-api/v6"
	"kz_bot/internal/models"
)

const nickname = "Для того что бы БОТ мог Вас индентифицировать, создайте уникальный НикНей в настройках. Вы можете использовать a-z, 0-9 и символы подчеркивания. Минимальная длина - 5 символов."

func (t *Telegram) callback(cb *tgbotapi.CallbackQuery) {
	callback := tgbotapi.NewCallback(cb.ID, cb.Data)
	if _, err := t.t.Request(callback); err != nil {
		t.log.Println("ошибка запроса калбек телеги ", err)
	}
	ok, config := t.CheckChannelConfigTG(cb.Message.Chat.ID)
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

		t.ChanRsMessage <- in
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
