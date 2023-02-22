package WhatsappClient

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"kz_bot/internal/config"
	"kz_bot/internal/models"
	"kz_bot/internal/storage"
	"os"
	"sync"
	"time"

	"github.com/mdp/qrterminal"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
	waLog "go.mau.fi/whatsmeow/util/log"

	_ "modernc.org/sqlite" // needed for sqlite
)

type Whatsapp struct {
	startedAt   time.Time
	wc          *whatsmeow.Client
	contacts    map[types.JID]types.ContactInfo
	users       map[string]types.ContactInfo
	userAvatars map[string]string
	log         *logrus.Logger
	inbox       chan models.InMessage
	storage     *storage.Storage
	debug       bool
	mu          *sync.RWMutex
}

func NewWhatsapp(inbox chan models.InMessage, log *logrus.Logger, st *storage.Storage, cfg *config.ConfigBot) *Whatsapp {

	b := &Whatsapp{
		users:       make(map[string]types.ContactInfo),
		userAvatars: make(map[string]string),
		log:         log,
		inbox:       inbox,
		storage:     st,
		debug:       cfg.IsDebug,
	}

	err := b.connect(cfg.Token.NameDbWhatsapp)
	if err != nil {
		fmt.Println(err)
	}

	return b
}

// Connect to WhatsApp
func (b *Whatsapp) connect(dbAddres string) error {
	device, err := b.getDevice(dbAddres)
	if err != nil {
		return err
	}

	b.wc = whatsmeow.NewClient(device, waLog.Stdout("Client", "INFO", true))
	b.wc.AddEventHandler(b.eventHandler)

	firstlogin := false
	var qrChan <-chan whatsmeow.QRChannelItem
	if b.wc.Store.ID == nil {
		firstlogin = true
		qrChan, err = b.wc.GetQRChannel(context.Background())
		if err != nil && !errors.Is(err, whatsmeow.ErrQRStoreContainsID) {
			return errors.New("failed to to get QR channel:" + err.Error())
		}
	}

	err = b.wc.Connect()
	if err != nil {
		return errors.New("failed to connect to WhatsApp: " + err.Error())
	}

	if b.wc.Store.ID == nil {
		for evt := range qrChan {
			if evt.Event == "code" {
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
			} else {
				b.log.Infof("QR channel result: %s", evt.Event)
			}
		}
	}

	// disconnect and reconnect on our first login/pairing
	// for some reason the GetJoinedGroups in JoinChannel doesn't work on first login
	if firstlogin {
		b.wc.Disconnect()
		time.Sleep(time.Second)

		err = b.wc.Connect()
		if err != nil {
			return errors.New("failed to connect to WhatsApp: " + err.Error())
		}
	}

	fmt.Println("WhatsApp connection successful")

	b.contacts, err = b.wc.Store.Contacts.GetAllContacts()
	if err != nil {
		return errors.New("failed to get contacts: " + err.Error())
	}

	b.startedAt = time.Now()

	// map all the users
	for id, contact := range b.contacts {
		if !isGroupJid(id.String()) && id.String() != "status@broadcast" {
			// it is user
			b.users[id.String()] = contact
		}
	}
	return nil
}

func (b *Whatsapp) disconnect() error {
	b.wc.Disconnect()

	return nil
}
