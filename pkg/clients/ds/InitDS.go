package ds

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

func InitDS(TokenD string) *discordgo.Session {
	DSBot, err := discordgo.New("Bot " + TokenD)
	if err != nil {
		log.Println(err)
	}
	/*
		u, err := DSBot.User("@me")
		if err != nil {
			log.Println(err.Error())
		}
	*/
	DSBot.AddHandler(messageHandler)
	//DSBot.AddHandler(MessageReactionAdd)

	err = DSBot.Open()
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("Бот DISCORD запущен!!!")
	return DSBot
}
