package TelegramClient

import (
	"fmt"
	tgbotapi "github.com/musianisamuele/telegram-bot-api"
	"kz_bot/internal/models"
	"strconv"
)

const nickname = "Для того что бы БОТ мог Вас индентифицировать, создайте уникальный НикНей в настройках. Вы можете использовать a-z, 0-9 и символы подчеркивания. Минимальная длина - 5 символов."

func (t *Telegram) callback(cb *tgbotapi.CallbackQuery) {
	callback := tgbotapi.NewCallback(cb.ID, cb.Data)
	if _, err := t.t.Request(callback); err != nil {
		t.log.Println("ошибка запроса калбек телеги ", err)
	}
	ChatId := strconv.FormatInt(cb.Message.Chat.ID, 10) + fmt.Sprintf("/%d", cb.Message.MessageThreadID)
	ok, config := t.CheckChannelConfigTG(ChatId)
	if ok {
		name := t.nameNick(cb.From.UserName, cb.From.FirstName, ChatId)
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
	ChatId := strconv.FormatInt(member.Chat.ID, 10) + "/0"
	if member.NewChatMember.Status == "member" {
		t.SendChannelDelSecond(ChatId, fmt.Sprintf("@%s мне нужны права админа для коректной работы", member.From.UserName), 60)
	} else if member.NewChatMember.Status == "administrator" {
		t.SendChannelDelSecond(ChatId, fmt.Sprintf("@%s спасибо ... я готов к работе \nподтверди активацию .add", member.From.UserName), 60)
	}
}

func (t *Telegram) chatMember(chMember *tgbotapi.ChatMemberUpdated) {
	if chMember.NewChatMember.IsMember {
		ChatId := strconv.FormatInt(chMember.Chat.ID, 10) + "/0"
		t.SendChannelDelSecond(ChatId,
			fmt.Sprintf("%s Добро пожаловать в наш чат ", chMember.NewChatMember.User.FirstName),
			60)
	}

}
func (t *Telegram) nameNick(UserName, FirstName string, chatid string) (name string) {
	if UserName != "" {
		name = UserName
	} else {
		name = FirstName
		go t.SendChannelDelSecond(chatid, nickname, 60)
	}
	return name
}
