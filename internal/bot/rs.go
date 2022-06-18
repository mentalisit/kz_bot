package bot

import (
	"fmt"

	corpsConfig "kz_bot/internal/clients/corpConfig"
	"kz_bot/internal/models"
)

func (b *Bot) RsPlus() {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	go b.iftipdelete()
	if b.Db.СountName(b.in.Name, b.in.Lvlkz, b.in.Config.CorpName) == 1 { //проверяем есть ли игрок в очереди
		b.ifTipSendMentionText(" ты уже в очереди")
	} else {
		countQueue := b.Db.CountQueue(b.in.Lvlkz, b.in.Config.CorpName)                    //проверяем, есть ли кто-то в очереди
		numkzN := b.Db.CountNumberNameActive1(b.in.Lvlkz, b.in.Config.CorpName, b.in.Name) //проверяем количество боёв по уровню кз игрока
		numkzL := b.Db.NumberQueueLvl(b.in.Lvlkz, b.in.Config.CorpName)                    //проверяем какой номер боя определенной красной звезды

		dsmesid := ""
		tgmesid := 0
		wamesid := ""

		if countQueue == 0 {
			if b.in.Config.DsChannel != "" {
				name1 := fmt.Sprintf("%s  🕒  %s  (%d)", b.emReadName(b.in.Name, ds), b.in.Timekz, numkzN)
				name2 := ""
				name3 := ""
				name4 := ""
				lvlk := b.Ds.RoleToIdPing(b.in.Lvlkz, b.in.Ds.Guildid)
				emb := b.Ds.EmbedDS(name1, name2, name3, name4, lvlk, numkzL)
				dsmesid = b.Ds.SendComplexContent(b.in.Config.DsChannel, b.in.Name+" запустил очередь "+lvlk)
				b.Ds.EditComplex(dsmesid, b.in.Config.DsChannel, emb)
				b.Ds.AddEnojiRsQueue(b.in.Config.DsChannel, dsmesid)
			}
			if b.in.Config.TgChannel != 0 {
				text := fmt.Sprintf("Очередь кз%s (%d)\n1. %s - %sмин. (%d) \n\n%s++ - принудительный старт",
					b.in.Lvlkz, numkzL, b.in.Name, b.in.Timekz, numkzN, b.in.Lvlkz)
				tgmesid = b.Tg.SendEmded(b.in.Lvlkz, b.in.Config.TgChannel, text)
				b.SubscribePing(1)
			}
			if b.in.Config.WaChannel != "" {
				//Тут будет логика ватса
			}

			b.Db.InsertQueue(dsmesid, wamesid, b.in.Config.CorpName, b.in.Name, b.in.NameMention, b.in.Tip, b.in.Lvlkz, b.in.Timekz, tgmesid, numkzN)

		} else if countQueue == 1 {
			u := b.Db.ReadAll(b.in.Lvlkz, b.in.Config.CorpName)
			dsmesid = u.User1.Dsmesid

			if b.in.Config.DsChannel != "" {
				name1 := fmt.Sprintf("%s  🕒  %d  (%d)", b.emReadName(u.User1.Name, ds), u.User1.Timedown, u.User1.Numkzn)
				name2 := fmt.Sprintf("%s  🕒  %s  (%d)", b.emReadName(b.in.Name, ds), b.in.Timekz, numkzN)
				name3 := ""
				name4 := ""
				lvlk := b.Ds.RoleToIdPing(b.in.Lvlkz, b.in.Ds.Guildid)
				emb := b.Ds.EmbedDS(name1, name2, name3, name4, lvlk, numkzL)
				text := lvlk + " 2/4 " + b.in.Name + " присоединился к очереди"
				go b.Ds.SendChannelDelSecond(b.in.Config.DsChannel, text, 10)
				b.Ds.EditComplex(u.User1.Dsmesid, b.in.Config.DsChannel, emb)
			}
			if b.in.Config.TgChannel != 0 {
				text1 := fmt.Sprintf("Очередь кз%s (%d)\n", b.in.Lvlkz, numkzL)
				name1 := fmt.Sprintf("1. %s - %dмин. (%d) \n", u.User1.Name, u.User1.Timedown, u.User1.Numkzn)
				name2 := fmt.Sprintf("2. %s - %sмин. (%d) \n", b.in.Name, b.in.Timekz, numkzN)
				text2 := fmt.Sprintf("\n%s++ - принудительный старт", b.in.Lvlkz)
				text := fmt.Sprintf("%s %s %s %s", text1, name1, name2, text2)
				tgmesid = b.Tg.SendEmded(b.in.Lvlkz, b.in.Config.TgChannel, text)
				go b.Tg.DelMessage(b.in.Config.TgChannel, u.User1.Tgmesid)
				b.Db.MesidTgUpdate(tgmesid, b.in.Lvlkz, b.in.Config.CorpName)
			}
			if b.in.Config.WaChannel != "" {
				//Тут будет логика ватса
			}
			b.Db.InsertQueue(dsmesid, wamesid, b.in.Config.CorpName, b.in.Name, b.in.NameMention, b.in.Tip, b.in.Lvlkz, b.in.Timekz, tgmesid, numkzN)

		} else if countQueue == 2 {
			u := b.Db.ReadAll(b.in.Lvlkz, b.in.Config.CorpName)
			dsmesid = u.User1.Dsmesid

			if b.in.Config.DsChannel != "" {
				name1 := fmt.Sprintf("%s  🕒  %d  (%d)", b.emReadName(u.User1.Name, b.in.Tip), u.User1.Timedown, u.User1.Numkzn)
				name2 := fmt.Sprintf("%s  🕒  %d  (%d)", b.emReadName(u.User2.Name, b.in.Tip), u.User2.Timedown, u.User2.Numkzn)
				name3 := fmt.Sprintf("%s  🕒  %s  (%d)", b.emReadName(b.in.Name, b.in.Tip), b.in.Timekz, numkzN)
				name4 := ""
				lvlk := b.Ds.RoleToIdPing(b.in.Lvlkz, b.in.Config.Config.Guildid)
				lvlk3 := b.Ds.RoleToIdPing(b.in.Lvlkz+"+", b.in.Config.Config.Guildid)
				emb := b.Ds.EmbedDS(name1, name2, name3, name4, lvlk, numkzL)
				text := lvlk + " 3/4 " + b.in.Name + " присоединился к очереди " + lvlk3 + " нужен еще один для фулки"
				go b.Ds.SendChannelDelSecond(b.in.Config.DsChannel, text, 10)
				b.Ds.EditComplex(u.User1.Dsmesid, b.in.Config.DsChannel, emb)
			}
			if b.in.Config.TgChannel != 0 {
				text1 := fmt.Sprintf("Очередь кз%s (%d)\n", b.in.Lvlkz, numkzL)
				name1 := fmt.Sprintf("1. %s - %dмин. (%d) \n", u.User1.Name, u.User1.Timedown, u.User1.Numkzn)
				name2 := fmt.Sprintf("2. %s - %dмин. (%d) \n", u.User2.Name, u.User2.Timedown, u.User2.Numkzn)
				name3 := fmt.Sprintf("3. %s - %sмин. (%d) \n", b.in.Name, b.in.Timekz, numkzN)
				text2 := fmt.Sprintf("\n%s++ - принудительный старт", b.in.Lvlkz)
				text := fmt.Sprintf("%s %s %s %s %s", text1, name1, name2, name3, text2)
				tgmesid = b.Tg.SendEmded(b.in.Lvlkz, b.in.Config.TgChannel, text)
				go b.Tg.DelMessage(b.in.Config.TgChannel, u.User1.Tgmesid)
				b.Db.MesidTgUpdate(tgmesid, b.in.Lvlkz, b.in.Config.CorpName)
				b.SubscribePing(3)
			}
			if b.in.Config.WaChannel != "" {
				//Тут будет логика ватса
			}
			b.Db.InsertQueue(dsmesid, wamesid, b.in.Config.CorpName, b.in.Name, b.in.NameMention, b.in.Tip, b.in.Lvlkz, b.in.Timekz, tgmesid, numkzN)

		} else if countQueue == 3 {
			u := b.Db.ReadAll(b.in.Lvlkz, b.in.Config.CorpName)
			textEvent, numkzEvent := "event(in)", 0
			numberevent := 0 //qweryNumevent1(in) //получаем номер ивета если он активен
			if numberevent > 0 {
				numkzL = numkzEvent
			}
			var name1, name2, name3, name4 string

			dsmesid = u.User1.Dsmesid

			if b.in.Config.DsChannel != "" {
				if u.User1.Tip == "ds" {
					name1 = u.User1.Mention
				} else {
					name1 = u.User1.Name
				}
				if u.User2.Tip == "ds" {
					name2 = u.User2.Mention
				} else {
					name2 = u.User2.Name
				}
				if u.User3.Tip == "ds" {
					name3 = u.User3.Mention
				} else {
					name3 = u.User3.Name
				}
				if b.in.Tip == "ds" {
					name4 = b.in.NameMention
				} else {
					name4 = b.in.Name
				}
				go b.Ds.DeleteMessage(b.in.Config.DsChannel, u.User1.Dsmesid)
				go b.Ds.SendChannelDelSecond(b.in.Config.DsChannel, " 4/4 "+b.in.Name+" присоединился к очереди", 10)
				text := fmt.Sprintf("4/4 Очередь КЗ%s сформирована\n %s\n %s\n %s\n %s \nВ ИГРУ %s",
					b.in.Lvlkz, b.emReadName(name1, ds), b.emReadName(name2, ds), b.emReadName(name3, ds), b.emReadName(name4, ds), textEvent)
				dsmesid = b.Ds.Send(b.in.Config.DsChannel, text)
				b.Db.MesidDsUpdate(dsmesid, b.in.Lvlkz, b.in.Config.DsChannel)
			}
			if b.in.Config.TgChannel != 0 {
				if u.User1.Tip == "tg" {
					name1 = u.User1.Mention
				} else {
					name1 = u.User1.Name
				}
				if u.User2.Tip == "tg" {
					name2 = u.User2.Mention
				} else {
					name2 = u.User2.Name
				}
				if u.User3.Tip == "tg" {
					name3 = u.User3.Mention
				} else {
					name3 = u.User3.Name
				}
				if b.in.Tip == "tg" {
					name4 = b.in.NameMention
				} else {
					name4 = b.in.Name
				}
				go b.Tg.DelMessage(b.in.Config.TgChannel, u.User1.Tgmesid)
				go b.Tg.SendChannelDelSecond(b.in.Config.TgChannel, b.in.Name+" закрыл очередь кз"+b.in.Lvlkz, 10)
				text := fmt.Sprintf("Очередь КЗ%s сформирована\n%s\n%s\n%s\n%s\n В ИГРУ \n%s",
					b.in.Lvlkz, name1, name2, name3, name4, textEvent)
				tgmesid = b.Tg.SendChannel(b.in.Config.TgChannel, text)
				b.Db.MesidTgUpdate(tgmesid, b.in.Lvlkz, b.in.Config.CorpName)
			}
			if b.in.Config.WaChannel != "" {
				//Тут будет логика ватса
			}

			b.Db.InsertQueue(dsmesid, wamesid, b.in.Config.CorpName, b.in.Name, b.in.NameMention, b.in.Tip, b.in.Lvlkz, b.in.Timekz, tgmesid, numkzN)
			b.Db.UpdateCompliteRS(b.in.Lvlkz, dsmesid, tgmesid, wamesid, numkzL, numberevent, b.in.Config.CorpName)

			//проверка есть ли игрок в других чатах
			go b.elseChat(u, b.in.Name)

		}

	}
}
func (b *Bot) RsMinus() {
	b.Mutex.Lock()
	b.callbackNo()
	CountNames := b.Db.СountName(b.in.Name, b.in.Lvlkz, b.in.Config.CorpName) //проверяем есть ли игрок в очереди
	if CountNames == 0 {
		b.ifTipSendMentionText(" ты не в очереди")
	} else if CountNames > 0 {
		//чтение айди очечреди
		u := b.Db.ReadAll(b.in.Lvlkz, b.in.Config.CorpName)
		//удаление с БД
		b.Db.DeleteQueue(b.in.Name, b.in.Lvlkz, b.in.Config.CorpName)
		//проверяем очередь
		countQueue := b.Db.CountQueue(b.in.Lvlkz, b.in.Config.CorpName)
		//numkzL := numberQueueLvl(in, lvlkz) + 1
		if b.in.Config.DsChannel != "" {
			go b.Ds.SendChannelDelSecond(b.in.Config.DsChannel, b.in.Name+" покинул очередь", 10)
			if countQueue == 0 {
				go b.Ds.SendChannelDelSecond(b.in.Config.DsChannel, "Очередь КЗ была удалена.", 10)
				b.Ds.DeleteMessage(b.in.Config.DsChannel, u.User1.Dsmesid)
			}
		}
		if b.in.Config.TgChannel != 0 {
			go b.Tg.SendChannelDelSecond(b.in.Config.TgChannel, b.in.Name+" покинул очередь", 10)
			if countQueue == 0 {
				go b.Tg.SendChannelDelSecond(b.in.Config.TgChannel, "Очередь КЗ была удалена.", 10)
				b.Tg.DelMessage(b.in.Config.TgChannel, u.User1.Tgmesid)
			}
		}
		if b.in.Config.WaChannel != "" {
			//тут логика ватса
		}
		b.Mutex.Unlock()
		if countQueue > 0 {
			b.QueueLevel()
		}
	}
}
func (b *Bot) QueueLevel() {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	b.callbackNo()
	count := b.Db.CountQueue(b.in.Lvlkz, b.in.Config.CorpName)
	numberLvl := b.Db.NumberQueueLvl(b.in.Lvlkz, b.in.Config.CorpName)
	// совподения количество  условие
	if count == 0 && !b.in.Option.Queue {
		text := "Очередь КЗ " + b.in.Lvlkz + " пуста "
		b.ifTipSendTextDelSecond(text, 10)
	} else if b.in.Option.Queue && count == 0 {
		b.ifTipSendTextDelSecond("Нет активных очередей ", 10)

	} else if count == 1 {
		u := b.Db.ReadAll(b.in.Lvlkz, b.in.Config.CorpName)
		if b.in.Config.DsChannel != "" {
			name1 := fmt.Sprintf("%s  🕒  %d  (%d)", b.emReadName(u.User1.Name, ds), u.User1.Timedown, u.User1.Numkzn)
			name2 := ""
			name3 := ""
			name4 := ""
			lvlk := b.Ds.RoleToIdPing(b.in.Lvlkz, b.in.Config.Config.Guildid)
			emb := b.Ds.EmbedDS(name1, name2, name3, name4, lvlk, numberLvl)
			if b.in.Option.Edit {
				b.Ds.EditComplex(u.User1.Dsmesid, b.in.Config.DsChannel, emb)
			} else if !b.in.Option.Edit {
				b.Ds.DeleteMessage(b.in.Config.DsChannel, u.User1.Dsmesid)
				dsmesid := b.Ds.SendComplex(b.in.Config.DsChannel, emb)

				b.Ds.AddEnojiRsQueue(b.in.Config.DsChannel, dsmesid)
				b.Db.MesidDsUpdate(dsmesid, b.in.Lvlkz, b.in.Config.CorpName)
			}
		}
		if b.in.Config.TgChannel != 0 {
			text1 := fmt.Sprintf("Очередь кз%s (%d)\n", b.in.Lvlkz, numberLvl)
			name1 := fmt.Sprintf("1. %s - %dмин. (%d) \n", u.User1.Name, u.User1.Timedown, u.User1.Numkzn)
			text2 := fmt.Sprintf("\n%s++ - принудительный старт", b.in.Lvlkz)
			text := fmt.Sprintf("%s %s %s", text1, name1, text2)
			if b.in.Option.Edit {
				b.Tg.EditMessageTextKey(b.in.Config.TgChannel, u.User1.Tgmesid, text, b.in.Lvlkz)
			} else if !b.in.Option.Edit {
				mesidTg := b.Tg.SendEmded(b.in.Lvlkz, b.in.Config.TgChannel, text)
				b.Db.MesidTgUpdate(mesidTg, b.in.Lvlkz, b.in.Config.CorpName)
				go b.Tg.DelMessage(b.in.Config.TgChannel, u.User1.Tgmesid)
			}
		}
		if b.in.Config.WaChannel != "" {

		}
	} else if count == 2 {
		u := b.Db.ReadAll(b.in.Lvlkz, b.in.Config.CorpName)

		if b.in.Config.DsChannel != "" {
			name1 := fmt.Sprintf("%s  🕒  %d  (%d)", b.emReadName(u.User1.Name, ds), u.User1.Timedown, u.User1.Numkzn)
			name2 := fmt.Sprintf("%s  🕒  %d  (%d)", b.emReadName(u.User2.Name, ds), u.User2.Timedown, u.User2.Numkzn)
			name3 := ""
			name4 := ""
			lvlk := b.Ds.RoleToIdPing(b.in.Lvlkz, b.in.Config.Config.Guildid)
			emb := b.Ds.EmbedDS(name1, name2, name3, name4, lvlk, numberLvl)
			if b.in.Option.Edit {
				b.Ds.EditComplex(u.User1.Dsmesid, b.in.Config.DsChannel, emb)
			} else if !b.in.Option.Edit {
				b.Ds.DeleteMessage(b.in.Config.DsChannel, u.User1.Dsmesid)
				dsmesid := b.Ds.SendComplex(b.in.Config.DsChannel, emb)

				b.Ds.AddEnojiRsQueue(b.in.Config.DsChannel, dsmesid)
				b.Db.MesidDsUpdate(dsmesid, b.in.Lvlkz, b.in.Config.CorpName)
			}
		}
		if b.in.Config.TgChannel != 0 {
			text1 := fmt.Sprintf("Очередь кз%s (%d)\n", b.in.Lvlkz, numberLvl)
			name1 := fmt.Sprintf("1. %s - %dмин. (%d) \n", u.User1.Name, u.User1.Timedown, u.User1.Numkzn)
			name2 := fmt.Sprintf("2. %s - %dмин. (%d) \n", u.User2.Name, u.User2.Timedown, u.User2.Numkzn)
			text2 := fmt.Sprintf("\n%s++ - принудительный старт", b.in.Lvlkz)
			text := fmt.Sprintf("%s %s %s %s", text1, name1, name2, text2)
			if b.in.Option.Edit {
				b.Tg.EditMessageTextKey(b.in.Config.TgChannel, u.User1.Tgmesid, text, b.in.Lvlkz)
			} else if !b.in.Option.Edit {
				mesidTg := b.Tg.SendEmded(b.in.Lvlkz, b.in.Config.TgChannel, text)
				b.Db.MesidTgUpdate(mesidTg, b.in.Lvlkz, b.in.Config.CorpName)
				go b.Tg.DelMessage(b.in.Config.TgChannel, u.User1.Tgmesid)
			}
		}
		if b.in.Config.WaChannel != "" {

		}
	} else if count == 3 {
		u := b.Db.ReadAll(b.in.Lvlkz, b.in.Config.CorpName)

		if b.in.Config.DsChannel != "" {
			name1 := fmt.Sprintf("%s  🕒  %d  (%d)", b.emReadName(u.User1.Name, ds), u.User1.Timedown, u.User1.Numkzn)
			name2 := fmt.Sprintf("%s  🕒  %d  (%d)", b.emReadName(u.User2.Name, ds), u.User2.Timedown, u.User2.Numkzn)
			name3 := fmt.Sprintf("%s  🕒  %d  (%d)", b.emReadName(u.User3.Name, ds), u.User3.Timedown, u.User3.Numkzn)
			name4 := ""
			lvlk := b.Ds.RoleToIdPing(b.in.Lvlkz, b.in.Config.Config.Guildid)
			emb := b.Ds.EmbedDS(name1, name2, name3, name4, lvlk, numberLvl)
			if b.in.Option.Edit {
				b.Ds.EditComplex(u.User1.Dsmesid, b.in.Config.DsChannel, emb)
			} else if !b.in.Option.Edit {
				b.Ds.DeleteMessage(b.in.Config.DsChannel, u.User1.Dsmesid)
				dsmesid := b.Ds.SendComplex(b.in.Config.DsChannel, emb)

				b.Ds.AddEnojiRsQueue(b.in.Config.DsChannel, dsmesid)
				b.Db.MesidDsUpdate(dsmesid, b.in.Lvlkz, b.in.Config.CorpName)
			}
		}
		if b.in.Config.TgChannel != 0 {
			text1 := fmt.Sprintf("Очередь кз%s (%d)\n", b.in.Lvlkz, numberLvl)
			name1 := fmt.Sprintf("1. %s - %dмин. (%d) \n", u.User1.Name, u.User1.Timedown, u.User1.Numkzn)
			name2 := fmt.Sprintf("2. %s - %dмин. (%d) \n", u.User2.Name, u.User2.Timedown, u.User2.Numkzn)
			name3 := fmt.Sprintf("3. %s - %dмин. (%d) \n", u.User3.Name, u.User3.Timedown, u.User3.Numkzn)
			text2 := fmt.Sprintf("\n%s++ - принудительный старт", b.in.Lvlkz)
			text := fmt.Sprintf("%s %s %s %s %s", text1, name1, name2, name3, text2)
			if b.in.Option.Edit {
				b.Tg.EditMessageTextKey(b.in.Config.TgChannel, u.User1.Tgmesid, text, b.in.Lvlkz)
			} else if !b.in.Option.Edit {
				mesidTg := b.Tg.SendEmded(b.in.Lvlkz, b.in.Config.TgChannel, text)
				b.Db.MesidTgUpdate(mesidTg, b.in.Lvlkz, b.in.Config.CorpName)
				go b.Tg.DelMessage(b.in.Config.TgChannel, u.User1.Tgmesid)
			}
		}
		if b.in.Config.WaChannel != "" {

		}
	}
}
func (b *Bot) QueueAll() {
	lvl := b.Db.Queue(b.in.Config.CorpName)
	lvlk := b.removeDuplicateElementString(lvl)
	var aa []string
	for _, corp := range lvlk {
		skip := false
		for _, u := range aa {
			if corp == u {
				skip = true
				break
			}
		}
		if !skip {
			b.in.Option.Queue = true
			b.in.Lvlkz = corp
			b.QueueLevel()
		}
	}
}
func (b *Bot) RsStart() {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	b.callbackNo()
	countName := b.Db.СountName(b.in.Name, b.in.Lvlkz, b.in.Config.CorpName)
	if countName == 0 {
		text := "Принудительный старт доступен участникам очереди."
		if b.in.Tip == ds {
			b.Ds.SendChannelDelSecond(b.in.Config.DsChannel, text, 10)
		} else if b.in.Tip == "tg" {
			b.Tg.SendChannelDelSecond(b.in.Config.TgChannel, text, 10)
		}
	} else if countName == 1 {
		numberkz := b.Db.NumberQueueLvl(b.in.Lvlkz, b.in.Config.CorpName)
		count := b.Db.CountQueue(b.in.Lvlkz, b.in.Config.CorpName)
		var name1, name2, name3 string
		dsmesid := ""
		tgmesid := 0
		wamesid := ""
		if count > 0 {
			u := b.Db.ReadAll(b.in.Lvlkz, b.in.Config.CorpName)
			textEvent, numkzEvent := "event(in)", 0
			numberevent := 0 //qweryNumevent1(in)
			if numberevent > 0 {
				numberkz = numkzEvent
			}
			if count == 1 {
				if b.in.Config.DsChannel != "" {
					if u.User1.Tip == "ds" {
						name1 = u.User1.Mention
					} else {
						name1 = u.User1.Name
					}
					text := fmt.Sprintf("Очередь кз%s (%d) была \nзапущена не полной \n\n1. %s\nВ игру %s",
						b.in.Lvlkz, numberkz, name1, textEvent)
					dsmesid = b.Ds.Send(b.in.Config.DsChannel, text)
					go b.Ds.DeleteMessage(b.in.Config.DsChannel, u.User1.Dsmesid)
					b.Db.MesidDsUpdate(dsmesid, b.in.Lvlkz, b.in.Config.CorpName)
				}
				if b.in.Config.TgChannel != 0 {
					if u.User1.Tip == "tg" {
						name1 = u.User1.Mention
					} else {
						name1 = u.User1.Name
					}
					go b.Tg.DelMessage(b.in.Config.TgChannel, u.User1.Tgmesid)
					text := fmt.Sprintf("Очередь кз%s (%d) была \nзапущена не полной \n\n1. %s\nВ игру %s",
						b.in.Lvlkz, numberkz, name1, textEvent)
					tgmesid = b.Tg.SendChannel(b.in.Config.TgChannel, text)
					b.Db.MesidTgUpdate(tgmesid, b.in.Lvlkz, b.in.Config.CorpName)
				}
				b.Db.UpdateCompliteRS(b.in.Lvlkz, dsmesid, tgmesid, wamesid, numberkz, numberevent, b.in.Config.CorpName)
				b.elseChat(u, b.in.Name)
			} else if count == 2 {
				if b.in.Config.DsChannel != "" { //discord
					if u.User1.Tip == "ds" {
						name1 = u.User1.Mention
					} else {
						name1 = u.User1.Name
					}
					if u.User2.Tip == "ds" {
						name2 = u.User2.Mention
					} else {
						name2 = u.User2.Name
					}
					text1 := fmt.Sprintf("Очередь кз%s (%d) была \nзапущена не полной \n", b.in.Lvlkz, numberkz)
					text2 := fmt.Sprintf("\n%s %s\nВ игру %s", name1, name2, textEvent)
					text := text1 + text2
					dsmesid = b.Ds.Send(b.in.Config.DsChannel, text)
					go b.Ds.DeleteMessage(b.in.Config.DsChannel, u.User1.Dsmesid)
					b.Db.MesidDsUpdate(dsmesid, b.in.Lvlkz, b.in.Config.CorpName)
				}
				if b.in.Config.TgChannel != 0 { //telegram
					if u.User1.Tip == "tg" {
						name1 = u.User1.Mention
					} else {
						name1 = u.User1.Name
					}
					if u.User2.Tip == "tg" {
						name2 = u.User2.Mention
					} else {
						name2 = u.User2.Name
					}
					go b.Tg.DelMessage(b.in.Config.TgChannel, u.User1.Tgmesid)
					text1 := fmt.Sprintf("Очередь кз%s (%d) была \nзапущена не полной \n", b.in.Lvlkz, numberkz)
					text2 := fmt.Sprintf("\n%s %s\nВ игру %s", name1, name2, textEvent)
					text := text1 + text2
					tgmesid = b.Tg.SendChannel(b.in.Config.TgChannel, text)
					b.Db.MesidTgUpdate(tgmesid, b.in.Lvlkz, b.in.Config.CorpName)
				}
				b.Db.UpdateCompliteRS(b.in.Lvlkz, dsmesid, tgmesid, wamesid, numberkz, numberevent, b.in.Config.CorpName)
				b.elseChat(u, b.in.Name)
			} else if count == 3 {
				if b.in.Config.DsChannel != "" { //discord
					if u.User1.Tip == "ds" {
						name1 = u.User1.Mention
					} else {
						name1 = u.User1.Name
					}
					if u.User2.Tip == "ds" {
						name2 = u.User2.Mention
					} else {
						name2 = u.User2.Name
					}
					if u.User3.Tip == "ds" {
						name3 = u.User3.Mention
					} else {
						name3 = u.User3.Name
					}
					text := fmt.Sprintf("Очередь кз%s (%d) была \nзапущена не полной \n\n%s %s %s\nВ игру %s",
						b.in.Lvlkz, numberkz, name1, name2, name3, textEvent)
					dsmesid = b.Ds.Send(b.in.Config.DsChannel, text)
					go b.Ds.DeleteMessage(b.in.Config.DsChannel, u.User1.Dsmesid)
					b.Db.MesidDsUpdate(dsmesid, b.in.Lvlkz, b.in.Config.CorpName)
				}
				if b.in.Config.TgChannel != 0 { //telegram
					if u.User1.Tip == "tg" {
						name1 = u.User1.Mention
					} else {
						name1 = u.User1.Name
					}
					if u.User2.Tip == "tg" {
						name2 = u.User2.Mention
					} else {
						name2 = u.User2.Name
					}
					if u.User3.Tip == "tg" {
						name3 = u.User3.Mention
					} else {
						name3 = u.User3.Name
					}
					go b.Tg.DelMessage(b.in.Config.TgChannel, u.User1.Tgmesid)
					text := fmt.Sprintf("Очередь кз%s (%d) была \nзапущена не полной \n\n%s %s %s\nВ игру %s",
						b.in.Lvlkz, numberkz, name1, name2, name3, textEvent)
					tgmesid = b.Tg.SendChannel(b.in.Config.TgChannel, text)
					b.Db.MesidTgUpdate(tgmesid, b.in.Lvlkz, b.in.Config.CorpName)
				}
				b.Db.UpdateCompliteRS(b.in.Lvlkz, dsmesid, tgmesid, wamesid, numberkz, numberevent, b.in.Config.CorpName)
				b.elseChat(u, b.in.Name)
			}
		}
	}
}
func (b *Bot) Pl30() {
	countName := b.Db.CountNameQueue(b.in.Name)
	text := ""
	if countName == 0 {
		text = b.in.NameMention + " ты не в очереди "
	} else if countName > 0 {
		timedown := b.Db.P30Pl(b.in.Lvlkz, b.in.Config.CorpName, b.in.Name)
		if timedown >= 150 {
			text = fmt.Sprintf("%s максимальное время в очереди ограничено на 180 минут\n твое время %d мин.  ",
				b.in.NameMention, timedown)
		} else {
			text = b.in.NameMention + " время обновлено +30"
			b.Db.UpdateTimedown(b.in.Lvlkz, b.in.Config.CorpName, b.in.Name)
			b.in.Option.Callback = true
			b.in.Option.Edit = true
			b.QueueLevel()
		}
	}
	b.ifTipSendTextDelSecond(text, 20)
}
func (b *Bot) Plus() bool {
	b.callbackNo()
	countName := b.Db.СountName(b.in.Name, b.in.Lvlkz, b.in.Config.CorpName)
	message := ""
	ins := true
	if countName == 0 {
		message = b.in.NameMention + " ты не в очереди"
		ins = false
	} else if countName > 0 {
		t := b.Db.UpdateMitutsQueue(b.in.Name, b.in.Config.CorpName)
		if t.Timedown > 3 {
			message = fmt.Sprintf("%s рановато плюсик жмешь, ты в очереди на кз%s будешь еще %dмин",
				t.Mention, t.Lvlkz, t.Timedown)
		} else if t.Timedown <= 3 {
			message = t.Mention + " время обновлено "
			b.in.Lvlkz = t.Lvlkz
			b.QueueLevel()
		}
	}
	if b.in.Tip == "ds" {
		go b.Ds.SendChannelDelSecond(b.in.Config.DsChannel, message, 10)
	} else if b.in.Tip == "tg" {
		go b.Tg.SendChannelDelSecond(b.in.Config.TgChannel, message, 10)
	}

	return ins
}
func (b *Bot) Minus() bool {
	b.callbackNo()
	message := ""
	countNames := b.Db.СountName(b.in.Name, b.in.Lvlkz, b.in.Config.CorpName)
	if countNames == 0 {
		message = b.in.NameMention + " ты не в очереди"
		return false
	} else if countNames > 0 {
		t := b.Db.UpdateMitutsQueue(b.in.Name, b.in.Config.CorpName)
		if t.Name == b.in.Name && t.Timedown > 3 {
			message = fmt.Sprintf("%s рановато минус жмешь, ты в очереди на кз%s будешь еще %dмин",
				t.Mention, t.Lvlkz, t.Timedown)
		} else if t.Name == b.in.Name && t.Timedown <= 3 {
			b.in.Lvlkz = t.Lvlkz
			b.RsMinus()
		}
	}
	if b.in.Tip == "ds" {
		go b.Ds.SendChannelDelSecond(b.in.Config.DsChannel, message, 10)
	} else if b.in.Tip == "tg" {
		go b.Tg.SendChannelDelSecond(b.in.Config.TgChannel, message, 10)
	}
	return true
}
func (b *Bot) Subscribe(tipPing int) {
	if b.in.Tip == "ds" {
		go b.Ds.DeleteMessage(b.in.Config.DsChannel, b.in.Ds.Mesid)
		argRoles := "кз" + b.in.Lvlkz
		if tipPing == 3 {
			argRoles = "кз" + b.in.Lvlkz + "+"
		}
		text := b.Ds.Subscribe(b.in.Ds.Nameid, argRoles, b.in.Config.Config.Guildid)
		b.Ds.SendChannelDelSecond(b.in.Config.DsChannel, text, 10)

	} else if b.in.Tip == "tg" {
		go b.Tg.DelMessage(b.in.Config.TgChannel, b.in.Tg.Mesid)
		//проверка активной подписки
		counts := b.Db.CheckSubscribe(b.in.Name, b.in.Lvlkz, b.in.Config.TgChannel, tipPing)
		if counts == 1 {
			text := fmt.Sprintf("%s ты уже подписан на кз%s %d/4\n для добавления в очередь напиши %s+",
				b.in.NameMention, b.in.Lvlkz, tipPing, b.in.Lvlkz)
			go b.Tg.SendChannelDelSecond(b.in.Config.TgChannel, text, 60)
		} else {
			//добавление в оочередь пинга
			b.Db.Subscribe(b.in.Name, b.in.NameMention, b.in.Lvlkz, tipPing, b.in.Config.TgChannel)
			text := fmt.Sprintf("%s вы подписались на пинг кз%s %d/4 \n для добавления в очередь напиши %s+",
				b.in.NameMention, b.in.Lvlkz, tipPing, b.in.Lvlkz)
			go b.Tg.SendChannelDelSecond(b.in.Config.TgChannel, text, 10)
		}
	}
}
func (b *Bot) Unsubscribe(tipPing int) {
	if b.in.Tip == "ds" {
		go b.Ds.DeleteMessage(b.in.Config.DsChannel, b.in.Ds.Mesid)
		argRoles := "кз" + b.in.Lvlkz
		if tipPing == 3 {
			argRoles = "кз" + b.in.Lvlkz + "+"
		}
		b.Ds.Unsubscribe(b.in.Ds.Nameid, argRoles, b.in.Config.Config.Guildid)
	} else if b.in.Tip == "tg" {
		go b.Tg.DelMessage(b.in.Config.TgChannel, b.in.Tg.Mesid)
		//проверка активной подписки
		var text string
		counts := b.Db.CheckSubscribe(b.in.Name, b.in.Lvlkz, b.in.Config.TgChannel, tipPing)
		if counts == 0 {
			text = fmt.Sprintf("%s ты не подписан на пинг кз%s %d/4", b.in.NameMention, b.in.Lvlkz, tipPing)
		} else if counts == 1 {
			//удаление с базы данных
			text = fmt.Sprintf("%s отписался от пинга кз%s %d/4", b.in.NameMention, b.in.Lvlkz, tipPing)
			b.Db.Unsubscribe(b.in.Name, b.in.Lvlkz, b.in.Config.TgChannel, tipPing)
		}
		b.Tg.SendChannelDelSecond(b.in.Config.TgChannel, text, 10)
	}
}

