package bot

import (
	"fmt"
	"kz_bot/internal/hades/ReservCopyPaste/ReservCopy"
	"regexp"
	"strconv"
)

// lang ok
// ivent not lang
func (b *Bot) lRsPlus() bool {
	var kzb string
	kz := false
	re := regexp.MustCompile(`^([3-9]|[1][0-2])([\+]|[-])(\d|\d{2}|\d{3})$`) //три переменные
	arr := re.FindAllStringSubmatch(b.in.Mtext, -1)
	if len(arr) > 0 {
		kz = true
		b.in.Lvlkz = arr[0][1]
		kzb = arr[0][2]
		timekzz, err := strconv.Atoi(arr[0][3])
		if err != nil {
			b.log.Println("Ошибка преобразования Atoi", err)
			timekzz = 0
		}
		if timekzz > 180 {
			timekzz = 180
		}
		b.in.Timekz = strconv.Itoa(timekzz)
	}
	re2 := regexp.MustCompile(`^([3-9]|[1][0-2])([\+]|[-])$`) // две переменные
	arr2 := (re2.FindAllStringSubmatch(b.in.Mtext, -1))
	if len(arr2) > 0 {
		kz = true
		b.in.Lvlkz = arr2[0][1]
		kzb = arr2[0][2]
		b.in.Timekz = "30"
	}
	switch kzb {
	case "+":
		b.RsPlus()
	case "-":
		b.RsMinus()
	default:
		kz = false
	}
	return kz
}

func (b *Bot) lSubs() (bb bool) {
	bb = false
	var subs string
	re3 := regexp.MustCompile(`^([\+]|[-])([3-9]|[1][0-2])$`) // две переменные для добавления или удаления подписок
	arr3 := (re3.FindAllStringSubmatch(b.in.Mtext, -1))
	if len(arr3) > 0 {
		b.in.Lvlkz = arr3[0][2]
		subs = arr3[0][1]
		bb = true
	}
	re3s := regexp.MustCompile(`^(Rs|rs)\s(S|s|u|U)\s([3-9]|[1][0-2])$`)
	arr3s := (re3s.FindAllStringSubmatch(b.in.Mtext, -1))
	if len(arr3s) > 0 {
		b.in.Lvlkz = arr3s[0][3]
		subs = arr3s[0][2]
		bb = true
		if subs == "S" || subs == "s" {
			subs = "+"
		} else if subs == "U" || subs == "u" {
			subs = "-"
		}
		//b.log.Println("Тестирование подписок совместимости")
	}
	re6 := regexp.MustCompile(`^([\+][\+]|[-][-])([3-9]|[1][0-2])$`) // две переменные
	arr6 := (re6.FindAllStringSubmatch(b.in.Mtext, -1))              // для добавления или удаления подписок 3/4
	if len(arr6) > 0 {
		bb = true
		b.in.Lvlkz = arr6[0][2]
		subs = arr6[0][1]
	} else {
		re6 = regexp.MustCompile(`^(Rs|rs)\s(S|s|u|U)\s([3-9]|[1][0-2])(\+)$`)
		arr6 = (re6.FindAllStringSubmatch(b.in.Mtext, -1))
		if len(arr6) > 0 {
			bb = true
			b.in.Lvlkz = arr6[0][3]
			subs = arr6[0][2]
			if subs == "S" || subs == "s" {
				subs = "++"
			} else if subs == "U" || subs == "u" {
				subs = "--"
			}
			//b.log.Println("проверка совместимости подписок 3 из 4")
		}
	}

	readd := regexp.MustCompile(`^(подписать)\s([3-9]|[1][0-2])\s(@\w+)\s([1]|[3])$`)
	arradd := (readd.FindAllStringSubmatch(b.in.Mtext, -1))
	if len(arradd) > 0 && b.client.Tg.CheckAdminTg(b.in.Config.TgChannel, b.in.Name) {
		bb = true
		atoi, err := strconv.Atoi(arradd[0][4])
		if err != nil {
			return false
		}
		a := arradd[0][3]
		b.in.NameMention = a
		b.in.Name = a[1 : len(a)-1]
		b.in.Lvlkz = arradd[0][2]
		b.Subscribe(atoi)

	}

	switch subs {
	case "+":
		b.Subscribe(1)
	case "++":
		b.Subscribe(3)
	case "-":
		b.Unsubscribe(1)
	case "--":
		b.Unsubscribe(3)
	}
	return bb
}

