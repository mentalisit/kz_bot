package postgres

import (
	"context"
	"fmt"
	"kz_bot/internal/models"
)

func (d *Db) UpdateMitutsQueue(ctx context.Context, name, CorpName string) models.Sborkz {
	if d.debug {
		fmt.Println("UpdateMitutsQueue", name, CorpName)
	}
	sel := "SELECT * FROM kzbot.sborkz WHERE name = $1 AND corpname = $2 AND active = 0"
	results, err := d.db.Query(ctx, sel, name, CorpName)
	if err != nil {
		d.log.Println("Ошибка проверки игрока в очереди для функции (-+) ", err)
	}
	var t models.Sborkz
	for results.Next() {

		err = results.Scan(&t.Id, &t.Corpname, &t.Name, &t.Mention, &t.Tip, &t.Dsmesid, &t.Tgmesid, &t.Wamesid, &t.Time,
			&t.Date, &t.Lvlkz, &t.Numkzn, &t.Numberkz, &t.Numberevent, &t.Eventpoints, &t.Active, &t.Timedown)

		if t.Name == name && t.Timedown <= 3 {
			upd := "update kzbot.sborkz set timedown = timedown + 30 where active = 0 AND name = $1 AND corpname = $2"
			_, err = d.db.Exec(ctx, upd, t.Name, t.Corpname)
			if err != nil {
				d.log.Println("Ошибка обновления времени игрока в очереди для функции (-+) ", err)
			}
		}
	}
	if d.debug {
		fmt.Println("UpdateMitutsQueue", name, CorpName, t)
	}
	return t
}

func (d *Db) MinusMin(ctx context.Context) []models.Sborkz {
	upd := `update kzbot.sborkz set timedown = timedown - 1 where active = 0`
	_, err := d.db.Exec(ctx, upd)
	if err != nil {
		d.log.Println("Ошибка удаления минуты ", err)
	}

	sel := "SELECT * FROM kzbot.sborkz WHERE active = 0"
	results, err := d.db.Query(ctx, sel)
	if err != nil {
		d.log.Println("Ошибка чтения после удаления минуты", err)
	}
	var tt []models.Sborkz
	for results.Next() {
		var t models.Sborkz
		err = results.Scan(&t.Id, &t.Corpname, &t.Name, &t.Mention, &t.Tip, &t.Dsmesid, &t.Tgmesid, &t.Wamesid, &t.Time, &t.Date, &t.Lvlkz, &t.Numkzn, &t.Numberkz, &t.Numberevent, &t.Eventpoints, &t.Active, &t.Timedown)
		tt = append(tt, t)

	}
	return tt
}
