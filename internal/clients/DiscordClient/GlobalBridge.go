package DiscordClient

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"kz_bot/internal/models"
	"kz_bot/internal/storage/memory"
)

func (d *Discord) logicMixGlobal(m *discordgo.MessageCreate) {
	ok, config := d.storage.CacheGlobal.CheckChannelConfigDS(m.ChannelID)
	if ok {
		if d.blackListFilter(m.Author.ID) {
			d.DeleteMesageSecond(m.ChannelID, m.ID, 5)
			return
		}
		username := m.Author.Username
		if m.Member.Nick != "" {
			username = m.Member.Nick
		}
		if len(m.Attachments) > 0 {
			for _, attach := range m.Attachments { //вложеные файлы
				m.Content = m.Content + "\n" + attach.URL
			}
		}
		fmt.Printf("	logicMixGlobal MentionEveryone %+v\n", m.MentionEveryone)
		fmt.Printf("	logicMixGlobal MentionRoles %+v\n", m.MentionRoles)
		fmt.Printf("	logicMixGlobal MentionChannels %+v\n", m.MentionChannels)
		fmt.Printf("	logicMixGlobal Mentions %+v\n", m.Mentions)
		//fmt.Printf("	logicMixGlobal ContentWithMentionsReplaced() %+v\n", m.ContentWithMentionsReplaced())
		fmt.Printf("	logicMixGlobal Content %s\n", m.Content)
		fmt.Printf("	logicMixGlobal m.Type %+v\n", m.Type)
		fmt.Printf("	logicMixGlobal m.Message.Type %+v\n", m.Message.Type)

		mes := models.InGlobalMessage{
			Content: d.replaceTextMessage(m.Content, m.GuildID),
			Tip:     "ds",
			Name:    username,
			Ds: struct {
				Mesid   string
				Nameid  string
				Guildid string
				Avatar  string
				ChatId  string
			}{
				Mesid:   m.ID,
				Nameid:  m.Author.ID,
				Guildid: m.GuildID,
				Avatar:  m.Author.AvatarURL("128"),
				ChatId:  m.ChannelID,
			},
			Config: config,
		}
		d.globalChat <- mes
	}

	//text:= cenzura m.Content

}
func (d *Discord) blackListFilter(userid string) bool {
	for _, s := range memory.BlackListNamesId {
		if s == userid {
			return true
		}
	}
	return false
}
