package DiscordClient

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"kz_bot/internal/clients/DiscordClient/transmitter"
	"kz_bot/internal/models"
	"kz_bot/internal/storage/CorpsConfig/hades"
)

const text = "text"

func (d *Discord) ifMessageForHades(m *discordgo.MessageCreate) bool {
	if d.ifComands(m) {
		return true
	}

	okA, corp := hades.HadesStorage.AllianceChat(m.ChannelID)
	if okA {
		d.sendToG(m, corp, 0)
	}
	okw1, corp := hades.HadesStorage.Ws1Chat(m.ChannelID)
	if okw1 {
		d.sendToG(m, corp, 1)
	}
	okw2, corp := hades.HadesStorage.Ws1Chat(m.ChannelID)
	if okw2 {
		d.sendToG(m, corp, 2)
	}
	return false
}
func (d *Discord) sendToG(m *discordgo.MessageCreate, corp models.Corporation, channelType int) {
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
	mes := models.Message{
		Text:        newText,
		Sender:      name,
		Avatar:      m.Author.AvatarURL("128"),
		ChannelType: channelType, //0alliancechat
		Corporation: corp.Corp,
		Command:     text,
		Messager:    "ds",
	}
	d.sendToGame <- mes
}

func (d *Discord) SendWebhookForHades(text, username, chatid, guildId, Avatar string) string {
	if text == "" {
		return ""
	}
	web := transmitter.New(d.s, guildId, "KzBot", true, d.log)
	pp := discordgo.WebhookParams{
		Content:   text,
		Username:  username,
		AvatarURL: Avatar,
	}
	m, err := web.Send(chatid, &pp)
	if err != nil {
		//d.log.Println("error create webhook message  " + err.Error())
		m, _ = d.s.ChannelMessageSend(chatid, "ошибка отправки вебхука..недостаточно разрешений")
		return m.ID
	}
	return m.ID
}
func (d *Discord) EditWebhookForHades(text, username, chatid, guildId, Avatar, mesid string) {
	if text == "" {
		return
	}
	web := transmitter.New(d.s, guildId, "KzBot", true, d.log)
	pp := discordgo.WebhookParams{
		Content:   text,
		Username:  username,
		AvatarURL: Avatar,
	}
	err := web.Edit(chatid, mesid, &pp)
	if err != nil {
		//d.log.Println("error create webhook message  " + err.Error())
		//_, _ = d.s.ChannelMessageSend(chatid, "ошибка edit вебхука..недостаточно разрешений")
		return
	}
	return
}
