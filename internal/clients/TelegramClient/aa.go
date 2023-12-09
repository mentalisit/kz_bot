package TelegramClient

import (
	"fmt"
	tgbotapi "github.com/musianisamuele/telegram-bot-api"
	"kz_bot/internal/models"
	"strconv"
	"strings"
	"time"
)

func (t *Telegram) DelMessage(chatid string, idSendMessage int) {
	a := strings.SplitN(chatid, "/", 2)
	chatId, err := strconv.ParseInt(a[0], 10, 64)
	if err != nil {
		t.log.Println("DelMessage " + err.Error())
	}
	_, _ = t.t.Request(tgbotapi.DeleteMessageConfig(tgbotapi.NewDeleteMessage(chatId, idSendMessage)))
	//if err != nil { t.log.Println("Ошибка удаления сообщения телеги ", err) }
}

func (t *Telegram) DelMessageSecond(chatid string, idSendMessage string, second int) {
	id, err := strconv.Atoi(idSendMessage)
	if err != nil {
		return
	}
	if second <= 60 {
		go func() {
			time.Sleep(time.Duration(second) * time.Second)
			t.DelMessage(chatid, id)
		}()
	} else {
		t.storage.TimeDeleteMessage.TimerInsert(models.Timer{
			Tgmesid:  strconv.Itoa(id),
			Tgchatid: chatid,
			Timed:    second,
		})
	}
}
func (t *Telegram) EditMessageTextKey(chatid string, editMesId int, textEdit string, lvlkz string) {
	a := strings.SplitN(chatid, "/", 2)
	chatId, err := strconv.ParseInt(a[0], 10, 64)
	if err != nil {
		t.log.Println(err)
	}

	var keyboardQueue = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"+", lvlkz+"+"),
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"-", lvlkz+"-"),
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"++", lvlkz+"++"),
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"+30", lvlkz+"+++"),
		),
	)
	mes := tgbotapi.EditMessageTextConfig{
		BaseEdit: tgbotapi.BaseEdit{
			ChatID:      chatId,
			MessageID:   editMesId,
			ReplyMarkup: &keyboardQueue,
		},
		Text: textEdit,
	}

	_, _ = t.t.Send(mes)
}
func (t *Telegram) EditText(chatid string, editMesId int, textEdit string) {
	a := strings.SplitN(chatid, "/", 2)
	chatId, err := strconv.ParseInt(a[0], 10, 64)
	if err != nil {
		t.log.Println(err)
	}
	_, err = t.t.Send(tgbotapi.NewEditMessageText(chatId, editMesId, textEdit))
	if err != nil {
		//t.log.Println("Ошибка редактирования EditText ", err)
	}
}
func (t *Telegram) CheckAdminTg(chatid string, name string) bool {
	admin := false
	a := strings.SplitN(chatid, "/", 2)
	chatId, err := strconv.ParseInt(a[0], 10, 64)
	if err != nil {
		t.log.Println(err)
	}
	admins, err := t.t.GetChatAdministrators(tgbotapi.ChatAdministratorsConfig{ChatConfig: struct {
		ChatID             int64
		SuperGroupUsername string
	}{ChatID: chatId, SuperGroupUsername: ""}})
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
	ChatId := strconv.FormatInt(c.Chat.ID, 10) + fmt.Sprintf("/%d", c.MessageThreadID)
	if c.Command() == "chatid" {
		t.SendChannelDelSecond(ChatId, ChatId, 20)
	}
	ok, config := t.CheckChannelConfigTG(ChatId)
	if ok {
		MessageID := strconv.Itoa(c.MessageID)
		switch c.Command() {
		case "help":
			t.help(config.TgChannel, MessageID)
		case "helpqueue":
			t.helpQueue(config.TgChannel, MessageID)
		case "helpnotification":
			t.helpNotification(config.TgChannel, MessageID)
		case "helpevent":
			t.helpEvent(config.TgChannel, MessageID)
		case "helptop":
			t.helpTop(config.TgChannel, MessageID)
		case "helpicon":
			t.helpIcon(config.TgChannel, MessageID)
		}
	} else {
		switch c.Command() {
		case "help":
			t.SendChannelDelSecond(ChatId, "Активируйте бота командой \n.add", 60)
		default:
			t.SendChannelDelSecond(ChatId, "Вам не доступна данная команда \n /help", 60)
		}
	}
}
func (t *Telegram) ChatName(chatid string) string {
	a := strings.SplitN(chatid, "/", 2)
	chatId, err := strconv.ParseInt(a[0], 10, 64)
	if err != nil {
		t.log.Println(err)
	}
	r, err := t.t.GetChat(tgbotapi.ChatInfoConfig{ChatConfig: struct {
		ChatID             int64
		SuperGroupUsername string
	}{ChatID: chatId}})
	if err != nil {
		t.log.Println("ошибка получения имени чата ", err)
	}
	return r.Title
}
func (t *Telegram) BotName() string {
	return t.t.Self.UserName
}
