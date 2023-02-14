package postgres

import (
	"context"
	"fmt"
)

func (d *Db) СountName(ctx context.Context, name, lvlkz, corpName string) (int, error) {
	if d.debug {
		fmt.Println("СountName name, lvlkz, corpName", name, lvlkz, corpName)
	}

	var countNames int
	sel := "SELECT  COUNT(*) as count FROM kzbot.sborkz WHERE name = $1 AND lvlkz = $2 AND corpname = $3 AND active = 0"
	row := d.db.QueryRow(ctx, sel, name, lvlkz, corpName)
	err := row.Scan(&countNames)
	if err != nil {
		d.log.Println("Ошибка проверки в очереди ли игрок  ", err)
		return 0, err
	}
	if d.debug {
		fmt.Println("СountName ", corpName)
	}
	return countNames, nil
}
func (d *Db) CountQueue(ctx context.Context, lvlkz, CorpName string) (int, error) { //проверка сколько игровок в очереди
	if d.debug {
		fmt.Println("CountQueue lvlkz, CorpName", lvlkz, CorpName)
	}
	var count int
	sel := "SELECT  COUNT(*) as count FROM kzbot.sborkz WHERE lvlkz = $1 AND corpname = $2 AND active = 0"
	row := d.db.QueryRow(ctx, sel, lvlkz, CorpName)
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
func (d *Db) CountNumberNameActive1(ctx context.Context, lvlkz, CorpName, name string) (int, error) { // выковыриваем из базы значение количества походов на кз
	if d.debug {
		fmt.Println("CountNumberNameActive1 lvlkz, CorpName, name", lvlkz, CorpName, name)
	}
	var countNumberNameActive1 int
	sel := "SELECT COALESCE(SUM(active),0) FROM kzbot.sborkz WHERE lvlkz = $1 AND corpname = $2 AND name = $3"
	//COALESCE(SUM(value), 0)
	row := d.db.QueryRow(ctx, sel, lvlkz, CorpName, name)
	err := row.Scan(&countNumberNameActive1)
	if err != nil {
		d.log.Println("Ошибка чтения количества игр", err)
		return 0, err
	}
	return countNumberNameActive1, nil
}

func (d *Db) CountNameQueue(ctx context.Context, name string) (countNames int) { //проверяем есть ли игрок в других очередях
	sel := "SELECT  COUNT(*) as count FROM kzbot.sborkz WHERE name = $1 AND active = 0"
	row := d.db.QueryRow(ctx, sel, name)
	err := row.Scan(&countNames)
	if err != nil {
		d.log.Println("Ошибка проверки игрока в других очередях ", err)
	}
	if d.debug {
		fmt.Println("CountNameQueue name", name, countNames)
	}
	return countNames
}
func (d *Db) CountNameQueueCorp(ctx context.Context, name, corp string) (countNames int) { //проверяем есть ли игрок в других очередях
	sel := "SELECT  COUNT(*) as count FROM kzbot.sborkz WHERE name = $1 AND corpname = $2 AND active = 0"
	row := d.db.QueryRow(ctx, sel, name, corp)
	err := row.Scan(&countNames)
	if err != nil {
		d.log.Println("Ошибка проверки игрока в других очередях этой корпы ", err)
		return 0
	}
	if d.debug {
		fmt.Println("CountNameQueueCorp name, corp", name, corp, countNames)
	}
	return countNames
}
