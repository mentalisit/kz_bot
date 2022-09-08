package dbasePostgres

import (
	"context"
	"fmt"
	"time"
)

type Top interface {
	TopTemp() string //временный топ
	TopTempEvent() string
	TopAll(CorpName string) bool
	TopAllEvent(CorpName string, numberevent int) bool
	TopLevel(CorpName, lvlkz string) bool                    //топ по уровню
	TopEventLevel(CorpName, lvlkz string, numEvent int) bool //топ по уровню во время ивента
	TopAllDay(CorpName string, oldDate string) bool
	TopLevelDay(CorpName, lvlkz string, oldDate string) bool
}

func (d *Db) TopLevel(CorpName, lvlkz string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var good = false
	sel := "SELECT name FROM kzbot.sborkz WHERE corpname=$1 AND active=1  AND lvlkz = $2 GROUP BY name LIMIT 40"
	results, err := d.Db.Query(ctx, sel, CorpName, lvlkz)
	if err != nil {
		d.log.Println("Ошибка получения списка участников топа ", err)
	}

	var name string
	for results.Next() {
		err = results.Scan(&name)
		if len(name) > 0 {
			good = true
			countNames, err1 := d.CountNumberNameActive1(lvlkz, CorpName, name)
			if err1 != nil {
				return false
			}

			insertTempTopEvent := `INSERT INTO kzbot.temptopevent(name,numkz,points) VALUES ($1,$2,$3)`
			_, err = d.Db.Exec(ctx, insertTempTopEvent, name, countNames, 0)
			if err != nil {
				d.log.Println("Ошибка внесения сохранения топа ", err.Error())
			}
		}
	}
	return good
}
func (d *Db) TopEventLevel(CorpName, lvlkz string, numEvent int) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var good = false
	sel := "SELECT name FROM kzbot.sborkz WHERE corpname=$1 AND active=1  AND lvlkz = $2 AND numberevent = $3 GROUP BY name LIMIT 40"
	results, err := d.Db.Query(ctx, sel, CorpName, lvlkz, numEvent)
	if err != nil {
		d.log.Println("Ошибка получения списка участников топа ивента ", err)
	}
	var name string
	for results.Next() {
		err = results.Scan(&name)
		if len(name) > 0 {
			good = true
			var countNames int
			selC := "SELECT  COUNT(*) as count FROM kzbot.sborkz WHERE name = $1 AND corpname = $2 AND active = 1 AND numberevent = $3 AND lvlkz = $4"
			row := d.Db.QueryRow(ctx, selC, name, CorpName, numEvent, lvlkz)
			err := row.Scan(&countNames)
			if err != nil {
				d.log.Println("Ошибка подсчета количества походов", err)
			}

			var points int
			selS := "SELECT  SUM(eventpoints) FROM kzbot.sborkz WHERE name = $1 AND corpname = $2 AND active = 1 AND numberevent = $3 AND lvlkz = $4"
			row4 := d.Db.QueryRow(ctx, selS, name, CorpName, numEvent, lvlkz)
			err4 := row4.Scan(&points)
			if err4 != nil {
				d.log.Println("Ошибка подсчета очков ивента ", err4)
			}

			insertTempTopEvent := `INSERT INTO kzbot.temptopevent(name,numkz,points) VALUES ($1,$2,$3)`
			_, err = d.Db.Exec(ctx, insertTempTopEvent, name, countNames, points)
			if err != nil {
				d.log.Println("Ошибка внесения д в в т ", err)
			}
		}
	}
	return good
}
func (d *Db) TopTemp() string {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	number := 1
	var (
		name, message2    string
		numkz, id, points int
	)

	sel := "SELECT * FROM kzbot.temptopevent ORDER BY numkz DESC"
	results, err := d.Db.Query(ctx, sel)
	if err != nil {
		d.log.Println("Ошибка чтения темпТоп ", err)
	}

	for results.Next() {
		results.Scan(&id, &name, &numkz, &points)
		message2 = message2 + fmt.Sprintf("%d. %s - %d \n", number, name, numkz)
		number = number + 1
	}

	_, err = d.Db.Exec(ctx, "DELETE FROM kzbot.temptopevent")
	if err != nil {
		d.log.Println("Ошибка удаления временной таблицы ", err)
	}
	return message2
}
func (d *Db) TopTempEvent() string {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	number := 1
	var (
		name, message2    string
		numkz, id, points int
	)

	sel := "SELECT * FROM kzbot.temptopevent ORDER BY points DESC"
	results, err := d.Db.Query(ctx, sel)
	if err != nil {
		d.log.Println("Ошибка чтения темпТопEvent ", err)
	}

	for results.Next() {
		results.Scan(&id, &name, &numkz, &points)
		message2 = message2 + fmt.Sprintf("%d. %s - %d (%d)\n", number, name, numkz, points)
		number = number + 1
	}

	_, err = d.Db.Exec(ctx, "DELETE FROM kzbot.temptopevent")
	if err != nil {
		d.log.Println("Ошибка удаления временной таблицы ", err)
	}
	return message2
}
func (d *Db) TopAll(CorpName string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	good := false
	sel := "SELECT name FROM kzbot.sborkz WHERE corpname=$1 AND active>0 GROUP BY name LIMIT 40"
	results, err := d.Db.Query(ctx, sel, CorpName)
	if err != nil {
		d.log.Println("Ошибка сканирования имен общего топа ", err)
	}
	var name string
	for results.Next() {
		err = results.Scan(&name)
		if len(name) > 0 {
			good = true
			var countNames int
			selC := "SELECT COALESCE(SUM(active),0) FROM kzbot.sborkz WHERE corpname = $1 AND name = $2 AND active>0"
			//selC := "SELECT  COUNT(*) as count FROM kzbot.sborkz WHERE name = $1 AND corpname = $2 AND active = 1"
			row := d.Db.QueryRow(ctx, selC, CorpName, name)
			err = row.Scan(&countNames)
			if err != nil {
				d.log.Println("Ошибка сканирования количества имен в общем топе ", err)
			}

			insertTempTopEvent := `INSERT INTO kzbot.temptopevent(name,numkz,points) VALUES ($1,$2,$3)`
			_, err = d.Db.Exec(ctx, insertTempTopEvent, name, countNames, 0)
			if err != nil {
				d.log.Println("Ошибка внесения общего топа в временную таблицуи ", err)
			}
		}
	}
	return good
}
func (d *Db) TopAllEvent(CorpName string, numberevent int) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	good := false

	//синтаксична помилка в або поблизу \"ASC\"
	sel := "SELECT name FROM kzbot.sborkz WHERE corpname=$1 AND numberevent = $2 AND active=1 GROUP BY name LIMIT 40"
	results, err := d.Db.Query(ctx, sel, CorpName, numberevent)
	if err != nil {
		d.log.Println("Ошибка запроса топалл эвент", err)
	}

	var name string
	for results.Next() {
		err = results.Scan(&name)
		if len(name) > 0 {
			good = true
			var countNames int
			selC := "SELECT  COUNT(*) as count FROM kzbot.sborkz WHERE name = $1 AND corpname = $2 AND active = 1 AND numberevent = $3"
			row := d.Db.QueryRow(ctx, selC, name, CorpName, numberevent)
			err = row.Scan(&countNames)
			if err != nil {
				d.log.Println("Ошибка запроса топалл эвент количество ", err)
			}
			var points int
			selS := "SELECT  SUM(eventpoints) FROM kzbot.sborkz WHERE name = $1 AND corpname = $2 AND active = 1 AND numberevent = $3"
			row4 := d.Db.QueryRow(ctx, selS, name, CorpName, numberevent)
			err4 := row4.Scan(&points)
			if err4 != nil {
				d.log.Println("Ошибка запроса топалл points", err)
			}

			insertTempTopEvent := `INSERT INTO kzbot.temptopevent(name,numkz,points) VALUES ($1,$2,$3)`
			_, err = d.Db.Exec(ctx, insertTempTopEvent, name, countNames, points)
			if err != nil {
				d.log.Println("Ошибка топалл внесение", err)
			}
		}
	}
	return good
}
func (d *Db) TopAllDay(CorpName string, oldDate string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	good := false
	sel := "SELECT name FROM kzbot.sborkz WHERE corpname=$1 AND date>$2 AND active=1 GROUP BY name LIMIT 40"
	results, err := d.Db.Query(ctx, sel, CorpName, oldDate)
	if err != nil {
		d.log.Println("Ошибка сканирования имен общего топа ", err)
	}
	var name string
	for results.Next() {
		err = results.Scan(&name)
		if len(name) > 0 {
			good = true
			var countNames int
			selC := "SELECT  COUNT(*) as count FROM kzbot.sborkz WHERE name = $1 AND corpname = $2 AND date>$3 AND active = 1"
			row := d.Db.QueryRow(ctx, selC, name, CorpName, oldDate)
			err = row.Scan(&countNames)
			if err != nil {
				d.log.Println("Ошибка сканирования количества имен в общем топе ", err)
			}

			insertTempTopEvent := `INSERT INTO kzbot.temptopevent(name,numkz,points) VALUES ($1,$2,$3)`
			_, err = d.Db.Exec(ctx, insertTempTopEvent, name, countNames, 0)
			if err != nil {
				d.log.Println("Ошибка внесения общего топа в временную таблицуи ", err)
			}
		}
	}
	return good
}
func (d *Db) TopLevelDay(CorpName, lvlkz string, oldDate string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var good = false
	sel := "SELECT name FROM kzbot.sborkz WHERE corpname=$1 AND active=1  AND lvlkz = $2 AND date>? GROUP BY name LIMIT 40"
	results, err := d.Db.Query(ctx, sel, CorpName, lvlkz, oldDate)
	if err != nil {
		d.log.Println("Ошибка получения списка участников топа ", err)
	}

	var name string
	for results.Next() {
		err = results.Scan(&name)
		if len(name) > 0 {
			good = true
			var countNames int
			selC := "SELECT  COUNT(*) as count FROM kzbot.sborkz WHERE name = $1 AND corpname = $2 AND lvlkz=$3 AND date>$4 AND active = 1"
			row := d.Db.QueryRow(ctx, selC, name, CorpName, lvlkz, oldDate)
			err = row.Scan(&countNames)
			if err != nil {
				d.log.Println("Ошибка сканирования количества имен в топе уровня за дату ", err)
			}

			insertTempTopEvent := `INSERT INTO kzbot.temptopevent(name,numkz,points) VALUES ($1,$2,$3)`
			_, err = d.Db.Exec(ctx, insertTempTopEvent, name, countNames, 0)
			if err != nil {
				d.log.Println("Ошибка внесения сохранения топа ", err.Error())
			}
		}
	}
	return good
}
