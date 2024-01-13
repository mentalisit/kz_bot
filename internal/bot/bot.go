package bot

import (
	"fmt"
	"kz_bot/internal/clients"
	"kz_bot/internal/config"
	"kz_bot/internal/models"
	"kz_bot/internal/storage"
	"kz_bot/pkg/logger"
	"strconv"
	"sync"
	"time"
)

const (
	ds = "ds"
	tg = "tg"
	wa = "wa"
)

// spravka
type Bot struct {
	storage    *storage.Storage
	client     *clients.Clients
	inbox      chan models.InMessage
	log        *logger.Logger
	debug      bool
	in         models.InMessage
	wg         sync.WaitGroup
	mu         sync.Mutex
	configCorp map[string]models.CorporationConfig
}

func NewBot(storage *storage.Storage, client *clients.Clients, log *logger.Logger, cfg *config.ConfigBot) *Bot {
	b := &Bot{
		storage:    storage,
		client:     client,
		log:        log,
		debug:      cfg.IsDebug,
		inbox:      make(chan models.InMessage, 10),
		configCorp: storage.CorpConfigRS,
	}
	go b.loadInbox()
	go b.RemoveMessage()

	return b
}

func (b *Bot) loadInbox() {
	b.log.Info("Бот загружен и готов к работе ")

	for {
		//ПОЛУЧЕНИЕ СООБЩЕНИЙ
		select {
		case in := <-b.client.Ds.ChanRsMessage:
			b.in = in
			b.LogicRs()
		case in := <-b.client.Tg.ChanRsMessage:
			b.in = in
			b.LogicRs()

		case in := <-b.inbox:
			b.in = in
			b.LogicRs()

			//fmt.Printf("\n\nin message %+v\n", in)
		}
	}
	b.log.Panic("Ошибка в боте")
}
func (b *Bot) RemoveMessage() { //цикл для удаления сообщений
	for {
		<-time.After(1 * time.Minute)
		b.MinusMin()             //ежеминутное обновление активной очереди
		b.client.Ds.Autohelpds() //автозапуск справки для дискорда
		time.Sleep(1 * time.Second)
	}

}

// LogicRs логика игры
func (b *Bot) LogicRs() {
	if len(b.in.Mtext) > 0 && b.in.Mtext != " `edit`" {
		if b.lRsPlus() {
		} else if b.lDarkRsPlus() {
		} else if b.lSubs() {
		} else if b.lDarkSubs() {
		} else if b.lQueue() {
		} else if b.lRsStart() {
		} else if b.lEvent() {
		} else if b.lTop() {
		} else if b.lEmoji() {
		} else if b.logicIfText() {
		} else if b.bridge() {
			//} else if b.lIfCommand() {
			//} else if b.SendALLChannel() {
		} else {
			b.cleanChat()
		}

	} else if b.in.Option.MinusMin {
		b.CheckTimeQueue()
	} else if b.in.Option.Update {
		b.QueueLevel()
	}
}

func (b *Bot) cleanChat() {
	if b.in.Tip == ds && b.in.Config.DelMesComplite == 0 && !b.in.Option.Edit {
		go b.client.Ds.CleanChat(b.in.Config.DsChannel, b.in.Ds.Mesid, b.in.Mtext)
	}
	// if hs ua
	if b.in.Tip == tg && b.in.Config.TgChannel == "-1002116077159/44" {
		go b.client.Tg.DelMessageSecond("-1002116077159/44", strconv.Itoa(b.in.Tg.Mesid), 180)
	}
}

func (b *Bot) logicIfText() bool {
	iftext := true
	switch b.in.Mtext {
	case "+":
		if !b.Plus() {
			iftext = false
		}
	case "-":
		if !b.Minus() {
			iftext = false
		}
	case "Справка":
		b.hhelp()
	default:
		iftext = false
	}
	return iftext
}

func (b *Bot) bridge() bool {
	if b.in.Tip == ds && b.in.Config.Forward {
		text := fmt.Sprintf("(DS)%s \n%s", b.in.Name, b.in.Mtext)
		b.client.Tg.SendChannelDelSecond(b.in.Config.TgChannel, text, 180)
		b.cleanChat()
	} else if b.in.Tip == tg && b.in.Config.Forward {
		text := fmt.Sprintf("(TG)%s \n%s", b.in.Name, b.in.Mtext)
		b.client.Ds.SendChannelDelSecond(b.in.Config.DsChannel, text, 180)
		b.cleanChat()
	}
	return b.in.Config.Forward
}
