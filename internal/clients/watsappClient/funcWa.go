package watsappClient

import (
	"strings"

	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	waLog "go.mau.fi/whatsmeow/util/log"
	"google.golang.org/protobuf/proto"
)

type Watsapp struct {
	cli *whatsmeow.Client
	log waLog.Logger
}

func (w *Watsapp) Send(args []string) {
	if len(args) < 2 {
		w.log.Errorf("Usage: send <jid> <text>")
		return
	}
	recipient, ok := w.parseJID(args[0])
	if !ok {
		return
	}
	msg := &waProto.Message{Conversation: proto.String(strings.Join(args[1:], " "))}
	ts, err := w.cli.SendMessage(recipient, "", msg)
	if err != nil {
		w.log.Errorf("Error sending message: %v", err)
	} else {
		w.log.Infof("Message sent (server timestamp: %s)", ts)
	}
}
