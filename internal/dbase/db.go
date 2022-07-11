package dbase

import (
	"github.com/sirupsen/logrus"
	"kz_bot/config"
	"kz_bot/internal/dbase/dbasePostgres"
)

type Db struct {
	CorpConfig dbasePostgres.CorpConfig
	Emoji      dbasePostgres.Emoji
	Event      dbasePostgres.Event
	Top        dbasePostgres.Top
	Subscribe  dbasePostgres.Subscribe
	Count      dbasePostgres.Count
	Update     dbasePostgres.Update
	dbasePostgres.DbInterface
}

func NewDb(cfg config.ConfigBot, log *logrus.Logger) (Db, error) {
	//db := dbaseMysql.Db{}
	//err := db.InitDB(log, cfg)
	//if err != nil { return Db{}, err }
	dbp := dbasePostgres.Db{}
	errp := dbp.InitPostrges(log, cfg)
	if errp != nil {
		return Db{}, errp
	}
	return Db{
		CorpConfig:  &dbp,
		Emoji:       &dbp,
		Event:       &dbp,
		Top:         &dbp,
		Subscribe:   &dbp,
		Count:       &dbp,
		Update:      &dbp,
		DbInterface: &dbp,
	}, nil
}
