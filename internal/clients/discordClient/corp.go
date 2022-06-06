package discordClient

import (
	"github.com/bwmarrin/discordgo"
	corpsConfig "kz_bot/internal/clients/corpConfig"
	"kz_bot/internal/dbase/dbaseMysql"
	"strings"
)

func (d *Ds) AccesChatDS(m *discordgo.MessageCreate) {
	res := strings.HasPrefix(m.Content, ".")
	if res == true && m.Content == ".add" {
		go d.DeleteMesageSecond(m.ChannelID, m.ID, 10)
		d.accessAddChannelDs(m.ChannelID, m.GuildID)
	} else if res == true && m.Content == ".del" {
		go d.DeleteMesageSecond(m.ChannelID, m.ID, 10)
		d.accessDelChannelDs(m.ChannelID)
	}
}

func (d *Ds) accessAddChannelDs(chatid, guildid string) { // внесение в дб и добавление в масив
	c := corpsConfig.CorpConfig{}
	ok, _ := c.CheckChannelConfigDS(chatid)
	if ok {
		go d.SendChannelDelSecond(chatid, "Я уже могу работать на вашем канале\n"+
			"повторная активация не требуется.\nнапиши Справка1", 30)
	} else {
		chatName := d.dsChatName(guildid)
		db := dbaseMysql.Db{}
		db.AddDsCorpConfig(chatName, chatid, guildid)
		go d.SendChannelDelSecond(chatid, "Спасибо за активацию.\nпиши Справка1", 60)
	}
}
func (d *Ds) accessDelChannelDs(chatid string) { //удаление с бд и масива для блокировки
	c := corpsConfig.CorpConfig{}
	ok, _ := c.CheckChannelConfigDS(chatid)
	if !ok {
		go d.SendChannelDelSecond(chatid, "ваш канал и так не подключен к логике бота ", 60)
	} else {
		db := dbaseMysql.Db{}
		db.DeleteDsChannel(chatid)
		c.ReloadConfig()
		db.ReadBotCorpConfig()
		go d.SendChannelDelSecond(chatid, "вы отключили мои возможности", 60)
	}
}
