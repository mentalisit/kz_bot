package DiscordClient

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func (d *Discord) messageCreate(m *discordgo.MessageCreate) {
	if m.Author.ID == d.s.State.User.ID {
		return
	}

	if m.Content == "Меню модулей" {
		// Создаем элементы для меню
		moduleMenu := createModuleSelectMenu()

		msg := &discordgo.MessageSend{
			Content: "Выберите модуль:",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{moduleMenu},
				},
			},
		}

		_, err := d.s.ChannelMessageSendComplex(m.ChannelID, msg)
		if err != nil {
			log.Println("Error sending message:", err)
		}
	} else if m.Content == "М" {
		// Регистрация слеш-команды при получении команды "!registerCommand"
		registerCommand(d.s, m.GuildID)
	}
}
func createModuleSelectMenu() *discordgo.SelectMenu {
	menu := &discordgo.SelectMenu{
		CustomID:    "moduleSelect",
		Placeholder: "Выберите модуль",
		Options: []discordgo.SelectMenuOption{
			{
				Label:       "RSE",
				Value:       "RSE",
				Description: "ингибитор кз",
				Emoji: discordgo.ComponentEmoji{
					Name: ":rse:",
					ID:   "1199068829511335946",
				},
			},
			{
				Label:       "GENESIS",
				Value:       "GENESIS",
				Description: "генезис",
				Emoji: discordgo.ComponentEmoji{
					Name: ":genesis:",
					ID:   "1199068748280242237",
				},
			},
			{
				Label:       "ENRICH",
				Value:       "ENRICH",
				Description: "обогащение",
				Emoji: discordgo.ComponentEmoji{
					Name: ":genesis:",
					ID:   "1199068793633251338",
				},
			},
		},
	}
	return menu
}
func createModuleSelectMenu1() discordgo.SelectMenu {
	// Создаем элементы для основного меню с вложенными опциями
	menu := discordgo.SelectMenu{
		CustomID:    "module_RSE",
		Placeholder: "Выберите уровень",
		Options: []discordgo.SelectMenuOption{
			{
				Label:       ":custom_emoji5: Опция 1",
				Value:       "option1",
				Description: "Описание для опции 1 :custom_emoji5:",
				Emoji: discordgo.ComponentEmoji{
					Name: ":rse:",
					ID:   "1199068829511335946",
				},
				//Options:     option1Nested, // Вложенные опции для опции 1
			},
			{
				Label:       ":custom_emoji6: Опция 2",
				Value:       "option2",
				Description: "Описание для опции 2 :custom_emoji6:",
				Emoji: discordgo.ComponentEmoji{
					Name: ":rse:",
					ID:   "1199068829511335946",
				},
				//Options:     option2Nested, // Вложенные опции для опции 2
			},
			{
				Label:       ":custom_emoji7: Опция 3",
				Value:       "option3",
				Description: "Описание для опции 3 :custom_emoji7:",
				Emoji: discordgo.ComponentEmoji{
					Name: ":rse:",
					ID:   "1199068829511335946",
				},
			},
		},
	}

	// Создаем сообщение с выбором
	//msg := &discordgo.MessageSend{
	//	Content: "Выберите опцию:",
	//	Components: []discordgo.MessageComponent{
	//		discordgo.ActionsRow{
	//			Components: []discordgo.MessageComponent{&menu},
	//		},
	//	},
	//}
	return menu

}
