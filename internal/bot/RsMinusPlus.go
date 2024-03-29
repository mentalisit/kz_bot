package bot

import (
	"context"
	"fmt"
	"time"
)

//lang ok

func (b *Bot) Plus() bool {
	if b.debug {
		fmt.Printf("in Plus %+v\n", b.in)
	}
	b.iftipdelete()
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	countName := b.storage.Count.CountNameQueueCorp(ctx, b.in.Name, b.in.Config.CorpName)
	message := ""
	ins := false
	if countName > 0 && b.in.Option.Reaction {
		b.iftipdelete()
		t := b.storage.Timers.UpdateMitutsQueue(ctx, b.in.Name, b.in.Config.CorpName)
		if t.Timedown > 3 {
			message = fmt.Sprintf("%s %s%s %s %d%s",
				t.Mention, b.GetLang("ranovatoPlysik"), t.Lvlkz, b.GetLang("budeshEshe"), t.Timedown, b.GetLang("min."))
		} else if t.Timedown <= 3 {
			ins = true
			message = t.Mention + b.GetLang("vremyaObnovleno")
			b.in.Lvlkz = t.Lvlkz
			b.in.Option.Reaction = false
			b.QueueLevel()
		}
		b.ifTipSendTextDelSecond(message, 10)
		//if b.in.Tip == ds {
		//	b.client.Ds.DeleteMessage(b.in.Config.DsChannel, b.in.Ds.Mesid)
		//} else if b.in.Tip == tg {
		//	b.client.Tg.DelMessage(b.in.Config.TgChannel, b.in.Tg.Mesid)
		//}
	}
	return ins
}
func (b *Bot) Minus() bool {
	if b.debug {
		fmt.Printf("in Minus %+v\n", b.in)
	}
	b.iftipdelete()
	bb := false
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	countNames := b.storage.Count.CountNameQueueCorp(ctx, b.in.Name, b.in.Config.CorpName)
	if countNames > 0 && b.in.Option.Reaction {
		b.iftipdelete()
		t := b.storage.Timers.UpdateMitutsQueue(ctx, b.in.Name, b.in.Config.CorpName)
		if t.Name == b.in.Name && t.Timedown > 3 {
			message := fmt.Sprintf("%s %s%s %s %d%s",
				t.Mention, b.GetLang("ranovatoMinus"), t.Lvlkz, b.GetLang("budeshEshe"), t.Timedown, b.GetLang("min."))
			b.ifTipSendTextDelSecond(message, 10)
		} else if t.Name == b.in.Name && t.Timedown <= 3 {
			b.in.Lvlkz = t.Lvlkz
			bb = true
			b.in.Option.Reaction = false
			b.in.Option.Update = true
			b.RsMinus()
		}
		if b.in.Tip == ds {
			b.client.Ds.DeleteMessage(b.in.Config.DsChannel, b.in.Ds.Mesid)
		} else if b.in.Tip == tg {
			b.client.Tg.DelMessage(b.in.Config.TgChannel, b.in.Tg.Mesid)
		}

	}
	return bb
}
