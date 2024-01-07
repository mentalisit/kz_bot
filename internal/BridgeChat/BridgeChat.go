package BridgeChat

import (
	"fmt"
	"kz_bot/internal/clients"
	"kz_bot/internal/models"
	"kz_bot/internal/storage"
	"kz_bot/pkg/logger"
	"strings"
)

type Bridge struct {
	log      *logger.Logger
	client   *clients.Clients
	in       models.BridgeMessage
	messages []models.BridgeTempMemory
	storage  *storage.Storage
	configs  map[string]models.BridgeConfig
}

func NewBridge(log *logger.Logger, client *clients.Clients, storage *storage.Storage) *Bridge {
	b := &Bridge{
		log:     log,
		client:  client,
		storage: storage,
		configs: storage.BridgeConfigs,
	}
	go b.inbox()
	return b
}
func (b *Bridge) inbox() {
	for {
		select {
		case b.in = <-b.client.Ds.ChanBridgeMessage:
			fmt.Printf(" in BridgeMessage ds  %+v  Config:%s\n", b.in, b.ifNoConfig())
			b.logic()
		case b.in = <-b.client.Tg.ChanBridgeMessage:
			fmt.Printf(" in BridgeMessage tg  %+v  Config:%s\n", b.in, b.ifNoConfig())
			b.logic()
		}
	}
}
func (b *Bridge) logic() {
	if strings.HasPrefix(b.in.Text, ".") {
		b.Command()
		return
	} else {
		b.logicMessage()
	}

}
func (b *Bridge) ifNoConfig() string {
	if b.in.Config != nil {
		return b.in.Config.NameRelay
	}
	return ""
}
