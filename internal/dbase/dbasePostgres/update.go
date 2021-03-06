package dbasePostgres

import (
	"context"
	"time"
)

type Update interface {
	MesidTgUpdate(mesidtg int, lvlkz string, corpname string)                                                                   //изменение ид сообщения в бд
	MesidDsUpdate(mesidds, lvlkz, corpname string)                                                                              //изменение ид сообщения в бд
	UpdateCompliteRS(lvlkz string, dsmesid string, tgmesid int, wamesid string, numberkz int, numberevent int, corpname string) //закрытие очереди кз
}

func (d *Db) MesidTgUpdate(mesidtg int, lvlkz string, corpname string) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	upd := `update kzbot.sborkz set tgmesid = $1 where lvlkz = $2 AND corpname = $3 `
	_, err := d.Db.Exec(ctx, upd, mesidtg, lvlkz, corpname)
	if err != nil {
		d.log.Println("Ошибка измениния месайди телеги", err)
	}
}
func (d *Db) MesidDsUpdate(mesidds, lvlkz, corpname string) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	upd := `update kzbot.sborkz set dsmesid = $1 where lvlkz = $2 AND corpname = $3 `
	_, err := d.Db.Exec(ctx, upd, mesidds, lvlkz, corpname)
	if err != nil {
		d.log.Println("Ошибка измениния месайди дискорда ", err)
	}
}
func (d *Db) UpdateCompliteRS(lvlkz string, dsmesid string, tgmesid int, wamesid string, numberkz int, numberevent int, corpname string) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	upd := `update kzbot.sborkz set active = 1,dsmesid = $1,tgmesid = $2,wamesid = $3,numberkz = $4,numberevent = $5 
				where lvlkz = $6 AND corpname = $7 AND active = 0`
	_, err := d.Db.Exec(ctx, upd, dsmesid, tgmesid, wamesid, numberkz, numberevent, lvlkz, corpname)
	if err != nil {
		d.log.Println("Ошибка сохранения закрытия очереди", err)
	}
	updN := `update kzbot.numkz set number=number+1 where lvlkz = $1 AND corpname = $2`
	_, err = d.Db.Exec(ctx, updN, lvlkz, corpname)
	if err != nil {
		d.log.Println("Ошибка обновления нумкзз", err)
	}
	if numberevent > 0 {
		updE := `update kzbot.rsevent set number = number+1  where corpname = $1 AND activeevent = 1`
		_, err = d.Db.Exec(ctx, updE, corpname)
		if err != nil {
			d.log.Println("Ошибка обновления номера катки ивента ", err)
		}
	}
}
