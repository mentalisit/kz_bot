package TelegramClient

import (
	"fmt"
	tgbotapi "github.com/samuelemusiani/telegram-bot-api"
	"regexp"
	"strconv"
	"strings"
)

func (t *Telegram) accesChatTg(m *tgbotapi.Message) {
	res := strings.HasPrefix(m.Text, ".")
	ThreadID := m.MessageThreadID
	if !m.IsTopicMessage && m.MessageThreadID != 0 {
		ThreadID = 0
	}
	ChatId := strconv.FormatInt(m.Chat.ID, 10) + fmt.Sprintf("/%d", ThreadID)
	if res {
		switch m.Text {
		case ".add":
			go t.DelMessageSecond(ChatId, strconv.Itoa(m.MessageID), 10)
			t.accessAddChannelTg(ChatId, "en", m)
		case ".добавить":
			go t.DelMessageSecond(ChatId, strconv.Itoa(m.MessageID), 10)
			t.accessAddChannelTg(ChatId, "ru", m)
		case ".добавитьт":
			go t.DelMessageSecond(ChatId, strconv.Itoa(m.MessageID), 10)
			t.accessAddChannelTg(ChatId, "dru", m)
		case ".додати":
			go t.DelMessageSecond(ChatId, strconv.Itoa(m.MessageID), 10)
			t.accessAddChannelTg(ChatId, "ua", m)
		case ".del":
			go t.DelMessageSecond(ChatId, strconv.Itoa(m.MessageID), 10)
			t.accessDelChannelTg(ChatId, m)
		case ".удалить":
			go t.DelMessageSecond(ChatId, strconv.Itoa(m.MessageID), 10)
			t.accessDelChannelTg(ChatId, m)
		case ".видалити":
			go t.DelMessageSecond(ChatId, strconv.Itoa(m.MessageID), 10)
			t.accessDelChannelTg(ChatId, m)
		case ".паника":
			t.log.Panic("перезагрузка по требованию")
		default:
			if t.setLang(m, ChatId) {
				return
			}
		}
	}
}
func (t *Telegram) accessAddChannelTg(chatid, lang string, m *tgbotapi.Message) { // внесение в дб и добавление в масив
	ok, _ := t.CheckChannelConfigTG(chatid)
	if ok {
		go t.SendChannelDelSecond(chatid, t.storage.Words.GetWords(lang, "accessAlready"), 20)
	} else {
		chatName := t.ChatName(chatid)
		if m.IsTopicMessage && m.ReplyToMessage != nil && m.ReplyToMessage.ForumTopicCreated != nil {
			chatName = fmt.Sprintf(" %s/%s", chatName, m.ReplyToMessage.ForumTopicCreated.Name)
		}
		t.AddTgCorpConfig(chatName, chatid, lang)
		t.log.Info("новая активация корпорации " + chatName)
		go t.SendChannelDelSecond(chatid, t.storage.Words.GetWords(lang, "accessTY"), 60)
	}
}
func (t *Telegram) accessDelChannelTg(chatid string, m *tgbotapi.Message) { //удаление с бд и масива для блокировки
	ok, config := t.CheckChannelConfigTG(chatid)
	if !ok {
		go t.SendChannelDelSecond(chatid, t.storage.Words.GetWords("ru", "accessYourChannel"), 60)
	} else {
		t.storage.ConfigRs.DeleteConfigRs(config)
		t.storage.ReloadDbArray()
		t.corpConfigRS = t.storage.CorpConfigRS
		chatName := t.ChatName(chatid)
		if m.IsTopicMessage && m.ReplyToMessage != nil && m.ReplyToMessage.ForumTopicCreated != nil {
			chatName = fmt.Sprintf(" %s/%s", chatName, m.ReplyToMessage.ForumTopicCreated.Name)
		}
		t.log.Info("отключение корпорации " + chatName)
		go t.SendChannelDelSecond(chatid, t.storage.Words.GetWords(config.Country, "YouDisabledMyFeatures"), 60)
	}
}
func (t *Telegram) setLang(m *tgbotapi.Message, chatid string) bool {
	re := regexp.MustCompile(`^\.set lang (ru|en|ua)$`)
	matches := re.FindStringSubmatch(m.Text)
	if len(matches) > 0 {
		langUpdate := matches[1]
		ok, config := t.CheckChannelConfigTG(chatid)
		if ok {
			go t.DelMessageSecond(chatid, strconv.Itoa(m.MessageID), 10)
			config.Country = langUpdate
			t.corpConfigRS[config.CorpName] = config
			t.storage.ConfigRs.AutoHelpUpdateMesid(config)
			go t.SendChannelDelSecond(chatid, t.storage.Words.GetWords(config.Country, "vashLanguage"), 20)
			t.log.Info(fmt.Sprintf("замена языка в %s на %s", config.CorpName, config.Country))
		}

		return true
	}
	return false
}
