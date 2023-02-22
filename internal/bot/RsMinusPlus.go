package bot

import (
	"context"
	"fmt"
	"time"
)

//lang ok

func (b *Bot) Plus() bool {
	if b.debug {
		fmt.Println("in Plus", b.in)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	countName := b.storage.Count.CountNameQueueCorp(ctx, b.in.Name, b.in.Config.CorpName)
	message := ""
	ins := false
	if countName > 0 && b.in.Option.Reaction {
		b.iftipdelete()
		ins = true
		t := b.storage.Timers.UpdateMitutsQueue(ctx, b.in.Name, b.in.Config.CorpName)
		if t.Timedown > 3 {
			message = fmt.Sprintf("%s %s%s %s %d%s",
				t.Mention, b.GetLang("ranovatoPlysik"), t.Lvlkz, b.GetLang("budeshEshe"), t.Timedown, b.GetLang("min."))
		} else if t.Timedown <= 3 {
			message = t.Mention + b.GetLang("vremyaObnovleno")
			b.in.Lvlkz = t.Lvlkz

			b.QueueLevel()
		}
		b.ifTipSendTextDelSecond(message, 10)
	}
	return ins
}
func (b *Bot) Minus() bool {
	if b.debug {
		fmt.Println("in Minus", b.in)
	}
	bb := false
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	countNames := b.storage.Count.CountNameQueueCorp(ctx, b.in.Name, b.in.Config.CorpName)
	if countNames > 0 && b.in.Option.Reaction {
		b.iftipdelete()
		bb = true
		t := b.storage.Timers.UpdateMitutsQueue(ctx, b.in.Name, b.in.Config.CorpName)
		if t.Name == b.in.Name && t.Timedown > 3 {
			message := fmt.Sprintf("%s %s%s %s %d%s",
				t.Mention, b.GetLang("ranovatoMinus"), t.Lvlkz, b.GetLang("budeshEshe"), t.Timedown, b.GetLang("min."))
			b.ifTipSendTextDelSecond(message, 10)
		} else if t.Name == b.in.Name && t.Timedown <= 3 {
			b.in.Lvlkz = t.Lvlkz
			b.RsMinus()
		}

	}
	return bb
}
