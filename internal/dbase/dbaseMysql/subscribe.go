package dbaseMysql

import (
	"kz_bot/internal/models"
)

type Subscribe interface {
	CheckSubscribe(name, lvlkz string, TgChannel int64, tipPing int) int                //проверка активной подписки
	Subscribe(name, nameMention, lvlkz string, tipPing int, TgChannel int64)            //подписка
	Unsubscribe(name, lvlkz string, TgChannel int64, tipPing int)                       //отписка
	SubscPing(nameMention, lvlkz, CorpName string, tipPing int, TgChannel int64) string //чтение для пинга игроков в телеграм
}

func (d *Db) SubscPing(nameMention, lvlkz, CorpName string, tipPing int, TgChannel int64) string {
	var name1, names, men string
	var u models.Users
	if tipPing == 3 {
		u = d.ReadAll(lvlkz, CorpName)
	}

	if rows, err := d.Db.Query("SELECT nameid FROM subscribe WHERE lvlkz = ? AND chatid = ? AND tip = ?",
		lvlkz, TgChannel, tipPing); err == nil {
		for rows.Next() {
			rows.Scan(&name1)
			if nameMention == name1 || u.User1.Mention == name1 || u.User2.Mention == name1 || u.User3.Mention == name1 {
				continue
			}
			names = name1 + " "
			men = names + men
		}
		rows.Close()
	}
	return men
}
func (d *Db) CheckSubscribe(name, lvlkz string, TgChannel int64, tipPing int) int {
	var counts int
	row := d.Db.QueryRow("SELECT  COUNT(*) as count FROM subscribe WHERE name = ? AND lvlkz = ? AND chatid = ? AND tip = ?",
		name, lvlkz, TgChannel, tipPing)
	err := row.Scan(&counts)
	if err != nil {
		d.log.Println("Ошибка проврки активной подписки ", err)
	}
	return counts
}
func (d *Db) Subscribe(name, nameMention, lvlkz string, tipPing int, TgChannel int64) {
	insertSubscribe := `INSERT INTO subscribe (name, nameid, lvlkz, tip, chatid, timestart, timeend) VALUES (?,?,?,?,?,?,?)`
	statement, err := d.Db.Prepare(insertSubscribe)
	_, err = statement.Exec(name, nameMention, lvlkz, tipPing, TgChannel, 0, 0)
	if err != nil {
		d.log.Println("Ошибка внесения в таблицу подписок ", err)
	}
}
func (d *Db) Unsubscribe(name, lvlkz string, TgChannel int64, tipPing int) {
	_, err := d.Db.Exec("delete from subscribe where name = ? AND lvlkz = ? AND chatid = ? AND tip = ?",
		name, lvlkz, TgChannel, tipPing)
	if err != nil {
		d.log.Println("Ошибка удаления подписки с БД", err)
	}
}
