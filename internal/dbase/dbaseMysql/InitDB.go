package dbaseMysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	cfg "kz_bot/config"
	corpsConfig "kz_bot/internal/clients/corpConfig"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Db struct {
	Db *sql.DB
	corpsConfig.CorpConfig
	log *logrus.Logger
}

func dsn(dbName string, conf cfg.ConfigBot) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", conf.Dbusername, conf.DbPassword, conf.DBHostname, dbName)
}

func (d *Db) InitDB(log *logrus.Logger, conf cfg.ConfigBot) {
	d.log = log
	db, err := sql.Open("mysql", dsn("", conf))
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return
	}
	defer db.Close()

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+conf.Dbname)
	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
		return
	}

	no, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when fetching rows", err)
		return
	}

	db.Close()
	db, err = sql.Open("mysql", dsn(conf.Dbname, conf))
	if err != nil {
		log.Printf("Error %s when opening DB", err)
		return
	}
	if no == 1 {
		d.log.Println("Создание таблиц в БД")
		c := CreateTable{db: db}
		c.AllTable()
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("Errors %s pinging DB", err)
		return
	}
	d.Db = db
	//log.Printf("Connected to DB %s successfully\n", dbname)
}
