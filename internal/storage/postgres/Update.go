package postgres

import (
	"context"
	"fmt"
)

func (d *Db) MesidTgUpdate(ctx context.Context, mesidtg int, lvlkz string, corpname string) {
	if d.debug {
		fmt.Println("MesidTgUpdate", "mesidtg", mesidtg, "lvlkz", lvlkz, "corpname", corpname)
	}
	upd := `update kzbot.sborkz set tgmesid = $1 where lvlkz = $2 AND corpname = $3 `
	_, err := d.db.Exec(ctx, upd, mesidtg, lvlkz, corpname)
	if err != nil {
		d.log.Error(err.Error())
	}
}
func (d *Db) MesidDsUpdate(ctx context.Context, mesidds, lvlkz, corpname string) {
	if d.debug {
		fmt.Println("MesidDsUpdate", "mesidds", mesidds, "lvlkz", lvlkz, "corpname", corpname)
	}
	ctx = context.Background()
	upd := `update kzbot.sborkz set dsmesid = $1 where lvlkz = $2 AND corpname = $3 `
	_, err := d.db.Exec(ctx, upd, mesidds, lvlkz, corpname)
	if err != nil {
		d.log.Error(err.Error())
	}
}
func (d *Db) UpdateCompliteRS(ctx context.Context, lvlkz string, dsmesid string, tgmesid int, wamesid string, numberkz int, numberevent int, corpname string) {
	if d.debug {
		fmt.Println("UpdateCompliteRS", lvlkz, dsmesid, tgmesid, wamesid, numberkz, numberevent, corpname)
	}
	upd := `update kzbot.sborkz set active = 1,dsmesid = $1,tgmesid = $2,wamesid = $3,numberkz = $4,numberevent = $5 
				where lvlkz = $6 AND corpname = $7 AND active = 0`
	_, err := d.db.Exec(ctx, upd, dsmesid, tgmesid, wamesid, numberkz, numberevent, lvlkz, corpname)
	if err != nil {
		d.log.Error(err.Error())
	}

	updN := `update kzbot.numkz set number=number+1 where lvlkz = $1 AND corpname = $2`
	_, err = d.db.Exec(ctx, updN, lvlkz, corpname)
	if err != nil {
		d.log.Error(err.Error())
	}
	if numberevent > 0 {
		updE := `update kzbot.rsevent set number = number+1  where corpname = $1 AND activeevent = 1`
		_, err = d.db.Exec(ctx, updE, corpname)
		if err != nil {
			d.log.Error(err.Error())
		}
	}
}
