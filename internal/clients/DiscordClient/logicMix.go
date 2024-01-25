package DiscordClient

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"kz_bot/internal/models"
	"net/url"
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
		d.log.Error(err.Error())
	}
	if user.ID != message.Author.ID {
		ok, config := d.CheckChannelConfigDS(r.ChannelID)
		if ok {
			in := models.InMessage{
				Tip:         "ds",
				Name:        user.Username,
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
					Avatar:  user.AvatarURL("128"),
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
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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
	if d.ifMentionBot(m) {
		return
	}
	if d.avatar(m) {
		return
	}
	go d.latinOrNot(m) //пытаемся переводить гостевой чат
	d.AccesChatDS(m)

	//filter Rs
	ok, config := d.CheckChannelConfigDS(m.ChannelID)
	if ok {
		d.SendToRsFilter(m, config)
		return
	}

	//bridge
	ds, bridgeConfig := d.BridgeCheckChannelConfigDS(m.ChannelID)
	if ds || strings.HasPrefix(m.Content, ".") {
		go d.SendToBridgeChatFilter(m, bridgeConfig)
	}

}

func (d *Discord) SendToRsFilter(m *discordgo.MessageCreate, config models.CorporationConfig) {
	//if m.Content == "registerCommandModules" {
	//	// Регистрация слеш-команды при получении команды
	//	d.registerCommand(d.s, m.GuildID)
	//}
	in := models.InMessage{
		Mtext:       m.Content,
		Tip:         "ds",
		Name:        m.Author.Username,
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

func (d *Discord) SendToBridgeChatFilter(m *discordgo.MessageCreate, config models.BridgeConfig) {
	mes := models.BridgeMessage{
		Text:          d.replaceTextMessage(m.Content, m.GuildID),
		Sender:        d.getAuthorName(m),
		Tip:           "ds",
		Avatar:        m.Author.AvatarURL("128"),
		ChatId:        m.ChannelID,
		MesId:         m.ID,
		GuildId:       m.GuildID,
		TimestampUnix: m.Timestamp.Unix(),
		Config:        &config,
	}
	if len(m.StickerItems) > 0 {
		mes.Text = fmt.Sprintf("https://cdn.discordapp.com/stickers/%s.png", m.Message.StickerItems[0].ID)
	}

	if m.ReferencedMessage != nil {
		usernameR := m.ReferencedMessage.Author.String() //.Username
		if m.ReferencedMessage.Member != nil && m.ReferencedMessage.Member.Nick != "" {
			usernameR = m.ReferencedMessage.Member.Nick
		}
		mes.Reply = &models.BridgeMessageReply{
			TimeMessage: m.ReferencedMessage.Timestamp.Unix(),
			Text:        d.replaceTextMessage(m.ReferencedMessage.Content, m.GuildID),
			Avatar:      m.ReferencedMessage.Author.AvatarURL("128"),
			UserName:    usernameR,
		}
	}
	if len(m.Attachments) > 0 {
		for _, a := range m.Attachments {
			parsedURL, err := url.Parse(a.URL)
			if err != nil {
				d.log.Error(err.Error())
			}
			parsedURL.RawQuery = ""
			parsedURL.Fragment = ""
			mes.FileUrl = parsedURL.String()
			d.ChanBridgeMessage <- mes
		}
	} else {
		d.ChanBridgeMessage <- mes
	}
}

func (d *Discord) readReactionTranslate(r *discordgo.MessageReactionAdd, m *discordgo.Message) {
	user, err := d.s.User(r.UserID)
	if err != nil {
		d.log.Error(err.Error())
	}
	if user.ID != m.Author.ID {

		switch r.Emoji.Name {
		case "🇺🇸":
			d.transtale(m, "en")
		case "🇷🇺":
			d.transtale(m, "ru")
		case "🇺🇦":
			d.transtale(m, "ua")
		case "🇬🇧":
			d.transtale(m, "en")
		case "🇧🇾":
			d.transtale(m, "be")
		case "🇩🇪":
			d.transtale(m, "de")
		case "🇵🇱":
			d.transtale(m, "pl")
		}
	}
}
