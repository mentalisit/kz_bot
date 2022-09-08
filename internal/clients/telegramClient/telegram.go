package telegramClient

import (
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Tg interface {
	SendChannel(chatid int64, text string) int
	SendChannelDelSecond(chatid int64, text string, second int)
	SendEmded(lvlkz string, chatid int64, text string) int
	SendEmbedTime(chatid int64, text string) int
	EditText(chatid int64, editMesId int, textEdit string)
	EditMessageTextKey(chatid int64, editMesId int, textEdit string, lvlkz string)
	DelMessage(chatid int64, idSendMessage int)
	DelMessageSecond(chatid int64, idSendMessage int, second int)
	CheckAdminTg(chatid int64, name string) bool
	RemoveDuplicateElementInt(mesididid []int) []int
	ChatName(chatid int64) string
	BotName() string
	Help(Channel int64)
}

func (t *Telegram) SendEmded(lvlkz string, chatid int64, text string) int {
	var keyboardQueue = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"+", lvlkz+"+"),
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"-", lvlkz+"-"),
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"++", lvlkz+"++"),
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"+30", lvlkz+"+++"),
		),
	)
	msg := tgbotapi.NewMessage(chatid, text)
	msg.ReplyMarkup = keyboardQueue
	message, _ := t.t.Send(msg)

	return message.MessageID

}
func (t *Telegram) SendEmbedTime(chatid int64, text string) int {
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
func (t *Telegram) SendChannel(chatid int64, text string) int {
	tMessage, _ := t.t.Send(tgbotapi.NewMessage(chatid, text))
	return tMessage.MessageID
}
func (t *Telegram) SendChannelDelSecond(chatid int64, text string, second int) {
	tMessage, err := t.t.Send(tgbotapi.NewMessage(chatid, text))
	if err != nil {
		t.log.Println(err)
	}
	if second <= 60 {
		go func() {
			time.Sleep(time.Duration(second) * time.Second)
			t.DelMessage(chatid, tMessage.MessageID)
		}()
	} else {
		t.dbase.TimerInsert("", "", tMessage.MessageID, chatid, second)
	}
}
func (t *Telegram) DelMessage(chatid int64, idSendMessage int) {
	_, _ = t.t.Request(tgbotapi.DeleteMessageConfig(tgbotapi.NewDeleteMessage(chatid, idSendMessage)))
	//if err != nil { t.log.Println("Ошибка удаления сообщения телеги ", err) }
}
func (t *Telegram) DelMessageSecond(chatid int64, idSendMessage int, second int) {
	if second <= 60 {
		go func() {
			time.Sleep(time.Duration(second) * time.Second)
			t.DelMessage(chatid, idSendMessage)
		}()
	} else {
		t.dbase.TimerInsert("", "", idSendMessage, chatid, second)
	}
}
func (t *Telegram) EditMessageTextKey(chatid int64, editMesId int, textEdit string, lvlkz string) {
	var keyboardQueue = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"+", lvlkz+"+"),
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"-", lvlkz+"-"),
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"++", lvlkz+"++"),
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"+30", lvlkz+"+++"),
		),
	)
	_, _ = t.t.Send(&tgbotapi.EditMessageTextConfig{
		BaseEdit: tgbotapi.BaseEdit{
			ChatID:      chatid,
			MessageID:   editMesId,
			ReplyMarkup: &keyboardQueue,
		},
		Text: textEdit,
	})
}
func (t *Telegram) EditText(chatid int64, editMesId int, textEdit string) {
	_, err := t.t.Send(tgbotapi.NewEditMessageText(chatid, editMesId, textEdit))
	if err != nil {
		t.log.Println("Ошибка редактирования EditText ", err)
	}
}
func (t *Telegram) CheckAdminTg(chatid int64, name string) bool {
	admin := false
	admins, err := t.t.GetChatAdministrators(tgbotapi.ChatAdministratorsConfig{ChatConfig: struct {
		ChatID             int64
		SuperGroupUsername string
	}{ChatID: chatid, SuperGroupUsername: ""}})
	if err != nil {
		t.log.Println("Ошибка проверки админа телеги ", err)
	}
	for _, ad := range admins {
		if name == ad.User.UserName && (ad.IsAdministrator() || ad.IsCreator()) {
			admin = true
			break
		}
	}
	return admin
}
func (t *Telegram) RemoveDuplicateElementInt(mesididid []int) []int {
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
func (t *Telegram) updatesComand(c *tgbotapi.Message) {
	ok, config := t.CorpConfig.CheckChannelConfigTG(c.Chat.ID)
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
		case "help":
			t.SendChannelDelSecond(c.Chat.ID, "Активируйте бота командой \n.add", 60)
		default:
			t.SendChannelDelSecond(c.Chat.ID, "Вам не доступна данная команда \n /help", 60)
		}
	}
}
func (t *Telegram) ChatName(chatid int64) string {
	r, err := t.t.GetChat(tgbotapi.ChatInfoConfig{ChatConfig: struct {
		ChatID             int64
		SuperGroupUsername string
	}{ChatID: chatid}})
	if err != nil {
		t.log.Println("ошибка получения имени чата ", err)
	}
	return r.Title
}
func (t *Telegram) BotName() string {
	return t.t.Self.UserName
}
