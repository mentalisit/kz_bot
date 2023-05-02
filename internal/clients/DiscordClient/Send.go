package DiscordClient

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io"
	"kz_bot/internal/clients/DiscordClient/transmitter"
	"kz_bot/internal/models"
	"time"
)

var mesContentNil string

func (d *Discord) SendEmbedText(chatid, title, text string) *discordgo.Message {
	Emb := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       16711680,
		Description: text,
		Title:       title,
	}
	m, err := d.s.ChannelMessageSendEmbed(chatid, Emb)
	if err != nil {
		d.log.Println("Ошибка отправки сообщения со вставкой ", err)
	}
	return m
}
func (d *Discord) SendChannelDelSecond(chatid, text string, second int) {
	if text != "" {
		message, err := d.s.ChannelMessageSend(chatid, text)
		if err != nil {
			d.log.Println("ошибка отправки сообщения SendChannelDelSecond", err)
		}
		if second <= 60 {
			go func() {
				time.Sleep(time.Duration(second) * time.Second)
				_ = d.s.ChannelMessageDelete(chatid, message.ID)
			}()
		} else {
			ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
			defer cancel()
			d.storage.Timers.TimerInsert(ctx, message.ID, chatid, 0, 0, second)
		}
	}
}
func (d *Discord) SendComplexContent(chatid, text string) (mesId string) { //отправка текста комплексного сообщения
	mesCompl, err := d.s.ChannelMessageSendComplex(chatid, &discordgo.MessageSend{
		Content: text})
	if err != nil {
		d.log.Println("Ошибка отправки комплексного сообщения ", err)
	}
	return mesCompl.ID
}
func (d *Discord) SendComplex(chatid string, embeds discordgo.MessageEmbed) (mesId string) { //отправка текста комплексного сообщения
	mesCompl, err := d.s.ChannelMessageSendComplex(chatid, &discordgo.MessageSend{
		Content: mesContentNil,
		Embed:   &embeds,
	})
	if err != nil {
		d.log.Println("Ошибка отправки комплексного сообщения ", err)
	}
	return mesCompl.ID
}
func (d *Discord) Send(chatid, text string) (mesId string) { //отправка текста
	message, err := d.s.ChannelMessageSend(chatid, text)
	if err != nil {
		d.log.Println("ошибка отправки текста ", err)
	}
	return message.ID
}
func (d *Discord) SendFiles(chatid, fileName string, r io.Reader) (mesId string) {
	send, err := d.s.ChannelFileSend(chatid, fileName, r)
	if err != nil {
		return ""
	}
	return send.ID
}
func (d *Discord) SendEmbedTime(chatid, text string) (mesId string) { //отправка текста с двумя реакциями
	message, err := d.s.ChannelMessageSend(chatid, text)
	if err != nil {
		d.log.Println("ошибка отправки текста ", err)
	}
	err = d.s.MessageReactionAdd(chatid, message.ID, emPlus)
	if err != nil {
		d.log.Println("Ошибка добавления эмоджи ", emPlus, err)
	}
	err = d.s.MessageReactionAdd(chatid, message.ID, emMinus)
	if err != nil {
		d.log.Println("Ошибка добавления эмоджи ", emMinus, err)
	}
	return message.ID
}
func (d *Discord) SendWebhook(text, username, chatid, guildId, Avatar string) (mesId string) {
	web := transmitter.New(d.s, guildId, "KzBot", true, d.log)
	pp := discordgo.WebhookParams{
		Content:   text,
		Username:  username,
		AvatarURL: Avatar,
	}
	mes, err := web.Send(chatid, &pp)
	if err != nil {
		fmt.Println(err)
		d.Send(chatid, "ошибка отправки вебхука..недостаточно разрешений")
		return ""
	}
	return mes.ID
}

func (d *Discord) SendWebhookReply(m models.InGlobalMessage) (mesId string) {
	web := transmitter.New(d.s, m.Ds.GuildId, "KzBot", true, d.log)
	var embeds []*discordgo.MessageEmbed
	e := discordgo.MessageEmbed{
		Description: m.Ds.Reply.Text,
		Timestamp:   m.Ds.Reply.TimeMessage.Format(time.RFC3339),
		Color:       14232643,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    m.Ds.Reply.UserName,
			IconURL: m.Ds.Reply.Avatar,
		},
	}

	embeds = append(embeds, &e)

	pp := &discordgo.WebhookParams{
		Content:   m.Content,
		Username:  m.Name,
		AvatarURL: m.Ds.Avatar,
		Embeds:    embeds,
	}
	mes, err := web.Send(m.Ds.ChatId, pp)
	if err != nil {
		d.log.Println(err)
		d.Send(m.Ds.ChatId, "ошибка отправки вебхука..недостаточно разрешений"+err.Error())
		return ""
	}
	return mes.ID
}
