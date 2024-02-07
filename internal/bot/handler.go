package bot

import (
	"context"
	"fmt"
	"kz_bot/internal/compendiumCli"
	"kz_bot/internal/models"
	"kz_bot/pkg/translator"
	"kz_bot/pkg/utils"
	"strings"
)

// lang
func (b *Bot) iftipdelete() {
	if b.in.Tip == ds && !b.in.Option.Reaction && !b.in.Option.Update && !b.in.Option.Edit {
		go b.client.Ds.DeleteMessage(b.in.Config.DsChannel, b.in.Ds.Mesid)
		go b.client.Ds.ChannelTyping(b.in.Config.DsChannel)
	} else if b.in.Tip == tg && !b.in.Option.Reaction && !b.in.Option.Update {
		go b.client.Tg.ChatTyping(b.in.Config.TgChannel)
		go b.client.Tg.DelMessage(b.in.Config.TgChannel, b.in.Tg.Mesid)
	}
}
func (b *Bot) ifTipSendMentionText(text string) {
	if b.in.Tip == ds {
		go b.client.Ds.SendChannelDelSecond(b.in.Config.DsChannel, b.in.NameMention+text, 10)
	} else if b.in.Tip == tg {
		go b.client.Tg.SendChannelDelSecond(b.in.Config.TgChannel, b.in.NameMention+text, 10)
	}
}
func (b *Bot) ifTipSendTextDelSecond(text string, time int) {
	if b.in.Tip == ds {
		go b.client.Ds.SendChannelDelSecond(b.in.Config.DsChannel, text, time)
	} else if b.in.Tip == tg {
		go b.client.Tg.SendChannelDelSecond(b.in.Config.TgChannel, text, time)
	}
}

func (b *Bot) emReadName(name, nameMention, tip string) string { // склеиваем имя и эмоджи
	t := b.storage.Emoji.EmojiModuleReadUsers(context.Background(), name, tip)
	newName := name
	if tip == ds {
		newName = nameMention
	} else {
		newName = name
	}

	if len(t.Name) > 0 {
		if tip == ds && tip == t.Tip {
			newName = fmt.Sprintf("%s %s %s %s %s %s%s%s%s", nameMention, t.Module1, t.Module2, t.Module3, t.Weapon, t.Em1, t.Em2, t.Em3, t.Em4)
		} else if tip == tg && tip == t.Tip {
			newName = fmt.Sprintf("%s %s%s%s%s", name, t.Em1, t.Em2, t.Em3, t.Em4)
			if t.Weapon != "" {
				newName = fmt.Sprintf("%s [%s] %s%s%s%s", name, t.Weapon, t.Em1, t.Em2, t.Em3, t.Em4)
			}
		}
	} else if b.in.Tip == ds && b.in.Config.Guildid == "716771579278917702" {
		genesis, enrich, rsextender := compendiumCli.GetUserId(b.in.Ds.Nameid)
		b.storage.Emoji.EmInsertEmpty(context.Background(), "ds", name)
		one := fmt.Sprintf("<:rse:1199068829511335946> %d ", rsextender)
		two := fmt.Sprintf("<:genesis:1199068748280242237> %d ", genesis)
		three := fmt.Sprintf("<:enrich:1199068793633251338> %d ", enrich)
		newName = fmt.Sprintf("%s ", nameMention)
		if rsextender != 0 {
			b.storage.Emoji.ModuleUpdate(context.Background(), name, "ds", "1", one)
			newName += one
		}
		if genesis != 0 {
			b.storage.Emoji.ModuleUpdate(context.Background(), name, "ds", "2", two)
			newName += two
		}
		if enrich != 0 {
			b.storage.Emoji.ModuleUpdate(context.Background(), name, "ds", "3", three)
			newName += three
		}

	}
	return newName
}
func (b *Bot) emReadMention(name, nameMention, tip string) string { // склеиваем имя и эмоджи
	t := b.storage.Emoji.EmojiModuleReadUsers(context.Background(), name, tip)
	newName := nameMention

	if len(t.Name) > 0 {
		if tip == ds && tip == t.Tip {
			newName = fmt.Sprintf("%s %s %s %s %s %s%s%s%s", nameMention, t.Module1, t.Module2, t.Module3, t.Weapon, t.Em1, t.Em2, t.Em3, t.Em4)
		} else if tip == tg && tip == t.Tip {
			newName = fmt.Sprintf("%s %s%s%s%s", nameMention, t.Em1, t.Em2, t.Em3, t.Em4)
			if t.Weapon != "" {
				newName = fmt.Sprintf("%s [%s] %s%s%s%s", nameMention, t.Weapon, t.Em1, t.Em2, t.Em3, t.Em4)
			}
		}
	} else if b.in.Tip == ds && b.in.Config.Guildid == "716771579278917702" {
		genesis, enrich, rsextender := compendiumCli.GetUserId(b.in.Ds.Nameid)
		b.storage.Emoji.EmInsertEmpty(context.Background(), "ds", name)
		one := fmt.Sprintf("<:rse:1199068829511335946> %d ", rsextender)
		two := fmt.Sprintf("<:genesis:1199068748280242237> %d ", genesis)
		three := fmt.Sprintf("<:enrich:1199068793633251338> %d ", enrich)
		newName = fmt.Sprintf("%s ", nameMention)
		if rsextender != 0 {
			b.storage.Emoji.ModuleUpdate(context.Background(), name, "ds", "1", one)
			newName += one
		}
		if genesis != 0 {
			b.storage.Emoji.ModuleUpdate(context.Background(), name, "ds", "2", two)
			newName += two
		}
		if enrich != 0 {
			b.storage.Emoji.ModuleUpdate(context.Background(), name, "ds", "3", three)
			newName += three
		}

	}
	return newName
}

