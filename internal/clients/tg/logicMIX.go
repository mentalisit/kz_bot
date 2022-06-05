package Tg

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	corpsConfig "kz_bot/internal/clients/corpConfig"
	"kz_bot/internal/models"
)

func (t Telegram) logicMixTelegram(m *tgbotapi.Message) {
	// тут я передаю чат айди и проверяю должен ли бот реагировать на этот чат
	c := corpsConfig.CorpConfig{}
	ok, config := c.CheckChannelConfigTG(m.Chat.ID)
	fmt.Println(ok)
	t.accesChatTg(m) //это была начальная функция при добавлени бота в группу
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
			Config: config,
			Option: models.Option{
				Callback: false,
				Edit:     false,
				Update:   false,
			},
		}
		//logicRs(in)
		//тут нужно передавать в логику бота
		fmt.Println(in)
		models.ChTg <- in
	}
}
