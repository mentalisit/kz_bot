package relay

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"kz_bot/internal/clients"
	"kz_bot/internal/models"
	"kz_bot/internal/storage"
	Relay2 "kz_bot/internal/storage/CorpsConfig/Relay"
)

type Relay struct {
	log      *logrus.Logger
	storage  *storage.Storage
	client   *clients.Clients
	in       models.RelayMessage
	messages []models.RelayMessageMemory
}

func NewRelay(l *logrus.Logger, s *storage.Storage, c *clients.Clients) *Relay {
	fmt.Println("loadrelay")
	r := &Relay{
		log:     l,
		storage: s,
		client:  c,
	}
	go r.inbox()
	go r.removeIfTimeDay()
	r.loadRelayConfig()

	return r
}
func (r *Relay) inbox() {
	for {
		select {
		case in := <-r.client.Ds.ChanRelay:
			r.in = in
			fmt.Printf("in relay ds  %+v\n", r.in)
			r.logic()
		}
	}
}

func (r *Relay) loadRelayConfig() {
	if len(*Relay2.R) == 0 {
		relayConfig := r.storage.CorpsConfig.RelayDB.ReadAllRelay()
		for _, q := range relayConfig {
			*Relay2.R = append(*Relay2.R, q)
		}
	}
}
