package TelegramClient

import (
	"fmt"
	tgbotapi "github.com/samuelemusiani/telegram-bot-api"
	"strconv"
	"strings"
)

func (t *Telegram) accesChatTg(m *tgbotapi.Message) {
	res := strings.HasPrefix(m.Text, ".")
	ThreadID := 0
	if !m.IsTopicMessage && m.MessageThreadID != 0 {
		ThreadID = 0
	}
	ChatId := strconv.FormatInt(m.Chat.ID, 10) + fmt.Sprintf("/%d", ThreadID)
	if res {
		switch m.Text {
		case ".add":
			go t.DelMessageSecond(ChatId, strconv.Itoa(m.MessageID), 10)
			t.accessAddChannelTg(ChatId, "en")
		case ".добавить":
			go t.DelMessageSecond(ChatId, strconv.Itoa(m.MessageID), 10)
			t.accessAddChannelTg(ChatId, "ru")
		case ".добавитьт":
			go t.DelMessageSecond(ChatId, strconv.Itoa(m.MessageID), 10)
			t.accessAddChannelTg(ChatId, "dru")
		case ".додати":
			go t.DelMessageSecond(ChatId, strconv.Itoa(m.MessageID), 10)
			t.accessAddChannelTg(ChatId, "ua")
		case ".del":
			go t.DelMessageSecond(ChatId, strconv.Itoa(m.MessageID), 10)
			t.accessDelChannelTg(ChatId)
		case ".удалить":
			go t.DelMessageSecond(ChatId, strconv.Itoa(m.MessageID), 10)
			t.accessDelChannelTg(ChatId)
		case ".видалити":
			go t.DelMessageSecond(ChatId, strconv.Itoa(m.MessageID), 10)
			t.accessDelChannelTg(ChatId)
		}
	}
}
func (t *Telegram) accessAddChannelTg(chatid, lang string) { // внесение в дб и добавление в масив
	ok, _ := t.CheckChannelConfigTG(chatid)
	if ok {
		go t.SendChannelDelSecond(chatid, t.storage.Words.GetWords(lang, "accessAlready"), 20)
	} else {
		chatName := t.ChatName(chatid)
		t.AddTgCorpConfig(chatName, chatid, lang)
		t.log.Println("новая активация корпорации ", chatName)
		go t.SendChannelDelSecond(chatid, t.storage.Words.GetWords(lang, "accessTY"), 60)
	}
}
func (t *Telegram) accessDelChannelTg(chatid string) { //удаление с бд и масива для блокировки
	ok, config := t.CheckChannelConfigTG(chatid)
	if !ok {
		go t.SendChannelDelSecond(chatid, t.storage.Words.GetWords("ru", "accessYourChannel"), 60)
	} else {
		t.storage.ConfigRs.DeleteConfigRs(config)
		t.storage.ReloadDbArray()
		t.log.Println("отключение корпорации ", t.ChatName(chatid))
		go t.SendChannelDelSecond(chatid, t.storage.Words.GetWords(config.Country, "YouDisabledMyFeatures"), 60)
	}
}
