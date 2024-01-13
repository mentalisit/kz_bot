package TelegramClient

import (
	"fmt"
	tgbotapi "github.com/samuelemusiani/telegram-bot-api"
	"kz_bot/internal/models"
	"strconv"
)

const nickname = "Для того что бы БОТ мог Вас индентифицировать, создайте уникальный НикНей в настройках. Вы можете использовать a-z, 0-9 и символы подчеркивания. Минимальная длина - 5 символов."

func (t *Telegram) callback(cb *tgbotapi.CallbackQuery) {
	callback := tgbotapi.NewCallback(cb.ID, cb.Data)
	if _, err := t.t.Request(callback); err != nil {
		t.log.Error(err.Error())
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
				Mesid int
			}{
				Mesid: cb.Message.MessageID,
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
		t.SendChannelDelSecond(ChatId, fmt.Sprintf("@%s спасибо ... я готов к работе \nАктивируй нужный режим бота,\n если сложности пиши мне @Mentalisit", member.From.UserName), 60)
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

func (t *Telegram) handleDownload(message *tgbotapi.Message) (url string) {
	size := int64(0)
	switch {
	case message.Sticker != nil:
		url, _ = t.t.GetFileDirectURL(message.Sticker.FileID)
		size = int64(message.Sticker.FileSize)
	case message.Voice != nil:
		url, _ = t.t.GetFileDirectURL(message.Voice.FileID)
		size = message.Voice.FileSize
	case message.Video != nil:
		url, _ = t.t.GetFileDirectURL(message.Video.FileID)
		size = message.Video.FileSize
	case message.Audio != nil:
		url, _ = t.t.GetFileDirectURL(message.Audio.FileID)
		size = message.Audio.FileSize
	case message.Document != nil:
		url, _ = t.t.GetFileDirectURL(message.Document.FileID)
		size = message.Document.FileSize
	case message.Photo != nil:
		photos := message.Photo
		size = int64(photos[len(photos)-1].FileSize)
		url, _ = t.t.GetFileDirectURL(photos[len(photos)-1].FileID)

	}
	if size > 25000000 {
		fmt.Println("big size")
		message.Text += " файл слишком большой для пересылки"
		return ""
	}

	return url
}

func (t *Telegram) handlePoll(message *tgbotapi.Message) {
	if message.Poll != nil {
		text := "Запущен  ОПРОС  \n"
		text += message.Poll.Question
		text += "\nВарианты ответа:\n"
		for _, o := range message.Poll.Options {
			text += fmt.Sprintf(" %s\n", o.Text)
		}
		message.Text = text
	}
}
