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
func (b *Bot) SendIF() {
	name := b.Tg.BotName()

	fmt.Println(" tg", name)
	fmt.Println(" ds", b.Ds.BotName())
}
func (b *Bot) InitBot() {
	in := <-models.ChTg
	fmt.Println(in.Tip)
}

func LogicRs() {

}
