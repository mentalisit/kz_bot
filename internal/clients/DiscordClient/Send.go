package DiscordClient

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"kz_bot/internal/clients/DiscordClient/transmitter"
	"kz_bot/internal/models"
	"net/http"
	"path/filepath"
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
			d.log.Println("ошибка отправки сообщения SendChannelDelSecond "+chatid+text, err)
			d.log.Println("SendChannelDelSecond " + chatid + "  " + text)
			return
		}
		if second <= 60 {
			go func() {
				time.Sleep(time.Duration(second) * time.Second)
				_ = d.s.ChannelMessageDelete(chatid, message.ID)
			}()
		} else {
			d.storage.TimeDeleteMessage.TimerInsert(models.Timer{
				Dsmesid:  message.ID,
				Dschatid: chatid,
				Timed:    second,
			})
		}
	}
}
func (d *Discord) SendComplexContent(chatid, text string) (mesId string) { //отправка текста комплексного сообщения
	mesCompl, err := d.s.ChannelMessageSendComplex(chatid, &discordgo.MessageSend{
		Content: text})
	if err != nil {
		channel, _ := d.s.Channel(chatid)
		d.log.Println("Ошибка отправки комплексного сообщения text "+channel.Name+" ", err)
		mesCompl, err = d.s.ChannelMessageSendComplex(chatid, &discordgo.MessageSend{
			Content: text})
		if err == nil {
			return mesCompl.ID
		}
		return ""
	}
	return mesCompl.ID
}
func (d *Discord) SendComplex(chatid string, embeds discordgo.MessageEmbed) (mesId string) { //отправка текста комплексного сообщения
	mesCompl, err := d.s.ChannelMessageSendComplex(chatid, &discordgo.MessageSend{
		Content: mesContentNil,
		Embed:   &embeds,
	})
	if err != nil {
		channel, _ := d.s.Channel(chatid)
		d.log.Println("Ошибка отправки комплексного сообщения embed "+channel.Name+" ", err)
		mesCompl, err = d.s.ChannelMessageSendComplex(chatid, &discordgo.MessageSend{
			Content: mesContentNil,
			Embed:   &embeds,
		})
		if err == nil {
			return mesCompl.ID
		}
		return ""
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
func (d *Discord) SendFiles(channelID, fileURL string) (mesId string) {
	// Получаем содержимое файла из интернета
	resp, err := http.Get(fileURL)
	if err != nil {
		d.log.Println("Error downloading file:", err)
		return
	}
	defer resp.Body.Close()

	send, errs := d.s.ChannelFileSend(channelID, filepath.Base(fileURL), resp.Body)
	if errs != nil {
		d.log.Println("Error sending file:", errs)
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
	if text == "" {
		return ""
	}
	web := transmitter.New(d.s, guildId, "KzBot", true, d.log)
	pp := discordgo.WebhookParams{
		Content:   text,
		Username:  username,
		AvatarURL: Avatar,
	}
	mes, err := web.Send(chatid, &pp)
	if err != nil {
		fmt.Println(err)
		d.Send(chatid, text)
		return ""
	}
	return mes.ID
}

func (d *Discord) SendWebhookReply(text, username, chatid, guildId, Avatar string, replytext, replyAvatar, replyName string, replyTime time.Time) (mesId string) {
	if text == "" {
		return ""
	}
	web := transmitter.New(d.s, guildId, "KzBot", true, d.log)
	var embeds []*discordgo.MessageEmbed
	e := discordgo.MessageEmbed{
		Description: replytext,
		Timestamp:   replyTime.Format(time.RFC3339),
		Color:       14232643,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    replyName,
			IconURL: replyAvatar,
		},
	}

	embeds = append(embeds, &e)

	pp := &discordgo.WebhookParams{
		Content:   text,
		Username:  username,
		AvatarURL: Avatar,
		Embeds:    embeds,
	}
	mes, err := web.Send(chatid, pp)
	if err != nil {
		d.log.Println(err)
		d.Send(chatid, "ошибка отправки вебхука..недостаточно разрешений"+err.Error())
		return ""
	}
	return mes.ID
}
func (d *Discord) Name() {
	fmt.Println(d.s.State.User.Username)
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
		m, _ = d.s.ChannelMessageSend(chatid, text)
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