func (b *Bot) emodjiadd(slot, emo string) {
	b.iftipdelete()
	if b.in.Tip == "ds" {
		t := b.Db.EmReadUsers(b.in.Name, b.in.Tip)
		if len(t.Name) > 0 {
			b.Db.EmUpdateEmodji(b.in.Name, b.in.Tip, slot, emo)
		} else {
			b.Db.EmInsertEmpty(b.in.Tip, b.in.Name)
			b.Db.EmUpdateEmodji(b.in.Name, b.in.Tip, slot, emo)
		}
	} else if b.in.Tip == "tg" {
		//нужно реализовать логику эмоджи для тг
	}
}
func (b *Bot) emodjis() {
	b.iftipdelete()
	if b.in.Tip == ds {
		e := b.Db.EmReadUsers(b.in.Name, ds)
		text := "	Для установки эмоджи пиши текст \n" +
			"Эмоджи пробел (номер ячейки1-4) пробел эмоджи \n" +
			"	пример \n" +
			"Эмоджи 1 🚀\n" +
			"	Ваши слоты" +
			"\n1" + e.Em1 +
			"\n2" + e.Em2 +
			"\n3" + e.Em3 +
			"\n4" + e.Em4
		b.Ds.SendEmbedText(b.in.Config.DsChannel, "Ваши эмоджи", text)
	} else if b.in.Tip == tg {

	}
}

