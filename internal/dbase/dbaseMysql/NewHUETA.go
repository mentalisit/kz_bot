package dbaseMysql

import (
	"fmt"

	"kz_bot/internal/models"
)

func (d *Db) ReadAllTable() (users string) {
	results, err := d.Db.Query("SELECT name,tip,time FROM sborkz")
	if err != nil {
		fmt.Println("Ошибка чтения активной очереди readall", err)
	}
	for results.Next() {
		var t models.Sborkz
		err = results.Scan(&t.Name, &t.Tip, &t.Time)
		s := fmt.Sprintf("\n"+
			"Имя %s Тип %s время %s"+
			"\n", t.Name, t.Tip, t.Time)
		users = users + s
	}
	return users
}
