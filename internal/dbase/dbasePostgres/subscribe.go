package dbasePostgres

import (
	"context"
	"kz_bot/internal/models"
	"strconv"
	"strings"
	"time"
)

type Subscribe interface {
	CheckSubscribe(name, lvlkz string, TgChannel int64, tipPing int) int                //проверка активной подписки
	Subscribe(name, nameMention, lvlkz string, tipPing int, TgChannel int64)            //подписка
	Unsubscribe(name, lvlkz string, TgChannel int64, tipPing int)                       //отписка
	SubscPing(nameMention, lvlkz, CorpName string, tipPing int, TgChannel int64) string //чтение для пинга игроков в телеграм
}

func (d *Db) SubscPing(nameMention, lvlkz, CorpName string, tipPing int, TgChannel int64) string {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var name1, names, men string
	var u models.Users
	if tipPing == 3 {
		u = d.ReadAll(lvlkz, CorpName)
	}

	TgChannelt := strconv.FormatInt(TgChannel, 10)
	sel := "SELECT nameid FROM kzbot.subscribe WHERE lvlkz = $1 AND chatid = $2 AND tip = $3"
	if rows, err := d.Db.Query(ctx, sel, lvlkz, TgChannelt, tipPing); err == nil {
		for rows.Next() {
			rows.Scan(&name1)
			if nameMention == name1 || u.User1.Mention == name1 || u.User2.Mention == name1 || u.User3.Mention == name1 {
				continue
			}
			names = name1 + ", "
			men = names + men
		}
		rows.Close()
	}
	men = strings.TrimSuffix(men, ", ")
	return men
}
func (d *Db) CheckSubscribe(name, lvlkz string, TgChannel int64, tipPing int) int {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var counts int
	TgChannelt := strconv.FormatInt(TgChannel, 10)
	sel := "SELECT  COUNT(*) as count FROM kzbot.subscribe WHERE name = $1 AND lvlkz = $2 AND chatid = $3 AND tip = $4"
	row := d.Db.QueryRow(ctx, sel, name, lvlkz, TgChannelt, tipPing)
	err := row.Scan(&counts)
	if err != nil {
		d.log.Println("Ошибка проврки активной подписки ", err)
	}
	return counts
}
func (d *Db) Subscribe(name, nameMention, lvlkz string, tipPing int, TgChannel int64) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	TgChannelt := strconv.FormatInt(TgChannel, 10)
	insertSubscribe := `INSERT INTO kzbot.subscribe (name, nameid, lvlkz, tip, chatid, timestart, timeend) VALUES ($1,$2,$3,$4,$5,$6,$7)`
	_, err := d.Db.Exec(ctx, insertSubscribe, name, nameMention, lvlkz, tipPing, TgChannelt, "0", "0")
	if err != nil {
		d.log.Println("Ошибка внесения в таблицу подписок ", err)
	}
}
func (d *Db) Unsubscribe(name, lvlkz string, TgChannel int64, tipPing int) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	TgChannelt := strconv.FormatInt(TgChannel, 10)
	del := "delete from kzbot.subscribe where name = $1 AND lvlkz = $2 AND chatid = $3 AND tip = $4"
	_, err := d.Db.Exec(ctx, del, name, lvlkz, TgChannelt, tipPing)
	if err != nil {
		d.log.Println("Ошибка удаления подписки с БД", err)
	}
}
