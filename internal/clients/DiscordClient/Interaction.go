package DiscordClient

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"kz_bot/internal/models"
	"strconv"
	"time"
)

//// register slash command module
//func (d *Discord) registerCommand(guildID string) {
//	// Регистрация слеш-команды с параметрами "module" и "level"
//	cmd := &discordgo.ApplicationCommand{
//		Name:        "module",
//		Description: "Выберите нужный модуль и уровень / Select the desired module and level",
//		Options: []*discordgo.ApplicationCommandOption{
//			{
//				Type:        discordgo.ApplicationCommandOptionString,
//				Name:        "module",
//				Description: "Выберите модуль / Select module",
//				Required:    true,
//				Choices: []*discordgo.ApplicationCommandOptionChoice{
//					{
//						Name:  "Ингибитор КЗ / RSE",
//						Value: "RSE",
//					},
//					{
//						Name:  "Генезис / Genesis",
//						Value: "GENESIS",
//					},
//					{
//						Name:  "Обогатить / Enrich",
//						Value: "ENRICH",
//					},
//					// Добавьте другие модули по мере необходимости
//				},
//			},
//			{
//				Type:        discordgo.ApplicationCommandOptionInteger,
//				Name:        "level",
//				Description: "Выберите уровень / Select level",
//				Required:    true,
//				Choices: []*discordgo.ApplicationCommandOptionChoice{
//					{
//						Name:  "Уровень / Level 0",
//						Value: 0,
//					},
//					{
//						Name:  "Уровень / Level 1",
//						Value: 1,
//					}, {
//						Name:  "Уровень / Level 2",
//						Value: 2,
//					}, {
//						Name:  "Уровень / Level 3",
//						Value: 3,
//					}, {
//						Name:  "Уровень / Level 4",
//						Value: 4,
//					}, {
//						Name:  "Уровень / Level 5",
//						Value: 5,
//					}, {
//						Name:  "Уровень / Level 6",
//						Value: 6,
//					}, {
//						Name:  "Уровень / Level 7",
//						Value: 7,
//					}, {
//						Name:  "Уровень / Level 8",
//						Value: 8,
//					}, {
//						Name:  "Уровень / Level 9",
//						Value: 9,
//					}, {
//						Name:  "Уровень / Level 10",
//						Value: 10,
//					}, {
//						Name:  "Уровень / Level 11",
//						Value: 11,
//					}, {
//						Name:  "Уровень / Level 12",
//						Value: 12,
//					}, {
//						Name:  "Уровень / Level 13",
//						Value: 13,
//					}, {
//						Name:  "Уровень / Level 14",
//						Value: 14,
//					}, {
//						Name:  "Уровень / Level 15",
//						Value: 15,
//					},
//					// Добавьте другие уровни по мере необходимости
//				},
//			},
//		},
//	}
//
//	_, err := d.s.ApplicationCommandCreate(d.s.State.User.ID, guildID, cmd)
//	if err != nil {
//		d.log.Error("Error registering command: " + err.Error())
//		return
//	}
//	// Регистрация слеш-команды оружие
//	cmd = &discordgo.ApplicationCommand{
//		Name:        "weapon",
//		Description: "Выберите основное оружие / Select your main weapon",
//		Options: []*discordgo.ApplicationCommandOption{
//			{
//				Type:        discordgo.ApplicationCommandOptionString,
//				Name:        "weapon",
//				Description: "Выберите оружие / Select weapon",
//				Required:    true,
//				Choices: []*discordgo.ApplicationCommandOptionChoice{
//					{
//						Name:  "Артобстрел / Barrage",
//						Value: "barrage",
//					},
//					{
//						Name:  "Лазер / Laser",
//						Value: "laser",
//					},
//					{
//						Name:  "Цепной луч / Chain ray",
//						Value: "chainray",
//					},
//					{
//						Name:  "Батарея / Battery",
//						Value: "battery",
//					},
//					{
//						Name:  "Залповая батарея / Mass battery",
//						Value: "massbattery",
//					},
//					{
//						Name:  "Пусковая установка / Dart launcher",
//						Value: "dartlauncher",
//					},
//					{
//						Name:  "Ракетная установка / Rocket launcher",
//						Value: "rocketlauncher",
//					},
//					// Добавьте другие модули по мере необходимости
//				},
//			},
//		},
//	}
//
//	_, err = d.s.ApplicationCommandCreate(d.s.State.User.ID, guildID, cmd)
//	if err != nil {
//		d.log.Error("Error registering command:" + err.Error())
//		return
//	}
//
//	fmt.Println("Command registered successfully.")
//}

