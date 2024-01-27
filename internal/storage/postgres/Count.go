package postgres

import (
	"context"
	"fmt"
	"kz_bot/internal/models"
	"time"
)

func (d *Db) OptimizationSborkz() {
	// Подсчет активных записей и сортировка по имени
	query := `SELECT mention,corpname,lvlkz, SUM(active) AS active_sum FROM kzbot.sborkz GROUP BY corpname, mention,lvlkz ORDER BY mention`
	rows, err := d.db.Query(context.Background(), query)
	if err != nil {
		d.log.Info(err.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		var mention string
		var activeCount int
		var corpname string
		var level string
		if err := rows.Scan(&mention, &corpname, &level, &activeCount); err != nil {
			d.log.Info(err.Error())
			return
		}
		var countNames int
		sel := "SELECT  COUNT(*) as count FROM kzbot.sborkz WHERE mention = $1 AND lvlkz = $2 AND corpname = $3 AND active > 0"
		row := d.db.QueryRow(context.Background(), sel, mention, level, corpname)
		err := row.Scan(&countNames)
		if err != nil {
			d.log.Info(err.Error())
			return
		}
		if countNames > 5 {
			sel := "SELECT * FROM kzbot.sborkz WHERE lvlkz = $1 AND corpname = $2 AND mention = $3"
			results, err := d.db.Query(context.Background(), sel, level, corpname, mention)
			if err != nil {
				d.log.Error(err.Error())
			}
			var t models.Sborkz
			for results.Next() {

				err = results.Scan(&t.Id, &t.Corpname, &t.Name, &t.Mention, &t.Tip, &t.Dsmesid,
					&t.Tgmesid, &t.Wamesid, &t.Time, &t.Date, &t.Lvlkz, &t.Numkzn, &t.Numberkz,
					&t.Numberevent, &t.Eventpoints, &t.Active, &t.Timedown)
			}
			del := "delete from kzbot.sborkz where mention = $1 and corpname = $2 and lvlkz = $3"
			_, err = d.db.Exec(context.Background(), del, mention, corpname, level)
			if err != nil {
				d.log.Error(err.Error())
			}
			tm := time.Now()
			mdate := (tm.Format("2006-01-02"))
			mtime := (tm.Format("15:04"))
			insertSborkztg1 := `INSERT INTO kzbot.sborkz(corpname,name,mention,tip,dsmesid,tgmesid,wamesid,time,date,lvlkz,
		          numkzn,numberkz,numberevent,eventpoints,active,timedown)
				VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16)`
			_, err = d.db.Exec(context.Background(), insertSborkztg1, t.Corpname, t.Name, t.Mention, t.Tip, t.Dsmesid, t.Tgmesid,
				t.Wamesid, mtime, mdate, t.Lvlkz, t.Numkzn, t.Numberkz, t.Numberevent, t.Eventpoints, activeCount, t.Timedown)
			if err != nil {
				d.log.Error(err.Error())
			}
			d.log.Info(fmt.Sprintf("Выполнено сжатие данных игрока %s в корпорации %s кз%s изза %d записей", t.Name, t.Corpname, level, activeCount))
		}
	}

	if err := rows.Err(); err != nil {
		d.log.Info(err.Error())
	}
}
func (d *Db) СountName(ctx context.Context, name, lvlkz, corpName string) (int, error) {
	if d.debug {
		fmt.Println("СountName name, lvlkz, corpName", name, lvlkz, corpName)
	}

	var countNames int
	sel := "SELECT  COUNT(*) as count FROM kzbot.sborkz WHERE name = $1 AND lvlkz = $2 AND corpname = $3 AND active = 0"
	row := d.db.QueryRow(ctx, sel, name, lvlkz, corpName)
	err := row.Scan(&countNames)
	if err != nil {
		d.log.Error(err.Error())
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
		d.log.Error(err.Error())
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
		d.log.Error(err.Error())
		return 0, err
	}
	return countNumberNameActive1, nil
}

func (d *Db) CountNameQueue(ctx context.Context, name string) (countNames int) { //проверяем есть ли игрок в других очередях
	sel := "SELECT  COUNT(*) as count FROM kzbot.sborkz WHERE name = $1 AND active = 0"
	row := d.db.QueryRow(ctx, sel, name)
	err := row.Scan(&countNames)
	if err != nil {
		d.log.Error(err.Error())
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
		d.log.Error(err.Error())
		return 0
	}
	if d.debug {
		fmt.Println("CountNameQueueCorp name, corp", name, corp, countNames)
	}
	return countNames
}
