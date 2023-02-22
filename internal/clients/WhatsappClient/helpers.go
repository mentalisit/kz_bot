package WhatsappClient

import (
	"fmt"
	"strings"

	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
)

func (b *Whatsapp) reloadContacts() {
	if _, err := b.wc.Store.Contacts.GetAllContacts(); err != nil {
		b.log.Errorf("error on update of contacts: %v", err)
	}

	allcontacts, err := b.wc.Store.Contacts.GetAllContacts()
	if err != nil {
		b.log.Errorf("error on update of contacts: %v", err)
	}

	if len(allcontacts) > 0 {
		b.contacts = allcontacts
	}
}

func (b *Whatsapp) getSenderName(info types.MessageInfo) string {
	// Parse AD JID
	var senderJid types.JID
	senderJid.User, senderJid.Server = info.Sender.User, info.Sender.Server

	sender, exists := b.contacts[senderJid]

	if !exists || (sender.FullName == "" && sender.FirstName == "") {
		b.reloadContacts() // Contacts may need to be reloaded
		sender, exists = b.contacts[senderJid]
	}

	if exists && sender.FullName != "" {
		return sender.FullName
	}

	if info.PushName != "" {
		return info.PushName
	}

	if exists && sender.FirstName != "" {
		return sender.FirstName
	}

	return "Кто тот, бот еще не понял"
}

func (b *Whatsapp) getSenderNotify(senderJid types.JID) string {
	sender, exists := b.contacts[senderJid]

	if !exists || (sender.FullName == "" && sender.PushName == "" && sender.FirstName == "") {
		b.reloadContacts() // Contacts may need to be reloaded
		sender, exists = b.contacts[senderJid]
	}

	if !exists {
		return "someone"
	}

	if exists && sender.FullName != "" {
		return sender.FullName
	}

	if exists && sender.PushName != "" {
		return sender.PushName
	}

	if exists && sender.FirstName != "" {
		return sender.FirstName
	}

	return "someone"
}

func isGroupJid(identifier string) bool {
	return strings.HasSuffix(identifier, "@g.us") ||
		strings.HasSuffix(identifier, "@temp") ||
		strings.HasSuffix(identifier, "@broadcast")
}

func (b *Whatsapp) getDevice(dbAddres string) (*store.Device, error) {
	device := &store.Device{}
	addres := fmt.Sprintf("file:%s?_foreign_keys=on&_pragma=busy_timeout=10000", dbAddres)
	storeContainer, err := sqlstore.New("sqlite", addres, nil)
	if err != nil {
		return device, fmt.Errorf("failed to connect to database: %v", err)
	}

	device, err = storeContainer.GetFirstDevice()
	if err != nil {
		return device, fmt.Errorf("failed to get device: %v", err)
	}

	return device, nil
}

func (b *Whatsapp) getGroupName(channel string) (group bool, groupsName string) {
	byJid := isGroupJid(channel)

	groups, err := b.wc.GetJoinedGroups()
	if err != nil {
		b.log.Println(err)
	}

	// verify if we are member of the given group
	if byJid {
		gJID, err := types.ParseJID(channel)
		if err != nil {
			fmt.Println(err)
		}

		for _, group := range groups {
			if group.JID == gJID {
				return true, group.Name
			}
		}
	}

	return false, ""
}
