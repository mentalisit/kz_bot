package bot

import (
	"fmt"
	"sync"
	"time"

	"kz_bot/internal/clients"
	"kz_bot/internal/dbase"
	"kz_bot/internal/dbase/dbaseMysql"
	"kz_bot/internal/models"
)

type Bot struct {
	Tg    clients.TelegramInterface
	Ds    clients.DiscordInterface
	Db    dbase.DbInterface
	in    models.InMessage
	Mutex sync.Mutex
}

func NewBot(tg clients.TelegramInterface, ds clients.DiscordInterface, db *dbaseMysql.Db) *Bot {
	return &Bot{Tg: tg, Ds: ds, Db: db}
}
func (b *Bot) InitBot() {
	fmt.Println("Бот загружен и готов к работе ")
	go func() {
		for {
			if time.Now().Second() == 0 {
				tt := b.Db.TimerDeleteMessage() //получаем ид сообщения для удаления
				for _, t := range tt {
					if t.Dsmesid != "" {
						b.Ds.DeleteMesageSecond(t.Dschatid, t.Dsmesid, t.Timed)
					}
					if t.Tgmesid != 0 {
						b.Tg.DelMessageSecond(t.Tgchatid, t.Tgmesid, t.Timed)
					}
				}
				b.MinusMin()
			}
			b.autohelp()

			time.Sleep(1 * time.Second)
		}

	}()

	for {
		select {
		case in := <-models.ChTg:
			b.in = in
			b.LogicRs()
		case in := <-models.ChDs:
			b.in = in
			b.LogicRs()
		case in := <-models.ChWa:
			//b.in = in
			fmt.Println(in)
		}
	}
}

func (b *Bot) LogicRs() {
	if len(b.in.Mtext) > 0 {

		if b.lRsPlus() {
		} else if b.lSubs() {
		} else if b.lQueue() {
		} else if b.lRsStart() {
		} else if b.lEvent() {
		} else if b.lTop() {
		} else if b.lEmoji() {
		} else if b.logicIfText() {
			//пробуем мост между месенджерами
		} else if b.in.Config.TgChannel != 0 && b.in.Config.DsChannel != "" {
			//	go bridge(in)
		}
	}
}

func (b *Bot) logicIfText() bool {
	iftext := true
	switch b.in.Mtext {
	case "+":
		b.Plus()
	case "-":
		b.Minus()
	case "Справка":
		b.hhelp()
	default:
		iftext = false
	}
	return iftext
}
