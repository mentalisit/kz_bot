package watsappClient

import (
	"context"
	"fmt"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	waTypes "go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"
	"time"
)

type Reactions string

const (
	ReactLove          Reactions = "❤️"
	ReactHandLike      Reactions = "👍"
	ReactHandBad       Reactions = "👎"
	ReactHandFolded    Reactions = "🙏"
	ReactFaceHot       Reactions = "🥵"
	ReactFacePalm      Reactions = "🤦‍♂️"
	ReactFaceClown     Reactions = "🤡"
	ReactFaceZany      Reactions = "🤪"
	ReactFaceVomiting  Reactions = "🤮"
	ReactFaceTired     Reactions = "😫"
	ReactFaceLying     Reactions = "🤥"
	ReactFaceTear      Reactions = "🥲"
	ReactFaceLove      Reactions = "🥰"
	ReactFaceMoney     Reactions = "🤑"
	ReactFaceInnoncent Reactions = "😇"
	ReactFaceWow       Reactions = "😮"
	ReactFaceJoy       Reactions = "😂"
	ReactFaceSad       Reactions = "😥"
	ReactFaceHugging   Reactions = "🤗"
	ReactFlagIndonesia Reactions = "🇮🇩"
	ReactHundred       Reactions = "💯"
	ReactMedalGold     Reactions = "🥇"
	ReactMedalSilver   Reactions = "🥈"
	ReactMedalBronze   Reactions = "🥉"
	ReactAirplane      Reactions = "✈️"
	ReactPlester       Reactions = "🩹"
	ReactAlarm         Reactions = "⏰"
	ReactBadminton     Reactions = "🏸"
	ReactNotEntry      Reactions = "⛔"
	ReactRocket        Reactions = "🚀"
)

type HandlerSentMessage func(waTypes.JID, string, *waProto.Message)

var handlersSentMessage = []HandlerSentMessage{}

func (w *Watsapp) LogicMIXwa(text, name, nameid, chatid, mesid string) {
	//ok, config := w.CorpsConfig.CheckChannelConfigWA(chatid)
	//w.AccesChatWA(text, chatid)
	//if ok {
	//	in := models.InMessage{
	//		Mtext:       text,
	//		Tip:         "wa",
	//		Name:        name,
	//		NameMention: "name",
	//		Wa: struct {
	//			Nameid string
	//			Mesid  string
	//		}{
	//			Nameid: nameid,
	//			Mesid:  mesid},
	//		Config: config,
	//		Option: models.Option{InClient: true},
	//	}
	//	models.ChWa <- in
	fmt.Println("text", text)
	send, err := w.Send(chatid, text)
	if err != nil {
		fmt.Println(err)
		return
	}
	time.Sleep(5 * time.Second)
	w.DeleteMessage(chatid, send)
	fmt.Println("del", send)

	//}
}

func SendReactMessage(event *events.Message, react Reactions, client *whatsmeow.Client) (whatsmeow.SendResponse, error) {
	this_message := &waProto.Message{
		ReactionMessage: &waProto.ReactionMessage{
			Key: &waProto.MessageKey{
				RemoteJid:   proto.String(event.Info.Chat.ToNonAD().String()),
				Participant: proto.String(event.Info.Sender.ToNonAD().String()),
				FromMe:      proto.Bool(event.Info.IsFromMe),
				Id:          &event.Info.ID,
			},
			Text: proto.String(string(react)),
		},
	}

	if resp, err := client.SendMessage(context.Background(), event.Info.Chat, whatsmeow.GenerateMessageID(), this_message); err != nil {
		return resp, err
	} else {
		for _, handler := range handlersSentMessage {
			handler(event.Info.Chat, resp.ID, this_message)
		}

		return resp, err
	}
}
func NewButtons(header interface{}, content, footer string, buttons []*waProto.ButtonsMessage_Button, ctx *waProto.ContextInfo) (*waProto.ButtonsMessage, error) {

	var message = &waProto.ButtonsMessage{
		ContentText: &content,
		FooterText:  &footer,
		Buttons:     buttons,
		ContextInfo: ctx,
	}

	switch hd := header.(type) {
	case *waProto.ButtonsMessage_DocumentMessage:
		message.HeaderType = waProto.ButtonsMessage_DOCUMENT.Enum()
		message.Header = hd

	case *waProto.ButtonsMessage_ImageMessage:
		message.HeaderType = waProto.ButtonsMessage_IMAGE.Enum()
		message.Header = hd

	case *waProto.ButtonsMessage_VideoMessage:
		message.HeaderType = waProto.ButtonsMessage_VIDEO.Enum()
		message.Header = hd

	case *waProto.ButtonsMessage_Text:
		message.HeaderType = waProto.ButtonsMessage_TEXT.Enum()
		message.Header = hd

	case *waProto.ButtonsMessage_LocationMessage:
		message.HeaderType = waProto.ButtonsMessage_LOCATION.Enum()
		message.Header = hd

	}

	return message, nil

}
