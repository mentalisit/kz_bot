package TelegramClient

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"kz_bot/internal/models"
	"kz_bot/internal/storage/CorpsConfig/hades"
	"regexp"
	"strings"
)

func (t *Telegram) ifComands(m *tgbotapi.Message) (command bool) {
	str, ok := strings.CutPrefix(m.Text, ". ")
	if ok {
		str = strings.ToLower(str)
		arr := strings.Split(str, " ")
		arrlen := len(arr)
		if arrlen == 1 {

		} else if arrlen == 2 {
			if t.lastWs(arr, m) {
				return true
			}
			if t.replayId(arr, m) {
				return true
			}
			if t.historyWs(arr, m) {
				return true
			}

		}
	}
	return false
}
func (t *Telegram) lastWs(arg []string, m *tgbotapi.Message) bool {
	if arg[0] == "повтор" && arg[1] == "бз" {
		_, corporation := hades.HadesStorage.AllianceChatTg(m.Chat.ID)
		mes := models.Message{
			Text:        "",
			Sender:      t.nameOrNick(m.From.UserName, m.From.FirstName),
			Avatar:      t.GetAvatar(m.From.ID),
			ChannelType: 0,
			Corporation: corporation.Corp,
			Command:     "повтор бз",
			Messager:    "tg",
		}
		fmt.Printf("lastWs %+v\n", mes)
		t.toGame <- mes
		go t.SendChannelDelSecond(m.Chat.ID, "отправка повтора последней бз", 10)
		go t.DelMessageSecond(m.Chat.ID, m.MessageID, 180)
		return true
	}
	return false
}
func (t *Telegram) replayId(arg []string, m *tgbotapi.Message) bool {
	if arg[0] == "повтор" {
		match, _ := regexp.MatchString("^[0-9]+$", arg[1])
		if match {
			_, corporation := hades.HadesStorage.AllianceChatTg(m.Chat.ID)
			mes := models.Message{
				Text:        arg[1],
				Sender:      t.nameOrNick(m.From.UserName, m.From.FirstName),
				Avatar:      t.GetAvatar(m.From.ID),
				ChannelType: 0,
				Corporation: corporation.Corp,
				Command:     "повтор",
				Messager:    "tg",
			}
			fmt.Printf("replayId %+v\n", mes)
			t.toGame <- mes
			go t.SendChannelDelSecond(m.Chat.ID, "отправка повтора "+arg[1], 10)
			go t.DelMessageSecond(m.Chat.ID, m.MessageID, 180)
			return true
		}
	}
	return false
}
func (t *Telegram) historyWs(arg []string, m *tgbotapi.Message) bool {
	if arg[0] == "история" && arg[1] == "бз" {
		_, corporation := hades.HadesStorage.AllianceChatTg(m.Chat.ID)
		mes := models.Message{
			Text:        "",
			Sender:      t.nameOrNick(m.From.UserName, m.From.FirstName),
			Avatar:      t.GetAvatar(m.From.ID),
			ChannelType: 0,
			Corporation: corporation.Corp,
			Command:     "история бз",
			Messager:    "tg",
		}
		fmt.Printf("historyWs %+v\n", mes)
		t.toGame <- mes
		go t.SendChannelDelSecond(m.Chat.ID, "готовлю список  бз", 10)
		go t.DelMessageSecond(m.Chat.ID, m.MessageID, 180)
		return true
	}
	return false
}
func (t *Telegram) letInId(arg []string, m *tgbotapi.Message) bool {
	if arg[0] == "впустить" {
		match, _ := regexp.MatchString("^[0-9]+$", arg[1])
		if match {
			_, corporation := hades.HadesStorage.AllianceChatTg(m.Chat.ID)
			mes := models.Message{
				Text:        arg[1],
				Sender:      t.nameOrNick(m.From.UserName, m.From.FirstName),
				Avatar:      t.GetAvatar(m.From.ID),
				ChannelType: 0,
				Corporation: corporation.Corp,
				Command:     "впустить",
				Messager:    "tg",
			}
			fmt.Printf("letInId %+v\n", mes)
			t.toGame <- mes
			go t.SendChannelDelSecond(m.Chat.ID, "впустить отправленно "+arg[1], 10)
			go t.DelMessageSecond(m.Chat.ID, m.MessageID, 180)
			return true
		}
	}
	return false
}
