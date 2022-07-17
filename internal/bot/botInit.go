package bot

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"sync"
	"time"

	"kz_bot/internal/clients"
	"kz_bot/internal/dbase"
	"kz_bot/internal/models"
)

type Bot struct {
	clients.Client
	Db  dbase.Db
	in  models.InMessage
	Mu  sync.Mutex
	log *logrus.Logger
}

func NewBot(cl clients.Client, db dbase.Db, log *logrus.Logger) *Bot {
	return &Bot{Client: cl, Db: db, log: log}
}
func (b *Bot) InitBot() {
	b.log.Println("Бот загружен и готов к работе ")
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
				b.Ds.Autohelp()
			}

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
		} else if b.SendALLChannel() {
			//пробуем мост между месенджерами
		} else if b.in.Config.TgChannel != 0 && b.in.Config.DsChannel != "" {
			b.bridge()
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
	case "Статистика":
		b.Statistic()
	case "Статистика.":
		b.log.Println("Case StatisticA")
		b.StatisticA()
	default:
		iftext = false
	}
	return iftext
}
func (b *Bot) bridge() {
	if b.in.Tip == ds {
		text := fmt.Sprintf("(DS)%s \n%s", b.in.Name, b.in.Mtext)
		b.Tg.SendChannelDelSecond(b.in.Config.TgChannel, text, 180)
	} else if b.in.Tip == tg {
		text := fmt.Sprintf("(TG)%s \n%s", b.in.Name, b.in.Mtext)
		b.Ds.SendChannelDelSecond(b.in.Config.DsChannel, text, 180)
	}
}
