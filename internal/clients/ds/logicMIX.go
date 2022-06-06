package ds

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"kz_bot/internal/clients/corpConfig"
	"kz_bot/internal/models"
	"log"
)

func (d Ds) readReactionQueue(r *discordgo.MessageReactionAdd, message *discordgo.Message) {
	user, err := d.d.User(r.UserID)
	if err != nil {
		fmt.Println("Ошибка получения Юзера по реакции ", err)
	}
	if user.ID != message.Author.ID {
		/*
			ok, config := checkChannelConfigDS(r.ChannelID)
			if ok {
				member, e := DSBot.GuildMember(config.Config.Guildid, user.ID)
				if e != nil {
					logrus.Println("ошибка в функдиск стр57", e)
				}
				name := user.Username
				if member.Nick != "" {
					name = member.Nick
				}

				in := inMessage{
					mtext:       "",
					tip:         "ds",
					name:        name,
					nameMention: user.Mention(),
					Ds: Ds{
						mesid:   r.MessageID,
						nameid:  user.ID,
						guildid: message.GuildID,
					},
					Tg: Tg{
						mesid:  0,
						nameid: 0,
					},
					config: config,
					option: Option{
						callback: true,
						edit:     true,
						update:   false,
					},
				}
				reactionUserRemove(r)
				if r.Emoji.Name == emPlus {
					if in.Plus() {
						dsDeleteMesage5s(in.config.DsChannel, in.Ds.mesid)
					}
				} else if r.Emoji.Name == emMinus {
					if in.Minus() {
						dsDelMessage(in.config.DsChannel, in.Ds.mesid)
					}
				} else if r.Emoji.Name == emOK || r.Emoji.Name == emCancel || r.Emoji.Name == emRsStart || r.Emoji.Name == emPl30 {
					in.lvlkz, err = readMesID(r.MessageID)
					if err == nil && in.lvlkz != "" {
						if r.Emoji.Name == emOK {
							in.timekz = "30"
							in.RsPlus()
						} else if r.Emoji.Name == emCancel {
							in.RsMinus()
						} else if r.Emoji.Name == emRsStart {
							in.RsStart()
						} else if r.Emoji.Name == emPl30 {
							in.Pl30()
						}
					}
				}
			}

		*/
	}
}

func (d Ds) reactionUserRemove(r *discordgo.MessageReactionAdd) {
	err := d.d.MessageReactionRemove(r.ChannelID, r.MessageID, r.Emoji.Name, r.UserID)
	if err != nil {
		log.Println("Ошибка удаления эмоджи", err)
	}
}

func (d Ds) logicMixDiscord(m *discordgo.MessageCreate) {
	c := corpsConfig.CorpConfig{}
	ok, config := c.CheckChannelConfigDS(m.ChannelID)
	d.AccesChatDS(m)
	if ok {
		if len(m.Attachments) > 0 {
			for _, attach := range m.Attachments { //вложеные файлы
				m.Content = m.Content + "\n" + attach.URL
			}
		}
		member, e := d.d.GuildMember(m.GuildID, m.Author.ID) //проверка есть ли изменения имени в этом дискорде
		if e != nil {
			fmt.Println("Ошибка получения ника пользователя", e)
		}
		name := m.Author.Username
		if member.Nick != "" {
			name = member.Nick
		}

		in := models.InMessage{
			Mtext:       m.Content,
			Tip:         "ds",
			Name:        name,
			NameMention: m.Author.Mention(),
			Ds: models.Ds{
				Mesid:   m.ID,
				Nameid:  m.Author.ID,
				Guildid: m.GuildID,
			},
			Tg:     models.Tg{},
			Config: config,
			Option: models.Option{
				Callback: false,
				Edit:     false,
				Update:   false,
			},
		}
		//logicRs(in)
		//тут нужно передавать в логику бота
		models.ChDs <- in

	}
}
