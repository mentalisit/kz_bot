package DiscordClient

import (
	"github.com/bwmarrin/discordgo"
	"kz_bot/internal/models"
	"kz_bot/internal/storage/memory"
	"strings"
	"time"
)

func (d *Discord) logicMixGlobal(m *discordgo.MessageCreate) {
	ok, config := d.storage.CacheGlobal.CheckChannelConfigDS(m.ChannelID)
	if ok {
		if d.blackListFilter(m.Author.ID) {
			d.DeleteMesageSecond(m.ChannelID, m.ID, 5)
			return
		}
		if ifPrefix(m.Content) {
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
			Ds: models.InGlobalMessageDs{
				MesId:         m.ID,
				NameId:        m.Author.ID,
				ChatId:        m.ChannelID,
				GuildId:       m.GuildID,
				Avatar:        m.Author.AvatarURL("128"),
				TimestampUnix: m.Timestamp.Unix(),
				Reply: struct {
					TimeMessage time.Time
					Text        string
					Avatar      string
					UserName    string
				}{},
			},
			Config: config,
		}
		if m.MessageReference != nil {
			usernameR := m.ReferencedMessage.Author.Username
			if m.ReferencedMessage.Member != nil {
				usernameR = m.Member.Nick
			}
			mes.Ds.Reply.UserName = usernameR
			mes.Ds.Reply.Text = m.ReferencedMessage.Content
			mes.Ds.Reply.Avatar = m.ReferencedMessage.Author.AvatarURL("128")
			mes.Ds.Reply.TimeMessage = m.ReferencedMessage.Timestamp
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
func ifPrefix(s string) (prefixBool bool) {
	prefix := []string{".", "!", "%", "-"}
	for _, p := range prefix {
		if strings.HasPrefix(s, p) {
			prefixBool = true
			break
		}
	}
	return prefixBool
}
func (d *Discord) deleteMessageGlobalChat(DelMessageId string) {
	command := models.InGlobalMessage{
		Content: DelMessageId,
		Tip:     "del",
	}
	d.globalChat <- command
}
