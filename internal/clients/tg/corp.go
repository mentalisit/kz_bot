package Tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

func accesChatTg(m *tgbotapi.Message) {
	res := strings.HasPrefix(m.Text, ".")
	if res == true && m.Text == ".add" {
		go tgDelMessage10s(m.Chat.ID, m.MessageID)
		accessAddChannelTg(m.Chat.ID)
	} else if res == true && m.Text == ".del" {
		go tgDelMessage10s(m.Chat.ID, m.MessageID)
		accessDelChannelTg(m.Chat.ID)
	}
}
func accessAddChannelTg(chatid int64) { // внесение в дб и добавление в масив
	ok, _ := checkChannelConfigTG(chatid)
	if ok {
		go tgSendChannelDel1m(chatid, "Я уже могу работать на вашем канале\n"+
			"повторная активация не требуется.\nнапиши Справка")
	} else {
		chatName := tgChatName(chatid)
		insertConfig := `INSERT INTO config (corpname,dschannel,tgchannel,wachannel,mesiddshelp,mesidtghelp,delmescomplite) VALUES (?,?,?,?,?,?,?)`
		statement, err := db.Prepare(insertConfig)
		if err != nil {
			logrus.Println(err)
		}
		_, err = statement.Exec(chatName, "", chatid, "", "", 0, 0)
		if err != nil {
			logrus.Println(err.Error())
		}
		addCorp(chatName, "", chatid, "", 1, "", 0, "")
		go tgSendChannelDel1m(chatid, "Спасибо за активацию.\nпиши Справка")
	}
}
func accessDelChannelTg(chatid int64) { //удаление с бд и масива для блокировки
	ok, _ := checkChannelConfigTG(chatid)
	if !ok {
		go tgSendChannelDel1m(chatid, "ваш канал и так не подключен к логике бота ")
	} else {
		_, err := db.Exec("delete from config where tgchannel = ? ", chatid)
		if err != nil {
			logrus.Println(err)
		}
		*P = *New()
		readBotConfig()
		go tgSendChannelDel1m(chatid, "вы отключили мои возможности")
	}
}
