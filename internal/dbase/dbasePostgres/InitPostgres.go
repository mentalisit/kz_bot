package dbasePostgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	cfg "kz_bot/config"
	corpsConfig "kz_bot/internal/clients/corpConfig"
	"os"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "postgres"
)

type Db struct {
	Db *pgxpool.Pool
	corpsConfig.CorpConfig
	log   *logrus.Logger
	debug bool
}

func (d *Db) InitPostrges(log *logrus.Logger, conf cfg.ConfigBot) error {

	d.log = log
	d.debug = conf.Debug

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), psqlconn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	err = conn.Ping(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	pool, err := pgxpool.Connect(context.Background(), psqlconn)
	if err != nil {
		return err
	}
	d.Db = pool

	return nil
}
