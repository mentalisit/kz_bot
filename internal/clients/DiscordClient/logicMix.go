package DiscordClient

import (
	"context"
	"fmt"
	gt "github.com/bas24/googletranslatefree"
	"github.com/bwmarrin/discordgo"
	"kz_bot/internal/models"
	"strings"
	"time"
)

const (
	emOK      = "✅"
	emCancel  = "❎"
	emRsStart = "🚀"
	emPl30    = "⌛"
	emPlus    = "➕"
	emMinus   = "➖"
)

func (d *Discord) readReactionQueue(r *discordgo.MessageReactionAdd, message *discordgo.Message) {
	user, err := d.s.User(r.UserID)
	if err != nil {
		d.log.Println("Ошибка получения Юзера по реакции ", err)
	}
	if user.ID != message.Author.ID {
		ok, config := d.CheckChannelConfigDS(r.ChannelID)
		if ok {
			member, e := d.s.GuildMember(config.Guildid, user.ID)
			if e != nil {
				d.log.Println("Oшибка получения участника ", e)
			}
			name := user.Username
			if member.Nick != "" {
				name = member.Nick
			}
			Avatar := "https://cdn.discordapp.com/avatars/" + user.ID + "/" + user.Avatar + ".jpg"

			in := models.InMessage{
				Mtext:       "",
				Tip:         "ds",
				Name:        name,
				NameMention: user.Mention(),
				Ds: struct {
					Mesid   string
					Nameid  string
					Guildid string
					Avatar  string
				}{
					Mesid:   r.MessageID,
					Nameid:  user.ID,
					Guildid: config.Guildid,
					Avatar:  Avatar,
				},

				Config: config,
				Option: models.Option{
					Reaction: true},
			}
			d.reactionUserRemove(r)

			if r.Emoji.Name == emPlus {
				in.Mtext = "+"
			} else if r.Emoji.Name == emMinus {
				in.Mtext = "-"
			} else if r.Emoji.Name == emOK || r.Emoji.Name == emCancel || r.Emoji.Name == emRsStart || r.Emoji.Name == emPl30 {
				ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
				defer cancel()
				in.Lvlkz, err = d.storage.DbFunc.ReadMesIdDS(ctx, r.MessageID)
				if err == nil && in.Lvlkz != "" {
					if r.Emoji.Name == emOK {
						in.Timekz = "30"
						in.Mtext = in.Lvlkz + "+"
					} else if r.Emoji.Name == emCancel {
						in.Mtext = in.Lvlkz + "-"
					} else if r.Emoji.Name == emRsStart {
						in.Mtext = in.Lvlkz + "++"
					} else if r.Emoji.Name == emPl30 {
						in.Mtext = in.Lvlkz + "+++"
					}
				}
			}
			d.ChanRsMessage <- in
		}
	}
}

func (d *Discord) reactionUserRemove(r *discordgo.MessageReactionAdd) {
	err := d.s.MessageReactionRemove(r.ChannelID, r.MessageID, r.Emoji.Name, r.UserID)
	if err != nil {
		fmt.Println("Ошибка удаления эмоджи", err)
	}
}

func (d *Discord) logicMix(m *discordgo.MessageCreate) {
	d.TranslateBelRus(m)
	if d.ifMentionBot(m) {
		return
	}
	if d.avatar(m) {
		return
	}

	d.AccesChatDS(m)

	//filter Rs
	ok, config := d.CheckChannelConfigDS(m.ChannelID)
	if ok {
		d.SendToRsFilter(m, config)
		return
	}
	//filter hs
	corpAlliance := d.getCorpHadesAlliance(m.ChannelID)
	if corpAlliance.Corp != "" {
		d.sendToFilterHades(m, corpAlliance, 0)
		return
	}
	corpWs1 := d.getCorpHadesWs1(m.ChannelID)
	if corpWs1.Corp != "" {
		d.sendToFilterHades(m, corpWs1, 1)
		return
	}

	//bridge
	ds, bridgeConfig := d.BridgeCheckChannelConfigDS(m.ChannelID)
	if ds || strings.HasPrefix(m.Content, ".") {
		go d.SendToBridgeChatFilter(m, bridgeConfig)
	}

}

