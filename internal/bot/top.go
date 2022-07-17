package bot

import (
	"fmt"
	"time"
)

const (
	scan      = "Сканирую базу данных"
	nohistory = " История не найдена "
	formlist  = "Формирую список "
)

func (b *Bot) TopLevel() {
	b.iftipdelete()
	numEvent := b.Db.Event.NumActiveEvent(b.in.Config.CorpName)
	if numEvent == 0 {
		mesage := "\xF0\x9F\x93\x96 ТОП Участников кз" + b.in.Lvlkz + "\n"
		b.ifTipSendTextDelSecond(scan, 5)
		good := b.Db.Top.TopLevel(b.in.Config.CorpName, b.in.Lvlkz)
		if !good {
			b.ifTipSendTextDelSecond(nohistory, 20)
		} else if good {
			b.ifTipSendTextDelSecond(formlist, 5)
			mest := b.Db.Top.TopTemp()
			if b.in.Tip == ds {
				m := b.Ds.SendEmbedText(b.in.Config.DsChannel, mesage, mest)
				b.Ds.DeleteMesageSecond(b.in.Config.DsChannel, m.ID, 60)
			} else if b.in.Tip == tg {
				b.Tg.SendChannelDelSecond(b.in.Config.TgChannel, mesage+mest, 60)
			}
		}
	} else {
		mesage := "\xF0\x9F\x93\x96 ТОП Участников ивента кз:" + b.in.Lvlkz + "\n"
		b.ifTipSendTextDelSecond(scan, 5)
		good := b.Db.Top.TopEventLevel(b.in.Config.CorpName, b.in.Lvlkz, numEvent)
		if !good {
			b.ifTipSendTextDelSecond(nohistory, 20)
		} else {
			b.ifTipSendTextDelSecond(formlist, 5)
			mest := b.Db.Top.TopTempEvent()
			if b.in.Tip == ds {
				m := b.Ds.SendEmbedText(b.in.Config.DsChannel, mesage, mest)
				b.Ds.DeleteMesageSecond(b.in.Config.DsChannel, m.ID, 60)
			} else if b.in.Tip == tg {
				b.Tg.SendChannelDelSecond(b.in.Config.TgChannel, mesage+mest, 60)
			}
		}
	}
}
func (b *Bot) TopAll() {
	b.iftipdelete()
	numEvent := b.Db.Event.NumActiveEvent(b.in.Config.CorpName)
	if numEvent == 0 {
		mesage := "\xF0\x9F\x93\x96 ТОП Участников:\n"
		b.ifTipSendTextDelSecond(scan, 5)
		good := b.Db.Top.TopAll(b.in.Config.CorpName)
		if good {
			b.ifTipSendTextDelSecond(formlist, 5)
			message2 := b.Db.Top.TopTemp()
			if b.in.Tip == ds {
				m := b.Ds.SendEmbedText(b.in.Config.DsChannel, mesage, message2)
				b.Ds.DeleteMesageSecond(b.in.Config.DsChannel, m.ID, 60)
			} else if b.in.Tip == tg {
				b.Tg.SendChannelDelSecond(b.in.Config.TgChannel, mesage+message2, 60)
			}
		} else if !good {
			b.ifTipSendTextDelSecond(nohistory, 10)
		}
	} else if numEvent > 0 {
		mesage := "\xF0\x9F\x93\x96 ТОП Участников Ивента:\n"
		b.ifTipSendTextDelSecond(scan, 10)
		good := b.Db.Top.TopAllEvent(b.in.Config.CorpName, numEvent)
		if good {
			b.ifTipSendTextDelSecond(formlist, 5)
			message2 := b.Db.Top.TopTempEvent()
			mesage = mesage + message2
			b.ifTipSendTextDelSecond(mesage, 60)
		} else if !good {
			b.ifTipSendTextDelSecond(nohistory, 10)
		}
	}
}
func (b *Bot) TopDate(oldDate string) {
	b.iftipdelete()
	mesage := fmt.Sprintf("\xF0\x9F\x93\x96 ТОП Участников начиная с %s \n", oldDate)
	b.ifTipSendTextDelSecond(scan, 5)
	good := b.Db.Top.TopAllDay(b.in.Config.CorpName, oldDate)
	if good {
		b.ifTipSendTextDelSecond(formlist, 5)
		message2 := b.Db.Top.TopTemp()
		mesage = mesage + message2
		b.ifTipSendTextDelSecond(mesage, 60)
	} else if !good {
		b.ifTipSendTextDelSecond(nohistory, 10)
	}

}
func (b *Bot) TopDateLevel(oldDate string) {
	b.iftipdelete()
	mesage := "\xF0\x9F\x93\x96 ТОП Участников кз" + b.in.Lvlkz + "\n"
	b.ifTipSendTextDelSecond(scan, 5)
	good := b.Db.Top.TopLevelDay(b.in.Config.CorpName, b.in.Lvlkz, oldDate)
	if !good {
		go b.ifTipSendTextDelSecond(nohistory, 20)
	} else if good {
		go b.ifTipSendTextDelSecond(formlist, 5)
		mest := b.Db.Top.TopTemp()
		if b.in.Tip == ds {
			m := b.Ds.SendEmbedText(b.in.Config.DsChannel, mesage, mest)
			b.Ds.DeleteMesageSecond(b.in.Config.DsChannel, m.ID, 60)
		} else if b.in.Tip == tg {
			b.Tg.SendChannelDelSecond(b.in.Config.TgChannel, mesage+mest, 60)
		}
	}
}

func (Bot) t7() string {
	tm := time.Now().Unix() - 604800 //минус  7 суток
	tu := time.Unix(tm, 0)
	t7 := tu.Format("2006-01-02")
	return t7
}
func (Bot) t1() string {
	tm := time.Now().Unix() - 86400 //минус сутки
	tu := time.Unix(tm, 0)
	t1 := tu.Format("2006-01-02")
	return t1
}
