package bot

import "fmt"

// lang ok
func (b *Bot) EventText() (text string, numE int) {
	//проверяем, есть ли активный ивент
	numberevent := b.storage.Event.NumActiveEvent(b.in.Config.CorpName)
	if numberevent == 0 { //ивент не активен
		return "", 0
	} else if numberevent > 0 { //активный ивент
		numE = b.storage.Event.NumberQueueEvents(b.in.Config.CorpName) //номер кз number FROM rsevent
		text = fmt.Sprintf("\nID %d %s", numE, b.GetLang("dly iventa"))
		return text, numE
	}
	return text, numE
}
func (b *Bot) EventStart() {
	if b.debug {
		fmt.Println("in EventStart", b.in)
	}
	b.iftipdelete()
	//проверяем, есть ли активный ивент
	event1 := b.storage.Event.NumActiveEvent(b.in.Config.CorpName)
	text := b.GetLang("iventZapushen")
	if event1 > 0 {
		b.ifTipSendTextDelSecond(b.GetLang("rejimIventaUje"), 10)
	} else {
		if b.in.Tip == ds && (b.in.Name == "Mentalisit" || b.client.Ds.CheckAdmin(b.in.Ds.Nameid, b.in.Config.DsChannel)) {
			b.storage.Event.EventStartInsert(b.in.Config.CorpName)
			if b.in.Config.TgChannel != "" {
				b.client.Tg.SendChannel(b.in.Config.TgChannel, text)
				b.client.Ds.Send(b.in.Config.DsChannel, text)
			} else {
				b.client.Ds.Send(b.in.Config.DsChannel, text)
			}
		} else if b.in.Tip == tg && (b.in.Name == "Mentalisit" || b.client.Tg.CheckAdminTg(b.in.Config.TgChannel, b.in.Name)) {
			b.storage.Event.EventStartInsert(b.in.Config.CorpName)
			if b.in.Config.DsChannel != "" {
				b.client.Ds.Send(b.in.Config.DsChannel, text)
				b.client.Tg.SendChannel(b.in.Config.TgChannel, text)
			} else {
				b.client.Tg.SendChannel(b.in.Config.TgChannel, text)
			}
		} else {
			text = b.GetLang("zapuskIostanovka")
			b.ifTipSendTextDelSecond(text, 60)
		}
	}
}
func (b *Bot) EventStop() {
	if b.debug {
		fmt.Println("in EventStop", b.in)
	}
	b.iftipdelete()
	event1 := b.storage.Event.NumActiveEvent(b.in.Config.CorpName)
	eventStop := b.GetLang("IventOstanovlen")
	eventNull := b.GetLang("iventItakAktiven")
	if b.in.Tip == "ds" && (b.in.Name == "Mentalisit" || b.client.Ds.CheckAdmin(b.in.Ds.Nameid, b.in.Config.DsChannel)) {
		if event1 > 0 {
			b.storage.Event.UpdateActiveEvent0(b.in.Config.CorpName, event1)
			go b.client.Ds.SendChannelDelSecond(b.in.Config.DsChannel, eventStop, 60)
		} else {
			go b.client.Ds.SendChannelDelSecond(b.in.Config.DsChannel, eventNull, 10)
		}
	} else if b.in.Tip == tg && (b.in.Name == "Mentalisit" || b.client.Tg.CheckAdminTg(b.in.Config.TgChannel, b.in.Name)) {
		if event1 > 0 {
			b.storage.Event.UpdateActiveEvent0(b.in.Config.CorpName, event1)
			go b.client.Tg.SendChannelDelSecond(b.in.Config.TgChannel, eventStop, 60)
		} else {
			go b.client.Tg.SendChannelDelSecond(b.in.Config.TgChannel, eventNull, 10)
		}
	} else {
		text := b.GetLang("zapuskIostanovka")
		b.ifTipSendTextDelSecond(text, 20)
	}
}
func (b *Bot) EventPoints(numKZ, points int) {
	if b.debug {
		fmt.Println("in EventPoints", b.in)
	}
	b.iftipdelete()
	// проверяем активен ли ивент
	event1 := b.storage.Event.NumActiveEvent(b.in.Config.CorpName)
	message := ""
	if event1 > 0 {
		CountEventNames := b.storage.Event.CountEventNames(b.in.Config.CorpName, b.in.Name, numKZ, event1)
		admin := b.checkAdmin()
		if CountEventNames > 0 || admin {
			pointsGood := b.storage.Event.CountEventsPoints(b.in.Config.CorpName, numKZ, event1)
			if pointsGood > 0 && !admin {
				message = b.GetLang("dannieKzUjeVneseni")
			} else if pointsGood == 0 || admin {
				countEvent := b.storage.Event.UpdatePoints(b.in.Config.CorpName, numKZ, points, event1)
				message = fmt.Sprintf("%s %d %s", b.in.Name, points, b.GetLang("ochki vnesen"))
				b.changeMessageEvent(points, countEvent, numKZ, event1)
			}
		} else {
			message = fmt.Sprintf("%s  %s %d", b.in.NameMention, b.GetLang("dobavlenieOchkovNevozmojno"), numKZ)
		}

	} else {
		message = b.GetLang("iventNeZapushen")
	}
	b.ifTipSendTextDelSecond(message, 20)
}
func (b *Bot) changeMessageEvent(points, countEvent, numberkz, numberEvent int) {
	if b.debug {
		fmt.Println("in changeMessageEvent ", b.in)
	}
	nd, nt, t := b.storage.Event.ReadNamesMessage(b.in.Config.CorpName, numberkz, numberEvent)
	mes1 := fmt.Sprintf("%s №%d\n", b.GetLang("iventIgra"), t.Numberkz)
	mesOld := fmt.Sprintf("%s %d", b.GetLang("vneseno"), points)
	if countEvent == 1 {
		if b.in.Config.DsChannel != "" {
			b.client.Ds.EditMessage(b.in.Config.DsChannel, t.Dsmesid, fmt.Sprintf("%s %s \n%s", mes1, nd.Name1, mesOld))
		}
		if b.in.Config.TgChannel != "" {
			b.client.Tg.EditText(b.in.Config.TgChannel, t.Tgmesid, fmt.Sprintf("%s %s \n%s", mes1, nt.Name1, mesOld))
		}
	} else if countEvent == 2 {
		if b.in.Config.DsChannel != "" {
			text := fmt.Sprintf("%s %s\n %s\n %s", mes1, nd.Name1, nd.Name2, mesOld)
			b.client.Ds.EditMessage(b.in.Config.DsChannel, t.Dsmesid, text)
		}
		if b.in.Config.TgChannel != "" {
			text := fmt.Sprintf("%s %s\n %s\n %s", mes1, nt.Name1, nt.Name2, mesOld)
			b.client.Tg.EditText(b.in.Config.TgChannel, t.Tgmesid, text)
		}
	} else if countEvent == 3 {
		if b.in.Config.DsChannel != "" {
			text := fmt.Sprintf("%s %s\n %s\n %s\n %s", mes1, nd.Name1, nd.Name2, nd.Name3, mesOld)
			b.client.Ds.EditMessage(b.in.Config.DsChannel, t.Dsmesid, text)
		}
		if b.in.Config.TgChannel != "" {
			text := fmt.Sprintf("%s %s\n %s\n %s\n %s", mes1, nt.Name1, nt.Name2, nt.Name3, mesOld)
			b.client.Tg.EditText(b.in.Config.TgChannel, t.Tgmesid, text)
		}
	} else if countEvent == 4 {
		if b.in.Config.DsChannel != "" {
			text := fmt.Sprintf("%s %s\n %s\n %s\n %s\n %s", mes1, nd.Name1, nd.Name2, nd.Name3, nd.Name4, mesOld)
			b.client.Ds.EditMessage(b.in.Config.DsChannel, t.Dsmesid, text)
		}
		if b.in.Config.TgChannel != "" {
			text := fmt.Sprintf("%s %s\n %s\n %s\n %s\n %s", mes1, nt.Name1, nt.Name2, nt.Name3, nt.Name4, mesOld)
			b.client.Tg.EditText(b.in.Config.TgChannel, t.Tgmesid, text)
		}
	}
}
