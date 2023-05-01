package bot

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"kz_bot/internal/clients"
	"kz_bot/internal/config"
	"kz_bot/internal/models"
	"kz_bot/internal/storage"
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
	storage *storage.Storage
	client  *clients.Clients
	inbox   chan models.InMessage
	log     *logrus.Logger
	debug   bool
	in      models.InMessage
	wg      sync.WaitGroup
	mu      sync.Mutex
}

func NewBot(storage *storage.Storage, client *clients.Clients, log *logrus.Logger, cfg *config.ConfigBot) *Bot {
	b := &Bot{storage: storage, client: client, log: log, debug: cfg.IsDebug, inbox: client.Inbox}
	go b.loadInbox()
	go b.RemoveMessage()

	return b
}

func (b *Bot) loadInbox() {
	b.log.Println("Бот загружен и готов к работе ")

	for {
		//ПОЛУЧЕНИЕ СООБЩЕНИЙ
		select {
		case in := <-b.client.Inbox:
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
		if time.Now().Second() == 0 {
			tt := b.storage.Timers.TimerDeleteMessage(context.Background()) //получаем ид сообщения для удаления
			for _, t := range tt {
				if t.Dsmesid != "" {
					b.client.Ds.DeleteMesageSecond(t.Dschatid, t.Dsmesid, t.Timed)
				}
				if t.Tgmesid != 0 {
					b.client.Tg.DelMessageSecond(t.Tgchatid, t.Tgmesid, t.Timed)
				}
			}
			b.MinusMin()             //ежеминутное обновление активной очереди
			b.client.Ds.Autohelpds() //автозапуск справки для дискорда
		}

		time.Sleep(1 * time.Second)
	}

}

// LogicRs логика игры
func (b *Bot) LogicRs() {
	if len(b.in.Mtext) > 0 && b.in.Mtext != " edit" {
		if b.lRsPlus() {
		} else if b.lSubs() {
		} else if b.lQueue() {
		} else if b.lRsStart() {
		} else if b.lEvent() {
		} else if b.lTop() {
		} else if b.lEmoji() {
		} else if b.logicIfText() {
		} else if b.SendALLChannel() {
			//пробуем мост между месенджерами
		} else if b.in.Config.TgChannel != 0 && b.in.Config.DsChannel != "" {
			b.bridge()
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
		b.client.Ds.CleanChat(b.in.Config.DsChannel, b.in.Ds.Mesid, b.in.Mtext)
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

func (b *Bot) bridge() {
	if b.in.Tip == ds {
		text := fmt.Sprintf("(DS)%s \n%s", b.in.Name, b.in.Mtext)
		b.client.Tg.SendChannelDelSecond(b.in.Config.TgChannel, text, 181)
		b.cleanChat()
	} else if b.in.Tip == tg {
		text := fmt.Sprintf("(TG)%s \n%s", b.in.Name, b.in.Mtext)
		b.client.Ds.SendChannelDelSecond(b.in.Config.DsChannel, text, 181)
	}
}
