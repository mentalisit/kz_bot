package discordClient

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func (d *Ds) InitDS(TokenD string, db *sql.DB) {
	DSBot, err := discordgo.New("Bot " + TokenD)
	if err != nil {
		fmt.Println(err)
		return
		//panic(err)
	}

	DSBot.AddHandler(d.messageHandler)
	DSBot.AddHandler(d.MessageReactionAdd)

	err = DSBot.Open()
	if err != nil {
		log.Println("Ошибка открытия ДС", err)
		fmt.Println(err)
		return
		//panic(err)
	}
	fmt.Println("Бот DISCORD запущен!!!")
	d.d = *DSBot
	d.dbase.Db = db
	return
}

func (d *Ds) messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	d.logicMixDiscord(m)

}

func (d *Ds) MessageReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	message, err := d.d.ChannelMessage(r.ChannelID, r.MessageID)
	if err != nil {
		fmt.Println("Ошибка чтения реакции в ДС", err)
	}
	if message.Author.ID == s.State.User.ID {
		d.readReactionQueue(r, message)
	}
}
