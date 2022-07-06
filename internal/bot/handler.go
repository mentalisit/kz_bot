package bot

import "C"
import (
	"fmt"
	"kz_bot/internal/telegraf"
	"regexp"
	"time"

	corpsConfig "kz_bot/internal/clients/corpConfig"
	"kz_bot/internal/models"
)

const (
	ds       = "ds"
	tg       = "tg"
	nickname = "Для того что бы БОТ мог Вас индентифицировать, создайте уникальный НикНей в настройках. Вы можете использовать a-z, 0-9 и символы подчеркивания. Минимальная длина - 5 символов."
)

func (b *Bot) EventText() (string, int) {
	text := ""
	numE := 0
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
			Ds: struct {
				Mesid   string
				Nameid  string
				Guildid string
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
			Option: struct {
				Callback bool
				Edit     bool
				Update   bool
				Queue    bool
			}{
				Callback: true,
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
	men := b.Db.Subscribe.SubscPing(b.in.NameMention, b.in.Lvlkz, b.in.Config.CorpName, tipPing, b.in.Config.TgChannel)
	if len(men) > 0 {
		men = fmt.Sprintf("Сбор на кз%s\n%s", b.in.Lvlkz, men)
		b.Tg.SendChannelDelSecond(b.in.Config.TgChannel, men, 600)
	}
}
func (b *Bot) checkAdmin() bool {
	admin := false
	if b.in.Tip == "ds" {
		admin = b.Ds.CheckAdmin(b.in.Ds.Nameid, b.in.Config.DsChannel)
	} else if b.in.Tip == "tg" {
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
	if b.in.Tip == "ds" {
		b.Ds.Help(b.in.Config.DsChannel)
	} else if b.in.Tip == "tg" {
		b.Tg.Help(b.in.Config.TgChannel)
	}
}
func (b *Bot) Statistic() {
	b.Mutex.Lock()
	b.iftipdelete()
	tf := telegraf.Telegraf{}
	tf.InitTelegraf(b.log)
	st := fmt.Sprintf("Статистика игрока %s", b.in.Name)
	content := b.Db.ReadStatistic(b.in.Name)
	b.log.Println("content", content)
	b.log.Println("st", st)
	mes := tf.CreatePageUserStatistic(st, content)
	if mes != "" {
		b.ifTipSendTextDelSecond(mes, 30)
	}
	b.Mutex.Unlock()
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
