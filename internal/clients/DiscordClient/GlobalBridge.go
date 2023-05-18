package DiscordClient

import (
	"github.com/bwmarrin/discordgo"
	"kz_bot/internal/models"
	"kz_bot/internal/storage/memory"
	"regexp"
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
		if d.ifAsksForRoleRs(m) {
			go d.DeleteMessage(m.ChannelID, m.ID)
			return
		}
		if ifPrefix(m.Content) {
			return
		}
		username := m.Author.Username
		if m.Member != nil && m.Member.Nick != "" {
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
			if m.ReferencedMessage.Member != nil && m.ReferencedMessage.Member.Nick != "" {
				usernameR = m.ReferencedMessage.Member.Nick
			}
			mes.Ds.Reply.UserName = usernameR
			mes.Ds.Reply.Text = d.replaceTextMessage(m.ReferencedMessage.Content, m.GuildID)
			mes.Ds.Reply.Avatar = m.ReferencedMessage.Author.AvatarURL("128")
			mes.Ds.Reply.TimeMessage = m.ReferencedMessage.Timestamp
		}

		d.ChanGlobalChat <- mes
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
func (d *Discord) ifAsksForRoleRs(m *discordgo.MessageCreate) bool {
	var role bool
	after, found := strings.CutPrefix(m.Content, ".")
	if found {
		re := regexp.MustCompile(`^rs([4-9]|1[0-2])$`)
		after = re.ReplaceAllStringFunc(after, func(s string) string {
			r := d.Subscribe(m.Author.ID, s, m.GuildID)
			if r == 0 {
				role = true
				d.SendChannelDelSecond(m.ChannelID, m.Author.Mention()+" подписался на "+s, 10)
			} else if r == 1 {
				role = true
				d.SendChannelDelSecond(m.ChannelID, m.Author.Mention()+" уже подписан на "+s, 10)
			} else if r == 2 {
				d.SendChannelDelSecond(m.ChannelID, m.Author.Mention()+" ошибка выдачи роли "+s, 10)
				d.log.Println("error add globalRsRole ")
			}
			return after
		})
		if after == "список" || after == "Список" {
			var corps = "Список подключеных серверов"
			for _, global := range *memory.G {
				corps = corps + "\n" + global.CorpName
			}
			role = true
			d.SendChannelDelSecond(m.ChannelID, corps, 60)
		}
	}
	return role
}
func (d *Discord) deleteMessageGlobalChat(DelMessageId string) {
	command := models.InGlobalMessage{
		Content: DelMessageId,
		Tip:     "del",
	}
	d.ChanGlobalChat <- command
}
