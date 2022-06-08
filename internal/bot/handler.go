package bot

import "C"
import (
	"fmt"
	corpsConfig "kz_bot/internal/clients/corpConfig"
	"kz_bot/internal/models"
)

const (
	ds       = "ds"
	tg       = "tg"
	nickname = "Для того что бы БОТ мог Вас индентифицировать, создайте уникальный НикНей в настройках. Вы можете использовать a-z, 0-9 и символы подчеркивания. Минимальная длина - 5 символов."
)

func (b *Bot) iftipdelete() {
	if b.in.Tip == ds && !b.in.Option.Callback {
		b.Ds.DeleteMessage(b.in.Config.DsChannel, b.in.Ds.Mesid)
	} else if b.in.Tip == tg && !b.in.Option.Callback {
		b.Tg.DelMessage(b.in.Config.TgChannel, b.in.Tg.Mesid)
		if b.in.NameMention == "@" {
			b.Tg.SendChannelDelSecond(b.in.Config.TgChannel, nickname, 60)
		}
	}
}
func (b *Bot) ifTipSendMentionText(text string) {
	if b.in.Tip == ds {
		go b.Ds.SendChannelDelSecond(b.in.Config.DsChannel, b.in.NameMention+text, 10)
	} else if b.in.Tip == tg {
		go b.Tg.SendChannelDelSecond(b.in.Config.TgChannel, b.in.NameMention+text, 10)
	}
}
func (b *Bot) emReadName(name, tip string) string { // склеиваем имя и эмоджи
	t := b.Db.EmReadUsers(name, tip)
	newName := name
	if tip == t.Tip {
		newName = fmt.Sprintf("%s %s%s%s%s", name, t.Em1, t.Em2, t.Em3, t.Em4)
	}
	return newName
}
func (b *Bot) elseChat(u models.Users, name4 string) { //проверяем всех игроков этой очереди на присутствие в других очередях или корпорациях
	if b.Db.CountNameQueue(u.User1.Name) > 0 {
		b.elsetrue(u.User1.Name)
	}
	if b.Db.CountNameQueue(u.User2.Name) > 0 {
		b.elsetrue(u.User2.Name)
	}
	if b.Db.CountNameQueue(u.User3.Name) > 0 {
		b.elsetrue(u.User3.Name)
	}
	if b.Db.CountNameQueue(name4) > 0 {
		b.elsetrue(name4)
	}
}
func (b *Bot) elsetrue(name string) { //удаляем игрока с очереди
	t := b.Db.ElseTrue(name)
	c := corpsConfig.CorpConfig{}
	ok, config := c.CheckCorpNameConfig(t.Corpname)
	if ok {
		in := models.InMessage{
			Tip:         t.Tip,
			Name:        t.Name,
			NameMention: t.Mention,
			Lvlkz:       t.Lvlkz,
			Timekz:      string(t.Timedown),
			Ds: models.Ds{
				Mesid: t.Dsmesid,
			},
			Tg: models.Tg{
				Mesid: t.Tgmesid,
			},
			Config: config,
			Option: models.Option{
				Callback: false,
				Edit:     true,
				Update:   false,
			},
		}
		b.in = in
		b.RsMinus()

	}
}
func (b *Bot) callbackNo() {
	if b.in.Tip == "ds" && !b.in.Option.Callback {
		go b.Ds.DeleteMessage(b.in.Config.DsChannel, b.in.Ds.Mesid)
	} else if b.in.Tip == "tg" && !b.in.Option.Callback {
		go b.Tg.DelMessage(b.in.Config.TgChannel, b.in.Tg.Mesid)
	}
}
func (b *Bot) SubscribePing(tipPing int) {
	men := b.Db.SubscPing(b.in.NameMention, b.in.Lvlkz, b.in.Config.CorpName, tipPing, b.in.Config.TgChannel)
	if len(men) > 0 {
		men = fmt.Sprintf("Сбор на кз%s\n%s", b.in.Lvlkz, men)
		b.Tg.SendChannelDelSecond(b.in.Config.TgChannel, men, 600)
	}
}
