package bot

import (
	"context"
	"fmt"
	"kz_bot/internal/models"
	"strconv"
)

//lang ok
//wats lang not ok

func (b *Bot) MinusMin() {
	tt := b.storage.Timers.MinusMin(context.Background())

	if len(tt) > 0 {
		for _, t := range tt {
			if t.Corpname != "" {
				ok, config := b.CheckCorpNameConfig(t.Corpname)
				if ok {
					time := strconv.Itoa(t.Timedown)

					in := models.InMessage{
						Mtext:       "",
						Tip:         t.Tip,
						Name:        t.Name,
						NameMention: t.Mention,
						Lvlkz:       t.Lvlkz,
						Timekz:      time,
						Ds: struct {
							Mesid   string
							Nameid  string
							Guildid string
							Avatar  string
						}{
							Mesid:   t.Dsmesid,
							Nameid:  "",
							Guildid: config.Guildid,
						},
						Tg: struct {
							Mesid  int
							Nameid int64
						}{
							Mesid:  t.Tgmesid,
							Nameid: 0,
						},
						Wa: struct {
							Nameid string
							Mesid  string
						}{
							Nameid: "",
							Mesid:  t.Wamesid,
						},
						Config: config,
						Option: models.Option{
							MinusMin: true,
							Edit:     true},
					}
					b.inbox <- in

					if b.debug {
						fmt.Printf("\n  MinusMin []models.Sborkz %+v\n\n", t)
					}
				}
			}
		}
		b.UpdateMessage()
	}
}
func (b *Bot) UpdateMessage() {
	corpActive0 := b.storage.DbFunc.OneMinutsTimer(context.Background())
	for _, corp := range corpActive0 {

		_, config := b.CheckCorpNameConfig(corp)

		dss, tgs, _ := b.storage.DbFunc.MessageUpdateMin(context.Background(), corp)

		if config.DsChannel != "" {
			for _, d := range dss {
				a := b.storage.DbFunc.MessageupdateDS(context.Background(), d, config)
				b.inbox <- a
			}
		}
		if config.TgChannel != 0 {
			for _, t := range tgs {
				a := b.storage.DbFunc.MessageupdateTG(context.Background(), t, config)
				b.inbox <- a
			}
		}
		if config.WaChannel != "" {
			//тут будет логика ватса
			//fmt.Println(was)
			//не будет пока нет редактировать сообщение
		}
	}
}

func (b *Bot) CheckTimeQueue() {
	atoi, err := strconv.Atoi(b.in.Timekz)
	if err != nil {
		b.log.Println("ошибка CheckTimeQueue Atoi", err)
		return
	}
	if atoi == 3 {
		text := b.in.NameMention + b.GetLang("VremyaPochtiVishlo")
		if b.in.Tip == ds {
			mID := b.client.Ds.SendEmbedTime(b.in.Config.DsChannel, text)
			go b.client.Ds.DeleteMesageSecond(b.in.Config.DsChannel, mID, 182)
		} else if b.in.Tip == tg {
			mID := b.client.Tg.SendEmbedTime(b.in.Config.TgChannel, text)
			go b.client.Tg.DelMessageSecond(b.in.Config.TgChannel, mID, 182)
		} else if b.in.Tip == wa {
			b.client.Wa.SendMention(b.in.Config.WaChannel, b.in.NameMention+" время почти вышло...\n"+
				"Для продления времени ожидания на 30м пиши +\n"+
				"Для выхода из очереди пиши -", []string{b.in.Wa.Nameid})
		}
	} else if atoi == 0 {
		b.RsMinus()
	} else if atoi < -1 {
		b.RsMinus()
	} else if atoi < 0 {
		b.RsMinus()
	}
}