func (b *Bot) lQueue() (bb bool) {
	re4 := regexp.MustCompile(`^([о]|[О]|[q]|[Q]|[Ч]|[ч])([3-9]|[1][0-2])$`) // две переменные для чтения  очереди
	arr4 := (re4.FindAllStringSubmatch(b.in.Mtext, -1))
	if len(arr4) > 0 {
		b.in.Lvlkz = arr4[0][2]
		bb = true
		b.QueueLevel()
	}
	//rus
	if b.in.Mtext == "Очередь" || b.in.Mtext == "очередь" {
		bb = true
		b.QueueAll()
	}
	//ukr
	if b.in.Mtext == "Черга" || b.in.Mtext == "черга" {
		bb = true
		b.QueueAll()
	}
	//eng
	if b.in.Mtext == "Queue" || b.in.Mtext == "queue" {
		bb = true
		b.QueueAll()
	}

	re4s := regexp.MustCompile(`^(Rs|rs)\s(Q|q)$`) // две переменные для чтения  очереди
	arr4s := (re4s.FindAllStringSubmatch(b.in.Mtext, -1))
	if len(arr4s) > 0 {
		bb = true
		b.QueueAll() //проверка совместимости
	}

	re4s = regexp.MustCompile(`^(Rs|rs)\s(Q|q)\s([3-9]|[1][0-2])$`)
	arr4s = (re4s.FindAllStringSubmatch(b.in.Mtext, -1))
	if len(arr4s) > 0 {
		bb = true
		b.in.Lvlkz = arr4s[0][3]
		b.QueueLevel()
	}
	return bb
}

func (b *Bot) lRsStart() (bb bool) {
	var rss string
	re5 := regexp.MustCompile(`^([3-9]|[1][0-2])([\+][\+])$`) //rs start
	arr5 := (re5.FindAllStringSubmatch(b.in.Mtext, -1))
	if len(arr5) > 0 {
		bb = true
		b.in.Lvlkz = arr5[0][1]
		rss = arr5[0][2]
	} else {
		re5 = regexp.MustCompile(`^(Rs|rs)\s(Start|start)\s([3-9]|[1][0-2])$`) //rs start
		arr5 = (re5.FindAllStringSubmatch(b.in.Mtext, -1))
		if len(arr5) > 0 {
			bb = true
			b.in.Lvlkz = arr5[0][3]
			rss = "++"
			//b.log.Println("Проверка совместимости принудительного старта ")
		}
	}
	reP := regexp.MustCompile(`^([3-9]|[1][0-2])([\+][\+][\+])$`) //p30pl
	arrP := (reP.FindAllStringSubmatch(b.in.Mtext, -1))
	if len(arrP) > 0 {
		b.in.Lvlkz = arrP[0][1]
		bb = true
		b.Pl30()
	}
	if rss == "++" {
		b.RsStart()
	}
	return bb
}

// ivent not lang
func (b *Bot) lEvent() (bb bool) {
	re7 := regexp.MustCompile(`^(["К"]|["к"])\s([0-9]+)\s([0-9]+)$`) // ивент
	arr7 := (re7.FindAllStringSubmatch(b.in.Mtext, -1))
	if len(arr7) > 0 {
		bb = true
		points, err := strconv.Atoi(arr7[0][3])
		if err != nil {
			b.log.Println("Ошибка преобразования Аtoi", err)
		}
		numkz, err := strconv.Atoi(arr7[0][2])
		if err != nil {
			b.log.Println("Ошибка преобразования Аtoi", err)
		}
		b.EventPoints(numkz, points)
	}
	re7s := regexp.MustCompile(`^(rs|Rs)\s(p|P)\s([0-9]+)\s([0-9]+)$`)
	arr7s := (re7s.FindAllStringSubmatch(b.in.Mtext, -1))
	if len(arr7s) > 0 {
		bb = true
		points, err := strconv.Atoi(arr7[0][4])
		if err != nil {
			b.log.Println("Ошибка преобразования Аtoi", err)
		}
		numkz, err := strconv.Atoi(arr7[0][3])
		if err != nil {
			b.log.Println("Ошибка преобразования Аtoi", err)
		}
		b.EventPoints(numkz, points)
	}
	switch b.in.Mtext {
	case "Ивент старт":
		b.EventStart()
		bb = true
	case "Ивент стоп":
		b.EventStop()
		bb = true
	}
	return bb
}

func (b *Bot) lTop() (bb bool) {
	re8 := regexp.MustCompile(`^(Топ)\s([3-9]|[1][0-2])$`) // запрос топа по уровню
	arr8 := (re8.FindAllStringSubmatch(b.in.Mtext, -1))
	if len(arr8) > 0 {
		b.in.Lvlkz = arr8[0][2]
		b.TopLevel()
		bb = true
		return bb
	}
	//eng
	re8e := regexp.MustCompile(`^(Top)\s([3-9]|[1][0-2])$`) // запрос топа по уровню
	arr8e := (re8e.FindAllStringSubmatch(b.in.Mtext, -1))
	if len(arr8e) > 0 {
		b.in.Lvlkz = arr8[0][2]
		b.TopLevel()
		bb = true
		return bb
	}

	switch b.in.Mtext {
	case "Топ":
		bb = true
		b.TopAll()

	case "Top":

		bb = true
		b.TopAll()

	}

	return bb
}

