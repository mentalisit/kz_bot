package postgresqlS

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"kz_bot/internal/config"
	"kz_bot/pkg/utils"
	"time"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, log *logrus.Logger, maxAttempts int, conf *config.ConfigBot) (*pgxpool.Pool, error) {
	var pool *pgxpool.Pool
	var err error
	dns := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		conf.Supabase.Username, conf.Supabase.Password,
		conf.Supabase.Host, conf.Supabase.Port, conf.Supabase.Name)

	err = utils.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.Connect(ctx, dns)
		if err != nil {
			log.Println("Errror Connect DoWithTries postgresqlS ", err)
			return err
		}
		return nil
	}, maxAttempts, 5*time.Second)
	if err != nil {
		log.Println("Error Connect postgresqlS", err)
		return nil, err
	}

	return pool, nil
}
