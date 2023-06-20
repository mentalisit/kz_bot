package HadesClient

import (
	"github.com/sirupsen/logrus"
	"kz_bot/internal/HadesClient/server"
	"kz_bot/internal/clients"
	"kz_bot/internal/models"
	"kz_bot/internal/storage"
	"sort"
	"time"
)

var mes map[string][]models.MessageHadesClient

type Hades struct {
	cl           *clients.Clients
	storage      *storage.Storage
	toGame       chan models.MessageHadesClient
	ToMessager   chan models.MessageHadesClient
	log          *logrus.Logger
	idMessage    map[string]int64
	idMessageWs1 map[int64]int64
	corporation  map[string]models.CorporationHadesClient
	member       map[string]models.AllianceMember
	in           models.MessageHadesClient
}

func NewHades(log *logrus.Logger, client *clients.Clients, storage *storage.Storage) *Hades {
	h := &Hades{

		toGame:       make(chan models.MessageHadesClient, 100),
		ToMessager:   make(chan models.MessageHadesClient, 200),
		log:          log,
		idMessage:    make(map[string]int64),
		idMessageWs1: make(map[int64]int64),
		corporation:  make(map[string]models.CorporationHadesClient),
		member:       make(map[string]models.AllianceMember),
		cl:           client,
		storage:      storage,
	}
	h.loadDB()
	server.NewServer(h.toGame, h.ToMessager)
	mes = make(map[string][]models.MessageHadesClient)

	go h.inbox()

	return h
}

func (h *Hades) inbox() {
	for {
		select {
		case in := <-h.ToMessager:
			h.sortIncomingMessage(in)
		case in := <-h.cl.Ds.ChanToGame:
			h.filterDs(in)
		case in := <-h.cl.Tg.ChanToGame:
			h.filterTg(in)

		default:
			if len(mes) > 0 {
				h.readmes()
			} else {
				time.Sleep(500 * time.Millisecond)
			}
		}
	}
}

func (h *Hades) sortIncomingMessage(in models.MessageHadesClient) {
	mes[in.Corporation] = append(mes[in.Corporation], in)
	sort.SliceStable(mes[in.Corporation], func(i, j int) bool {
		return mes[in.Corporation][i].MessageId < mes[in.Corporation][j].MessageId
	})
}
func (h *Hades) readmes() {
	for corp, hadesClients := range mes {
		if len(hadesClients) > 0 { // []
			for g, client := range hadesClients {
				h.filterGame(client)
				if g+1 == len(hadesClients) {
					mes[corp] = nil
				}
			}
		}
	}
}
