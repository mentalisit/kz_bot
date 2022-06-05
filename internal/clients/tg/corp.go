package Tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	corpsConfig "kz_bot/internal/clients/corpConfig"
	db_Mysql "kz_bot/internal/dbase/dbaseMysql"
	"strings"
)

func (t Telegram) accesChatTg(m *tgbotapi.Message) {
	res := strings.HasPrefix(m.Text, ".")
	if res == true && m.Text == ".add" {
		go t.DelMessageSecond(m.Chat.ID, m.MessageID, 10)
		t.accessAddChannelTg(m.Chat.ID)
	} else if res == true && m.Text == ".del" {
		go t.DelMessageSecond(m.Chat.ID, m.MessageID, 10)
		t.accessDelChannelTg(m.Chat.ID)
	}
}
func (t Telegram) accessAddChannelTg(chatid int64) { // внесение в дб и добавление в масив
	c := corpsConfig.CorpConfig{}
	ok, _ := c.CheckChannelConfigTG(chatid)
	if ok {
		go t.SendChannelDelSecond(chatid, "Я уже могу работать на вашем канале\n"+
			"повторная активация не требуется.\nнапиши Справка", 20)
	} else {
		chatName := t.ChatName(chatid)
		db := db_Mysql.Db{}
		db.AddTgCorpConfig(chatName, chatid)
		go t.SendChannelDelSecond(chatid, "Спасибо за активацию.\nпиши Справка", 60)
	}
}
func (t Telegram) accessDelChannelTg(chatid int64) { //удаление с бд и масива для блокировки
	c := corpsConfig.CorpConfig{}
	ok, _ := c.CheckChannelConfigTG(chatid)
	if !ok {
		go t.SendChannelDelSecond(chatid, "ваш канал и так не подключен к логике бота ", 60)
	} else {
		db := db_Mysql.Db{}
		db.DeleteTgchannel(chatid)
		c.ReloadConfig()
		db.ReadBotCorpConfig()
		go t.SendChannelDelSecond(chatid, "вы отключили мои возможности", 60)
	}
}
