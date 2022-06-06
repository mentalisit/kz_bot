package bot

import (
	"fmt"
	"kz_bot/internal/clients/ds"
	Tg "kz_bot/internal/clients/tg"
	"kz_bot/internal/dbase/dbaseMysql"
	"kz_bot/internal/models"
)

type Bot struct {
	Tg *Tg.Telegram
	Ds *ds.Ds
	Db *dbaseMysql.Db
}

func NewBot(tg Tg.Telegram, ds ds.Ds, db dbaseMysql.Db) *Bot {
	return &Bot{Tg: &tg, Ds: &ds, Db: &db}
}
func (b *Bot) InitBot() {
	for {
		select {
		case in := <-models.ChTg:
			b.LogicRs(in)
		case in := <-models.ChDs:
			b.LogicRs(in)

		}
	}
}

func (b Bot) LogicRs(in models.InMessage) {
	fmt.Println(in.Name, "пишет")

}