// slash command module respond
func (d *Discord) handleModuleCommand(i *discordgo.InteractionCreate) {
	module := i.ApplicationCommandData().Options[0].StringValue()
	level := i.ApplicationCommandData().Options[1].IntValue()

	response := fmt.Sprintf("Выбран модуль: %s, уровень: %d", module, level)
	if level == 0 {
		response = fmt.Sprintf("Удален модуль: %s, уровень: %d", module, level)
	}
	// Отправка ответа
	err := d.s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: response,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		time.Sleep(20 * time.Second)
		err = d.s.InteractionResponseDelete(i.Interaction)
		if err != nil {
			return
		}
	}()
	d.updateModuleOrWeapon(i.Interaction.Member.User.Username, module, strconv.FormatInt(level, 10))
}

// slash command weapon respond
func (d *Discord) handleWeaponCommand(i *discordgo.InteractionCreate) {
	weapon := i.ApplicationCommandData().Options[0].StringValue()

	response := fmt.Sprintf("Установлено оружие: %s", weapon)

	// Отправка ответа
	err := d.s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: response,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		time.Sleep(20 * time.Second)
		err = d.s.InteractionResponseDelete(i.Interaction)
		if err != nil {
			return
		}
	}()
	d.updateModuleOrWeapon(i.Interaction.Member.User.Username, weapon, "")
}

func (d *Discord) updateModuleOrWeapon(username, module, level string) {
	rse := "<:rse:1199068829511335946> " + level
	genesis := "<:genesis:1199068748280242237> " + level
	enrich := "<:enrich:1199068793633251338> " + level
	if level == "0" {
		rse, genesis, enrich = "", "", ""
	}

	barrage := "<:barrage:1199084425393225782>"
	laser := "<:laser:1199084197571207339>"
	chainray := "<:chainray:1199073579577376888>"
	battery := "<:batteryw:1199072534562345021>"
	massbattery := "<:massbattery:1199072493760151593>"
	dartlauncher := "<:dartlauncher:1199072434674991145>"
	rocketlauncher := "<:rocketlauncher:1199071677548605562>"
	t := d.storage.Emoji.EmojiModuleReadUsers(context.Background(), username, "ds")
	if len(t.Name) == 0 {
		d.storage.Emoji.EmInsertEmpty(context.Background(), "ds", username)
	}
	switch module {
	case "RSE":
		d.storage.Emoji.ModuleUpdate(context.Background(), username, "ds", "1", rse)
	case "GENESIS":
		d.storage.Emoji.ModuleUpdate(context.Background(), username, "ds", "2", genesis)
	case "ENRICH":
		d.storage.Emoji.ModuleUpdate(context.Background(), username, "ds", "3", enrich)
	case "barrage":
		d.storage.Emoji.WeaponUpdate(context.Background(), username, "ds", barrage)
	case "laser":
		d.storage.Emoji.WeaponUpdate(context.Background(), username, "ds", laser)
	case "chainray":
		d.storage.Emoji.WeaponUpdate(context.Background(), username, "ds", chainray)
	case "battery":
		d.storage.Emoji.WeaponUpdate(context.Background(), username, "ds", battery)
	case "massbattery":
		d.storage.Emoji.WeaponUpdate(context.Background(), username, "ds", massbattery)
	case "dartlauncher":
		d.storage.Emoji.WeaponUpdate(context.Background(), username, "ds", dartlauncher)
	case "rocketlauncher":
		d.storage.Emoji.WeaponUpdate(context.Background(), username, "ds", rocketlauncher)
	case "Remove":
		d.storage.Emoji.WeaponUpdate(context.Background(), username, "ds", "")
	}
}

func (d *Discord) handleButtonPressed(i *discordgo.InteractionCreate) {
	ok, config := d.CheckChannelConfigDS(i.ChannelID)
	if ok {
		in := models.InMessage{
			Mtext:       i.MessageComponentData().CustomID,
			Tip:         "ds",
			Name:        i.Interaction.Member.User.Username,
			NameMention: i.Interaction.Member.User.Mention(),
			Ds: struct {
				Mesid   string
				Nameid  string
				Guildid string
				Avatar  string
			}{
				Mesid:   i.Interaction.Message.ID,
				Nameid:  i.Interaction.Member.User.ID,
				Guildid: i.Interaction.GuildID,
				Avatar:  i.Interaction.Member.User.AvatarURL("128"),
			},
			Config: config,
			Option: models.Option{Reaction: true},
		}
		d.ChanRsMessage <- in
		err := d.s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseDeferredMessageUpdate,
		})
		if err != nil {
			d.log.ErrorErr(err)
			return
		}
	}
}
