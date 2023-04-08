package DiscordClient

import (
	"context"
	"github.com/bwmarrin/discordgo"
	"strings"
	"time"
)

//nujno sdelat lang

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
	ok, _ := d.storage.Cache.CheckChannelConfigDS(chatid)
	if ok {
		go d.SendChannelDelSecond(chatid, "Я уже могу работать на вашем канале\n"+
			"повторная активация не требуется.\nнапиши Справка", 30)
	} else {
		chatName := d.dsChatName(chatid, guildid)
		d.log.Println("новая активация корпорации ", chatName)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		d.storage.CorpsConfig.AddDsCorpConfig(ctx, chatName, chatid, guildid)
		go d.SendChannelDelSecond(chatid, "Спасибо за активацию.", 10)
		d.HelpChannelUpdate(chatid)
	}
}
func (d *Discord) accessDelChannelDs(chatid, guildid string) { //удаление с бд и масива для блокировки
	ok, config := d.storage.Cache.CheckChannelConfigDS(chatid)
	d.DeleteMessage(chatid, config.MesidDsHelp)
	if !ok {
		go d.SendChannelDelSecond(chatid, "ваш канал и так не подключен к логике бота ", 60)
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		d.storage.CorpsConfig.DeleteDs(ctx, chatid)
		d.log.Println("отключение корпорации ", d.dsChatName(chatid, guildid))
		d.storage.Cache.ReloadConfig()
		d.storage.CorpsConfig.ReadCorps()
		go d.SendChannelDelSecond(chatid, "вы отключили мои возможности", 60)
	}
}