func (b *Bot) updateCompendiumModules() {
	b.iftipdelete()
	genesis, enrich, rsextender := compendiumCli.GetUserId(b.in.Ds.Nameid)
	one := fmt.Sprintf("<:rse:1199068829511335946> %d ", rsextender)
	two := fmt.Sprintf("<:genesis:1199068748280242237> %d ", genesis)
	three := fmt.Sprintf("<:enrich:1199068793633251338> %d ", enrich)
	if rsextender != 0 {
		b.storage.Emoji.ModuleUpdate(context.Background(), b.in.Name, "ds", "1", one)
	}
	if genesis != 0 {
		b.storage.Emoji.ModuleUpdate(context.Background(), b.in.Name, "ds", "2", two)
	}
	if enrich != 0 {
		b.storage.Emoji.ModuleUpdate(context.Background(), b.in.Name, "ds", "3", three)
	}
	b.ifTipSendMentionText(" загружено из компендиум бота " + one + two + three)
}
func (b *Bot) checkAdmin() bool {
	admin := false
	if b.in.Tip == ds {
		admin = b.client.Ds.CheckAdmin(b.in.Ds.Nameid, b.in.Config.DsChannel)
	} else if b.in.Tip == tg {
		admin = b.client.Tg.CheckAdminTg(b.in.Config.TgChannel, b.in.Name)
	}
	return admin
}

func (b *Bot) nameMention(u models.Users, tip string) (n1, n2, n3, n4 string) {
	if u.User1.Tip == tip {
		n1 = b.emReadMention(u.User1.Name, u.User1.Mention, tip)
	} else {
		n1 = u.User1.Name
	}
	if u.User2.Tip == tip {
		n2 = b.emReadMention(u.User2.Name, u.User2.Mention, tip)
	} else {
		n2 = u.User2.Name
	}
	if u.User3.Tip == tip {
		n3 = b.emReadMention(u.User3.Name, u.User3.Mention, tip)
	} else {
		n3 = u.User3.Name
	}
	if b.in.Tip == tip {
		n4 = b.emReadMention(b.in.Name, b.in.NameMention, tip)
	} else {
		n4 = b.in.Name
	}
	return
}

func (b *Bot) elseChat(user []string) { //проверяем всех игроков этой очереди на присутствие в других очередях или корпорациях
	user = utils.RemoveDuplicateElementString(user)
	for _, u := range user {
		if b.storage.Count.CountNameQueue(context.Background(), u) > 0 {
			b.elsetrue(u)
		}
	}
}
func (b *Bot) elsetrue(name string) { //удаляем игрока с очереди
	tt := b.storage.DbFunc.ElseTrue(context.Background(), name)
	for _, t := range tt {
		ok, config := b.CheckCorpNameConfig(t.Corpname)
		if ok {
			in := models.InMessage{
				Mtext:       t.Lvlkz + "-",
				Tip:         t.Tip,
				Name:        t.Name,
				NameMention: t.Mention,
				Lvlkz:       t.Lvlkz,
				Ds: struct {
					Mesid   string
					Nameid  string
					Guildid string
					Avatar  string
				}{
					Mesid:   t.Dsmesid,
					Nameid:  "",
					Guildid: ""},
				Tg: struct {
					Mesid int
				}{
					Mesid: t.Tgmesid,
				},
				Config: config,
				Option: models.Option{Elsetrue: true},
			}
			b.inbox <- in
		}
	}
}

func (b *Bot) hhelp() {
	b.iftipdelete()
	if b.in.Tip == ds {
		go b.client.Ds.Help(b.in.Config.DsChannel)
	} else if b.in.Tip == tg {
		go b.client.Tg.Help(b.in.Config.TgChannel, b.in.Config.Country)
	}
}

func (b *Bot) GetLang(key string) string {
	return b.storage.Words.GetWords(b.in.Config.Country, key)
}

func containsSymbolD(s string) (dark bool, result string) {
	for _, char := range s {
		if char == 'd' {
			dark = true
		}
	}
	if dark {
		result = strings.Replace(s, "d", "", -1)
	}

	return dark, result
}

func (b *Bot) Transtale() {
	text2 := translator.TranslateAnswer(b.in.Mtext, b.in.Config.Country)
	if b.in.Mtext != text2 {
		if b.in.Tip == ds {
			go func() {
				m := b.client.Ds.SendWebhook(text2, b.in.Name, b.in.Config.DsChannel, b.in.Config.Guildid, b.in.Ds.Avatar)
				b.client.Ds.DeleteMesageSecond(b.in.Config.DsChannel, m, 90)
			}()
		} else if b.in.Tip == tg {
			go b.client.Tg.SendChannelDelSecond(b.in.Config.TgChannel, text2, 90)
		}
	}
}
