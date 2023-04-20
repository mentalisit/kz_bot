package DiscordClient

import (
	"regexp"
	"strings"
	"unicode"
)

var (
	channelMentionRE = regexp.MustCompile("<#[0-9]+>")
	userMentionRE    = regexp.MustCompile("<@(\\d+)>")
	//emoteRE          = regexp.MustCompile(`<a?(:\w+:)\d+>`)
)

func (d *Discord) replaceTextMessage(text string, guildid string) (newtext string) {
	newtext = d.replaceChannelMentions(text, guildid)
	newtext = d.replaceUserMentions(newtext, guildid)
	return newtext
}

func (d *Discord) replaceChannelMentions(text string, guildid string) string {
	replaceChannelMentionFunc := func(match string) string {
		channelID := match[2 : len(match)-1]
		channels, err := d.s.GuildChannels(guildid)
		if err != nil {
			return "#unknownchannel"
		}
		channelName := ""
		for _, channel := range channels {
			if channelID == channel.ID {
				channelName = channel.Name
			}
		}
		return "#" + channelName
	}
	return channelMentionRE.ReplaceAllStringFunc(text, replaceChannelMentionFunc)
}

func (d *Discord) replaceUserMentions(text string, guildid string) string {
	mentionIds := userMentionRE.FindAllStringSubmatch(text, -1)
	for _, match := range mentionIds {
		mention := match[0]
		userId := match[1]
		username := d.getUserById(userId, guildid)
		text = strings.Replace(text, mention, "@"+username, 1)
	}
	return text
}
func (d *Discord) getUserById(userId string, guildId string) string {
	members, err := d.s.GuildMembers(guildId, "", 999)
	if err != nil {
		d.log.Println("error getGuildMember " + err.Error())
	}
	for _, member := range members {
		if member.User.ID == userId {
			if member.Nick != "" {
				return member.Nick
			} else {
				return member.User.Username
			}
		}
	}
	return "Unknown user"
}

func enumerateUsernames(s string) []string {
	onlySpace := true
	for _, r := range s {
		if !unicode.IsSpace(r) {
			onlySpace = false
			break
		}
	}
	if onlySpace {
		return nil
	}

	var username, endSpace string
	var usernames []string
	skippingSpace := true
	for _, r := range s {
		if unicode.IsSpace(r) {
			if !skippingSpace {
				usernames = append(usernames, username)
				skippingSpace = true
			}
			endSpace += string(r)
			username += string(r)
		} else {
			endSpace = ""
			username += string(r)
			skippingSpace = false
		}
	}
	if endSpace == "" {
		usernames = append(usernames, username)
	}
	return usernames
}
