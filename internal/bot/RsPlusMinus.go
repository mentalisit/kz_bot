package bot

import (
	"context"
	"fmt"
	"time"
)

//lang ok

func (b *Bot) RsPlus() {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.debug {
		fmt.Printf("\n\nin RsPlus %+v\n", b.in)
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
		wamesid := ""
		var n map[string]string
		n = make(map[string]string)
		n["lang"] = b.in.Config.Country
		if b.in.Config.DsChannel != "" {
			n["lvlkz"], err = b.client.Ds.RoleToIdPing(b.GetLang("kz")+b.in.Lvlkz, b.in.Config.Guildid)
			if err != nil {
				b.log.Info(fmt.Sprintf("RoleToIdPing %+v lvl %s", b.in.Config, b.in.Lvlkz[1:]))
			}
		}

		if countQueue == 0 {
			if b.in.Config.DsChannel != "" {
				b.wg.Add(1)
				go func() {
					n["name1"] = fmt.Sprintf("%s  🕒  %s  (%d)", b.emReadName(b.in.NameMention, ds), b.in.Timekz, numkzN)
					emb := b.client.Ds.EmbedDS(n, numkzL, 1, false)
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
						b.GetLang("ocheredKz"), b.in.Lvlkz, numkzL,
						b.emReadName(b.in.Name, tg), b.in.Timekz, b.GetLang("min."), numkzN,
						b.in.Lvlkz, b.GetLang("prinuditelniStart"))
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
					n["name1"] = fmt.Sprintf("%s  🕒  %d  (%d)", b.emReadName(u.User1.Mention, ds), u.User1.Timedown, u.User1.Numkzn)
					n["name2"] = fmt.Sprintf("%s  🕒  %s  (%d)", b.emReadName(b.in.NameMention, ds), b.in.Timekz, numkzN)
					emb := b.client.Ds.EmbedDS(n, numkzL, 2, false)
					text := n["lvlkz"] + " 2/4 " + b.in.Name + b.GetLang("prisoedenilsyKocheredi")
					go b.client.Ds.SendChannelDelSecond(b.in.Config.DsChannel, text, 10)
					b.client.Ds.EditComplex(u.User1.Dsmesid, b.in.Config.DsChannel, emb)
					b.wg.Done()
				}()
			}
			if b.in.Config.TgChannel != "" {
				b.wg.Add(1)
				go func() {
					text1 := fmt.Sprintf("%s%s (%d)\n", b.GetLang("ocheredKz"), b.in.Lvlkz, numkzL)
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
		} else if countQueue == 2 {
			dsmesid = u.User1.Dsmesid

			if b.in.Config.DsChannel != "" {
				b.wg.Add(1)
				go func() {
					n["name1"] = fmt.Sprintf("%s  🕒  %d  (%d)", b.emReadName(u.User1.Mention, b.in.Tip), u.User1.Timedown, u.User1.Numkzn)
					n["name2"] = fmt.Sprintf("%s  🕒  %d  (%d)", b.emReadName(u.User2.Mention, b.in.Tip), u.User2.Timedown, u.User2.Numkzn)
					n["name3"] = fmt.Sprintf("%s  🕒  %s  (%d)", b.emReadName(b.in.NameMention, b.in.Tip), b.in.Timekz, numkzN)
					lvlk3, err4 := b.client.Ds.RoleToIdPing(b.GetLang("kz")+b.in.Lvlkz+"+", b.in.Config.Guildid)
					if err4 != nil {
						b.log.Info(fmt.Sprintf("RoleToIdPing %+v lvl %s", b.in.Config, b.in.Lvlkz[1:]))
					}
					emb := b.client.Ds.EmbedDS(n, numkzL, 3, false)
					text := fmt.Sprintf("%s  3/4 %s %s %s %s",
						n["lvlkz"], b.in.Name, b.GetLang("prisoedenilsyKocheredi"), lvlk3, b.GetLang("nujenEsheOdinDlyFulki"))
					go b.client.Ds.SendChannelDelSecond(b.in.Config.DsChannel, text, 10)
					b.client.Ds.EditComplex(u.User1.Dsmesid, b.in.Config.DsChannel, emb)
					b.wg.Done()
				}()
			}
			if b.in.Config.TgChannel != "" {
				b.wg.Add(1)
				go func() {
					text1 := fmt.Sprintf("%s%s (%d)\n", b.GetLang("ocheredKz"), b.in.Lvlkz, numkzL)
					name1 := fmt.Sprintf("1. %s - %d%s (%d) \n",
						b.emReadName(u.User1.Name, tg), u.User1.Timedown, b.GetLang("min."), u.User1.Numkzn)
					name2 := fmt.Sprintf("2. %s - %d%s (%d) \n",
						b.emReadName(u.User2.Name, tg), u.User2.Timedown, b.GetLang("min."), u.User2.Numkzn)
					name3 := fmt.Sprintf("3. %s - %s%s (%d) \n",
						b.emReadName(b.in.Name, tg), b.in.Timekz, b.GetLang("min."), numkzN)
					text2 := fmt.Sprintf("\n%s++ - %s", b.in.Lvlkz, b.GetLang("prinuditelniStart"))
					text := fmt.Sprintf("%s %s %s %s %s", text1, name1, name2, name3, text2)
					tgmesid = b.client.Tg.SendEmded(b.in.Lvlkz, b.in.Config.TgChannel, text)
					go b.client.Tg.DelMessage(b.in.Config.TgChannel, u.User1.Tgmesid)
					b.storage.Update.MesidTgUpdate(ctx, tgmesid, b.in.Lvlkz, b.in.Config.CorpName)
					b.SubscribePing(3)
					b.wg.Done()
				}()
			}
		}
		if countQueue <= 2 {
			b.wg.Wait()
			b.storage.DbFunc.InsertQueue(ctx, dsmesid, wamesid, b.in.Config.CorpName, b.in.Name, b.in.NameMention, b.in.Tip, b.in.Lvlkz, b.in.Timekz, tgmesid, numkzN)
		}

		if countQueue == 3 {
			dsmesid = u.User1.Dsmesid

			textEvent, numkzEvent := b.EventText()
			numberevent := b.storage.Event.NumActiveEvent(b.in.Config.CorpName) //получаем номер ивета если он активен
			if numberevent > 0 {
				numkzL = numkzEvent
			}

			if b.in.Config.DsChannel != "" {
				b.wg.Add(1)
				go func() {
					n1, n2, n3, n4 := b.nameMention(u, ds)
					go b.client.Ds.DeleteMessage(b.in.Config.DsChannel, u.User1.Dsmesid)
					go b.client.Ds.SendChannelDelSecond(b.in.Config.DsChannel,
						" 4/4 "+b.in.Name+" "+b.GetLang("prisoedenilsyKocheredi"), 10)
					text := fmt.Sprintf("4/4 %s%s %s\n"+
						" %s\n"+
						" %s\n"+
						" %s\n"+
						" %s\n"+
						"%s %s",
						b.GetLang("ocheredKz"), b.in.Lvlkz, b.GetLang("sformirovana"),
						n1,
						n2,
						n3,
						n4,
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
					n1, n2, n3, n4 := b.nameMention(u, tg)
					go b.client.Tg.DelMessage(b.in.Config.TgChannel, u.User1.Tgmesid)
					go b.client.Tg.SendChannelDelSecond(b.in.Config.TgChannel,
						b.in.Name+b.GetLang("zakrilOcheredKz")+b.in.Lvlkz, 10)
					text := fmt.Sprintf("%s%s %s\n"+
						"%s\n"+
						"%s\n"+
						"%s\n"+
						"%s\n"+
						" %s \n"+
						"%s",
						b.GetLang("ocheredKz"), b.in.Lvlkz, b.GetLang("sformirovana"),
						n1, n2, n3, n4,
						b.GetLang("Vigru"), textEvent)
					tgmesid = b.client.Tg.SendChannel(b.in.Config.TgChannel, text)
					b.storage.Update.MesidTgUpdate(ctx, tgmesid, b.in.Lvlkz, b.in.Config.CorpName)
					b.wg.Done()
				}()
			}

			b.wg.Wait()
			b.storage.DbFunc.InsertQueue(ctx, dsmesid, wamesid, b.in.Config.CorpName, b.in.Name, b.in.NameMention, b.in.Tip, b.in.Lvlkz, b.in.Timekz, tgmesid, numkzN)
			b.storage.Update.UpdateCompliteRS(ctx, b.in.Lvlkz, dsmesid, tgmesid, wamesid, numkzL, numberevent, b.in.Config.CorpName)

			//проверка есть ли игрок в других чатах
			user := []string{u.User1.Name, u.User2.Name, u.User3.Name, b.in.Name}
			go b.elseChat(user)

		}

	}
}
func (b *Bot) RsMinus() {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.debug {
		fmt.Printf("\n in RsMinus %+v\n", b.in)
	}
	b.iftipdelete()
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	CountNames, err := b.storage.Count.СountName(ctx, b.in.Name, b.in.Lvlkz, b.in.Config.CorpName) //проверяем есть ли игрок в очереди
	if err != nil {
		b.log.Error(err.Error())
		return
	}
	if CountNames == 0 {
		b.ifTipSendMentionText(b.GetLang("tiNeVOcheredi"))
	} else if CountNames > 0 {
		//чтение айди очечреди
		u := b.storage.DbFunc.ReadAll(ctx, b.in.Lvlkz, b.in.Config.CorpName)
		//удаление с БД
		b.storage.DbFunc.DeleteQueue(ctx, b.in.Name, b.in.Lvlkz, b.in.Config.CorpName)
		//проверяем очередь
		countQueue, err2 := b.storage.Count.CountQueue(ctx, b.in.Lvlkz, b.in.Config.CorpName)
		if err2 != nil {
			b.log.Error(err2.Error())
			return
		}
		//numkzL := numberQueueLvl(in, lvlkz) + 1
		if b.in.Config.DsChannel != "" {
			go b.client.Ds.SendChannelDelSecond(b.in.Config.DsChannel, b.in.Name+b.GetLang("pokinulOchered"), 10)
			if countQueue == 0 {
				go b.client.Ds.SendChannelDelSecond(b.in.Config.DsChannel,
					fmt.Sprintf("%s%s %s.", b.GetLang("ocheredKz"), b.in.Lvlkz, b.GetLang("bilaUdalena")), 10)
				go b.client.Ds.DeleteMessage(b.in.Config.DsChannel, u.User1.Dsmesid)
			}
		}
		if b.in.Config.TgChannel != "" {
			go b.client.Tg.SendChannelDelSecond(b.in.Config.TgChannel, b.in.Name+b.GetLang("pokinulOchered"), 10)
			if countQueue == 0 {
				go b.client.Tg.SendChannelDelSecond(b.in.Config.TgChannel,
					fmt.Sprintf("%s%s %s.", b.GetLang("ocheredKz"), b.in.Lvlkz, b.GetLang("bilaUdalena")), 10)
				go b.client.Tg.DelMessage(b.in.Config.TgChannel, u.User1.Tgmesid)
			}
		}
		if countQueue > 0 {

			b.QueueLevel()
		}
	}
}
