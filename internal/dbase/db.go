package dbase

import (
	"github.com/sirupsen/logrus"
	"kz_bot/config"
	"kz_bot/internal/dbase/dbaseMysql"
)

type Db struct {
	CorpConfig dbaseMysql.CorpConfig
	Emoji      dbaseMysql.Emoji
	Event      dbaseMysql.Event
	Top        dbaseMysql.Top
	Subscribe  dbaseMysql.Subscribe
	dbaseMysql.DbInterface
}

func NewDb(cfg config.ConfigBot, log *logrus.Logger) (Db, error) {
	db := dbaseMysql.Db{}
	err := db.InitDB(log, cfg)
	if err != nil {
		return Db{}, err
	}
	return Db{
		CorpConfig:  &db,
		Emoji:       &db,
		Event:       &db,
		Top:         &db,
		Subscribe:   &db,
		DbInterface: &db,
	}, nil
}
