package BridgeChat

import (
	"fmt"
	"kz_bot/internal/models"
	"strings"
	"sync"
)

func (b *Bridge) logicMessage() {
	if b.checkingForIdenticalMessage() {
		return
	}
	if b.in.Tip == "delDs" {
		b.RemoveMessage()
		return
	}
	if b.in.Tip == "dse" {
		b.EditMessageDS()
		return
	}
	if b.in.Tip == "tge" {
		//b.EditMessageTG()// нужно исправить
		return
	}

	var memory models.BridgeTempMemory

	memory.Wg.Add(1)

	// Создаем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup
	chatIdsTG, chatIdsDS := b.Channels()

	lenChTG := len(chatIdsTG)
	resultChannelTg := make(chan models.MessageTg, lenChTG)
	if lenChTG > 0 {
		wg.Add(lenChTG)
		b.sendTg(chatIdsTG, resultChannelTg, &wg)
	}

	// DS
	lenChannels := len(chatIdsDS)
	resultChannelDs := make(chan models.MessageDs, lenChannels)
	if b.in.Reply != nil && b.in.Reply.Text != "" && b.in.Reply.UserName == "gote1st_bot" {
		at := strings.SplitN(b.in.Reply.Text, "\n", 2)
		b.in.Reply.UserName = at[0]
		b.in.Reply.Text = at[1]
	}
	if lenChannels > 0 {
		wg.Add(lenChannels)
		b.client.Ds.SendBridgeAsync(b.in.Text, b.GetSenderName(), chatIdsDS, b.in.FileUrl, b.in.Avatar, b.in.Reply, resultChannelDs, &wg)
	}

	go func() {
		wg.Wait()
		close(resultChannelTg)
		close(resultChannelDs)
		for value := range resultChannelTg {
			memory.MessageTg = append(memory.MessageTg, value)
		}
		for value := range resultChannelDs {
			memory.MessageDs = append(memory.MessageDs, value)
		}
		memory.Wg.Done()
	}()
	memory.Wg.Wait()
	b.in = models.BridgeMessage{}
	b.messages = append(b.messages, memory)
}

func (b *Bridge) sendTg(chatIdsTG []string, resultChannelTg chan<- models.MessageTg, wg *sync.WaitGroup) {
	textTg := fmt.Sprintf("%s\n%s", b.GetSenderName(), b.in.Text)
	if b.in.Reply != nil && (b.in.Reply.FileUrl != "" || b.in.Reply.Text != "") {
		textTg = fmt.Sprintf("%s\n%s\nReply: %s", b.GetSenderName(), b.in.Text, b.in.Reply.Text)
		go b.client.Tg.SendBridgeAsync(chatIdsTG, textTg, b.in.Reply.FileUrl, resultChannelTg, wg)
	} else {
		go b.client.Tg.SendBridgeAsync(chatIdsTG, textTg, b.in.FileUrl, resultChannelTg, wg)
	}
}
