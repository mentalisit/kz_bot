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
	Db    dbase.Db
	in    models.InMessage
	Mu    sync.Mutex
	log   *logrus.Logger
	debug bool
	wg    sync.WaitGroup
}

func NewBot(cl clients.Client, db dbase.Db, log *logrus.Logger, debug bool) *Bot {
	return &Bot{Client: cl, Db: db, log: log, debug: debug}
}
func (b *Bot) InitBot() {
	b.log.Println("Бот загружен и готов к работе ")
	go func() { //цикл для удаления сообщений
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
				b.MinusMin()    //ежеминутное обновление активной очереди
				b.Ds.Autohelp() //автозапуск справки для дискорда
			}

			time.Sleep(1 * time.Second)
		}

	}()

	for {
		//ПОЛУЧЕНИЕ СООБЩЕНИЙ ПО ГЛОБАЛЬНЫМ КАНАЛАМ ... НУЖНО ПЕРЕДЕЛАТЬ
		select {
		case in := <-models.ChTg: //получение с телеги
			b.in = in
			b.LogicRs()
		case in := <-models.ChDs: //получение с дискорда
			b.in = in
			b.LogicRs()
		case in := <-models.ChWa: //получение с ватса
			b.in = in
			fmt.Printf("\n\nin Watsapp %+v\n", in)
			b.LogicRs()
		}
	}
	b.log.Panic("Ошибка в боте")

}

// LogicRs логика игры
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
	if b.in.Tip == ds {
		b.Ds.CleanChat(b.in.Config.DsChannel, b.in.Ds.Mesid, b.in.Mtext)
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
		b.cleanChat()
	} else if b.in.Tip == tg {
		text := fmt.Sprintf("(TG)%s \n%s", b.in.Name, b.in.Mtext)
		b.Ds.SendChannelDelSecond(b.in.Config.DsChannel, text, 180)
	}
}
