package discordClient

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (d *Discord) AccesChatDS(m *discordgo.MessageCreate) {
	res := strings.HasPrefix(m.Content, ".")
	if res == true && m.Content == ".add" {
		go d.DeleteMesageSecond(m.ChannelID, m.ID, 10)
		d.accessAddChannelDs(m.ChannelID, m.GuildID)
	} else if res == true && m.Content == ".del" {
		go d.DeleteMesageSecond(m.ChannelID, m.ID, 10)
		d.accessDelChannelDs(m.ChannelID, m.GuildID)
	}
}

func (d *Discord) accessAddChannelDs(chatid, guildid string) { // внесение в дб и добавление в масив
	ok, _ := d.CorpConfig.CheckChannelConfigDS(chatid)
	if ok {
		go d.SendChannelDelSecond(chatid, "Я уже могу работать на вашем канале\n"+
			"повторная активация не требуется.\nнапиши Справка", 30)
	} else {
		chatName := d.dsChatName(guildid)
		d.log.Println("новая активация корпорации ", chatName)
		d.dbase.CorpConfig.AddDsCorpConfig(chatName, chatid, guildid)
		go d.SendChannelDelSecond(chatid, "Спасибо за активацию.", 60)
		d.HelpChannelUpdate(chatid)
	}
}
func (d *Discord) accessDelChannelDs(chatid, guildid string) { //удаление с бд и масива для блокировки
	ok, _ := d.CorpConfig.CheckChannelConfigDS(chatid)
	if !ok {
		go d.SendChannelDelSecond(chatid, "ваш канал и так не подключен к логике бота ", 60)
	} else {
		d.dbase.CorpConfig.DeleteDsChannel(chatid)
		d.log.Println("отключение корпорации ", d.dsChatName(guildid))
		d.CorpConfig.ReloadConfig()
		d.dbase.CorpConfig.ReadBotCorpConfig()
		go d.SendChannelDelSecond(chatid, "вы отключили мои возможности", 60)
	}
}
