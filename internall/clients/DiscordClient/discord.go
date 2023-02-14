package DiscordClient

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"kz_bot/config"
	"kz_bot/internall/models"
	"kz_bot/internall/storage"
	"kz_bot/pkg/clientDiscord"
)

type Discord struct {
	inbox   chan models.InMessage
	s       *discordgo.Session
	log     *logrus.Logger
	storage *storage.Storage
}

func NewDiscord(inbox chan models.InMessage, log *logrus.Logger, st *storage.Storage, cfg config.ConfigBot) *Discord {
	ds, err := clientDiscord.NewDiscord(log, cfg)
	if err != nil {
		log.Println("error running Discord " + err.Error())
	}
	DS := &Discord{
		inbox:   inbox,
		s:       ds,
		log:     log,
		storage: st,
	}
	ds.AddHandler(DS.messageHandler)
	ds.AddHandler(DS.messageReactionAdd)
	ds.AddHandler(DS.slash)
	ds.AddHandler(DS.ready)

	return DS
}
