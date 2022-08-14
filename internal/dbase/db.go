package dbase

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"kz_bot/config"
	"kz_bot/internal/dbase/dbaseMysql"
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

func NewDb(cfg config.ConfigBot, log *logrus.Logger, debug bool) (Db, error) {
	//db := dbaseMysql.Db{}
	//err := db.InitDB(log, cfg)
	//if err != nil {return Db{}, err}
	dbp := dbasePostgres.Db{}
	errp := dbp.InitPostrges(log, cfg, debug)
	if errp != nil {
		return Db{}, errp
	}

	//migrate(db, dbp)

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
func migrate(dbm dbaseMysql.Db, dbp dbasePostgres.Db) {
	results, err := dbm.Db.Query("SELECT name,lvlkz,corpname,tip FROM sborkz WHERE active=1 GROUP BY name,lvlkz")
	if err != nil {
		fmt.Println(err)
	}
	var name string
	var level string
	var numActive int
	var Corpname string
	var tip string
	for results.Next() {
		err = results.Scan(&name, &level, &Corpname, &tip)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(name, level)

		row := dbm.Db.QueryRow(
			"SELECT SUM(active) FROM sborkz WHERE name = ? AND lvlkz = ? AND active = 1", name, level)
		err = row.Scan(&numActive)
		if err != nil {
			fmt.Println("Ошибка сканирования количества иMен", err)
		}
		fmt.Printf("%s был на Кз%s  %dраз ", name, level, numActive)

		insertSborkztg1 := `INSERT INTO kzbot.sborkz(corpname,name,mention,tip,dsmesid,tgmesid,wamesid,time,date,lvlkz,
                   numkzn,numberkz,numberevent,eventpoints,active,timedown) 
				VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16)`
		_, err = dbp.Db.Exec(context.Background(), insertSborkztg1, Corpname, name, "", tip, "", 0,
			"", "old", "до 11.07.2022", level, 0, 0, 0, 0, numActive, 0)
		if err != nil {
			fmt.Println("Ошибка записи старта очереди", err)
		}
	}
}
