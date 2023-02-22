package bot

import (
	"context"
	"fmt"
	"time"
)

//lang ok

func (b *Bot) RsStart() {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.debug {
		fmt.Println("in RsStart", b.in)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	b.iftipdelete()
	countName, err := b.storage.Count.СountName(ctx, b.in.Name, b.in.Lvlkz, b.in.Config.CorpName)
	if err != nil {
		return
	}
	if countName == 0 {
		text := b.GetLang("prinuditelniStartDostupen")
		b.ifTipSendTextDelSecond(text, 10)
	} else if countName == 1 {
		numberkz, err1 := b.storage.DbFunc.NumberQueueLvl(ctx, b.in.Lvlkz, b.in.Config.CorpName)
		if err1 != nil {
			return
		}
		count, err2 := b.storage.Count.CountQueue(ctx, b.in.Lvlkz, b.in.Config.CorpName)
		if err2 != nil {
			return
		}
		//var name1, name2, name3 string
		dsmesid := ""
		tgmesid := 0
		wamesid := ""
		if count > 0 {
			u := b.storage.DbFunc.ReadAll(ctx, b.in.Lvlkz, b.in.Config.CorpName)
			textEvent, numkzEvent := b.EventText()
			numberevent := b.storage.Event.NumActiveEvent(b.in.Config.CorpName)
			if numberevent > 0 {
				numberkz = numkzEvent
			}
			if count == 1 {
				if b.in.Config.DsChannel != "" {
					b.wg.Add(1)
					go func() {
						name1, _, _, _ := b.nameMention(u, ds)
						text := fmt.Sprintf("%s%s (%d) %s \n\n1. %s\n%s %s",
							b.GetLang("ocheredKz"), b.in.Lvlkz, numberkz,
							b.GetLang("bilaZapushenaNe"), name1, b.GetLang("Vigru"), textEvent)

						if b.in.Tip == ds {
							dsmesid = b.client.Ds.SendWebhook(text, "КзБот", b.in.Config.DsChannel, b.in.Config.Guildid, b.in.Ds.Avatar)

						} else {
							dsmesid = b.client.Ds.Send(b.in.Config.DsChannel, text)
						}

						go b.client.Ds.DeleteMessage(b.in.Config.DsChannel, u.User1.Dsmesid)
						b.storage.Update.MesidDsUpdate(ctx, dsmesid, b.in.Lvlkz, b.in.Config.CorpName)
						b.wg.Done()
					}()

				}
				if b.in.Config.TgChannel != 0 {
					b.wg.Add(1)
					go func() {
						name1, _, _, _ := b.nameMention(u, tg)
						go b.client.Tg.DelMessage(b.in.Config.TgChannel, u.User1.Tgmesid)
						text := fmt.Sprintf("%s%s (%d) %s \n\n1. %s\n%s %s",
							b.GetLang("ocheredKz"), b.in.Lvlkz, numberkz,
							b.GetLang("bilaZapushenaNe"), name1, b.GetLang("Vigru"), textEvent)
						tgmesid = b.client.Tg.SendChannel(b.in.Config.TgChannel, text)
						b.storage.Update.MesidTgUpdate(ctx, tgmesid, b.in.Lvlkz, b.in.Config.CorpName)
						b.wg.Done()
					}()

				}
			} else if count == 2 {
				if b.in.Config.DsChannel != "" { //discord
					b.wg.Add(1)
					go func() {
						name1, name2, _, _ := b.nameMention(u, ds)
						text1 := fmt.Sprintf("%s%s (%d) %s \n\n",
							b.GetLang("ocheredKz"), b.in.Lvlkz, numberkz, b.GetLang("bilaZapushenaNe"))
						text2 := fmt.Sprintf("%s\n%s\n%s %s", name1, name2, b.GetLang("Vigru"), textEvent)
						text := text1 + text2
						if b.in.Tip == ds {
							dsmesid = b.client.Ds.SendWebhook(text, "КзБот", b.in.Config.DsChannel, b.in.Config.Guildid, b.in.Ds.Avatar)
						} else {
							dsmesid = b.client.Ds.Send(b.in.Config.DsChannel, text)
						}
						go b.client.Ds.DeleteMessage(b.in.Config.DsChannel, u.User1.Dsmesid)
						b.storage.Update.MesidDsUpdate(ctx, dsmesid, b.in.Lvlkz, b.in.Config.CorpName)
						b.wg.Done()
					}()

				}
				if b.in.Config.TgChannel != 0 { //telegram
					b.wg.Add(1)
					go func() {
						name1, name2, _, _ := b.nameMention(u, tg)
						go b.client.Tg.DelMessage(b.in.Config.TgChannel, u.User1.Tgmesid)
						text1 := fmt.Sprintf("%s%s (%d) %s \n\n",
							b.GetLang("ocheredKz"), b.in.Lvlkz, numberkz, b.GetLang("bilaZapushenaNe"))
						text2 := fmt.Sprintf("%s\n%s\n%s %s", name1, name2, b.GetLang("Vigru"), textEvent)
						text := text1 + text2
						tgmesid = b.client.Tg.SendChannel(b.in.Config.TgChannel, text)
						b.storage.Update.MesidTgUpdate(ctx, tgmesid, b.in.Lvlkz, b.in.Config.CorpName)
						b.wg.Done()
					}()

				}
			} else if count == 3 {
				if b.in.Config.DsChannel != "" { //discord
					b.wg.Add(1)
					go func() {
						name1, name2, name3, _ := b.nameMention(u, ds)
						text := fmt.Sprintf("%s%s (%d) %s \n\n%s\n%s\n%s\n%s %s",
							b.GetLang("ocheredKz"), b.in.Lvlkz, numberkz, b.GetLang("bilaZapushenaNe"),
							name1, name2, name3, b.GetLang("Vigru"), textEvent)
						if b.in.Tip == ds {
							dsmesid = b.client.Ds.SendWebhook(text, "КзБот", b.in.Config.DsChannel, b.in.Config.Guildid, b.in.Ds.Avatar)
						} else {
							dsmesid = b.client.Ds.Send(b.in.Config.DsChannel, text)
						}
						go b.client.Ds.DeleteMessage(b.in.Config.DsChannel, u.User1.Dsmesid)
						b.storage.Update.MesidDsUpdate(ctx, dsmesid, b.in.Lvlkz, b.in.Config.CorpName)
						b.wg.Done()
					}()
				}
				if b.in.Config.TgChannel != 0 { //telegram
					b.wg.Add(1)
					go func() {
						name1, name2, name3, _ := b.nameMention(u, tg)
						go b.client.Tg.DelMessage(b.in.Config.TgChannel, u.User1.Tgmesid)
						text := fmt.Sprintf("%s%s (%d) %s \n\n%s\n%s\n%s\n%s %s",
							b.GetLang("ocheredKz"), b.in.Lvlkz, numberkz, b.GetLang("bilaZapushenaNe"),
							name1, name2, name3, b.GetLang("Vigru"), textEvent)
						tgmesid = b.client.Tg.SendChannel(b.in.Config.TgChannel, text)
						b.storage.Update.MesidTgUpdate(ctx, tgmesid, b.in.Lvlkz, b.in.Config.CorpName)
						b.wg.Done()
					}()

				}
			}
			b.wg.Wait()
			b.storage.Update.UpdateCompliteRS(ctx, b.in.Lvlkz, dsmesid, tgmesid, wamesid, numberkz, numberevent, b.in.Config.CorpName)
			user := []string{u.User1.Name, u.User2.Name, u.User3.Name, b.in.Name}
			b.elseChat(user)
		}
	}
}
func (b *Bot) Pl30() {
	if b.debug {
		fmt.Println("in Pl30", b.in)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	countName := b.storage.Count.CountNameQueue(ctx, b.in.Name)
	text := ""
	if countName == 0 {
		text = b.in.NameMention + b.GetLang("tiNeVOcheredi")
	} else if countName > 0 {
		timedown := b.storage.DbFunc.P30Pl(ctx, b.in.Lvlkz, b.in.Config.CorpName, b.in.Name)
		if timedown >= 150 {
			text = fmt.Sprintf("%s %s %d %s",
				b.in.NameMention, b.GetLang("maksimalnoeVremya"), timedown, b.GetLang("min."))
		} else {
			text = b.in.NameMention + b.GetLang("vremyaObnovleno")
			b.storage.DbFunc.UpdateTimedown(ctx, b.in.Lvlkz, b.in.Config.CorpName, b.in.Name)
			b.in.Option.Pl30 = true
			b.in.Option.Edit = true
			b.QueueLevel()
		}
	}
	b.ifTipSendTextDelSecond(text, 20)
}
