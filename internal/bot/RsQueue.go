package bot

import (
	"context"
	"fmt"
	"kz_bot/pkg/utils"
	"time"
)

//lang ok

func (b *Bot) QueueLevel() {
	b.iftipdelete()
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	count, err := b.storage.Count.CountQueue(ctx, b.in.Lvlkz, b.in.Config.CorpName)
	if err != nil {
		return
	}
	numberLvl, err2 := b.storage.DbFunc.NumberQueueLvl(ctx, b.in.Lvlkz, b.in.Config.CorpName)
	if err2 != nil {
		return
	}
	// совподения количество  условие
	if count == 0 {
		if !b.in.Option.Queue {
			text := b.GetLang("ocheredKz") + b.in.Lvlkz + b.GetLang("pusta")
			b.ifTipSendTextDelSecond(text, 10)
		} else if b.in.Option.Queue {
			b.ifTipSendTextDelSecond(b.GetLang("netAktivnuh"), 10)
		}
		return
	}

	u := b.storage.DbFunc.ReadAll(ctx, b.in.Lvlkz, b.in.Config.CorpName)
	var n map[string]string
	n = make(map[string]string)
	n["lang"] = b.in.Config.Country

	if count == 1 {

		if b.in.Config.DsChannel != "" {
			b.wg.Add(1)
			go func() {
				n["name1"] = fmt.Sprintf("%s  🕒  %d  (%d)", b.emReadName(u.User1.Name, ds), u.User1.Timedown, u.User1.Numkzn)
				n["lvlkz"] = b.client.Ds.RoleToIdPing(b.GetLang("kz")+b.in.Lvlkz, b.in.Config.Guildid)
				emb := b.client.Ds.EmbedDS(n, numberLvl, 1, false)
				if b.in.Option.Edit {
					errr := b.client.Ds.EditComplex(u.User1.Dsmesid, b.in.Config.DsChannel, emb)
					if errr != nil {
						b.in.Option.Edit = false
					}
				}
				if !b.in.Option.Edit {
					b.client.Ds.DeleteMessage(b.in.Config.DsChannel, u.User1.Dsmesid)
					dsmesid := b.client.Ds.SendComplex(b.in.Config.DsChannel, emb)

					b.client.Ds.AddEnojiRsQueue(b.in.Config.DsChannel, dsmesid)
					b.storage.Update.MesidDsUpdate(ctx, dsmesid, b.in.Lvlkz, b.in.Config.CorpName)
				}
				b.wg.Done()
			}()
		}
		if b.in.Config.TgChannel != "" {
			b.wg.Add(1)
			go func() {
				text1 := fmt.Sprintf("%s%s (%d)\n", b.GetLang("ocheredKz"), b.in.Lvlkz, numberLvl)
				name1 := fmt.Sprintf("1. %s - %d%s (%d) \n",
					b.emReadName(u.User1.Name, tg), u.User1.Timedown, b.GetLang("min."), u.User1.Numkzn)
				text2 := fmt.Sprintf("\n%s++ - %s", b.in.Lvlkz, b.GetLang("prinuditelniStart"))
				text := fmt.Sprintf("%s %s %s", text1, name1, text2)
				if b.in.Option.Edit {
					b.client.Tg.EditMessageTextKey(b.in.Config.TgChannel, u.User1.Tgmesid, text, b.in.Lvlkz)
				} else if !b.in.Option.Edit {
					mesidTg := b.client.Tg.SendEmded(b.in.Lvlkz, b.in.Config.TgChannel, text)
					b.storage.Update.MesidTgUpdate(ctx, mesidTg, b.in.Lvlkz, b.in.Config.CorpName)
					b.client.Tg.DelMessage(b.in.Config.TgChannel, u.User1.Tgmesid)
				}
				b.wg.Done()
			}()
		}
		if b.in.Config.WaChannel != "" {
			//b.wg.Add(1)
			//go func() {
			//	text1 := fmt.Sprintf("Очередь кз%s (%d)\n", b.in.Lvlkz, numberLvl)
			//	name1 := fmt.Sprintf("1. %s - %dмин. (%d) \n", b.emReadName(u.User1.Name, wa), u.User1.Timedown, u.User1.Numkzn)
			//	text2 := fmt.Sprintf("\n%s++ - принудительный старт", b.in.Lvlkz)
			//	text := fmt.Sprintf("%s %s %s", text1, name1, text2)
			//	wamesid, errs := b.Wa.Send(b.in.Config.WaChannel, text)
			//	if errs != nil {
			//		b.log.Println("error sending rsQueue1")
			//	}
			//	b.Db.Update.MesidWaUpdate(wamesid, b.in.Lvlkz, b.in.Config.CorpName)
			//	b.Wa.DeleteMessage(b.in.Config.WaChannel, u.User1.Wamesid)
			//	b.wg.Done()
			//}()
		}

	} else if count == 2 {

		if b.in.Config.DsChannel != "" {
			b.wg.Add(1)
			go func() {
				n["name1"] = fmt.Sprintf("%s  🕒  %d  (%d)", b.emReadName(u.User1.Name, ds), u.User1.Timedown, u.User1.Numkzn)
				n["name2"] = fmt.Sprintf("%s  🕒  %d  (%d)", b.emReadName(u.User2.Name, ds), u.User2.Timedown, u.User2.Numkzn)
				n["lvlkz"] = b.client.Ds.RoleToIdPing(b.GetLang("kz")+b.in.Lvlkz, b.in.Config.Guildid)
				emb := b.client.Ds.EmbedDS(n, numberLvl, 2, false)
				if b.in.Option.Edit {
					b.client.Ds.EditComplex(u.User1.Dsmesid, b.in.Config.DsChannel, emb)
				} else if !b.in.Option.Edit {
					b.client.Ds.DeleteMessage(b.in.Config.DsChannel, u.User1.Dsmesid)
					dsmesid := b.client.Ds.SendComplex(b.in.Config.DsChannel, emb)

					b.client.Ds.AddEnojiRsQueue(b.in.Config.DsChannel, dsmesid)
					b.storage.Update.MesidDsUpdate(ctx, dsmesid, b.in.Lvlkz, b.in.Config.CorpName)
				}
				b.wg.Done()
			}()
		}
		if b.in.Config.TgChannel != "" {
			b.wg.Add(1)
			go func() {
				text1 := fmt.Sprintf("%s%s (%d)\n", b.GetLang("ocheredKz"), b.in.Lvlkz, numberLvl)
				name1 := fmt.Sprintf("1. %s - %d%s (%d) \n",
					b.emReadName(u.User1.Name, tg), u.User1.Timedown, b.GetLang("min."), u.User1.Numkzn)
				name2 := fmt.Sprintf("2. %s - %d%s (%d) \n",
					b.emReadName(u.User2.Name, tg), u.User2.Timedown, b.GetLang("min."), u.User2.Numkzn)
				text2 := fmt.Sprintf("\n%s++ - %s", b.in.Lvlkz, b.GetLang("prinuditelniStart"))
				text := fmt.Sprintf("%s %s %s %s", text1, name1, name2, text2)
				if b.in.Option.Edit {
					b.client.Tg.EditMessageTextKey(b.in.Config.TgChannel, u.User1.Tgmesid, text, b.in.Lvlkz)
				} else if !b.in.Option.Edit {
					mesidTg := b.client.Tg.SendEmded(b.in.Lvlkz, b.in.Config.TgChannel, text)
					b.storage.Update.MesidTgUpdate(ctx, mesidTg, b.in.Lvlkz, b.in.Config.CorpName)
					b.client.Tg.DelMessage(b.in.Config.TgChannel, u.User1.Tgmesid)
				}
				b.wg.Done()
			}()
		}
		if b.in.Config.WaChannel != "" {

		}

	} else if count == 3 {

		if b.in.Config.DsChannel != "" {
			b.wg.Add(1)
			go func() {
				n["name1"] = fmt.Sprintf("%s  🕒  %d  (%d)", b.emReadName(u.User1.Name, ds), u.User1.Timedown, u.User1.Numkzn)
				n["name2"] = fmt.Sprintf("%s  🕒  %d  (%d)", b.emReadName(u.User2.Name, ds), u.User2.Timedown, u.User2.Numkzn)
				n["name3"] = fmt.Sprintf("%s  🕒  %d  (%d)", b.emReadName(u.User3.Name, ds), u.User3.Timedown, u.User3.Numkzn)
				n["lvlkz"] = b.client.Ds.RoleToIdPing(b.GetLang("kz")+b.in.Lvlkz, b.in.Config.Guildid)
				emb := b.client.Ds.EmbedDS(n, numberLvl, 3, false)
				if b.in.Option.Edit {
					b.client.Ds.EditComplex(u.User1.Dsmesid, b.in.Config.DsChannel, emb)
				} else if !b.in.Option.Edit {
					b.client.Ds.DeleteMessage(b.in.Config.DsChannel, u.User1.Dsmesid)
					dsmesid := b.client.Ds.SendComplex(b.in.Config.DsChannel, emb)

					b.client.Ds.AddEnojiRsQueue(b.in.Config.DsChannel, dsmesid)
					b.storage.Update.MesidDsUpdate(ctx, dsmesid, b.in.Lvlkz, b.in.Config.CorpName)
				}
				b.wg.Done()
			}()
		}
		if b.in.Config.TgChannel != "" {
			b.wg.Add(1)
			go func() {
				text1 := fmt.Sprintf("%s%s (%d)\n", b.GetLang("ocheredKz"), b.in.Lvlkz, numberLvl)
				name1 := fmt.Sprintf("1. %s - %d%s (%d) \n",
					b.emReadName(u.User1.Name, tg), u.User1.Timedown, b.GetLang("min."), u.User1.Numkzn)
				name2 := fmt.Sprintf("2. %s - %d%s (%d) \n",
					b.emReadName(u.User2.Name, tg), u.User2.Timedown, b.GetLang("min."), u.User2.Numkzn)
				name3 := fmt.Sprintf("3. %s - %d%s (%d) \n",
					b.emReadName(u.User3.Name, tg), u.User3.Timedown, b.GetLang("min."), u.User3.Numkzn)
				text2 := fmt.Sprintf("\n%s++ - %s", b.in.Lvlkz, b.GetLang("prinuditelniStart"))
				text := fmt.Sprintf("%s %s %s %s %s", text1, name1, name2, name3, text2)
				if b.in.Option.Edit {
					b.client.Tg.EditMessageTextKey(b.in.Config.TgChannel, u.User1.Tgmesid, text, b.in.Lvlkz)
				} else if !b.in.Option.Edit {
					mesidTg := b.client.Tg.SendEmded(b.in.Lvlkz, b.in.Config.TgChannel, text)
					b.storage.Update.MesidTgUpdate(ctx, mesidTg, b.in.Lvlkz, b.in.Config.CorpName)
					b.client.Tg.DelMessage(b.in.Config.TgChannel, u.User1.Tgmesid)
				}
				b.wg.Done()
			}()
		}
		if b.in.Config.WaChannel != "" {

		}
	}
	b.wg.Wait()
}
func (b *Bot) QueueAll() {
	if b.debug {
		fmt.Printf("in QueueAll %+v", b.in)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	lvl := b.storage.DbFunc.Queue(ctx, b.in.Config.CorpName)
	lvlk := utils.RemoveDuplicateElementString(lvl)
	if len(lvlk) > 0 {
		for _, corp := range lvlk {
			if corp != "" {
				b.in.Option.Queue = true
				b.in.Lvlkz = corp
				b.QueueLevel()

			}
		}
	} else {
		b.ifTipSendTextDelSecond(b.GetLang("netAktivnuh"), 10)
		b.iftipdelete()
	}

}
