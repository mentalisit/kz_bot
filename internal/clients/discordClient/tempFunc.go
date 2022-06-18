package discordClient

import (
	"fmt"
	"regexp"

	"github.com/bwmarrin/discordgo"
)

var (
	// See https://discordapp.com/developers/docs/reference#message-formatting.
	channelMentionRE = regexp.MustCompile("<#[0-9]+>")
	userMentionRE    = regexp.MustCompile("@[^@\n]{1,32}")
	emoteRE          = regexp.MustCompile(`<a?(:\w+:)\d+>`)
)

func (d *Ds) replaceChannelMentions(text string, guildID string) string {
	replaceChannelMentionFunc := func(match string) string {
		channelID := match[2 : len(match)-1]
		channel := d.getChannel(channelID, guildID)

		if channel.Name == "" {
			return "#unknownchannel"
		}
		return "#" + channel.Name
	}
	return channelMentionRE.ReplaceAllStringFunc(text, replaceChannelMentionFunc)
}

func (d *Ds) getChannel(channelId, guildId string) *discordgo.Channel {
	g, err := d.d.Guild(guildId)
	if err != nil {
		fmt.Println(err)
	}
	c := g.Channels
	for _, i := range c {
		if i.ID == channelId {
			return i
		}

	}
	return nil
}

/*
func replaceUserMentions(text string) string {
	replaceUserMentionFunc := func(match string) string {
		var (
			err      error
			member   *discordgo.Member
			username string
		)

		usernames := enumerateUsernames(match[1:])
		for _, username = range usernames {
			b.Log.Debugf("Testing mention: '%s'", username)
			member, err = b.getGuildMemberByNick(username)
			if err == nil {
				break
			}
		}
		if member == nil {
			return match
		}
		return strings.Replace(match, "@"+username, member.User.Mention(), 1)
	}
	return userMentionRE.ReplaceAllStringFunc(text, replaceUserMentionFunc)
}

func replaceEmotes(text string) string {
	return emoteRE.ReplaceAllString(text, "$1")
}

func replaceAction(text string) (string, bool) {
	length := len(text)
	if length > 1 && text[0] == '_' && text[length-1] == '_' {
		return text[1 : length-1], true
	}
	return text, false
}


*/
