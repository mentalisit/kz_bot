package dbasePostgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
)

func (d *Db) name() {
	exec, err := d.Db.Exec(context.Background(), "SELECT * FROM users")
	if pgErr, ok := err.(*pgconn.PgError); ok {
		fmt.Println(pgErr.Code)
	}
	exec.Select()
}
