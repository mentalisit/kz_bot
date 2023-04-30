package DiscordClient

import (
	"github.com/bwmarrin/discordgo"
	"kz_bot/internal/models"
	"kz_bot/internal/storage/memory"
)

func (d *Discord) logicMixGlobal(m *discordgo.MessageCreate) {
	ok, config := d.storage.CacheGlobal.CheckChannelConfigDS(m.ChannelID)
	if ok {
		if d.blackListFilter(m.Author.ID) {
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

	//if !blacklist m.Author.ID
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
