package dbasePostgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"time"
)

type Count interface {
	СountName(name, lvlkz, corpName string) (int, error)
	CountNameQueue(name string) (countNames int)
	CountNameQueueCorp(name, corp string) (countNames int)
	CountQueue(lvlkz, CorpName string) (int, error)
	CountNumberNameActive1(lvlkz, CorpName, name string) (int, error)
}

func (d *Db) СountName(name, lvlkz, corpName string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if d.debug {
		fmt.Println("СountName name, lvlkz, corpName", name, lvlkz, corpName)
	}

	var countNames int
	sel := "SELECT  COUNT(*) as count FROM kzbot.sborkz WHERE name = $1 AND lvlkz = $2 AND corpname = $3 AND active = 0"
	row := d.Db.QueryRow(ctx, sel, name, lvlkz, corpName)
	err := row.Scan(&countNames)
	if err != nil {
		d.log.Println("Ошибка проверки в очереди ли игрок  ", err)
		d.log.Println("name, lvlkz, corpName", name, lvlkz, corpName)
		return 0, err
	}
	if d.debug {
		fmt.Println("СountName ", corpName)
	}
	return countNames, nil
}
func (d *Db) CountQueue(lvlkz, CorpName string) (int, error) { //проверка сколько игровок в очереди
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if d.debug {
		fmt.Println("CountQueue lvlkz, CorpName", lvlkz, CorpName)
	}
	var count int
	sel := "SELECT  COUNT(*) as count FROM kzbot.sborkz WHERE lvlkz = $1 AND corpname = $2 AND active = 0"
	row := d.Db.QueryRow(ctx, sel, lvlkz, CorpName)
	err := row.Scan(&count)
	if err != nil {
		d.log.Println("Ошибка проверки количества игроков в очереди", err)
		return 0, err
	}
	if d.debug {
		fmt.Println("CountQueue ", count)
	}
	return count, nil
}
func (d *Db) CountNumberNameActive1(lvlkz, CorpName, name string) (int, error) { // выковыриваем из базы значение количества походов на кз
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if d.debug {
		fmt.Println("CountNumberNameActive1 lvlkz, CorpName, name", lvlkz, CorpName, name)
	}
	var countNumberNameActive1 int
	sel := "SELECT COALESCE(SUM(active),0) FROM kzbot.sborkz WHERE lvlkz = $1 AND corpname = $2 AND name = $3"
	//COALESCE(SUM(value), 0)
	row := d.Db.QueryRow(ctx, sel, lvlkz, CorpName, name)
	err := row.Scan(&countNumberNameActive1)
	if err != nil {
		d.log.Println("Ошибка чтения количества игр", err)
		return 0, err
	}
	return countNumberNameActive1, nil
}

func (d *Db) CountNameQueue(name string) (countNames int) { //проверяем есть ли игрок в других очередях
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	sel := "SELECT  COUNT(*) as count FROM kzbot.sborkz WHERE name = $1 AND active = 0"
	row := d.Db.QueryRow(ctx, sel, name)
	err := row.Scan(&countNames)
	if err != nil {
		d.log.Println("Ошибка проверки игрока в других очередях ", err)
	}
	if d.debug {
		fmt.Println("CountNameQueue name", name, countNames)
	}
	return countNames
}
func (d *Db) CountNameQueueCorp(name, corp string) (countNames int) { //проверяем есть ли игрок в других очередях
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	sel := "SELECT  COUNT(*) as count FROM kzbot.sborkz WHERE name = $1 AND corpname = $2 AND active = 0"
	row := d.Db.QueryRow(ctx, sel, name, corp)
	err := row.Scan(&countNames)
	if err != nil {
		d.log.Println("Ошибка проверки игрока в других очередях этой корпы ", err)
	}
	if d.debug {
		fmt.Println("CountNameQueueCorp name, corp", name, corp, countNames)
	}
	return countNames
}
func (d *Db) selectEerrrorPGX(err error) {
	if err == pgx.ErrNoRows {
		fmt.Println("err==pgx.ErrNoRows")
	} else if err == pgx.ErrTxClosed {
		fmt.Println("err==pgx.ErrTxClosed")
	} else if err == pgx.ErrTxCommitRollback {
		fmt.Println("err==pgx.ErrTxCommitRollback")
	} else if err == pgx.ErrInvalidLogLevel {
		fmt.Println("err==pgx.ErrInvalidLogLevel")
	} else {
		fmt.Println("ошибка отсутствует в pgc.err ", err)
	}
}
