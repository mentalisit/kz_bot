package bot

import "C"
import (
	"fmt"
	"kz_bot/pkg/telegraf"
	"regexp"
	"time"

	corpsConfig "kz_bot/internal/clients/corpConfig"
	"kz_bot/internal/models"
)

const (
	ds = "ds"
	tg = "tg"
	wa = "wa"
)

func (b *Bot) EventText() (text string, numE int) {
	//проверяем, есть ли активный ивент
	numberevent := b.Db.Event.NumActiveEvent(b.in.Config.CorpName)
	if numberevent == 0 { //ивент не активен
		return "", 0
	} else if numberevent > 0 { //активный ивент
		numE = b.Db.Event.NumberQueueEvents(b.in.Config.CorpName) //номер кз number FROM rsevent
		text = fmt.Sprintf("\nID %d для ивента ", numE)
		return text, numE
	}
	return text, numE
}

func (b *Bot) iftipdelete() bool {
	if b.in.Tip == ds && !b.in.Option.Reaction && !b.in.Option.Update {
		go b.Ds.DeleteMessage(b.in.Config.DsChannel, b.in.Ds.Mesid)
	} else if b.in.Tip == tg && !b.in.Option.Reaction && !b.in.Option.Update {
		go b.Tg.DelMessage(b.in.Config.TgChannel, b.in.Tg.Mesid)
	}
	return true
}
func (b *Bot) ifTipSendMentionText(text string) {
	if b.in.Tip == ds {
		go b.Ds.SendChannelDelSecond(b.in.Config.DsChannel, b.in.NameMention+text, 10)
	} else if b.in.Tip == tg {
		go b.Tg.SendChannelDelSecond(b.in.Config.TgChannel, b.in.NameMention+text, 10)
	} else if b.in.Tip == wa {
		go b.Wa.Send(b.in.Config.WaChannel, b.in.NameMention+text)
	}
}
func (b *Bot) ifTipSendTextDelSecond(text string, time int) {
	if b.in.Tip == ds {
		go b.Ds.SendChannelDelSecond(b.in.Config.DsChannel, text, time)
	} else if b.in.Tip == tg {
		go b.Tg.SendChannelDelSecond(b.in.Config.TgChannel, text, time)
	}
}
func (b *Bot) emReadName(name, tip string) string { // склеиваем имя и эмоджи
	t := b.Db.Emoji.EmReadUsers(name, tip)
	newName := name
	if tip == t.Tip {
		newName = fmt.Sprintf("%s %s%s%s%s", name, t.Em1, t.Em2, t.Em3, t.Em4)
	}
	return newName
}
func (b *Bot) elseChat(user []string) { //проверяем всех игроков этой очереди на присутствие в других очередях или корпорациях
	user = b.removeDuplicateElementString(user)
	for _, u := range user {
		if b.Db.Count.CountNameQueue(u) > 0 {
			b.elsetrue(u)
		}
	}
}

