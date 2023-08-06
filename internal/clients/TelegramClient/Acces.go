package TelegramClient

import (
	tgbotapi "github.com/matterbridge/telegram-bot-api/v6"
	"strconv"
	"strings"
)

func (t *Telegram) accesChatTg(m *tgbotapi.Message) {
	res := strings.HasPrefix(m.Text, ".")
	ChatId := strconv.FormatInt(m.Chat.ID, 10) + "/" + string(rune(m.MessageThreadID))
	if res == true && m.Text == ".add" {
		go t.DelMessageSecond(ChatId, m.MessageID, 10)
		t.accessAddChannelTg(ChatId)
	} else if res == true && m.Text == ".del" {
		go t.DelMessageSecond(ChatId, m.MessageID, 10)
		t.accessDelChannelTg(ChatId)
	}
}
func (t *Telegram) accessAddChannelTg(chatid string) { // внесение в дб и добавление в масив
	ok, _ := t.CheckChannelConfigTG(chatid)
	if ok {
		go t.SendChannelDelSecond(chatid, "Я уже могу работать на вашем канале\n"+
			"повторная активация не требуется.\nнапиши Справка", 20)
	} else {
		chatName := t.ChatName(chatid)
		t.AddTgCorpConfig(chatName, chatid)
		t.log.Println("новая активация корпорации ", chatName)
		go t.SendChannelDelSecond(chatid, "Спасибо за активацию.\nпиши Справка", 60)
	}
}
func (t *Telegram) accessDelChannelTg(chatid string) { //удаление с бд и масива для блокировки
	ok, config := t.CheckChannelConfigTG(chatid)
	if !ok {
		go t.SendChannelDelSecond(chatid, "ваш канал и так не подключен к логике бота ", 60)
	} else {
		t.storage.ConfigRs.DeleteConfigRs(config)
		t.storage.ReloadDbArray()
		t.log.Println("отключение корпорации ", t.ChatName(chatid))
		go t.SendChannelDelSecond(chatid, "вы отключили мои возможности", 60)
	}
}
