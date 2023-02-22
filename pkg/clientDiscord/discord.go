package clientDiscord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"kz_bot/internal/config"
)

func NewDiscord(log *logrus.Logger, cfg *config.ConfigBot) (*discordgo.Session, error) {
	DSBot, err := discordgo.New("Bot " + cfg.Token.TokenDiscord)
	if err != nil {
		log.Panic("Ошибка запуска дискорда", err)
		return nil, err
	}

	err = DSBot.Open()
	if err != nil {
		log.Panic("Ошибка открытия ДС ", err)
	}
	fmt.Println("Бот Дискорд загружен ")
	return DSBot, nil
}
