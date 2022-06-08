package telegramClient

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	corpsConfig "kz_bot/internal/clients/corpConfig"
	"time"
)

type Telegram struct {
	t tgbotapi.BotAPI
}

func (t Telegram) SendEmded(lvlkz string, chatid int64, text string) int {
	var keyboardQueue = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"+", lvlkz+"+"),
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"-", lvlkz+"-"),
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"++", lvlkz+"++"),
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"+30", lvlkz+"+30"),
		),
	)
	msg := tgbotapi.NewMessage(chatid, text)
	msg.ReplyMarkup = keyboardQueue
	message, _ := t.t.Send(msg)

	return message.MessageID

}
func (t Telegram) SendEmbedTime(chatid int64, text string) int {
	var keyboardQueue = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("+", "+"),
			tgbotapi.NewInlineKeyboardButtonData("-", "-"),
		),
	)
	msg := tgbotapi.NewMessage(chatid, text)
	msg.ReplyMarkup = keyboardQueue
	message, _ := t.t.Send(msg)

	return message.MessageID

}

// отправка сообщения в телегу
func (t Telegram) SendChannel(chatid int64, text string) int {
	tMessage, _ := t.t.Send(tgbotapi.NewMessage(chatid, text))
	return tMessage.MessageID
}
func (t Telegram) SendChannelDelSecond(chatid int64, text string, second int) {
	tMessage, _ := t.t.Send(tgbotapi.NewMessage(chatid, text))
	if second < 60 {
		go func() {
			time.Sleep(time.Duration(second) * time.Second)
			t.t.Request(tgbotapi.DeleteMessageConfig(tgbotapi.NewDeleteMessage(chatid, tMessage.MessageID)))
		}()
	} else {
		fmt.Println("нужно удалять через бд")
	}

}
func (t Telegram) DelMessage(chatid int64, idSendMessage int) {
	t.t.Request(tgbotapi.DeleteMessageConfig(tgbotapi.NewDeleteMessage(chatid, idSendMessage)))
}
func (t Telegram) DelMessageSecond(chatid int64, idSendMessage int, second int) {
	go func() {
		time.Sleep(time.Duration(second) * time.Second)
		t.t.Request(tgbotapi.DeleteMessageConfig(tgbotapi.NewDeleteMessage(chatid, idSendMessage)))
	}()
}
func (t Telegram) EditMessageTextKey(chatid int64, editMesId int, textEdit string, lvlkz string) {
	var keyboardQueue = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"+", lvlkz+"+"),
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"-", lvlkz+"-"),
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"++", lvlkz+"++"),
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"+30", lvlkz+"+30"),
		),
	)
	tgbotapi.NewEditMessageText(chatid, editMesId, textEdit)
	t.t.Send(&tgbotapi.EditMessageTextConfig{
		BaseEdit: tgbotapi.BaseEdit{
			ChatID:          chatid,
			ChannelUsername: "",
			MessageID:       editMesId,
			InlineMessageID: "",
			ReplyMarkup:     &keyboardQueue,
		},
		Text:                  textEdit,
		ParseMode:             "",
		DisableWebPagePreview: false,
	})
}
func (t Telegram) EditText(chatid int64, editMesId int, textEdit string) {
	t.t.Send(tgbotapi.NewEditMessageText(chatid, editMesId, textEdit))
}
func (t Telegram) CheckAdminTg(chatid int64, name string) bool {
	admin := false
	admins, err := t.t.GetChatAdministrators(tgbotapi.ChatAdministratorsConfig{struct {
		ChatID             int64
		SuperGroupUsername string
	}{ChatID: chatid, SuperGroupUsername: ""}})
	if err != nil {
		fmt.Println("Ошибка проверки админа телеги ", err)
	}
	for _, ad := range admins {
		if name == ad.User.UserName && (ad.IsAdministrator() || ad.IsCreator()) {
			admin = true
			break
		}
	}
	return admin
}
func (t Telegram) RemoveDuplicateElementInt(mesididid []int) []int {
	result := make([]int, 0, len(mesididid))
	temp := map[int]struct{}{}
	for _, item := range mesididid {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
func (t Telegram) updatesComand(c *tgbotapi.Message) {
	conf := corpsConfig.CorpConfig{}
	ok, config := conf.CheckChannelConfigTG(c.Chat.ID)
	if ok {
		switch c.Command() {
		case "help":
			t.help(config.TgChannel, c.MessageID)
		case "helpqueue":
			t.helpQueue(config.TgChannel, c.MessageID)
		case "helpnotification":
			t.helpNotification(config.TgChannel, c.MessageID)
		case "helpevent":
			t.helpEvent(config.TgChannel, c.MessageID)
		case "helptop":
			t.helpTop(config.TgChannel, c.MessageID)
		case "helpicon":
			t.helpIcon(config.TgChannel, c.MessageID)
		}
	} else {
		switch c.Command() {
		case "help": //отправить справку о боте и способ активации
		default: //отправить спарвка в этом чате не доступна

		}
	}
}
func (t Telegram) ChatName(chatid int64) string {
	r, err := t.t.GetChat(tgbotapi.ChatInfoConfig{struct {
		ChatID             int64
		SuperGroupUsername string
	}{ChatID: chatid}})
	if err != nil {
		fmt.Println("ошибка получения имени чата ", err)
	}
	return r.Title
}
func (t Telegram) BotName() string {
	return t.t.Self.UserName
}
