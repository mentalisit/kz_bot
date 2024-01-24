package bot

import (
	"context"
	"fmt"
	"time"
)

// lang ok
func (b *Bot) emodjiadd(slot, emo string) {
	if b.debug {
		fmt.Println("in emodjiadd", b.in)
	}
	b.iftipdelete()
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	t := b.storage.Emoji.EmojiModuleReadUsers(ctx, b.in.Name, b.in.Tip)
	if len(t.Name) == 0 {
		b.storage.Emoji.EmInsertEmpty(ctx, b.in.Tip, b.in.Name)
	}
	text := b.storage.Emoji.EmojiUpdate(ctx, b.in.Name, b.in.Tip, slot, emo)
	b.ifTipSendTextDelSecond(text, 20)
}
func (b *Bot) emodjis() {
	if b.debug {
		fmt.Println("in emodjis", b.in)
	}
	b.iftipdelete()
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	e := b.storage.Emoji.EmojiModuleReadUsers(ctx, b.in.Name, b.in.Tip)

	text := b.GetLang("dly ustanovki") +
		"\n1 " + e.Em1 +
		"\n2 " + e.Em2 +
		"\n3 " + e.Em3 +
		"\n4 " + e.Em4
	if b.in.Tip == ds {
		text += fmt.Sprintf("\n %s %s %s %s", e.Module1, e.Module2, e.Module3, e.Weapon)
	}
	b.ifTipSendTextDelSecond(b.GetLang("vashiEmodji")+text, 20)
}
