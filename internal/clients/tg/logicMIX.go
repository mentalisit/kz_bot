package Tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"kz_bot/internal/models"
)

func logicMixTelegram(m *tgbotapi.Message) {
	// тут я передаю чат айди и проверяю должен ли бот реагировать на этот чат
	//ok, config := checkChannelConfigTG(m.Chat.ID)
	//accesChatTg(m) //это была начальная функция при добавлени бота в группу
	var ok = true
	if ok {
		in := models.InMessage{
			Mtext:       m.Text,
			Tip:         "tg",
			Name:        m.From.UserName,
			NameMention: "@" + m.From.UserName,
			Ds:          models.Ds{},
			Tg: models.Tg{
				Mesid:  m.MessageID,
				Nameid: m.From.ID,
			},
			Option: models.Option{
				Callback: false,
				Edit:     false,
				Update:   false,
			},
		}
		//logicRs(in)
		//тут нужно передавать в логику бота
		//fmt.Println(in)
		models.ChTg <- in
	}
}
