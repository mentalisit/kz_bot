package DiscordClient

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"kz_bot/internal/models"
	"kz_bot/internal/storage/CorpsConfig/hades"
	"regexp"
	"strings"
)

func (d *Discord) ifComands(m *discordgo.MessageCreate) (command bool) {
	str, ok := strings.CutPrefix(m.Content, ". ")
	if ok {
		str = strings.ToLower(str)
		arr := strings.Split(str, " ")
		arrlen := len(arr)
		if arrlen == 1 {

		} else if arrlen == 2 {
			if d.avatar(arr, m) {
				return true
			}
			if d.lastWs(arr, m) {
				return true
			}
			if d.replayId(arr, m) {
				return true
			}
			if d.historyWs(arr, m) {
				return true
			}
			if d.letInId(arr, m) {
				return true
			}

		}
	}
	return false
}
func (d *Discord) avatar(arg []string, m *discordgo.MessageCreate) bool {
	if arg[0] == "ава" {
		mentionIds := userMentionRE.FindAllStringSubmatch(arg[1], -1)
		if len(mentionIds) > 0 {
			members, err := d.s.GuildMembers(m.GuildID, "", 999)
			if err != nil {
				d.log.Println("error getGuildMember " + err.Error())
			}
			for _, member := range members {
				if member.User.ID == mentionIds[0][1] {
					em := &discordgo.MessageEmbed{
						Title: "Аватар по запросу " + m.Author.Username,
						Color: 14232643,
						Image: &discordgo.MessageEmbedImage{
							URL: member.AvatarURL("2048"),
						},
						Author: nil,
					}
					embed, err := d.s.ChannelMessageSendEmbed(m.ChannelID, em)
					if err != nil {
						fmt.Println(err.Error())
						return false
					}
					go d.DeleteMesageSecond(m.ChannelID, embed.ID, 180)
					return true
				}
			}
		}
	}
	return false
}
func (d *Discord) lastWs(arg []string, m *discordgo.MessageCreate) bool {
	if arg[0] == "повтор" && arg[1] == "бз" {
		_, corporation := hades.HadesStorage.AllianceChat(m.ChannelID)
		mes := models.Message{
			Text:        "",
			Sender:      m.Author.Username,
			Avatar:      "",
			ChannelType: 0,
			Corporation: corporation.Corp,
			Command:     "повтор бз",
			Messager:    "ds",
		}
		fmt.Printf("lastWs %+v\n", mes)
		d.sendToGame <- mes
		go d.SendChannelDelSecond(m.ChannelID, "отправка повтора последней бз", 10)
		go d.DeleteMesageSecond(m.ChannelID, m.ID, 180)
		return true
	}
	return false
}
func (d *Discord) replayId(arg []string, m *discordgo.MessageCreate) bool {
	if arg[0] == "повтор" {
		match, _ := regexp.MatchString("^[0-9]+$", arg[1])
		if match {
			_, corporation := hades.HadesStorage.AllianceChat(m.ChannelID)
			mes := models.Message{
				Text:        arg[1],
				Sender:      m.Author.Username,
				Avatar:      "",
				ChannelType: 0,
				Corporation: corporation.Corp,
				Command:     "повтор",
				Messager:    "ds",
			}
			fmt.Printf("replayId %+v\n", mes)
			d.sendToGame <- mes
			go d.SendChannelDelSecond(m.ChannelID, "отправка повтора "+arg[1], 10)
			go d.DeleteMesageSecond(m.ChannelID, m.ID, 180)
			return true
		}
	}
	return false
}
func (d *Discord) historyWs(arg []string, m *discordgo.MessageCreate) bool {
	if arg[0] == "история" && arg[1] == "бз" {
		_, corporation := hades.HadesStorage.AllianceChat(m.ChannelID)
		mes := models.Message{
			Text:        "",
			Sender:      m.Author.Username,
			Avatar:      "",
			ChannelType: 0,
			Corporation: corporation.Corp,
			Command:     "история бз",
			Messager:    "ds",
		}
		fmt.Printf("historyWs %+v\n", mes)
		d.sendToGame <- mes
		go d.SendChannelDelSecond(m.ChannelID, "готовлю список  бз", 10)
		go d.DeleteMesageSecond(m.ChannelID, m.ID, 180)
		return true
	}
	return false
}
func (d *Discord) letInId(arg []string, m *discordgo.MessageCreate) bool {
	if arg[0] == "впустить" {
		match, _ := regexp.MatchString("^[0-9]+$", arg[1])
		if match {
			_, corporation := hades.HadesStorage.AllianceChat(m.ChannelID)
			mes := models.Message{
				Text:        arg[1],
				Sender:      m.Author.Username,
				Avatar:      "",
				ChannelType: 0,
				Corporation: corporation.Corp,
				Command:     "впустить",
				Messager:    "ds",
			}
			fmt.Printf("letInId %+v\n", mes)
			d.sendToGame <- mes
			go d.SendChannelDelSecond(m.ChannelID, "впустить отправленно  "+arg[1], 10)
			go d.DeleteMesageSecond(m.ChannelID, m.ID, 180)
			return true
		}
	}
	return false
}
