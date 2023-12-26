package postgres

import (
	"context"
	"fmt"
	"kz_bot/internal/models"
)

func (d *Db) EmReadUsers(ctx context.Context, name, tip string) models.EmodjiUser {
	selec := "SELECT * FROM kzbot.users WHERE name = $1 AND tip = $2"
	results, err := d.db.Query(ctx, selec, name, tip)
	if err != nil {
		d.log.Println("Ощибка чтения эмоджи с БД", err)
	}
	var t models.EmodjiUser
	for results.Next() {
		err = results.Scan(&t.Id, &t.Tip, &t.Name, &t.Em1, &t.Em2, &t.Em3, &t.Em4)
		if err != nil {
			d.log.Println("EmReadUsers", err)
		}
	}
	return t
}
func (d *Db) EmUpdateEmodji(ctx context.Context, name, tip, slot, emo string) string {
	sqlUpd := fmt.Sprintf(`update kzbot.users set em%s = $1 where name = $2 AND tip = $3`, slot)
	_, err := d.db.Exec(ctx, sqlUpd, emo, name, tip)
	if err != nil {
		d.log.Println(fmt.Sprintf("Ошибка обновления слота эмоджи %s", slot), err)
	}
	return fmt.Sprintf("Слот %s обновлен\n%s", slot, emo)
}
func (d *Db) EmInsertEmpty(ctx context.Context, tip, name string) {
	insert := `INSERT INTO kzbot.users(tip,name,em1,em2,em3,em4) VALUES ($1,$2,$3,$4,$5,$6)`
	_, err := d.db.Exec(ctx, insert, tip, name, "", "", "", "")
	if err != nil {
		d.log.Println("Ошибка внесения пользователя для эмоджи ", err)
	}
}
