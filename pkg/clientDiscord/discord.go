package clientDiscord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"kz_bot/config"
)

func NewDiscord(log *logrus.Logger, cfg config.ConfigBot) (*discordgo.Session, error) {
	DSBot, err := discordgo.New("Bot " + cfg.TokenD)
	if err != nil {
		log.Panic("Ошибка запуска дискорда", err)
		return nil, err
	}

	err = DSBot.Open()
	if err != nil {
		log.Panic("Ошибка открытия ДС", err)
	}
	return DSBot, nil
}
