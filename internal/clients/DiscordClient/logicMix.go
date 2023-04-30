package DiscordClient

import (
	"context"
	"github.com/bwmarrin/discordgo"
	"kz_bot/internal/models"
	"time"
)

const (
	emOK      = "✅"
	emCancel  = "❎"
	emRsStart = "🚀"
	emPl30    = "⌛"
	emPlus    = "➕"
	emMinus   = "➖"
)

func (d *Discord) readReactionQueue(r *discordgo.MessageReactionAdd, message *discordgo.Message) {
	user, err := d.s.User(r.UserID)
	if err != nil {
		d.log.Println("Ошибка получения Юзера по реакции ", err)
	}
	if user.ID != message.Author.ID {
		ok, config := d.storage.Cache.CheckChannelConfigDS(r.ChannelID)
		if ok {
			member, e := d.s.GuildMember(config.Guildid, user.ID)
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
					Guildid: config.Guildid,
					Avatar:  Avatar,
				},

				Config: config,
				Option: models.Option{
					Reaction: true},
			}
			d.reactionUserRemove(r)

			if r.Emoji.Name == emPlus {
				in.Mtext = "+"
			} else if r.Emoji.Name == emMinus {
				in.Mtext = "-"
			} else if r.Emoji.Name == emOK || r.Emoji.Name == emCancel || r.Emoji.Name == emRsStart || r.Emoji.Name == emPl30 {
				ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
				defer cancel()
				in.Lvlkz, err = d.storage.DbFunc.ReadMesIdDS(ctx, r.MessageID)
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
			d.inbox <- in
		}
	}
}

func (d *Discord) reactionUserRemove(r *discordgo.MessageReactionAdd) {
	err := d.s.MessageReactionRemove(r.ChannelID, r.MessageID, r.Emoji.Name, r.UserID)
	if err != nil {
		d.log.Println("Ошибка удаления эмоджи", err)
	}
}

func (d *Discord) logicMix(m *discordgo.MessageCreate) {
	if d.ifMessageForHades(m) {
		return
	}
	ok, config := d.storage.Cache.CheckChannelConfigDS(m.ChannelID)
	d.AccesChatDS(m)
	if ok {
		if len(m.Attachments) > 0 {
			for _, attach := range m.Attachments { //вложеные файлы
				m.Content = m.Content + "\n" + attach.URL
			}
		}
		member, e := d.s.GuildMember(m.GuildID, m.Author.ID) //проверка есть ли изменения имени в этом дискорде
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
			Option: models.Option{InClient: true},
		}
		d.inbox <- in
	}
	if !ok {
		go d.logicMixGlobal(m)
	}
}
