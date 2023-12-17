package DiscordClient

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"regexp"
	"strings"
)

//nujno sdelat lang

func (d *Discord) AccesChatDS(m *discordgo.MessageCreate) {
	res := strings.HasPrefix(m.Content, ".")
	if res {
		switch m.Content {
		case ".add":
			go d.DeleteMesageSecond(m.ChannelID, m.ID, 10)
			d.accessAddChannelDs(m.ChannelID, m.GuildID, "en")
		case ".добавить":
			go d.DeleteMesageSecond(m.ChannelID, m.ID, 10)
			d.accessAddChannelDs(m.ChannelID, m.GuildID, "ru")
		case ".добавитьт":
			go d.DeleteMesageSecond(m.ChannelID, m.ID, 10)
			d.accessAddChannelDs(m.ChannelID, m.GuildID, "dru")
		case ".додати":
			go d.DeleteMesageSecond(m.ChannelID, m.ID, 10)
			d.accessAddChannelDs(m.ChannelID, m.GuildID, "ua")
		case ".del":
			go d.DeleteMesageSecond(m.ChannelID, m.ID, 10)
			d.accessDelChannelDs(m.ChannelID, m.GuildID)
		case ".удалить":
			go d.DeleteMesageSecond(m.ChannelID, m.ID, 10)
			d.accessDelChannelDs(m.ChannelID, m.GuildID)
		case ".видалити":
			go d.DeleteMesageSecond(m.ChannelID, m.ID, 10)
			d.accessDelChannelDs(m.ChannelID, m.GuildID)

		default:
			if d.CleanOldMessage(m) {
				return
			}
		}
	}
}

func (d *Discord) accessAddChannelDs(chatid, guildid, lang string) { // внесение в дб и добавление в масив
	ok, _ := d.CheckChannelConfigDS(chatid)
	if ok {
		go d.SendChannelDelSecond(chatid, d.storage.Words.GetWords(lang, "accessAlready"), 30)
	} else {
		chatName := d.GuildChatName(chatid, guildid)
		d.log.Println("новая активация корпорации ", chatName)
		d.AddDsCorpConfig(chatName, chatid, guildid, lang)
		go d.SendChannelDelSecond(chatid, d.storage.Words.GetWords(lang, "accessTY"), 10)

	}
}
func (d *Discord) accessDelChannelDs(chatid, guildid string) { //удаление с бд и масива для блокировки
	ok, config := d.CheckChannelConfigDS(chatid)
	d.DeleteMessage(chatid, config.MesidDsHelp)
	if !ok {
		go d.SendChannelDelSecond(chatid, d.storage.Words.GetWords("ru", "accessYourChannel"), 60)
	} else {
		d.SendChannelDelSecond(chatid, d.getLang(chatid, "YouDisabledMyFeatures"), 60)
		d.storage.ConfigRs.DeleteConfigRs(config)
		d.storage.ReloadDbArray()
		d.log.Println("отключение корпорации ", d.GuildChatName(chatid, guildid))
	}
}

func (d *Discord) CleanOldMessage(m *discordgo.MessageCreate) bool {
	re := regexp.MustCompile(`^\.очистка (\d{1,2}|100)`)
	matches := re.FindStringSubmatch(m.Content)
	if len(matches) > 0 {
		fmt.Println("limitMessage " + matches[1])
		d.CleanOldMessageChannel(m.ChannelID, matches[1])
		return true
	}
	return false
}
