package watsappClient

import (
	"context"
	"fmt"
	"go.mau.fi/whatsmeow/types"
	corpsConfig "kz_bot/internal/clients/corpConfig"
	"kz_bot/internal/dbase"
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
	Send(chatid, text string) (string, error)
	DeleteMessage(chatid, mesid string)
	ChatName(chatid string) string
}

func (w *Watsapp) ChatName(chatid string) string {
	chatName := ""
	group, ok := w.parseJID(chatid)
	if !ok {
		return ""
	} else if group.Server != types.GroupServer {
		//w.log.Errorf("Input must be a group JID (@%s)", types.GroupServer)
		return ""
	}
	resp, err := w.cli.GetGroupInfo(group)
	if err != nil {
		w.log.Errorf("Failed to get group info: %v", err)
	} else {
		//w.log.Infof("Group info: %+v", resp)
		chatName = resp.Name
	}
	return chatName
}

func (w *Watsapp) Send(chatid, text string) (string, error) {
	args := []string{chatid, text}
	if len(args) < 2 {

		return "", nil
	}
	recipient, ok := w.parseJID(args[0])
	if !ok {
		return "", nil
	}
	//var a []string
	//a = append(a, "380637157959@s.whatsapp.net")
	conversation := proto.String(strings.Join(args[1:], " "))
	//extendedTextMessage := &waProto.ExtendedTextMessage{Text: &text,		ContextInfo: &waProto.ContextInfo{MentionedJid: a,},}
	msg := &waProto.Message{
		Conversation: conversation,
		//ExtendedTextMessage: extendedTextMessage,
	}
	resp, err := w.cli.SendMessage(context.Background(), recipient, "", msg)
	if err != nil {
		w.log.Errorf("Error sending message: %v", err)
	}
	return resp.ID, nil
}

func (w *Watsapp) DeleteMessage(chatid, mesid string) {
	groupJID, _ := types.ParseJID(chatid)
	if mesid == "" {
		fmt.Println("mesid00000")
		return
	}

	_, err := w.cli.RevokeMessage(groupJID, mesid)
	if err != nil {
		fmt.Println("err", err)
		return
	}
}
