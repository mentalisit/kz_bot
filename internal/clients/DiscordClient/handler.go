package DiscordClient

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"kz_bot/internal/models"
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

func (d *Discord) messageReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	message, err := s.ChannelMessage(r.ChannelID, r.MessageID)
	if err != nil {
		channel, err1 := s.Channel(r.ChannelID)
		if err1 != nil {
			return
		}
		user, err2 := s.User(r.UserID)
		if err2 != nil {
			return
		}
		d.log.Println(fmt.Sprintln(channel.Name, r.Emoji, user.Username))
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

			d.log.Printf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}
}

func (d *Discord) removeComand(s *discordgo.Session) {
	fmt.Println("Removing commands...", "700238199070523412")
	registeredCommands, err := s.ApplicationCommands(s.State.User.ID, "700238199070523412")
	fmt.Println("registeredCommands", registeredCommands)
	if err != nil {
		d.log.Fatalf("Could not fetch registered commands: %v", err)
	}

	for _, v := range registeredCommands {
		err := s.ApplicationCommandDelete(s.State.User.ID, "700238199070523412", v.ID)
		if err != nil {
			d.log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
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
