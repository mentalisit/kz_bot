package supabasesql

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"kz_bot/config"
	corpsConfig "kz_bot/internal/clients/corpConfig"
	"kz_bot/pkg/utils"
	"time"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

const (
	user   = "postgres"
	host   = "db.zaieyyciriiknaixkiyc.supabase.co"
	port   = 5432
	dbname = "postgres"
)

type SupaDB struct {
	Db *pgxpool.Pool
	corpsConfig.CorpConfig
	log   *logrus.Logger
	debug bool
}

func (s *SupaDB) NewClientOld(log *logrus.Logger, conf config.ConfigBot) *SupaDB {
	dns := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, conf.SupabasePass, host, port, dbname)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.Connect(ctx, dns)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &SupaDB{
		Db:         pool,
		CorpConfig: corpsConfig.CorpConfig{},
		log:        log,
		debug:      conf.Debug,
	}

}

func (s *SupaDB) NewClient(ctx context.Context, log *logrus.Logger, maxAttempts int, conf config.ConfigBot) *SupaDB {
	dns := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, conf.SupabasePass, host, port, dbname)
	var pool *pgxpool.Pool
	var err error

	eroorConnect := utils.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.Connect(ctx, dns)
		if err != nil {
			log.Println("Errror Connect DoWithTries ", err)
		}
		return nil
	}, maxAttempts, 5*time.Second)
	if eroorConnect != nil {
		log.Println("Error Connect ", eroorConnect)
		return nil
	}

	return &SupaDB{
		Db:    pool,
		log:   log,
		debug: conf.Debug,
	}
}
