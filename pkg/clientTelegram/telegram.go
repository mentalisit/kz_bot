package clientTelegram

import (
	"fmt"
	tgbotapi "github.com/musianisamuele/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"kz_bot/internal/config"
)

func NewTelegram(log *logrus.Logger, cfg *config.ConfigBot) (*tgbotapi.BotAPI, error) {
	tgBot, err := tgbotapi.NewBotAPI(cfg.Token.TokenTelegram)
	if err != nil {
		log.Panic("ошибка подключения к телеграм ", err)
	}
	tgBot.Debug = false
	fmt.Printf("Бот TELEGRAM загружен  %s\n", tgBot.Self.UserName)

	return tgBot, err
}
