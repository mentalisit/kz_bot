package bot

import (
	"fmt"
	"kz_bot/internal/clients"
	"kz_bot/internal/dbase/dbaseMysql"
	"kz_bot/internal/models"
	"regexp"
	"strconv"
)

type Bot struct {
	Tg clients.TelegramInterface
	Ds clients.DiscordInterface
	Db *dbaseMysql.Db
	in models.InMessage
}

func NewBot(tg clients.TelegramInterface, ds clients.DiscordInterface, db *dbaseMysql.Db) *Bot {
	return &Bot{Tg: tg, Ds: ds, Db: db}
}
func (b *Bot) InitBot() {
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
	var rss, kzb, subs, subs3, qwery string
	if len(b.in.Mtext) > 0 {
		str := b.in.Mtext
		re := regexp.MustCompile(`^([4-9]|[1][0-1])([\+]|[-])(\d|\d{2}|\d{3})$`) //три переменные
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
		re2 := regexp.MustCompile(`^([4-9]|[1][0-1])([\+]|[-])$`) // две переменные
		arr2 := (re2.FindAllStringSubmatch(str, -1))
		if len(arr2) > 0 {
			b.in.Lvlkz = arr2[0][1]
			kzb = arr2[0][2]
			b.in.Timekz = "30"
		}

		re3 := regexp.MustCompile(`^([\+]|[-])([4-9]|[1][0-1])$`) // две переменные для добавления или удаления подписок
		arr3 := (re3.FindAllStringSubmatch(str, -1))
		if len(arr3) > 0 {
			b.in.Lvlkz = arr3[0][2]
			subs = arr3[0][1]
		} else {
			re3 := regexp.MustCompile(`^(Rs|rs)\s(S|s)\s([4-9]|[1][0-1])$`)
			arr3 := (re3.FindAllStringSubmatch(str, -1))
			if len(arr3) > 0 {
				b.in.Lvlkz = arr3[0][3]
				subs = arr3[0][2]
				if subs == "S" || subs == "s" {
					subs = "+"
				} else if subs == "U" || subs == "u" {
					subs = "-"
				}
				fmt.Println("Тестирование подписок совместимости")
			}
		}

		re4 := regexp.MustCompile(`^(["о"]|["О"]|["o"]|["O"])([4-9]|[1][0-1])$`) // две переменные для чтения  очереди
		arr4 := (re4.FindAllStringSubmatch(str, -1))
		if len(arr4) > 0 {
			qwery = arr4[0][1]
			b.in.Lvlkz = arr4[0][2]
		}

		re5 := regexp.MustCompile(`^([4-9]|[1][0-1])([\+][\+])$`) //rs start
		arr5 := (re5.FindAllStringSubmatch(str, -1))
		if len(arr5) > 0 {
			b.in.Lvlkz = arr5[0][1]
			rss = arr5[0][2]
		} else {
			re5 = regexp.MustCompile(`^(Rs|rs)\s(Start|start)\s([4-9]|[1][0-1])$`) //rs start
			arr5 = (re5.FindAllStringSubmatch(str, -1))
			if len(arr5) > 0 {
				b.in.Lvlkz = arr5[0][3]
				rss = "++"
				fmt.Println("Проверка совместимости принудительного старта ")
			}

			re6 := regexp.MustCompile(`^([\+][\+]|[-][-])([4-9]|[1][0-1])$`) // две переменные
			arr6 := (re6.FindAllStringSubmatch(str, -1))                     // для добавления или удаления подписок 3/4
			if len(arr6) > 0 {
				b.in.Lvlkz = arr6[0][2]
				subs3 = arr6[0][1]
			} else {
				re6 = regexp.MustCompile(`^(Rs|rs)\s(S|s)\s([4-9]|[1][0-1])(\+)$`)
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

			re8 := regexp.MustCompile(`^(["T"]|["t"]|["т"]|["Т"])([4-9]|[1][0-1])$`) // запрос топа по уровню
			arr8 := (re8.FindAllStringSubmatch(str, -1))
			if len(arr8) > 0 {
				b.in.Lvlkz = arr8[0][2]
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

			if kzb == "+" {
				b.RsPlus()
			} else if kzb == "-" {
				b.RsMinus()
			} else if len(qwery) > 0 {
				b.Queue()
			} else if len(rss) > 0 {
				b.RsStart()
			} else if subs == "+" {
				go b.Subscribe(1)
			} else if subs3 == "++" {
				go b.Subscribe(3)
			} else if subs == "-" {
				go b.Unsubscribe(1)
			} else if subs3 == "--" {
				go b.Unsubscribe(3)
			} else if len(arr8) > 0 {
				//go TopLevel(in, lvlkz)
			} else if len(slot) > 0 {
				emodjiadd(in, slot, emo)
			} else if b.logicIfText() {
			} else if str == "1" {
				//dsSendChannelDel1m(in.config.DsChannel, "test "+emReadName(in.name))
				//} else if in.config.TgChannel != 0 && in.config.DsChannel != "" {
				//	go bridge(in)
			}
		}
	}
}
func (b *Bot) logicIfText() bool {
	iftext := true
	switch b.in.Mtext {
	case "Ивент старт":
		EventStart(in)
	case "Ивент стоп":
		EventStop(in)
	case "+":
		b.Plus()
	case "-":
		b.Minus()
	case "Справка":
		hhelpName(in)
	case "Справка1":
		if in.tip == "ds" {
			dsDelMessage(in.config.DsChannel, in.Ds.mesid)
			helpChannelUpdate(in.config.DsChannel)
		}
	case "Топ":
		TopAll(in)
	case "Эмоджи":
		emodjis(in)

	default:
		iftext = false
	}
	return iftext
}
