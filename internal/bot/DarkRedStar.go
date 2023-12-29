package bot

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

const dark = "d"

func (b *Bot) lDarkRsPlus() bool {
	var kzb string
	kz := false
	re := regexp.MustCompile(`^([7-9]|[1][0-2])([\*]|[-])(\d|\d{2}|\d{3})$`) //три переменные
	arr := re.FindAllStringSubmatch(b.in.Mtext, -1)
	if len(arr) > 0 {
		kz = true
		b.in.Lvlkz = dark + arr[0][1]
		kzb = arr[0][2]
		timekzz, err := strconv.Atoi(arr[0][3])
		if err != nil {
			b.log.Error(err.Error())
			timekzz = 0
		}
		if timekzz > 180 {
			timekzz = 180
		}
		b.in.Timekz = strconv.Itoa(timekzz)
	}
	re2 := regexp.MustCompile(`^([7-9]|[1][0-2])([\*]|[-])$`) // две переменные
	arr2 := (re2.FindAllStringSubmatch(b.in.Mtext, -1))
	if len(arr2) > 0 {
		fmt.Println(b.in.Mtext)
		kz = true
		b.in.Lvlkz = dark + arr2[0][1]
		kzb = arr2[0][2]
		b.in.Timekz = "30"
	}
	re2d := regexp.MustCompile(`^(d)([7-9]|[1][0-2])([\+]|[-])$`) // две переменные
	arr2d := (re2d.FindAllStringSubmatch(b.in.Mtext, -1))
	if len(arr2d) > 0 {
		fmt.Println(b.in.Mtext)
		kz = true
		b.in.Lvlkz = dark + arr2d[0][2]
		kzb = arr2d[0][3]
		b.in.Timekz = "30"
	}
	switch kzb {
	case "*":
		b.RsDarkPlus()
	case "+":
		b.RsDarkPlus()
	case "-":
		b.RsMinus()
	default:
		kz = false
	}
	return kz
}
func (b *Bot) RsDarkPlus() {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.debug {
		fmt.Printf("\n\nin RsDarkPlus %+v\n", b.in)
	}
	b.iftipdelete()
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	CountName, err := b.storage.Count.СountName(ctx, b.in.Name, b.in.Lvlkz, b.in.Config.CorpName)
	if err != nil {
		return
	}
	if CountName == 1 { //проверяем есть ли игрок в очереди
		b.ifTipSendMentionText(b.GetLang("tiUjeVocheredi"))
	} else {
		countQueue, err1 := b.storage.Count.CountQueue(ctx, b.in.Lvlkz, b.in.Config.CorpName) //проверяем, есть ли кто-то в очереди
		if err1 != nil {
			return
		}
		numkzN, err2 := b.storage.Count.CountNumberNameActive1(ctx, b.in.Lvlkz, b.in.Config.CorpName, b.in.Name) //проверяем количество боёв по уровню кз игрока
		if err2 != nil {
			return
		}
		numkzL, err3 := b.storage.DbFunc.NumberQueueLvl(ctx, b.in.Lvlkz, b.in.Config.CorpName) //проверяем какой номер боя определенной красной звезды
		if err3 != nil {
			return
		}

		dsmesid := ""
		tgmesid := 0
		var n map[string]string
		n = make(map[string]string)
		n["lang"] = b.in.Config.Country
		n["lvlkz"] = b.client.Ds.RoleToIdPing(b.GetLang("dkz")+b.in.Lvlkz[1:], b.in.Config.Guildid)
		if countQueue == 0 {
			if b.in.Config.DsChannel != "" {
				b.wg.Add(1)
				go func() {
					n["name1"] = fmt.Sprintf("%s  🕒  %s  (%d)", b.emReadName(b.in.Name, ds), b.in.Timekz, numkzN)
					emb := b.client.Ds.EmbedDS(n, numkzL, 1, true)
					dsmesid = b.client.Ds.SendComplexContent(b.in.Config.DsChannel, b.in.Name+b.GetLang("zapustilOchered")+n["lvlkz"])
					time.Sleep(1 * time.Second)
					b.client.Ds.EditComplex(dsmesid, b.in.Config.DsChannel, emb)
					b.client.Ds.AddEnojiRsQueue(b.in.Config.DsChannel, dsmesid)
					b.wg.Done()
				}()
			}
			if b.in.Config.TgChannel != "" {
				b.wg.Add(1)
				go func() {
					text := fmt.Sprintf("%s%s (%d)\n"+
						"1. %s - %s%s (%d) \n\n"+
						"%s++ - %s",
						b.GetLang("ocheredTKz"), b.in.Lvlkz[1:], numkzL,
						b.emReadName(b.in.Name, tg), b.in.Timekz, b.GetLang("min."), numkzN,
						b.in.Lvlkz[1:], b.GetLang("prinuditelniStart"))
					tgmesid = b.client.Tg.SendEmded(b.in.Lvlkz, b.in.Config.TgChannel, text)
					b.SubscribePing(1)
					b.wg.Done()
				}()
			}
		}

		u := b.storage.DbFunc.ReadAll(ctx, b.in.Lvlkz, b.in.Config.CorpName)

		if countQueue == 1 {
			dsmesid = u.User1.Dsmesid

			if b.in.Config.DsChannel != "" {
				b.wg.Add(1)
				go func() {
					n["name1"] = fmt.Sprintf("%s  🕒  %d  (%d)", b.emReadName(u.User1.Name, ds), u.User1.Timedown, u.User1.Numkzn)
					n["name2"] = fmt.Sprintf("%s  🕒  %s  (%d)", b.emReadName(b.in.Name, ds), b.in.Timekz, numkzN)
					emb := b.client.Ds.EmbedDS(n, numkzL, 2, true)
					text := n["lvlkz"] + " 2/3 " + b.in.Name + b.GetLang("prisoedenilsyKocheredi")
					go b.client.Ds.SendChannelDelSecond(b.in.Config.DsChannel, text, 10)
					b.client.Ds.EditComplex(u.User1.Dsmesid, b.in.Config.DsChannel, emb)
					b.wg.Done()
				}()
			}
			if b.in.Config.TgChannel != "" {
				b.wg.Add(1)
				go func() {
					text1 := fmt.Sprintf("%s%s (%d)\n", b.GetLang("ocheredTKz"), b.in.Lvlkz, numkzL)
					name1 := fmt.Sprintf("1. %s - %d%s (%d) \n",
						b.emReadName(u.User1.Name, tg), u.User1.Timedown, b.GetLang("min."), u.User1.Numkzn)
					name2 := fmt.Sprintf("2. %s - %s%s (%d) \n",
						b.emReadName(b.in.Name, tg), b.in.Timekz, b.GetLang("min."), numkzN)
					text2 := fmt.Sprintf("\n%s++ - %s", b.in.Lvlkz, b.GetLang("prinuditelniStart"))
					text := fmt.Sprintf("%s %s %s %s", text1, name1, name2, text2)
					tgmesid = b.client.Tg.SendEmded(b.in.Lvlkz, b.in.Config.TgChannel, text)
					go b.client.Tg.DelMessage(b.in.Config.TgChannel, u.User1.Tgmesid)
					b.storage.Update.MesidTgUpdate(ctx, tgmesid, b.in.Lvlkz, b.in.Config.CorpName)
					b.wg.Done()
				}()
			}
		}
		if countQueue < 2 {
			b.wg.Wait()
			b.storage.DbFunc.InsertQueue(ctx, dsmesid, "", b.in.Config.CorpName, b.in.Name, b.in.NameMention, b.in.Tip, b.in.Lvlkz, b.in.Timekz, tgmesid, numkzN)
		}

		if countQueue == 2 {
			dsmesid = u.User1.Dsmesid

			textEvent, numkzEvent := b.EventText()
			numberevent := b.storage.Event.NumActiveEvent(b.in.Config.CorpName) //получаем номер ивета если он активен
			if numberevent > 0 {
				numkzL = numkzEvent
			}

			if b.in.Config.DsChannel != "" {
				b.wg.Add(1)
				go func() {
					n1, n2, _, _ := b.nameMention(u, ds)
					go b.client.Ds.DeleteMessage(b.in.Config.DsChannel, u.User1.Dsmesid)
					go b.client.Ds.SendChannelDelSecond(b.in.Config.DsChannel,
						" 3/3 "+b.in.Name+" "+b.GetLang("prisoedenilsyKocheredi"), 10)
					text := fmt.Sprintf("3/3 %s%s %s\n"+
						" %s\n"+
						" %s\n"+
						" %s\n"+
						"%s %s",
						b.GetLang("ocheredTKz"), b.in.Lvlkz[1:], b.GetLang("sformirovana"),
						n1,
						n2,
						b.in.NameMention,
						b.GetLang("Vigru"), textEvent)

					if b.in.Tip == ds {
						dsmesid = b.client.Ds.SendWebhook(text, "КзБот", b.in.Config.DsChannel, b.in.Config.Guildid, b.in.Ds.Avatar)
					} else {
						dsmesid = b.client.Ds.Send(b.in.Config.DsChannel, text)
					}
					b.storage.Update.MesidDsUpdate(ctx, dsmesid, b.in.Lvlkz, b.in.Config.DsChannel)
					b.wg.Done()
				}()
			}
			if b.in.Config.TgChannel != "" {
				b.wg.Add(1)
				go func() {
					n1, n2, _, _ := b.nameMention(u, tg)
					go b.client.Tg.DelMessage(b.in.Config.TgChannel, u.User1.Tgmesid)
					go b.client.Tg.SendChannelDelSecond(b.in.Config.TgChannel,
						b.in.Name+b.GetLang("zakrilOcheredTKz")+b.in.Lvlkz[1:], 10)
					text := fmt.Sprintf("%s%s %s\n"+
						"%s\n"+
						"%s\n"+
						"%s\n"+
						" %s \n"+
						"%s",
						b.GetLang("ocheredTKz"), b.in.Lvlkz[1:], b.GetLang("sformirovana"),
						n1, n2, b.in.NameMention,
						b.GetLang("Vigru"), textEvent)
					tgmesid = b.client.Tg.SendChannel(b.in.Config.TgChannel, text)
					b.storage.Update.MesidTgUpdate(ctx, tgmesid, b.in.Lvlkz, b.in.Config.CorpName)
					b.wg.Done()
				}()
			}

			b.wg.Wait()
			b.storage.DbFunc.InsertQueue(ctx, dsmesid, "", b.in.Config.CorpName, b.in.Name, b.in.NameMention, b.in.Tip, b.in.Lvlkz, b.in.Timekz, tgmesid, numkzN)
			b.storage.Update.UpdateCompliteRS(ctx, b.in.Lvlkz, dsmesid, tgmesid, "", numkzL, numberevent, b.in.Config.CorpName)

			//проверка есть ли игрок в других чатах
			user := []string{u.User1.Name, u.User2.Name, b.in.Name}
			go b.elseChat(user)

		}

	}
}
func (b *Bot) lDarkSubs() (bb bool) {
	bb = false
	var subs string
	re3 := regexp.MustCompile(`^([\+]|[-])(d)([7-9]|[1][0-2])$`) // две переменные для добавления или удаления подписок
	arr3 := (re3.FindAllStringSubmatch(b.in.Mtext, -1))
	if len(arr3) > 0 {
		b.in.Lvlkz = "d" + arr3[0][3]
		subs = arr3[0][1]
		bb = true
	}

	re6 := regexp.MustCompile(`^([\+][\+]|[-][-])(d)([7-9]|[1][0-2])$`) // две переменные
	arr6 := (re6.FindAllStringSubmatch(b.in.Mtext, -1))                 // для добавления или удаления подписок 2/3
	if len(arr6) > 0 {
		bb = true
		b.in.Lvlkz = "d" + arr6[0][3]
		subs = arr6[0][1]
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
