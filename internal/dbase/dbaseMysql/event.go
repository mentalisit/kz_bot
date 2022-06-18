package dbaseMysql

import (
	"database/sql"
	"fmt"
	"kz_bot/internal/models"
	"log"
)

func (d *Db) UpdatePoints(CorpName string, numberkz, points, event1 int) int {
	// считаем количество участников КЗ опр уровня
	var countEvent int
	row := d.Db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE numberevent = ? AND corpname=? AND numberkz=?  AND active=1",
		event1, CorpName, numberkz)
	err := row.Scan(&countEvent)
	if err != nil {
		fmt.Println("Ошибка получения количествой участников катки ", err)
	}
	pointsq := points / countEvent
	//вносим очки
	_, err = d.Db.Exec(`update sborkz set eventpoints=? WHERE numberevent = ? AND corpname =? AND numberkz=? AND active=1`,
		pointsq, event1, CorpName, numberkz)
	if err != nil {
		log.Println("Ошибка внесения очков катки ", err)
	}
	return countEvent
}
func (d *Db) ReadNamesMessage(CorpName string, numberkz, numberEvent int) (nd, nt models.Names, t models.Sborkz) {
	var name string
	results, err := d.Db.Query("SELECT * FROM sborkz WHERE corpname=? AND numberkz=? AND numberevent = ? AND active=1",
		CorpName, numberkz, numberEvent)
	if err != nil {
		fmt.Println("ошибка извлечения для изменения сообщения катки ", err)
	}

	num := 1
	for results.Next() {
		err = results.Scan(&t.Id, &t.Corpname, &t.Name, &t.Mention, &t.Tip, &t.Dsmesid, &t.Tgmesid, &t.Wamesid, &t.Time, &t.Date, &t.Lvlkz, &t.Numkzn, &t.Numberkz, &t.Numberevent, &t.Eventpoints, &t.Active, &t.Timedown)
		if t.Tip == "ds" {
			name = t.Mention
		} else {
			name = t.Name
		}
		if num == 1 {
			nd.Name1 = name
		} else if num == 2 {
			nd.Name2 = name
		} else if num == 3 {
			nd.Name3 = name
		} else if num == 4 {
			nd.Name4 = name
		}
		if t.Tip == "tg" {
			name = t.Mention
		} else {
			name = t.Name
		}
		if num == 1 {
			nt.Name1 = name
		} else if num == 2 {
			nt.Name2 = name
		} else if num == 3 {
			nt.Name3 = name
		} else if num == 4 {
			nt.Name4 = name
		}
		num = num + 1
	}
	return nd, nt, t
}
func (d *Db) CountEventNames(CorpName, name string, numberkz, numEvent int) (countEventNames int) {
	row := d.Db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE corpname = ? AND numberkz=?  AND active=1 AND name=? AND numberevent = ?",
		CorpName, numberkz, name, numEvent)
	err := row.Scan(&countEventNames)
	if err != nil {
		fmt.Println("Ошибка получения количества участников определенной кз для ивента ", err)
	}
	return countEventNames
}
func (d *Db) CountEventsPoints(CorpName string, numberkz, numberEvent int) int {
	var countEventPoints int
	row := d.Db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE corpname=? AND numberkz=? AND numberevent = ? AND active=1 AND eventpoints > 0",
		CorpName, numberkz, numberEvent)
	err := row.Scan(&countEventPoints)
	if err != nil {
		fmt.Println("Ошибка проверки внесены ли очки по катке ивента ", err)
	}
	return countEventPoints
}
func (d *Db) NumActiveEvent(CorpName string) (event1 int) { //запрос номера ивента
	row := d.Db.QueryRow("SELECT numevent FROM rsevent WHERE corpname=? AND activeevent=1 ORDER BY numevent DESC LIMIT 1",
		CorpName)
	err := row.Scan(&event1)
	if err != nil {
		if err == sql.ErrNoRows {
			event1 = 0
		} else {
			fmt.Println("Ошибка получения номера ивента ", err)
		}
	}
	return event1
}
func (d *Db) NumDeactivEvent(CorpName string) (event0 int) { //запрос номера последнего ивента
	row := d.Db.QueryRow("SELECT numevent FROM rsevent WHERE corpname=? AND activeevent=0 ORDER BY numevent DESC LIMIT 1",
		CorpName)
	err := row.Scan(&event0)
	if err != nil {
		fmt.Println("Ошибка проверки прошлого номера ивента ", err)
	}
	return event0
}
func (d *Db) UpdateActiveEvent0(CorpName string, event1 int) {
	_, err := d.Db.Exec("UPDATE rsevent SET activeevent=0 WHERE corpname=? AND numevent=?",
		CorpName, event1)
	if err != nil {
		fmt.Println("Ошибка обновления активИвент ", err)
	}
}
func (d *Db) EventStartInsert(CorpName string) {
	event0 := d.NumDeactivEvent(CorpName)
	if event0 > 0 {
		numberevent := event0 + 1
		insertEvent := `INSERT INTO rsevent (corpname,numevent,activeevent,number) VALUES (?,?,?,?)`
		_, err := d.Db.Exec(insertEvent, CorpName, numberevent, 1, 1)
		if err != nil {
			log.Println("Ошибка внесения старта ивента ", err)
		}
	} else {
		insertEvent := `INSERT INTO rsevent (corpname,numevent,activeevent,number) VALUES (?,?,?,?)`
		_, err := d.Db.Exec(insertEvent, CorpName, 1, 1, 1)
		if err != nil {
			log.Println("Ошибка внесения старта ивента0 ", err)
		}
	}
}
