package dbaseMysql

import (
	"fmt"
	"log"
)

func (d *Db) TopLevel(CorpName, lvlkz string) bool {
	var good = false
	results, err := d.Db.Query("SELECT name FROM sborkz WHERE corpname=? AND active=1  AND lvlkz = ? GROUP BY name ASC LIMIT 40",
		CorpName, lvlkz)
	if err != nil {
		fmt.Println("Ошибка получения списка участников топа ", err)
	}

	var name string
	for results.Next() {
		err = results.Scan(&name)
		if len(name) > 0 {
			good = true
			countNames := d.CountNumberNameActive1(lvlkz, CorpName, name)

			insertTempTopEvent := `INSERT INTO temptopevent(name,numkz,points) VALUES (?,?,?)`
			statement, err := d.Db.Prepare(insertTempTopEvent)
			if err != nil {
				fmt.Println("Ошибка внесения топа ", err)
			}
			_, err = statement.Exec(name, countNames, 0)
			if err != nil {
				fmt.Println("Ошибка внесения сохранения топа ", err.Error())
			}
		}
	}
	return good
}
func (d *Db) TopEventLevel(CorpName, lvlkz string, numEvent int) bool {
	var good = false
	results, err := d.Db.Query(
		"SELECT name FROM sborkz WHERE corpname=? AND active=1  AND lvlkz = ? AND numberevent = ?GROUP BY name ASC LIMIT 40",
		CorpName, lvlkz, numEvent)
	if err != nil {
		fmt.Println("Ошибка получения списка участников топа ивента ", err)
	}
	var name string
	for results.Next() {
		err = results.Scan(&name)
		if len(name) > 0 {
			good = true
			var countNames int
			row := d.Db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE name = ? AND corpname = ? AND active = 1 AND numberevent = ? AND lvlkz = ?",
				name, CorpName, numEvent, lvlkz)
			err := row.Scan(&countNames)
			if err != nil {
				fmt.Println("Ошибка подсчета количества походов", err)
			}

			var points int
			row4 := d.Db.QueryRow(
				"SELECT  SUM(eventpoints) FROM sborkz WHERE name = ? AND corpname = ? AND active = 1 AND numberevent = ? AND lvlkz = ?",
				name, CorpName, numEvent, lvlkz)
			err4 := row4.Scan(&points)
			if err4 != nil {
				fmt.Println("Ошибка подсчета очков ивента ", err4)
			}

			insertTempTopEvent := `INSERT INTO temptopevent(name,numkz,points) VALUES (?,?,?)`
			statement, err := d.Db.Prepare(insertTempTopEvent)
			if err != nil {
				fmt.Println("Ошибка внесения данных во временную таблицу", err)
			}
			_, err = statement.Exec(name, countNames, points)
			if err != nil {
				fmt.Println("Ошибка внесения д в в т ", err)
			}
		}
	}
	return good
}
func (d *Db) TopTemp() string {
	number := 1
	var (
		name, message2    string
		numkz, id, points int
	)

	results, err := d.Db.Query("SELECT * FROM temptopevent ORDER BY numkz DESC")
	if err != nil {
		fmt.Println("Ошибка чтения темпТоп ", err)
	}

	for results.Next() {
		results.Scan(&id, &name, &numkz, &points)
		message2 = message2 + fmt.Sprintf("%d. %s - %d \n", number, name, numkz)
		number = number + 1
	}

	_, err = d.Db.Exec("DELETE FROM temptopevent")
	if err != nil {
		fmt.Println("Ошибка удаления временной таблицы ", err)
	}
	return message2
}
func (d *Db) TopTempEvent() string {
	number := 1
	var (
		name, message2    string
		numkz, id, points int
	)

	results, err := d.Db.Query("SELECT * FROM temptopevent ORDER BY poins DESC")
	if err != nil {
		fmt.Println("Ошибка чтения темпТопEvent ", err)
	}

	for results.Next() {
		results.Scan(&id, &name, &numkz, &points)
		message2 = message2 + fmt.Sprintf("%d. %s - %d (%d)\n", number, name, numkz, points)
		number = number + 1
	}

	_, err = d.Db.Exec("DELETE FROM temptopevent")
	if err != nil {
		fmt.Println("Ошибка удаления временной таблицы ", err)
	}
	return message2
}
func (d *Db) TopAll(CorpName string) bool {
	good := false
	results, err := d.Db.Query("SELECT name FROM sborkz WHERE corpname=? AND active=1 GROUP BY name ASC LIMIT 40",
		CorpName)
	if err != nil {
		log.Println("Ошибка сканирования имен общего топа ", err)
	}
	var name string
	for results.Next() {
		err = results.Scan(&name)
		if len(name) > 0 {
			good = true
			var countNames int
			row := d.Db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE name = ? AND corpname = ? AND active = 1",
				name, CorpName)
			err := row.Scan(&countNames)
			if err != nil {
				log.Println("Ошибка сканирования количества имен в общем топе ", err)
			}

			insertTempTopEvent := `INSERT INTO temptopevent(name,numkz,points) VALUES (?,?,?)`
			statement, err := d.Db.Prepare(insertTempTopEvent)
			if err != nil {
				log.Println("Ошибка внесения общего топа в временную таблицу ", err)
			}
			_, err = statement.Exec(name, countNames, 0)
			if err != nil {
				log.Println("Ошибка внесения общего топа в временную таблицуи ", err)
			}
		}
	}
	return good
}
func (d *Db) TopAllEvent(CorpName string, numberevent int) bool {
	good := false
	results, err := d.Db.Query("SELECT name FROM sborkz WHERE corpname=? AND numberevent = ? AND active=1 GROUP BY name ASC LIMIT 40",
		CorpName, numberevent)
	if err != nil {
		log.Println("Ошибка запроса топалл эвент", err)
	}

	var name string
	for results.Next() {
		err = results.Scan(&name)
		if len(name) > 0 {
			good = true
			var countNames int
			row := d.Db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE name = ? AND corpname = ? AND active = 1 AND numberevent = ?",
				name, CorpName, numberevent)
			err := row.Scan(&countNames)
			if err != nil {
				log.Println("Ошибка запроса топалл эвент количество ", err)
			}
			var points int
			row4 := d.Db.QueryRow("SELECT  SUM(eventpoints) FROM sborkz WHERE name = ? AND corpname = ? AND active = 1 AND numberevent = ?",
				name, CorpName, numberevent)
			err4 := row4.Scan(&points)
			if err4 != nil {
				log.Println("Ошибка запроса топалл points", err)
			}

			insertTempTopEvent := `INSERT INTO temptopevent(name,numkz,points) VALUES (?,?,?)`
			statement, err := d.Db.Prepare(insertTempTopEvent)
			if err != nil {
				fmt.Println("Ошибка запроса топалл эвент подготовка", err)
			}
			_, err = statement.Exec(name, countNames, points)
			if err != nil {
				fmt.Println("Ошибка топалл внесение", err)
			}
		}
	}
	return good
}
func (d *Db) TopAllDay(CorpName string, oldDate string) bool {
	good := false
	results, err := d.Db.Query("SELECT name FROM sborkz WHERE corpname=? AND date>? AND active=1 GROUP BY name ASC LIMIT 40",
		CorpName, oldDate)
	if err != nil {
		log.Println("Ошибка сканирования имен общего топа ", err)
	}
	var name string
	for results.Next() {
		err = results.Scan(&name)
		if len(name) > 0 {
			good = true
			var countNames int
			row := d.Db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE name = ? AND corpname = ? AND date>? AND active = 1",
				name, CorpName, oldDate)
			err := row.Scan(&countNames)
			if err != nil {
				log.Println("Ошибка сканирования количества имен в общем топе ", err)
			}

			insertTempTopEvent := `INSERT INTO temptopevent(name,numkz,points) VALUES (?,?,?)`
			statement, err := d.Db.Prepare(insertTempTopEvent)
			if err != nil {
				log.Println("Ошибка внесения общего топа в временную таблицу ", err)
			}
			_, err = statement.Exec(name, countNames, 0)
			if err != nil {
				log.Println("Ошибка внесения общего топа в временную таблицуи ", err)
			}
		}
	}
	return good
}
func (d *Db) TopLevelDay(CorpName, lvlkz string, oldDate string) bool {
	var good = false
	results, err := d.Db.Query("SELECT name FROM sborkz WHERE corpname=? AND active=1  AND lvlkz = ? AND date>? GROUP BY name ASC LIMIT 40",
		CorpName, lvlkz, oldDate)
	if err != nil {
		fmt.Println("Ошибка получения списка участников топа ", err)
	}

	var name string
	for results.Next() {
		err = results.Scan(&name)
		if len(name) > 0 {
			good = true
			var countNames int
			row := d.Db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE name = ? AND corpname = ? AND lvlkz=? AND date>? AND active = 1",
				name, CorpName, lvlkz, oldDate)
			err := row.Scan(&countNames)
			if err != nil {
				log.Println("Ошибка сканирования количества имен в топе уровня за дату ", err)
			}

			insertTempTopEvent := `INSERT INTO temptopevent(name,numkz,points) VALUES (?,?,?)`
			statement, err := d.Db.Prepare(insertTempTopEvent)
			if err != nil {
				fmt.Println("Ошибка внесения топа ", err)
			}
			_, err = statement.Exec(name, countNames, 0)
			if err != nil {
				fmt.Println("Ошибка внесения сохранения топа ", err.Error())
			}
		}
	}
	return good
}
