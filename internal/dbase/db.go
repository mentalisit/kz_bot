package dbase

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"kz_bot/config"
	"kz_bot/internal/dbase/dbaseMysql"
	"kz_bot/internal/dbase/dbasePostgres"
	"kz_bot/internal/dbase/supabasesql"
	"kz_bot/internal/models"
	"log"
)

type Db struct {
	CorpConfig *supabasesql.SupaDB
	Emoji      dbasePostgres.Emoji
	Event      dbasePostgres.Event
	Top        dbasePostgres.Top
	Subscribe  dbasePostgres.Subscribe
	Count      dbasePostgres.Count
	Update     dbasePostgres.Update
	dbasePostgres.DbInterface
}

func NewDb(cfg config.ConfigBot, log *logrus.Logger) (Db, error) {
	dbs := supabasesql.SupaDB{}
	//dbs.NewClientOld(log, cfg)
	dbs.NewClient(context.Background(), log, 5, cfg)

	dbp := dbasePostgres.Db{}
	dbp.InitPostrges(log, cfg)

	//migrate(db, dbp)
	//migrateConfig(db, dbp)
	//migrateConfigSupa(dbs, dbp)

	return Db{
		CorpConfig:  &dbs,
		Emoji:       &dbp,
		Event:       &dbp,
		Top:         &dbp,
		Subscribe:   &dbp,
		Count:       &dbp,
		Update:      &dbp,
		DbInterface: &dbp,
	}, nil
}
func migrateConfigSupa(dbs supabasesql.SupaDB, dbp dbasePostgres.Db) {
	results, err := dbp.Db.Query(context.Background(), `SELECT * FROM kzbot.config`)
	if err != nil {
		log.Println("Ошибка чтения крнфигурации корпораций", err)
	}
	var t models.TableConfig
	for results.Next() {
		err = results.Scan(&t.Id, &t.Corpname, &t.Dschannel, &t.Tgchannel, &t.Wachannel, &t.Mesiddshelp, &t.Mesidtghelp, &t.Delmescomplite, &t.Guildid)
		insertConfig := `INSERT INTO kzbot.config (corpname,dschannel,tgchannel,wachannel,mesiddshelp,mesidtghelp,delmescomplite,guildid)VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`
		_, err := dbs.Db.Exec(context.Background(), insertConfig, t.Corpname, t.Dschannel, t.Tgchannel, t.Wachannel, t.Mesiddshelp, t.Mesidtghelp, t.Delmescomplite, t.Guildid)
		if err != nil {
			log.Println("Ошибка внесения конфигурации ", err)
		}
	}
}

func migrateConfig(dbm dbaseMysql.Db, dbp dbasePostgres.Db) {
	results, err := dbp.Db.Query(context.Background(), `SELECT * FROM kzbot.config`)
	if err != nil {
		log.Println("Ошибка чтения крнфигурации корпораций", err)
	}
	var t models.TableConfig
	for results.Next() {
		err = results.Scan(&t.Id, &t.Corpname, &t.Dschannel, &t.Tgchannel, &t.Wachannel, &t.Mesiddshelp, &t.Mesidtghelp, &t.Delmescomplite, &t.Guildid)
		insertConfig := `INSERT INTO config (corpname,dschannel,tgchannel,wachannel,mesiddshelp,mesidtghelp,delmescomplite,guildid) 
						VALUES (?,?,?,?,?,?,?,?)`
		statement, err := dbm.Db.Prepare(insertConfig)
		if err != nil {
			log.Println("Ошибка подготовки внесения в бд конфигурации ", err)
		}
		_, err = statement.Exec(t.Corpname, t.Dschannel, t.Tgchannel, t.Wachannel, t.Mesiddshelp, t.Mesidtghelp, t.Delmescomplite, t.Guildid)
		if err != nil {
			log.Println("Ошибка внесения конфигурации ", err)
		}
	}
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
