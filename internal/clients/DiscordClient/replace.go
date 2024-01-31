package DiscordClient

import (
	"github.com/bwmarrin/discordgo"
	"regexp"
	"strings"
)

var (
	channelMentionRE = regexp.MustCompile("<#[0-9]+>")
	userMentionRE    = regexp.MustCompile("<@(\\d+)>")
	roleMentionRE    = regexp.MustCompile("<@&(\\d+)>")
)

func (d *Discord) replaceTextMessage(text string, guildid string) (newtext string) {
	newtext = d.replaceChannelMentions(text, guildid)
	newtext = d.replaceUserMentions(newtext, guildid)
	newtext = d.replaceRoleMentions(newtext, guildid)
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
func (d *Discord) replaceRoleMentions(text string, guildid string) string {
	mentionIds := roleMentionRE.FindAllStringSubmatch(text, -1)
	for _, match := range mentionIds {
		mention := match[0]
		roleId := match[1]
		role := d.getRoleById(roleId, guildid)
		if role != nil {
			text = strings.Replace(text, mention, "@&"+role.Name, 1)
		}
	}
	return text
}
func (d *Discord) getRoleById(roleId string, guildId string) *discordgo.Role {
	roles, _ := d.s.GuildRoles(guildId)

	for _, role := range roles {
		if role.ID == roleId {
			return role
		}
	}
	return nil
}
func (d *Discord) replaceUserMentions(text string, guildid string) string {
	mentionIds := userMentionRE.FindAllStringSubmatch(text, -1)
	for _, match := range mentionIds {
		mention := match[0]
		userId := match[1]
		username := d.getUserNameById(userId, guildid)
		text = strings.Replace(text, mention, "@"+username, 1)
	}
	return text
}
func (d *Discord) getUserNameById(userId string, guildId string) string {
	members, err := d.s.GuildMembers(guildId, "", 999)
	if err != nil {
		d.log.ErrorErr(err)
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
