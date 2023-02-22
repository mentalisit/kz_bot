package WhatsappClient

import (
	"go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	"kz_bot/internal/models"
	"strings"
)

// nolint:gocritic
func (b *Whatsapp) eventHandler(evt interface{}) {
	switch e := evt.(type) {
	case *events.Message:
		b.handleMessage(e)
	}
}

func (b *Whatsapp) handleMessage(message *events.Message) {
	msg := message.Message
	switch {
	case msg == nil, message.Info.IsFromMe, message.Info.Timestamp.Before(b.startedAt):
		return
	}

	switch {
	case msg.Conversation != nil || msg.ExtendedTextMessage != nil:
		b.handleTextMessage(message.Info, msg)
	}
}

// nolint:funlen
func (b *Whatsapp) handleTextMessage(messageInfo types.MessageInfo, msg *proto.Message) {
	senderJID := messageInfo.Sender
	channel := messageInfo.Chat

	senderName := b.getSenderName(messageInfo)

	if msg.GetExtendedTextMessage() == nil && msg.GetConversation() == "" {
		b.log.Debugf("message without text content? %#v", msg)
		return
	}

	var text string

	// nolint:nestif
	if msg.GetExtendedTextMessage() == nil {
		text = msg.GetConversation()
	} else {
		text = msg.GetExtendedTextMessage().GetText()
		ci := msg.GetExtendedTextMessage().GetContextInfo()

		if senderJID == (types.JID{}) && ci.Participant != nil {
			senderJID = types.NewJID(ci.GetParticipant(), types.DefaultUserServer)
		}

		if ci.MentionedJid != nil {
			// handle user mentions
			for _, mentionedJID := range ci.MentionedJid {
				numberAndSuffix := strings.SplitN(mentionedJID, "@", 2)

				// mentions comes as telephone numbers and we don't want to expose it to other bridges
				// replace it with something more meaninful to others
				mention := b.getSenderNotify(types.NewJID(numberAndSuffix[0], types.DefaultUserServer))

				text = strings.Replace(text, "@"+numberAndSuffix[0], "@"+mention, 1)
			}
		}
	}

	//senderJID.String() //380989033544@s.whatsapp.net
	//senderName       // Павел
	//channel.String() //380989033544-1616265986@g.us
	//messageInfo.ID		//566282D3E644DBB08BDEEA11A68E2DD7

	b.accesChatWA(text, channel.String())

	//b.SendMention(channel.String(), "@"+senderJID.User+"  vgdjgk", []string{messageInfo.Sender.ToNonAD().String()})
	ok, config := b.storage.Cache.CheckChannelConfigWA(channel.String())
	if ok {

		b.inbox <- models.InMessage{
			Mtext:       text,
			Tip:         "wa",
			Name:        senderName,
			NameMention: "@" + senderJID.User,
			Wa: struct {
				Nameid string
				Mesid  string
			}{
				Nameid: senderJID.String(),
				Mesid:  messageInfo.ID,
			},
			Config: config,
			Option: models.Option{
				InClient: true,
			},
		}
	}
}