func (d *Discord) sendToFilterHades(m *discordgo.MessageCreate, corp models.CorporationHadesClient, channelType int) {
	if len(m.Attachments) > 0 {
		for _, attach := range m.Attachments { //вложеные файлы
			m.Content = m.Content + "\n" + attach.URL
		}
	}
	if m.Content == "" || m.Message.EditedTimestamp != nil {
		return
	}
	name := m.Author.Username
	member, e := d.s.GuildMember(m.GuildID, m.Author.ID) //проверка есть ли изменения имени в этом дискорде
	if e != nil {
		fmt.Println("Ошибка получения ника пользователя", e, m.ID)
	} else if member != nil {
		if member.Nick != "" {
			name = member.Nick
		}
	}

	newText := d.replaceTextMessage(m.Content, m.GuildID)
	mes := models.MessageHades{
		Text:        newText,
		Sender:      name,
		Avatar:      m.Author.AvatarURL("128"),
		ChannelType: channelType, //0 AllianceChat
		Corporation: corp.Corp,
		Command:     "text",
		Messager:    "ds",
		Ds: models.MessageHadesDs{
			MessageId: m.ID,
		},
	}

	d.ChanToGame <- mes

}
func (d *Discord) SendToRsFilter(m *discordgo.MessageCreate, config models.CorporationConfig) {

	if len(m.Attachments) > 0 {
		for _, attach := range m.Attachments { //вложеные файлы
			m.Content = m.Content + "\n" + attach.URL
		}
	}
	member, e := d.s.GuildMember(m.GuildID, m.Author.ID) //проверка есть ли изменения имени в этом дискорде
	if e != nil {
		d.log.Println("Ошибка получения ника пользователя", e, m.ID)
	}
	name := m.Author.Username
	if member.Nick != "" {
		name = member.Nick
	}

	in := models.InMessage{
		Mtext:       m.Content,
		Tip:         "ds",
		Name:        name,
		NameMention: m.Author.Mention(),
		Ds: struct {
			Mesid   string
			Nameid  string
			Guildid string
			Avatar  string
		}{
			Mesid:   m.ID,
			Nameid:  m.Author.ID,
			Guildid: m.GuildID,
			Avatar:  m.Author.AvatarURL("128"),
		},
		Config: config,
		Option: models.Option{InClient: true},
	}
	d.ChanRsMessage <- in

}
func (d *Discord) ifMentionBot(m *discordgo.MessageCreate) bool {
	after, found := strings.CutPrefix(m.Content, d.s.State.User.Mention())
	if found {
		if len(after) > 0 {
			split := strings.Split(after, " ")
			if split[0] == "help" || split[0] == "справка" || split[0] == "довідка" {
				//nujno sdelat obshuu spravku
				d.SendChannelDelSecond(m.ChannelID, "сорян в разработке", 10)
				return true
			}
		}

		d.DeleteMesageSecond(m.ChannelID, m.ID, 30)
		goodRs, _ := d.CheckChannelConfigDS(m.ChannelID)
		//okAlliance, corp := hades.HadesStorage.AllianceChat(m.ChannelID)
		//okWs1, corpw := hades.HadesStorage.Ws1Chat(m.ChannelID)
		var text string
		if goodRs {
			text = fmt.Sprintf("%s че пингуешь? пиши Справка,или пиши создателю бота @Mentalisit#5159 ", m.Author.Mention())
			//} else if okAlliance {
			//	text = fmt.Sprintf("%s не балуйся бот занят пересылкой сообщений в игру в корпорацию %s", m.Author.Mention(), corp.Corp)
			//} else if okWs1 {
			//	text = fmt.Sprintf("%s не балуйся бот занят пересылкой сообщений в игру в корпорацию %s", m.Author.Mention(), corpw.Corp)
		} else {
			text = fmt.Sprintf("%s че пингуешь? я же многофункциональный бот, Префикс доступен только после активации нужного режима \n Для получения справки пиши %s help",
				m.Author.Mention(), d.s.State.User.Mention())
		}
		d.SendChannelDelSecond(m.ChannelID, text, 30)
	}
	return found
}
func (d *Discord) deleteMessageBridgeChat(DelMessageId string) {
	d.ChanBridgeMessage <- models.BridgeMessage{
		Tip: "del",
		Ds: models.BridgeMessageDs{
			MesId: DelMessageId,
		},
	}
}
func (d *Discord) SendToBridgeChatFilter(m *discordgo.MessageCreate, config models.BridgeConfig) {
	username := m.Author.Username
	if m.Member != nil && m.Member.Nick != "" {
		username = m.Member.Nick
	}
	if len(m.Attachments) > 0 {
		for _, attach := range m.Attachments { //вложеные файлы
			m.Content = m.Content + "\n" + attach.URL
		}
	}
	mes := models.BridgeMessage{
		Text:   d.replaceTextMessage(m.Content, m.GuildID),
		Sender: username,
		Tip:    "ds",
		Ds: models.BridgeMessageDs{
			ChatId:        m.ChannelID,
			MesId:         m.ID,
			Avatar:        m.Author.AvatarURL("128"),
			GuildId:       m.GuildID,
			TimestampUnix: m.Timestamp.Unix(),
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

	d.ChanBridgeMessage <- mes
}

var chatidTr string

func (d *Discord) TranslateBelRus(m *discordgo.MessageCreate) {
	if m.Content == "trbr" {
		chatidTr = m.ChannelID
	}
	if m.ChannelID == chatidTr {
		result, _ := gt.Translate(m.Content, "be", "ru")
		if result != m.Content {
			d.SendWebhook(result, m.Author.Username, chatidTr, m.GuildID, m.Author.AvatarURL("128"))
		}
	}
}
