package ds

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

var (
	DSBot *discordgo.Session
	err   error
)

func InitDS(TokenD string) *discordgo.Session {
	DSBot, err = discordgo.New("Bot " + TokenD)
	if err != nil {
		panic(err)
	}

	DSBot.AddHandler(messageHandler)
	DSBot.AddHandler(MessageReactionAdd)

	err = DSBot.Open()
	if err != nil {
		log.Println("Ошибка открытия ДС", err)
	}
	fmt.Println("Бот DISCORD запущен!!!")
	return DSBot
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	logicMixDiscord(m)

}

func MessageReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	message, err := DSBot.ChannelMessage(r.ChannelID, r.MessageID)
	if err != nil {
		fmt.Println("Ошибка чтения реакции в ДС", err)
	}
	if message.Author.ID == s.State.User.ID {
		readReactionQueue(r, message)
	}
}