func (b *Bot) elsetrue(name string) { //удаляем игрока с очереди
	tt := b.Db.ElseTrue(name)
	c := corpsConfig.CorpConfig{}
	for _, t := range tt {
		ok, config := c.CheckCorpNameConfig(t.Corpname)
		if ok {
			in := models.InMessage{
				Mtext:       t.Lvlkz + "-",
				Tip:         t.Tip,
				Name:        t.Name,
				NameMention: t.Mention,
				Lvlkz:       t.Lvlkz,
				Timekz:      string(t.Timedown),
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
					Mesid  int
					Nameid int64
				}{
					Mesid:  t.Tgmesid,
					Nameid: 0},
				Config: config,
				Option: models.Option{Elsetrue: true},
			}
			if t.Tip == ds {
				models.ChDs <- in
			} else if t.Tip == tg {
				models.ChTg <- in
			}
		}
	}
}
func (b *Bot) SubscribePing(tipPing int) {
	men := b.Db.Subscribe.SubscPing(b.in.NameMention, b.in.Lvlkz, b.in.Config.CorpName, tipPing, b.in.Config.TgChannel)
	if len(men) > 0 {
		men = fmt.Sprintf("Сбор на кз%s\n%s", b.in.Lvlkz, men)
		go b.Tg.SendChannelDelSecond(b.in.Config.TgChannel, men, 600)
	}
}
func (b *Bot) checkAdmin() bool {
	admin := false
	if b.in.Tip == ds {
		admin = b.Ds.CheckAdmin(b.in.Ds.Nameid, b.in.Config.DsChannel)
	} else if b.in.Tip == tg {
		admin = b.Tg.CheckAdminTg(b.in.Config.TgChannel, b.in.Name)
	}
	return admin
}
func (b *Bot) removeDuplicateElementString(mesididid []string) []string {
	result := make([]string, 0, len(mesididid))
	temp := map[string]struct{}{}
	for _, item := range mesididid {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// получаю текущее время
func (b *Bot) currentTime() (string, string) {
	tm := time.Now()
	mdate := (tm.Format("2006-01-02"))
	mtime := (tm.Format("15:04"))

	return mdate, mtime
}

func (b *Bot) SendALLChannel() (bb bool) {
	if b.in.Name == "Mentalisit" {
		re := regexp.MustCompile(`^(Всем|всем)\s([А-Яа-я\s.]+)$`)
		arr := (re.FindAllStringSubmatch(b.in.Mtext, -1))
		if len(arr) > 0 {
			fmt.Println(arr[0])
			bb = true

			text := arr[0][2]

			c := corpsConfig.CorpConfig{}
			d, t, w := c.ReadAllChannel()
			if len(d) > 0 {
				for _, chatds := range d {
					b.Ds.Send(chatds, text)
				}
			}
			if len(t) > 0 {
				for _, chattg := range t {
					b.Tg.SendChannel(chattg, text)
				}
			}
			if len(w) > 0 {
				for _, chatwa := range w {
					fmt.Println("тут нужно сделать отправку на ватс ", chatwa)
				}
			}
		}
	}
	return bb
}
func (b *Bot) hhelp() {
	b.iftipdelete()
	if b.in.Tip == ds {
		go b.Ds.Help(b.in.Config.DsChannel)
	} else if b.in.Tip == tg {
		go b.Tg.Help(b.in.Config.TgChannel)
	}
}
func (b *Bot) Statistic() {
	b.Mu.Lock()
	defer b.Mu.Unlock()

	b.iftipdelete()
	tf := telegraf.Telegraf{}
	tf.InitTelegraf(b.log)
	st := fmt.Sprintf("Статистика игрока %s", b.in.Name)
	content := b.Db.ReadStatistic(b.in.Name)
	mes := tf.CreatePageUserStatistic(st, content)
	if mes != "" {
		go b.ifTipSendTextDelSecond(mes, 30)
	}
	b.Mu.Unlock()
}
func (b *Bot) StatisticA() {
	b.iftipdelete()
	tf := telegraf.Telegraf{}
	tf.InitTelegraf(b.log)
	st := fmt.Sprintf("Статистика игрока %s", "ApplePie")
	content := b.Db.ReadStatistic("ApplePie")
	fmt.Println("contentLEN", content)
	b.log.Println("contentLEN", content)
	b.log.Println("st", st)
	mes := tf.CreatePageUserStatistic(st, content)
	if mes != "" {
		b.ifTipSendTextDelSecond(mes, 30)
	}
}
func (b *Bot) nameMention(u models.Users, tip string) (n1, n2, n3, n4 string) {
	if u.User1.Tip == tip {
		n1 = u.User1.Mention
	} else {
		n1 = u.User1.Name
	}
	if u.User2.Tip == tip {
		n2 = u.User2.Mention
	} else {
		n2 = u.User2.Name
	}
	if u.User3.Tip == tip {
		n3 = u.User3.Mention
	} else {
		n3 = u.User3.Name
	}
	if b.in.Tip == tip {
		n4 = b.in.NameMention
	} else {
		n4 = b.in.Name
	}
	return
}
