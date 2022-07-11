package dbasePostgres

import (
	"context"
	"fmt"
	"kz_bot/internal/models"
)

type Emoji interface {
	EmReadUsers(name, tip string) models.EmodjiUser    //чтение эмоджи игрока с бд
	EmUpdateEmodji(name, tip, slot, emo string) string //обновление эмоджи игрока
	EmInsertEmpty(tip, name string)                    // внесение имени для эмоджи
}

func (d *Db) EmReadUsers(name, tip string) models.EmodjiUser {
	selec := "SELECT * FROM kzbot.users WHERE name = $1 AND tip = $2"
	results, err := d.Db.Query(context.Background(), selec, name, tip)
	if err != nil {
		d.log.Println("Ощибка чтения эмоджи с БД", err)
	}
	var t models.EmodjiUser
	for results.Next() {
		err = results.Scan(&t.Id, &t.Tip, &t.Name, &t.Em1, &t.Em2, &t.Em3, &t.Em4)
		if err != nil {
			d.log.Println(err)
		}
	}
	return t
}
func (d *Db) EmUpdateEmodji(name, tip, slot, emo string) string {
	text := ""
	switch slot {
	case "1":
		_, err := d.Db.Exec(context.Background(), `update kzbot.users set em1 = $1 where name = $2 AND tip = $3`, emo, name, tip)
		if err != nil {
			d.log.Println("Ошибка обновления слота эмоджи 1", err)
		}
		text = fmt.Sprintf("Слот %s обновлен\n%s", slot, emo)
	case "2":
		_, err := d.Db.Exec(context.Background(), `update kzbot.users set em2 = $1 where name = $2 AND tip = $3`, emo, name, tip)
		if err != nil {
			d.log.Println("Ошибка обновления слота эмоджи 2", err)
		}
		text = fmt.Sprintf("Слот %s обновлен\n%s", slot, emo)
	case "3":
		_, err := d.Db.Exec(context.Background(), `update kzbot.users set em3 = $1 where name = $2 AND tip = $3`, emo, name, tip)
		if err != nil {
			d.log.Println("Ошибка обновления слота эмоджи ", err)
		}
		text = fmt.Sprintf("Слот %s обновлен\n%s", slot, emo)
	case "4":
		_, err := d.Db.Exec(context.Background(), `update kzbot.users set em4 = $1 where name = $2 AND tip = $3`, emo, name, tip)
		if err != nil {
			d.log.Println("Ошибка обновления слота эмоджи 4", err)
		}
		text = fmt.Sprintf("Слот %s обновлен\n%s", slot, emo)
	}
	return text
}
func (d *Db) EmInsertEmpty(tip, name string) {
	insert := `INSERT INTO kzbot.users(tip,name,em1,em2,em3,em4) VALUES ($1,$2,$3,$4,$5,$6)`
	_, err := d.Db.Exec(context.Background(), insert, tip, name, "", "", "", "")
	if err != nil {
		d.log.Println("Ошибка внесения пользователя для эмоджи ", err)
	}
}
