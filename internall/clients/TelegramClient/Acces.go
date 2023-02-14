package TelegramClient

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
	"time"
)

func (t *Telegram) accesChatTg(m *tgbotapi.Message) {
	res := strings.HasPrefix(m.Text, ".")
	if res == true && m.Text == ".add" {
		go t.DelMessageSecond(m.Chat.ID, m.MessageID, 10)
		t.accessAddChannelTg(m.Chat.ID)
	} else if res == true && m.Text == ".del" {
		go t.DelMessageSecond(m.Chat.ID, m.MessageID, 10)
		t.accessDelChannelTg(m.Chat.ID)
	}
}
func (t *Telegram) accessAddChannelTg(chatid int64) { // внесение в дб и добавление в масив
	ok, _ := t.storage.Cache.CheckChannelConfigTG(chatid)
	if ok {
		go t.SendChannelDelSecond(chatid, "Я уже могу работать на вашем канале\n"+
			"повторная активация не требуется.\nнапиши Справка", 20)
	} else {
		chatName := t.ChatName(chatid)
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		t.storage.CorpsConfig.AddTgCorpConfig(ctx, chatName, chatid)
		t.log.Println("новая активация корпорации ", chatName)
		go t.SendChannelDelSecond(chatid, "Спасибо за активацию.\nпиши Справка", 60)
	}
}
func (t *Telegram) accessDelChannelTg(chatid int64) { //удаление с бд и масива для блокировки
	ok, _ := t.storage.Cache.CheckChannelConfigTG(chatid)
	if !ok {
		go t.SendChannelDelSecond(chatid, "ваш канал и так не подключен к логике бота ", 60)
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		t.storage.CorpsConfig.DeleteTg(ctx, chatid)
		t.log.Println("отключение корпорации ", t.ChatName(chatid))
		t.storage.Cache.ReloadConfig()
		t.storage.CorpsConfig.ReadCorps()
		go t.SendChannelDelSecond(chatid, "вы отключили мои возможности", 60)
	}
}
