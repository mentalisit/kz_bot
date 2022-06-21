package watsappClient

/*
func (w *Watsapp) handleTextMessage(messageInfo types.MessageInfo, msg *proto.Message)  {
	senderJID := messageInfo.Sender
	channel := messageInfo.Chat

	senderName := messageInfo.PushName
	if senderName == "" {
		senderName = "Someone" // don't expose telephone number
	}

	if msg.GetExtendedTextMessage() == nil && msg.GetConversation() == "" {
		w.log.Debugf("message without text content? %#v", msg)
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
				mention := w.getSenderNotify(types.NewJID(numberAndSuffix[0], types.DefaultUserServer))
				if mention == "" {
					mention = "someone"
				}

				text = strings.Replace(text, "@"+numberAndSuffix[0], "@"+mention, 1)
			}
		}
	}
	in:=models.InMessage{
		Mtext:       text,
		Tip:         "wa",
		Name:        name,
		NameMention: psend,
	}
	models.ChWa<-in
}

*/
