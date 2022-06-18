package bot

import (
	"fmt"
	"regexp"
	"strconv"
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
				tt := b.Db.TimerDeleteMessage()
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
		}
	}
}

func (b *Bot) LogicRs() {
	var rss, p30pl, kzb, subs, subs3, qwery, top string
	if len(b.in.Mtext) > 0 {
		str := b.in.Mtext
		re := regexp.MustCompile(`^([4-9]|[1][0-2])([\+]|[-])(\d|\d{2}|\d{3})$`) //три переменные
		arr := (re.FindAllStringSubmatch(str, -1))
		if len(arr) > 0 {
			b.in.Lvlkz = arr[0][1]
			kzb = arr[0][2]
			timekzz, err := strconv.Atoi(arr[0][3])
			if err != nil {
				fmt.Println("Ошибка преоразования atoi", err)
			}
			if timekzz > 180 {
				timekzz = 180
			}
			b.in.Timekz = strconv.Itoa(timekzz)
		}
		re2 := regexp.MustCompile(`^([4-9]|[1][0-2])([\+]|[-])$`) // две переменные
		arr2 := (re2.FindAllStringSubmatch(str, -1))
		if len(arr2) > 0 {
			b.in.Lvlkz = arr2[0][1]
			kzb = arr2[0][2]
			b.in.Timekz = "30"
		}

		re3 := regexp.MustCompile(`^([\+]|[-])([4-9]|[1][0-2])$`) // две переменные для добавления или удаления подписок
		arr3 := (re3.FindAllStringSubmatch(str, -1))
		if len(arr3) > 0 {
			b.in.Lvlkz = arr3[0][2]
			subs = arr3[0][1]
		}
		re3s := regexp.MustCompile(`^(Rs|rs)\s(S|s)\s([4-9]|[1][0-2])$`)
		arr3s := (re3s.FindAllStringSubmatch(str, -1))
		if len(arr3s) > 0 {
			b.in.Lvlkz = arr3s[0][3]
			subs = arr3s[0][2]
			if subs == "S" || subs == "s" {
				subs = "+"
			} else if subs == "U" || subs == "u" {
				subs = "-"
			}
			fmt.Println("Тестирование подписок совместимости")
		}

		re4 := regexp.MustCompile(`^([о]|[О])([4-9]|[1][0-2])$`) // две переменные для чтения  очереди
		arr4 := (re4.FindAllStringSubmatch(str, -1))
		if len(arr4) > 0 {
			qwery = arr4[0][1]
			b.in.Lvlkz = arr4[0][2]
		}
		re4s := regexp.MustCompile(`^(Rs|rs)\s(Q|q)$`) // две переменные для чтения  очереди
		arr4s := (re4s.FindAllStringSubmatch(str, -1))
		//if len(arr4s) > 0 {
		//	b.QueueAll() //проверка совместимости
		//}

		re5 := regexp.MustCompile(`^([4-9]|[1][0-2])([\+][\+])$`) //rs start
		arr5 := (re5.FindAllStringSubmatch(str, -1))
		if len(arr5) > 0 {
			b.in.Lvlkz = arr5[0][1]
			rss = arr5[0][2]
		} else {
			re5 = regexp.MustCompile(`^(Rs|rs)\s(Start|start)\s([4-9]|[1][0-2])$`) //rs start
			arr5 = (re5.FindAllStringSubmatch(str, -1))
			if len(arr5) > 0 {
				b.in.Lvlkz = arr5[0][3]
				rss = "++"
				fmt.Println("Проверка совместимости принудительного старта ")
			}
		}

		reP := regexp.MustCompile(`^([4-9]|[1][0-2])([\+][\+][\+])$`) //p30pl
		arrP := (reP.FindAllStringSubmatch(str, -1))
		if len(arrP) > 0 {
			b.in.Lvlkz = arrP[0][1]
			p30pl = arrP[0][2]
		}

		re6 := regexp.MustCompile(`^([\+][\+]|[-][-])([4-9]|[1][0-2])$`) // две переменные
		arr6 := (re6.FindAllStringSubmatch(str, -1))                     // для добавления или удаления подписок 3/4
		if len(arr6) > 0 {
			b.in.Lvlkz = arr6[0][2]
			subs3 = arr6[0][1]
		} else {
			re6 = regexp.MustCompile(`^(Rs|rs)\s(S|s)\s([4-9]|[1][0-2])(\+)$`)
			arr6 = (re6.FindAllStringSubmatch(str, -1))
			if len(arr6) > 0 {
				b.in.Lvlkz = arr6[0][3]
				subs3 = arr6[0][2]
				if subs3 == "S" || subs3 == "s" {
					subs3 = "++"
				} else if subs3 == "U" || subs3 == "u" {
					subs3 = "--"
				}
				fmt.Println("проверка совместимости подписок 3 из 4")
			}
		}

		re7 := regexp.MustCompile(`^(["К"]|["к"])\s([0-9]+)\s([0-9]+)$`) // ивент
		arr7 := (re7.FindAllStringSubmatch(str, -1))
		if len(arr7) > 0 {
			points, err := strconv.Atoi(arr7[0][3])
			if err != nil {
				fmt.Println("Ошибка преоразования atoi", err)
			}
			numkz, err := strconv.Atoi(arr7[0][2])
			if err != nil {
				fmt.Println("Ошибка преоразования atoi", err)
			}
			fmt.Println(numkz, points) //EventPoints(in, numkz, points)

		}

		re8 := regexp.MustCompile(`^(Топ)\\s([4-9]|[1][0-2])$`) // запрос топа по уровню
		arr8 := (re8.FindAllStringSubmatch(str, -1))
		if len(arr8) > 0 {
			b.in.Lvlkz = arr8[0][2]
		}
		re8d := regexp.MustCompile(`^(Топ)\\s([4-9]|[1][0-2])\\s([неделя]|[день])$`) // запрос топа по уровню за период
		arr8d := (re8d.FindAllStringSubmatch(str, -1))
		if len(arr8d) > 0 {
			b.in.Lvlkz = arr8d[0][2]
			top = arr8d[0][3]
		}
		re8s := regexp.MustCompile(`^(Топ)\\s([4-9]|[1][0-2])\\s([неделя]|[день])$`) // запрос топа по уровню за период
		arr8s := (re8s.FindAllStringSubmatch(str, -1))
		if len(arr8s) > 0 {
			b.in.Lvlkz = arr8s[0][2]
			top = arr8s[0][3]
		}

		var slot, emo string
		reEmodji := regexp.MustCompile("^(Эмоджи)\\s([1-4])\\s(<:\\w+:\\d+>)$") //добавления внутрених эмоджи
		arrEmodji := (reEmodji.FindAllStringSubmatch(str, -1))
		if len(arrEmodji) > 0 {
			slot = arrEmodji[0][2]
			emo = arrEmodji[0][3]
		}
		reEmodji = regexp.MustCompile("^(Эмоджи)\\s([1-4])\\s(\\P{Greek})$") //добавления эмоджи
		arrEmodji = (reEmodji.FindAllStringSubmatch(str, -1))
		if len(arrEmodji) > 0 {
			slot = arrEmodji[0][2]
			emo = arrEmodji[0][3]
		}
		reEmodji = regexp.MustCompile("^(Эмоджи)\\s([1-4])$") //удаление эмоджи с ячейки
		arrEmodji = (reEmodji.FindAllStringSubmatch(str, -1))
		if len(arrEmodji) > 0 {
			slot = arrEmodji[0][2]
			emo = ""
		}
		reEmodji = regexp.MustCompile("^(Rs|rs)\\s(icon)\\s([1-4])\\s(del)$") //удаление эмоджи с ячейки совместимость
		arrEmodji = (reEmodji.FindAllStringSubmatch(str, -1))
		if len(arrEmodji) > 0 {
			slot = arrEmodji[0][3]
			emo = ""
		}

		reEmodji = regexp.MustCompile("^(Rs|rs)\\s(icon)\\s([1-4])\\s(\\&\\#[0-9]+\\;)$") //Эмоджи совместимость
		arrEmodji = (reEmodji.FindAllStringSubmatch(str, -1))
		if len(arrEmodji) > 0 {
			slot = arrEmodji[0][3]
			emo = arrEmodji[0][4]
		}

		if kzb == "+" {
			b.RsPlus()
		} else if kzb == "-" {
			b.RsMinus()
		} else if len(qwery) > 0 {
			b.QueueLevel()
		} else if len(rss) > 0 {
			b.RsStart()
		} else if len(p30pl) > 0 {
			b.Pl30()
		} else if len(arr4s) > 0 {
			b.QueueAll()
		} else if subs == "+" {
			go b.Subscribe(1)
		} else if subs3 == "++" {
			go b.Subscribe(3)
		} else if subs == "-" {
			go b.Unsubscribe(1)
		} else if subs3 == "--" {
			go b.Unsubscribe(3)
		} else if len(arr8) > 0 {
			go b.TopLevel()
		} else if top == "день" {
			go b.TopDateLevel(b.t1())
		} else if top == "неделя" {
			go b.TopDateLevel(b.t7())
		} else if len(slot) > 0 {
			b.emodjiadd(slot, emo)
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
	case "Ивент старт":
		b.EventStart()
	case "Ивент стоп":
		b.EventStop()
	case "+":
		b.Plus()
	case "-":
		b.Minus()
	case "Справка":
		b.hhelp()
	case "Топ":
		b.TopAll()
	case "Топ неделя":
		b.TopDate(b.t7())
	case "Топ сутки":
		b.TopDate(b.t1())
	case "Эмоджи":
		b.emodjis()
	default:
		iftext = false
	}
	return iftext
}
