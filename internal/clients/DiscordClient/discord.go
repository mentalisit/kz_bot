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
	ChanRsMessage  chan models.InMessage
	ChanToGame     chan models.MessageHades
	ChanGlobalChat chan models.InGlobalMessage
	ChanRelay      chan models.RelayMessage
	s              *discordgo.Session
	log            *logrus.Logger
	storage        *storage.Storage
}

func NewDiscord(log *logrus.Logger, st *storage.Storage, cfg *config.ConfigBot) *Discord {
	ds, err := clientDiscord.NewDiscord(log, cfg)
	if err != nil {
		log.Println("error running Discord " + err.Error())
	}

	DS := &Discord{
		s:              ds,
		log:            log,
		storage:        st,
		ChanRsMessage:  make(chan models.InMessage, 10),
		ChanToGame:     make(chan models.MessageHades, 10),
		ChanGlobalChat: make(chan models.InGlobalMessage, 20),
		ChanRelay:      make(chan models.RelayMessage, 20),
	}
	ds.AddHandler(DS.messageHandler)
	ds.AddHandler(DS.messageUpdate)
	ds.AddHandler(DS.messageReactionAdd)
	ds.AddHandler(DS.onMessageDelete)
	ds.AddHandler(DS.slash)
	ds.AddHandler(DS.ready)

	return DS
}