func (b *Bot) EventStart() {
	//проверяем, есть ли активный ивент
	event1 := b.Db.NumActiveEvent(b.in.Config.CorpName)
	text := "Ивент запущен. После каждого похода на КЗ, " +
		"один из участников КЗ вносит полученные очки в базу командой К (номер катки) (количество набраных очков)"
	if event1 > 0 {
		b.ifTipSendTextDelSecond("Режим ивента уже активирован.", 10)
	} else {
		if b.in.Tip == "ds" && (b.in.Name == "Mentalisit" || b.Ds.CheckAdmin(b.in.Ds.Nameid, b.in.Config.DsChannel)) {
			b.Db.EventStartInsert(b.in.Config.CorpName)
			if b.in.Config.TgChannel != 0 {
				b.Tg.SendChannel(b.in.Config.TgChannel, text)
				b.Ds.Send(b.in.Config.DsChannel, text)
			} else {
				b.Ds.Send(b.in.Config.DsChannel, text)
			}
		} else if b.in.Tip == "tg" && (b.in.Name == "Mentalisit" || b.Tg.CheckAdminTg(b.in.Config.TgChannel, b.in.Name)) {
			b.Db.EventStartInsert(b.in.Config.CorpName)
			if b.in.Config.DsChannel != "" {
				b.Ds.Send(b.in.Config.DsChannel, text)
				b.Tg.SendChannel(b.in.Config.TgChannel, text)
			} else {
				b.Tg.SendChannel(b.in.Config.TgChannel, text)
			}
		} else {
			text = "Запуск | Оcтановка Ивента доступен Администратору канала."
			b.ifTipSendTextDelSecond(text, 60)
		}
	}
}
func (b *Bot) EventStop() {
	event1 := b.Db.NumActiveEvent(b.in.Config.CorpName)
	eventStop := "Ивент остановлен."
	eventNull := "Ивент и так не активен. Нечего останавливать "
	if b.in.Tip == "ds" && (b.in.Name == "Mentalisit" || b.Ds.CheckAdmin(b.in.Ds.Nameid, b.in.Config.DsChannel)) {
		if event1 > 0 {
			b.Db.UpdateActiveEvent0(b.in.Config.CorpName, event1)
			go b.Ds.SendChannelDelSecond(b.in.Config.DsChannel, eventStop, 60)
		} else {
			go b.Ds.SendChannelDelSecond(b.in.Config.DsChannel, eventNull, 10)
		}
	} else if b.in.Tip == "tg" && (b.in.Name == "Mentalisit" || b.Tg.CheckAdminTg(b.in.Config.TgChannel, b.in.Name)) {
		if event1 > 0 {
			b.Db.UpdateActiveEvent0(b.in.Config.CorpName, event1)
			go b.Tg.SendChannelDelSecond(b.in.Config.TgChannel, eventStop, 60)
		} else {
			go b.Tg.SendChannelDelSecond(b.in.Config.TgChannel, eventNull, 10)
		}
	} else {
		text := "Запуск|Остановка Ивента доступен Администратору канала."
		b.ifTipSendTextDelSecond(text, 20)
	}
}
func (b *Bot) EventPoints(numKZ, points int) {
	b.iftipdelete()
	// проверяем активен ли ивент
	event1 := b.Db.NumActiveEvent(b.in.Config.CorpName)
	message := ""
	if event1 > 0 {
		CountEventNames := b.Db.CountEventNames(b.in.Config.CorpName, b.in.Name, numKZ, event1)
		admin := b.checkAdmin()
		if CountEventNames > 0 || admin {
			pointsGood := b.Db.CountEventsPoints(b.in.Config.CorpName, numKZ, event1)
			if pointsGood > 0 && !admin {
				message = "данные о кз уже внесены "
			} else if pointsGood == 0 || admin {
				countEvent := b.Db.UpdatePoints(b.in.Config.CorpName, numKZ, points, event1)
				message = fmt.Sprintf("%s Очки %d внесены в базу", b.in.Name, points)
				b.changeMessageEvent(points, countEvent, numKZ, event1)
			}
		} else {
			message = fmt.Sprintf("%s Вы не являетесь участником КЗ под номером %d добавление очков невозможно.", b.in.NameMention, numKZ)
		}

	} else {
		message = "Ивент не запущен."
	}
	b.ifTipSendTextDelSecond(message, 20)
}