func (b *Bot) lEmoji() (bb bool) {
	var slot, emo string
	reEmodji := regexp.MustCompile("^(Эмоджи)\\s([1-4])\\s(<:\\w+:\\d+>)$") //добавления внутрених эмоджи
	arrEmodji := (reEmodji.FindAllStringSubmatch(b.in.Mtext, -1))
	if len(arrEmodji) > 0 {
		slot = arrEmodji[0][2]
		emo = arrEmodji[0][3]
	}
	reEmodji = regexp.MustCompile("^(Эмоджи)\\s([1-4])\\s(\\P{Greek})$") //добавления эмоджи
	arrEmodji = (reEmodji.FindAllStringSubmatch(b.in.Mtext, -1))
	if len(arrEmodji) > 0 {
		slot = arrEmodji[0][2]
		emo = arrEmodji[0][3]
	}
	reEmodji = regexp.MustCompile("^(Эмоджи)\\s([1-4])$") //удаление эмоджи с ячейки
	arrEmodji = (reEmodji.FindAllStringSubmatch(b.in.Mtext, -1))
	if len(arrEmodji) > 0 {
		slot = arrEmodji[0][2]
		emo = ""
	}
	reEmodji = regexp.MustCompile("^(Rs|rs)\\s(icon)\\s([1-4])\\s(del)$") //удаление эмоджи с ячейки совместимость
	arrEmodji = (reEmodji.FindAllStringSubmatch(b.in.Mtext, -1))
	if len(arrEmodji) > 0 {
		slot = arrEmodji[0][3]
		emo = ""
	}

	reEmodji = regexp.MustCompile("^(Rs|rs)\\s(icon)\\s([1-4])\\s(\\&\\#[0-9]+\\;)$") //Эмоджи совместимость
	arrEmodji = (reEmodji.FindAllStringSubmatch(b.in.Mtext, -1))
	if len(arrEmodji) > 0 {
		slot = arrEmodji[0][3]
		emo = arrEmodji[0][4]
	}
	if slot != "" {
		b.emodjiadd(slot, emo)
		bb = true
	}
	if b.in.Mtext == "Эмоджи" || b.in.Mtext == "Emoji" {
		bb = true
		b.emodjis()
	}
	return bb
}

func (b *Bot) SendALLChannel() (bb bool) {
	if b.in.Name == "Mentalisit" {
		re := regexp.MustCompile(`^(Всем|всем)\s([А-Яа-я\s.]+)$`)
		arr := (re.FindAllStringSubmatch(b.in.Mtext, -1))
		if len(arr) > 0 {
			fmt.Println(arr[0])
			bb = true

			text := arr[0][2]

			d, t, w := b.storage.Cache.ReadAllChannel()
			if len(d) > 0 {
				for _, chatds := range d {
					b.client.Ds.Send(chatds, text)
				}
			}
			if len(t) > 0 {
				for _, chattg := range t {
					b.client.Tg.SendChannel(chattg, text)
				}
			}
			if len(w) > 0 {
				for _, chatwa := range w {
					b.client.Wa.SendText(chatwa, text)
				}
			}
		}
	}
	return bb
}

func (b *Bot) lIfCommand() bool {
	re := regexp.MustCompile(`^\. Добавить ([0-2]) (.+)`)
	matches := re.FindStringSubmatch(b.in.Mtext)
	if len(matches) > 0 {
		fmt.Println("rang " + matches[1])
		fmt.Println("name " + matches[2])
		d := ReservCopy.NewReservDB()
		rang, _ := strconv.Atoi(matches[1])
		d.UpdateMember([]ReservCopy.Member{ReservCopy.Member{
			CorpName: b.in.Config.CorpName,
			UserName: matches[2],
			Rang:     rang,
		}})
		return true
	}

	reclin := regexp.MustCompile(`^\. Очистка (\d{1,2}|100)`)
	matches = reclin.FindStringSubmatch(b.in.Mtext)
	if len(matches) > 0 {
		fmt.Println("Очистка " + matches[1])
		fmt.Println("limitMessage " + matches[2])
		if matches[1] == "Очистка" {
			b.client.Ds.CleanOldMessageChannel(b.in.Config.DsChannel, matches[2])
			return true
		}
	}
	return false
}
