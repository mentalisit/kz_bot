package WhatsappClient

import (
	"context"
	"fmt"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"time"
)

func (b *Whatsapp) SendChannelDelSecond(chatid, textPingNumber string, nameId []string, second int) {
	groupJID, _ := types.ParseJID(chatid)
	var ti uint32 = 30
	a := &waProto.Message{
		ExtendedTextMessage: &waProto.ExtendedTextMessage{
			Text: &textPingNumber,
			ContextInfo: &waProto.ContextInfo{
				MentionedJid: nameId,
				Expiration:   &ti,
			},
		},
	}

	resp, err := b.wc.SendMessage(context.Background(), groupJID, a)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Duration(second) * time.Second)
	b.DeleteMessage(chatid, resp.ID)

}
func (b *Whatsapp) SendText(chatid, text string) string {
	groupJID, _ := types.ParseJID(chatid)

	var msg = &waProto.Message{
		ViewOnceMessage: &waProto.FutureProofMessage{
			Message: &waProto.Message{
				Conversation: &text,
			},
		},
	}

	resp, err := b.wc.SendMessage(context.Background(), groupJID, msg)
	if err != nil {
		b.log.Println(err)
	}
	return resp.ID

}
func (b *Whatsapp) SendMention(chatid, textPingNumber string, nameId []string) string {
	groupJID, _ := types.ParseJID(chatid)
	var ti uint32 = 30
	a := &waProto.Message{
		ExtendedTextMessage: &waProto.ExtendedTextMessage{
			Text: &textPingNumber,
			ContextInfo: &waProto.ContextInfo{
				MentionedJid: nameId,
				Expiration:   &ti,
			},
		},
	}

	resp, err := b.wc.SendMessage(context.Background(), groupJID, a)
	if err != nil {
		fmt.Println(err)
	}
	return resp.ID
}

func (b *Whatsapp) DeleteMessage(chatId, mesId string) {
	groupJID, _ := types.ParseJID(chatId)

	_, err := b.wc.RevokeMessage(groupJID, mesId)
	if err != nil {
		b.log.Println(err)
	}
}
func (b *Whatsapp) DeleteMessageSecond(chatId, mesId string, second int) {
	go func() {
		time.Sleep(time.Duration(second) * time.Second)

		groupJID, _ := types.ParseJID(chatId)

		_, err := b.wc.RevokeMessage(groupJID, mesId)
		if err != nil {
			b.log.Println(err)
		}
	}()
}
