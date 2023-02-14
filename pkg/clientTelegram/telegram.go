package clientTelegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"kz_bot/config"
)

func NewTelegram(log *logrus.Logger, cfg config.ConfigBot) (*tgbotapi.BotAPI, error) {
	tgBot, err := tgbotapi.NewBotAPI(cfg.TokenT)
	if err != nil {
		log.Panic("ошибка подключения к телеграм ", err)
	}
	tgBot.Debug = false
	fmt.Printf("Бот TELEGRAM загружен  %s\n", tgBot.Self.UserName)

	return tgBot, err
}
