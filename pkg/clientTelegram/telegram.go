package clientTelegram

import (
	"fmt"
	tgbotapi "github.com/matterbridge/telegram-bot-api/v6"
	"github.com/mentalisit/logger"
	"kz_bot/internal/config"
)

func NewTelegram(log *logger.Logger, cfg *config.ConfigBot) (*tgbotapi.BotAPI, error) {
	tgBot, err := tgbotapi.NewBotAPI(cfg.Token.TokenTelegram)
	if err != nil {
		log.Panic("ошибка подключения к телеграм " + err.Error())
	}
	tgBot.Debug = false
	fmt.Printf("Бот TELEGRAM загружен  %s\n", tgBot.Self.UserName)

	return tgBot, err
}