func (b *Bot) changeMessageEvent(points, countEvent, numberkz, numberEvent int) {
	nd, nt, t := b.Db.ReadNamesMessage(b.in.Config.CorpName, numberkz, numberEvent)
	mes1 := fmt.Sprintf("ивент игра №%d\n", t.Numberkz)
	mesOld := fmt.Sprintf("внесено %d", points)
	if countEvent == 1 {
		if b.in.Config.DsChannel != "" {
			b.Ds.EditMessage(b.in.Config.DsChannel, t.Dsmesid, fmt.Sprintf("%s %s \n%s", mes1, nd.Name1, mesOld))
		}
		if b.in.Config.TgChannel != 0 {
			b.Tg.EditText(b.in.Config.TgChannel, t.Tgmesid, fmt.Sprintf("%s %s \n%s", mes1, nt.Name1, mesOld))
		}
	} else if countEvent == 2 {
		if b.in.Config.DsChannel != "" {
			text := fmt.Sprintf("%s %s\n %s\n %s", mes1, nd.Name1, nd.Name2, mesOld)
			b.Ds.EditMessage(b.in.Config.DsChannel, t.Dsmesid, text)
		}
		if b.in.Config.TgChannel != 0 {
			text := fmt.Sprintf("%s %s\n %s\n %s", mes1, nt.Name1, nt.Name2, mesOld)
			b.Tg.EditText(b.in.Config.TgChannel, t.Tgmesid, text)
		}
	} else if countEvent == 3 {
		if b.in.Config.DsChannel != "" {
			text := fmt.Sprintf("%s %s\n %s\n %s\n %s", mes1, nd.Name1, nd.Name2, nd.Name3, mesOld)
			b.Ds.EditMessage(b.in.Config.DsChannel, t.Dsmesid, text)
		}
		if b.in.Config.TgChannel != 0 {
			text := fmt.Sprintf("%s %s\n %s\n %s\n %s", mes1, nt.Name1, nt.Name2, nt.Name3, mesOld)
			b.Tg.EditText(b.in.Config.TgChannel, t.Tgmesid, text)
		}
	} else if countEvent == 4 {
		if b.in.Config.DsChannel != "" {
			text := fmt.Sprintf("%s %s\n %s\n %s\n %s\n %s", mes1, nd.Name1, nd.Name2, nd.Name3, nd.Name4, mesOld)
			b.Ds.EditMessage(b.in.Config.DsChannel, t.Dsmesid, text)
		}
		if b.in.Config.TgChannel != 0 {
			text := fmt.Sprintf("%s %s\n %s\n %s\n %s\n %s", mes1, nt.Name1, nt.Name2, nt.Name3, nt.Name4, mesOld)
			b.Tg.EditText(b.in.Config.TgChannel, t.Tgmesid, text)
		}
	}
}
func (b *Bot) MinusMin() {
	tt := b.Db.MinusMin()
	c := corpsConfig.CorpConfig{}
	if len(tt) > 0 {
		for _, t := range tt {
			if t.Corpname != "" {
				ok, config := c.CheckCorpNameConfig(t.Corpname)
				if ok {
					in := models.InMessage{
						Mtext:       "",
						Tip:         t.Tip,
						Name:        t.Name,
						NameMention: t.Mention,
						Lvlkz:       t.Lvlkz,
						Ds: struct {
							Mesid   string
							Nameid  string
							Guildid string
						}{
							Mesid:   t.Dsmesid,
							Nameid:  "",
							Guildid: config.Config.Guildid,
						},
						Tg: struct {
							Mesid  int
							Nameid int64
						}{
							Mesid:  t.Tgmesid,
							Nameid: 0,
						},
						Config: config,
					}
					b.in = in
				}
			}

			if t.Timedown == 3 {
				text := t.Mention + " время почти вышло  ...\n если ты еще тут пиши +"
				if t.Tip == "ds" {
					mID := b.Ds.SendEmbedTime(b.in.Config.DsChannel, text)
					go b.Ds.DeleteMesageSecond(b.in.Config.DsChannel, mID, 180)
				} else if t.Tip == "tg" {
					mID := b.Tg.SendEmbedTime(b.in.Config.TgChannel, text)
					go b.Tg.DelMessageSecond(b.in.Config.TgChannel, mID, 180)
				}
			} else if t.Timedown == 0 {
				b.RsMinus()
			} else if t.Timedown <= -1 {
				b.RsMinus()
			}

		}
	}
	corpActive0 := b.Db.OneMinutsTimer()
	for _, corp := range corpActive0 {

		_, config := c.CheckCorpNameConfig(corp)

		ds, tg, wa := b.Db.MessageUpdateMin(corp)

		if config.DsChannel != "" {
			var aa []string
			for _, dsmesid := range ds {
				skip := false
				for _, u := range aa {
					if dsmesid == u {
						skip = true
						break
					}
				}

				if !skip {
					in := b.Db.MessageupdateDS(dsmesid, config)
					b.in = in
					b.QueueLevel()
				}
			}

		}
		if config.TgChannel != 0 {
			var aa []int
			for _, tgmesid := range tg {
				skip := false
				for _, u := range aa {
					if tgmesid == u {
						skip = true
						break
					}
				}
				if !skip {
					b.Db.MessageupdateTG(tgmesid, config)
				}
			}
			if config.WaChannel != "" {
				//тут будет логика ватса
				fmt.Println(wa)
			}
		}
	}
}
