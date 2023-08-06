package postgres

import (
	"context"
	"kz_bot/internal/models"
	"strings"
)

func (d *Db) SubscribePing(ctx context.Context, nameMention, lvlkz, CorpName string, tipPing int, TgChannel string) string {
	var name1, names, men string
	var u models.Users

	if tipPing == 3 {
		//u = d.ReadAll(lvlkz, CorpName)
	}

	sel := "SELECT nameid FROM kzbot.subscribe WHERE lvlkz = $1 AND chatid = $2 AND tip = $3"
	if rows, err := d.db.Query(ctx, sel, lvlkz, TgChannel, tipPing); err == nil {
		for rows.Next() {
			rows.Scan(&name1)
			if nameMention == name1 || u.User1.Mention == name1 ||
				u.User2.Mention == name1 || u.User3.Mention == name1 {
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
func (d *Db) CheckSubscribe(ctx context.Context, name, lvlkz string, TgChannel string, tipPing int) int {
	var counts int
	sel := "SELECT  COUNT(*) as count FROM kzbot.subscribe " +
		"WHERE name = $1 AND lvlkz = $2 AND chatid = $3 AND tip = $4"
	row := d.db.QueryRow(ctx, sel, name, lvlkz, TgChannel, tipPing)
	err := row.Scan(&counts)
	if err != nil {
		d.log.Println("Ошибка проврки активной подписки ", err)
	}
	return counts
}
func (d *Db) Subscribe(ctx context.Context, name, nameMention, lvlkz string, tipPing int, TgChannel string) {
	insertSubscribe := `INSERT INTO kzbot.subscribe (name, nameid, lvlkz, tip, chatid, timestart, timeend) VALUES ($1,$2,$3,$4,$5,$6,$7)`
	_, err := d.db.Exec(ctx, insertSubscribe, name, nameMention, lvlkz, tipPing, TgChannel, "0", "0")
	if err != nil {
		d.log.Println("Ошибка внесения в таблицу подписок ", err)
	}
}
func (d *Db) Unsubscribe(ctx context.Context, name, lvlkz string, TgChannel string, tipPing int) {
	del := "delete from kzbot.subscribe where name = $1 AND lvlkz = $2 AND chatid = $3 AND tip = $4"
	_, err := d.db.Exec(ctx, del, name, lvlkz, TgChannel, tipPing)
	if err != nil {
		d.log.Println("Ошибка удаления подписки с БД", err)
	}
}
