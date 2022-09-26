package discordClient

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"kz_bot/internal/models"
)

func (d *Discord) readReactionQueue(r *discordgo.MessageReactionAdd, message *discordgo.Message) {
	user, err := d.d.User(r.UserID)
	if err != nil {
		d.log.Println("Ошибка получения Юзера по реакции ", err)
	}
	if user.ID != message.Author.ID {
		ok, config := d.CorpConfig.CheckChannelConfigDS(r.ChannelID)
		if ok {
			member, e := d.d.GuildMember(config.Config.Guildid, user.ID)
			if e != nil {
				d.log.Println("Oшибка получения участника ", e)
			}
			name := user.Username
			if member.Nick != "" {
				name = member.Nick
			}
			Avatar := "https://cdn.discordapp.com/avatars/" + user.ID + "/" + user.Avatar + ".jpg"

			in := models.InMessage{
				Mtext:       "",
				Tip:         "ds",
				Name:        name,
				NameMention: user.Mention(),
				Ds: struct {
					Mesid   string
					Nameid  string
					Guildid string
					Avatar  string
				}{
					Mesid:   r.MessageID,
					Nameid:  user.ID,
					Guildid: config.Config.Guildid,
					Avatar:  Avatar,
				},

				Config: config,
				Option: struct {
					Callback bool
					Edit     bool
					Update   bool
					Queue    bool
				}{
					Callback: true,
					Edit:     true,
					Update:   false,
				},
			}
			d.reactionUserRemove(r)

			if r.Emoji.Name == emPlus {
				in.Mtext = "+"
			} else if r.Emoji.Name == emMinus {
				in.Mtext = "-"
			} else if r.Emoji.Name == emOK || r.Emoji.Name == emCancel || r.Emoji.Name == emRsStart || r.Emoji.Name == emPl30 {
				in.Lvlkz, err = d.dbase.ReadMesIdDS(r.MessageID)
				if err == nil && in.Lvlkz != "" {
					if r.Emoji.Name == emOK {
						in.Timekz = "30"
						in.Mtext = in.Lvlkz + "+"
					} else if r.Emoji.Name == emCancel {
						in.Mtext = in.Lvlkz + "-"
					} else if r.Emoji.Name == emRsStart {
						in.Mtext = in.Lvlkz + "++"
					} else if r.Emoji.Name == emPl30 {
						in.Mtext = in.Lvlkz + "+++"
					}
				}
			}
			if d.debug {
				fmt.Printf("\n\nin readReactionQueue %+v\n ", in)
			}

			models.ChDs <- in
		}
	}
}

func (d *Discord) reactionUserRemove(r *discordgo.MessageReactionAdd) {
	err := d.d.MessageReactionRemove(r.ChannelID, r.MessageID, r.Emoji.Name, r.UserID)
	if err != nil {
		d.log.Println("Ошибка удаления эмоджи", err)
	}
}

func (d *Discord) logicMixDiscord(m *discordgo.MessageCreate) {
	ok, config := d.CorpConfig.CheckChannelConfigDS(m.ChannelID)
	d.AccesChatDS(m)
	if ok {
		if len(m.Attachments) > 0 {
			for _, attach := range m.Attachments { //вложеные файлы
				m.Content = m.Content + "\n" + attach.URL
			}
		}
		member, e := d.d.GuildMember(m.GuildID, m.Author.ID) //проверка есть ли изменения имени в этом дискорде
		if e != nil {
			d.log.Println("Ошибка получения ника пользователя", e, m.ID)
		}
		name := m.Author.Username
		if member.Nick != "" {
			name = member.Nick
		}
		Avatar := "https://cdn.discordapp.com/avatars/" + m.Author.ID + "/" + m.Author.Avatar + ".jpg"

		in := models.InMessage{
			Mtext:       m.Content,
			Tip:         "ds",
			Name:        name,
			NameMention: m.Author.Mention(),
			Ds: struct {
				Mesid   string
				Nameid  string
				Guildid string
				Avatar  string
			}{
				Mesid:   m.ID,
				Nameid:  m.Author.ID,
				Guildid: m.GuildID,
				Avatar:  Avatar,
			},
			Config: config,
			Option: struct {
				Callback bool
				Edit     bool
				Update   bool
				Queue    bool
			}{
				Callback: false,
				Edit:     false,
				Update:   false,
			},
		}
		if d.debug {
			fmt.Printf("\n\nin logicMixDiscord %+v\n", in)
		}

		models.ChDs <- in
	}
}
