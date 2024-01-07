package DiscordClient

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"kz_bot/internal/models"
	"net/url"
	"time"
)

func (d *Discord) messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Message.WebhookID != "" {
		return
	}

	d.logicMix(m)

}
func (d *Discord) messageUpdate(s *discordgo.Session, m *discordgo.MessageUpdate) { //nolint:unparam
	if m.Message.WebhookID != "" {
		return
	}

	if m.Message.EditedTimestamp != nil && m.Content != "" {
		good, config := d.BridgeCheckChannelConfigDS(m.ChannelID)
		if good {
			username := m.Author.Username
			if m.Member != nil && m.Member.Nick != "" {
				username = m.Member.Nick
			}
			mes := models.BridgeMessage{
				Text:   d.replaceTextMessage(m.Content, m.GuildID),
				Sender: username,
				Tip:    "dse",
				Ds: &models.BridgeMessageDs{
					ChatId:        m.ChannelID,
					MesId:         m.ID,
					Avatar:        m.Author.AvatarURL("128"),
					GuildId:       m.GuildID,
					TimestampUnix: m.Timestamp.Unix(),
				},
				Config: &config,
			}

			if len(m.Attachments) > 0 {
				if len(m.Attachments) != 1 {
					d.log.Info(fmt.Sprintf("вложение %d", len(m.Attachments)))
				}

				// Разбираем URL
				parsedURL, err := url.Parse(m.Attachments[0].URL)
				if err != nil {
					d.log.Error(err.Error())
				}

				// Очищаем параметры запроса (query parameters) и фрагмент
				parsedURL.RawQuery = ""
				parsedURL.Fragment = ""

				// Получаем очищенную ссылку
				mes.FileUrl = parsedURL.String()
			}

			if m.ReferencedMessage != nil {
				usernameR := m.ReferencedMessage.Author.String() //.Username
				if m.ReferencedMessage.Member != nil && m.ReferencedMessage.Member.Nick != "" {
					usernameR = m.ReferencedMessage.Member.Nick
				}
				mes.Ds.Reply = &models.ReplyDs{
					TimeMessage: m.ReferencedMessage.Timestamp.Unix(),
					Text:        d.replaceTextMessage(m.ReferencedMessage.Content, m.GuildID),
					Avatar:      m.ReferencedMessage.Author.AvatarURL("128"),
					UserName:    usernameR,
				}
			}

			d.ChanBridgeMessage <- mes
		}
	}
}
func (d *Discord) onMessageDelete(s *discordgo.Session, m *discordgo.MessageDelete) {
	good, config := d.BridgeCheckChannelConfigDS(m.ChannelID)
	if good {
		d.ChanBridgeMessage <- models.BridgeMessage{
			Tip: "del",
			Ds: &models.BridgeMessageDs{
				MesId: m.ID,
			},
			Config: &config,
		}
	}
}

func (d *Discord) messageReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	message, err := s.ChannelMessage(r.ChannelID, r.MessageID)
	if err != nil {
		channel, err1 := s.Channel(r.ChannelID)
		if err1 != nil {
			d.log.Error(err1.Error())
			return
		}
		user, err2 := s.User(r.UserID)
		if err2 != nil {
			d.log.Error(err2.Error())
			return
		}
		d.log.Info(fmt.Sprintln(channel.Name, r.Emoji.Name, user.Username, err.Error()))
		return
	}

	if message.Author.ID == s.State.User.ID {
		d.readReactionQueue(r, message)
	}
}

func (d *Discord) slash(s *discordgo.Session, i *discordgo.InteractionCreate) {
	commandHandlers := d.addSlashHandler()
	if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
		h(s, i)
	}
}
func (d *Discord) ready(s *discordgo.Session, r *discordgo.Ready) {
	commands := d.addSlashCommand()
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		if err != nil {

			d.log.Error(err.Error())
		}
		registeredCommands[i] = cmd
	}
}

func (d *Discord) removeComand(s *discordgo.Session) {
	registeredCommands, err := s.ApplicationCommands(s.State.User.ID, "700238199070523412")
	if err != nil {
		d.log.Fatal(err.Error())
	}

	for _, v := range registeredCommands {
		err := s.ApplicationCommandDelete(s.State.User.ID, "700238199070523412", v.ID)
		if err != nil {
			d.log.Error(err.Error())
		}
	}
	fmt.Println("удалены")
}
func (d *Discord) addSlashCommand() []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand{
		{
			Name:        "help",
			Description: "Общая справка",
		},
		{
			Name:        "helpqueue",
			Description: "Очередь КЗ",
		},
		{
			Name:        "helpnotification",
			Description: "Уведомления",
		},
		{
			Name:        "helpevent",
			Description: "Ивент КЗ",
		},
		{
			Name:        "helptop",
			Description: "ТОП лист",
		},
		{
			Name:        "helpicon",
			Description: "Работа с иконками",
		},
	}
}
func (d *Discord) addSlashHandler() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"help": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{

					Content: models.Help,
				},
			})
			go func() {
				time.Sleep(1 * time.Minute)
				s.InteractionResponseDelete(i.Interaction)
			}()
		},
		"helpqueue": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: models.HelpQueue,
				},
			})
			go func() {
				time.Sleep(1 * time.Minute)
				s.InteractionResponseDelete(i.Interaction)
			}()
		},
		"helpnotification": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Уведомления:\n" +
						"	Подписаться на уведомления о начале очереди: +[4-11]\n" +
						"+10 -подписаться на уведомления о начале очереди на КЗ 10ур.\n\n" +
						"	Подписаться на уведомление, если в очереди 3 человека: ++[4-11]\n" +
						"++10 -подписаться на уведомления о наличии 3х человек в очереди на КЗ 10ур.\n\n" +
						"	Отключить уведомления о начале сбора: -[5-11]\n" +
						"-9 -отключить уведомления о начале сборе на КЗ 9ур.\n\n" +
						"	Отключить уведомления 3/4 в очереди: --[5-11]\n" +
						"--9 -отключить уведомления о наличии 3х человек в очереди на КЗ 9ур.",
				},
			})
			go func() {
				time.Sleep(1 * time.Minute)
				s.InteractionResponseDelete(i.Interaction)
			}()
		},
		"helpevent": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: models.HelpEvent,
				},
			})
			go func() {
				time.Sleep(1 * time.Minute)
				s.InteractionResponseDelete(i.Interaction)
			}()
		},
		"helptop": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: models.HelpTop,
				},
			})
			go func() {
				time.Sleep(1 * time.Minute)
				s.InteractionResponseDelete(i.Interaction)
			}()
		},
		"helpicon": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: models.HelpIcon,
				},
			})
			go func() {
				time.Sleep(1 * time.Minute)
				s.InteractionResponseDelete(i.Interaction)
			}()
		},
	}
	return commandHandlers
}
