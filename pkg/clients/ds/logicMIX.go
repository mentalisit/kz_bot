package ds

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"kz_bot/pkg/models"
)

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	logicMixDiscord(m)

}

func logicMixDiscord(m *discordgo.MessageCreate) {
	var ok = true
	if ok {
		in := models.InMessage{
			Mtext:       m.Content,
			Tip:         "ds",
			Name:        m.Author.Username,
			NameMention: "@" + m.Author.Mention(),
			Ds: models.Ds{
				Mesid:   m.ID,
				Nameid:  m.Author.ID,
				Guildid: m.GuildID,
			},
			Tg: models.Tg{},
			Option: models.Option{
				Callback: false,
				Edit:     false,
				Update:   false,
			},
		}
		//logicRs(in)
		//тут нужно передавать в логику бота
		fmt.Println(in)

	}
}
