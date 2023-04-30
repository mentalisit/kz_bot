package DiscordClient

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"kz_bot/internal/config"
	"kz_bot/internal/models"
	"kz_bot/internal/storage"
	"kz_bot/pkg/clientDiscord"
)

type Discord struct {
	inbox      chan models.InMessage
	sendToGame chan models.Message
	globalChat chan models.InGlobalMessage
	s          *discordgo.Session
	log        *logrus.Logger
	storage    *storage.Storage
}

func NewDiscord(inbox chan models.InMessage, sendToGame chan models.Message, log *logrus.Logger, st *storage.Storage, cfg *config.ConfigBot, global chan models.InGlobalMessage) *Discord {
	ds, err := clientDiscord.NewDiscord(log, cfg)
	if err != nil {
		log.Println("error running Discord " + err.Error())
	}
	DS := &Discord{
		inbox:      inbox,
		globalChat: global,
		s:          ds,
		log:        log,
		storage:    st,
		sendToGame: sendToGame,
	}
	ds.AddHandler(DS.messageHandler)
	ds.AddHandler(DS.messageReactionAdd)
	ds.AddHandler(DS.slash)
	ds.AddHandler(DS.ready)

	return DS
}
