package watsappClient

import (
	"go.mau.fi/whatsmeow/types"
	corpsConfig "kz_bot/internal/clients/corpConfig"
	"kz_bot/internal/dbase"
	"kz_bot/internal/models"
	"strings"
	"time"

	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	waLog "go.mau.fi/whatsmeow/util/log"
	"google.golang.org/protobuf/proto"
)

type Watsapp struct {
	corpsConfig.CorpConfig
	dbase         dbase.Db
	cli           *whatsmeow.Client
	log           waLog.Logger
	historySyncID int32
	startupTime   time.Time
}
type Wa interface {
	Send(chatid, text string)
}

func (w *Watsapp) Send(chatid, text string) {
	args := []string{chatid, text}
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
func (w *Watsapp) LogicMIXwa(text, name, nameid, chatid string) {
	ok, config := w.CorpConfig.CheckChannelConfigWA(chatid)
	w.AccesChatWA(text, chatid)
	if ok {
		in := models.InMessage{
			Mtext:       text,
			Tip:         "wa",
			Name:        name,
			NameMention: name,
			Wa: struct {
				Nameid string
			}{
				Nameid: nameid},
			Config: config,
			Option: struct {
				Callback bool
				Edit     bool
				Update   bool
				Queue    bool
			}{},
		}
		models.ChWa <- in
	}

}
func (w *Watsapp) ChatName(chatid string) string {
	chatName := ""
	group, ok := w.parseJID(chatid)
	if !ok {
		return ""
	} else if group.Server != types.GroupServer {
		w.log.Errorf("Input must be a group JID (@%s)", types.GroupServer)
		return ""
	}
	resp, err := w.cli.GetGroupInfo(group)
	if err != nil {
		w.log.Errorf("Failed to get group info: %v", err)
	} else {
		w.log.Infof("Group info: %+v", resp)
		chatName = resp.Name
	}
	return chatName
}
